package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/GOpcy/Zdrapywacz/databaseconf"
	discordbot "github.com/GOpcy/Zdrapywacz/discordBot"
	"github.com/go-co-op/gocron/v2"
	"github.com/gocolly/colly"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func main() {
	//envs
	err := godotenv.Load()
  	if err != nil {
    log.Fatal("Error loading .env file")
  }
	//Database
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),
	)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	panic("failed to connect database")
	}else{
		fmt.Println("--!Database Migrated!--")
	}

	db.AutoMigrate(&databaseconf.Offer{})

	
	c := colly.NewCollector(
		colly.AllowedDomains("justjoin.it", "www.justjoin.it"),

		//colly.CacheDir("./justjoinit_cache"),

		colly.MaxDepth(2),
		colly.Async(true),
	)
	discordbot.RunBot();

	detailCollector := c.Clone()

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	

	j, err := s.NewJob(
		gocron.DurationJob(
			time.Minute,
		),
		gocron.NewTask(
			func() {
				runScrape(c, detailCollector, db)
			},
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	//runScrape(c, detailCollector, db);

	select {
	}
	
}

func runScrape(c *colly.Collector, detailCollector *colly.Collector, db *gorm.DB){
	var sum int64 = 0;
	var m sync.Mutex;

	//offers := make([]Offer, 0, 200)

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("https://api.justjoin.it/v2/user-panel/offers?experienceLevels[]=junior&page=2&sortBy=published&orderBy=DESC&perPage=100&salaryCurrencies=PLN", r.URL.String())
	})

	c.OnHTML("div[data-test-id=virtuoso-item-list]", func(e *colly.HTMLElement) {
		
		offers := e.ChildAttrs("a[href]", "href")
		for _, offer := range offers {
			finalURL := e.Request.AbsoluteURL(offer)
			detailCollector.Visit(finalURL)
		}
		
		//fmt.Println(offers)
	})

	detailCollector.OnHTML(`div.css-tnvghs`, func(e *colly.HTMLElement){
		
		title := e.ChildText("h1")
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		offer := databaseconf.Offer{
			Title: title,
			URL: e.Request.URL.String(),
			Company: e.ChildText("h2"),
			Location: e.ChildText(".css-1o4wo1x"),
		}
		
		e.ForEach(".css-if24yw > div", func(_ int, el *colly.HTMLElement){
			target := el.Text
			m.Lock()
			switch target{
			case "Experience": 
				offer.Experience = el.DOM.Next().Text()
			case "Operating mode": 
				offer.OperatingMode = el.DOM.Next().Text()
			}	
			m.Unlock()
		})
		m.Lock()
		sum++
		fmt.Println(offer.URL," | ", offer.Company, " | ", offer.Location, " | ", offer.Title, " | ", offer.Experience, " | ", offer.OperatingMode);
		result := db.FirstOrCreate(&offer)
		if result.Error != nil{
			panic("failed to create a record in db")
		}
		if result.RowsAffected > 0{
			discordbot.BotPing(&offer)
		}
		m.Unlock()
		fmt.Println(sum)
	})
	c.Visit("https://justjoin.it/job-offers/all-locations?experience-level=junior&orderBy=DESC&sortBy=published&from=0")

	c.Wait()
	detailCollector.Wait()
}