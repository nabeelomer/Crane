package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var token string
	flag.StringVar(&token, "token", "", "Discord Bot Token")
	flag.Parse()

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
		return
	}

	discord.AddHandler(handleStorm)

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Discord Up!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func handleStorm(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if strings.Contains(strings.ToLower(m.Content), "pingstorm") {
		if len(m.Mentions) > 0 {
			var v = ""
			for _, m := range m.Mentions {
				v = v + " " + m.Mention()
			}
			for i := 0; i < 5; i++ {
				go s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" has pinged "+v)
				time.Sleep(1 * time.Second)
			}
		} else {
			for i := 0; i < 5; i++ {
				go s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" has pinged @here")
				time.Sleep(1 * time.Second)
			}
		}
	}
}
