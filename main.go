package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"Bot"`
}

func ReadConfig() error {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}

var BotId string
var goBot *discordgo.Session

func Run() error {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	user, err := goBot.User("@me")

	if err != nil {
		log.Fatal(err)
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
	err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = Run()
	if err != nil {
		log.Fatal(err)
	}

	// keep the program running.
	select {}
}
