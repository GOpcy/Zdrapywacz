package discordbot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func RunBot(){
	discord, err := discordgo.New("Bot MTMzNzA5MzEyMTc3MDc4MjcyMA.GeDa7n.y1qoFImHJ8ugZOmUIH2onLNM_6fsPffmPHudZs")
	if err !=nil{
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

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}