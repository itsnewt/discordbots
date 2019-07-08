package main

// imports
import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// declare variables
var (
	commandPrefix string
	botID         string
)

// main function
func main() {
	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create a new bot instance
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD"))
	errCheck("Error creating the discord session", err)

	// get the bot user
	user, err := discord.User("@me")
	errCheck("Error retrieving account", err)

	botID = user.ID

	// add the command handler and ready event
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "with your emotions.")

		if err != nil {
			fmt.Println("Error attempting to set my status.")
		}

		servers := discord.State.Guilds
		fmt.Printf("Bot is ready and has started on %d servers", len(servers))
	})

	// open a connection to Discord
	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	// set the bot's prefix
	commandPrefix = "go!"

	// allow the function to idle
	<-make(chan struct{})
}

// error checker function
func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

// command handler
func commandHandler(discord *discordgo.Session, msg *discordgo.MessageCreate) {
	// if the user is a bot, ignore the messahe
	if msg.Author.ID == botID || msg.Author.Bot {
		return
	}

	// get the content of the message
	content := msg.Content

	// make a ping command
	if content == commandPrefix+"ping" {
		discord.ChannelMessageSend(msg.ChannelID, "pong!")
	}
}
