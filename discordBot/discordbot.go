package discordbot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/GOpcy/Zdrapywacz/databaseconf"
	"github.com/bwmarrin/discordgo"
)

func GetDiscord() (*discordgo.Session, error) {
	return discordgo.New(fmt.Sprintf("Bot %s",
		os.Getenv("BOT_ID"),
	))
}

func RunBot() {

	discord, err := GetDiscord()
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "World")
		}
	})

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	fmt.Println("bot is online")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

func BotPing(o *databaseconf.Offer) {
	techs := map[string]string{
		"vue":        "557644663213785089",
		"typescript": "557643834473971714",
		"golang":     "557653889705574400",
		"svelte":     "557654204856926257",
		"node":       "611509538100740112",
		"javascript": "557644165450694686",
		"nuxt":       "557644526269759513",
		"laravel":    "648942050854830083",
	}

	discord, err := GetDiscord()
	if err != nil {
		log.Fatal(err)
	}

	e := discordgo.MessageEmbed{
		URL:   o.URL,
		Title: o.Title,
		Color: 0xFF0000,
		Description: fmt.Sprintf("**%s**,\n %s, \n %s, \n %s",
			o.Tags,
			o.Company,
			o.Experience,
			o.OperatingMode),
	}

	checkTags := func() string {
		for tech := range techs {
			if strings.Contains(strings.ToLower(o.Tags), tech) {
				return tech
			}
		}
		return ""
	}

	result := checkTags()

	if result != "" {
		roleID := techs[result]
		discord.ChannelMessageSend("1337156063799279616", fmt.Sprintf("<@&%s>", roleID))
	}

	discord.ChannelMessageSend("1337156063799279616", "@everyone !!!NEW OFFER!!!:")
	discord.ChannelMessageSendEmbed("1337156063799279616", &e)
}
