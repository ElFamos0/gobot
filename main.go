package main

import (
	"fmt"
	"gobot/config"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Run() error {
	config.InitEnv()

	Token := os.Getenv("DISCORD_TOKEN")
	goBot, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error creating Discord session")
		return err
	}
	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println("Login error: " + err.Error())
		return err
	}
	BotId = user.ID
	// register the messageCreate func as a callback for MessageCreate events
	goBot.AddHandler(messageCreate)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	return nil
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotId {
		return
	}
	if message.Content == "!help" {
		session.ChannelMessageSend(message.ChannelID, "!help - displays this message\n!ping - pong")
	}
	if message.Content == "!ping" {
		session.ChannelMessageSend(message.ChannelID, "pong")
	}

}

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
	// Wait here until CTRL-C or other termnal signal is received.

	<-make(chan struct{})
}
