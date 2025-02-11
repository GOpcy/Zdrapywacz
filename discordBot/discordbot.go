package discordbot

import (
	"fmt"
	"log"
	"os"

	"github.com/GOpcy/Zdrapywacz/databaseconf"
	"github.com/bwmarrin/discordgo"
)

func GetDiscord() (*discordgo.Session, error){
	return discordgo.New(fmt.Sprintf("Bot %s",
	os.Getenv("BOT_ID"),
	))
}

func RunBot(){

	discord, err := GetDiscord()
	if err!=nil{
		log.Fatal(err)
	}
	
	
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
		if m.Author.ID == s.State.User.ID{
			return
		}

		if m.Content == "hello"{
			s.ChannelMessageSend(m.ChannelID, "World")
		}
	})

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil{
		log.Fatal(err)
	}
	defer discord.Close()

	fmt.Println("bot is online")


}

func BotPing(o *databaseconf.Offer){
	discord, err := GetDiscord()
	if err!=nil{
		log.Fatal(err)
	}

	e := discordgo.MessageEmbed{
		URL: o.URL,
		Title: o.Title,
		Color: 0xFF0000,
		Description: fmt.Sprintf("%s, \n %s, \n %s",
		o.Company,
		o.Experience,
		o.OperatingMode),
	}
	discord.ChannelMessageSend("1337156063799279616", "@everyone !!!NEW OFFER!!!:")
	discord.ChannelMessageSendEmbed("1337156063799279616", &e)
}