package main

import (
	"fmt"
	"gobot/config"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Embed struct {
	*discordgo.MessageEmbed
}

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
		session.ChannelMessageSend(message.ChannelID, "!help - displays this message\n!ping - pong\n!heure - affiche les heures des bros")
	}
	if message.Content == "!ping" {
		session.ChannelMessageSend(message.ChannelID, "pong")
	}
	if message.Content == "!heure" {
		gigamessage := discordgo.MessageEmbed{
			Title:       "Heure",
			Description: "L'heure de tes maxis potes",
			Color:       0x00ff00,
		}

		// get local time in Paris
		t := time.Now().In(time.FixedZone("CET", 7200))

		// get montréal time

		t2 := time.Now().In(time.FixedZone("EST", -18000))

		// get ireland time

		t3 := time.Now().In(time.FixedZone("IST", 3600))

		// add fields to embed
		gigamessage.Fields = []*discordgo.MessageEmbedField{
			{
				// put local time in Paris
				Name:   "France",
				Value:  t.Format("15:04:05"),
				Inline: true,
			},
			{
				Name:   "Canada",
				Value:  t2.Format("15:04:05"),
				Inline: true,
			},
			{
				Name:   "Irlande",
				Value:  t3.Format("15:04:05"),
				Inline: true,
			},
		}

		session.ChannelMessageSendEmbed(message.ChannelID, &gigamessage)

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
