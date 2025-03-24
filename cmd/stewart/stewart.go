package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/bwmarrin/discordgo"
	"github.com/servusdei2018/shards/v2"
	log "github.com/sirupsen/logrus"
)

var s * discordgo.Session
var mgr *shards.Manager
var accessToken string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not read .env")
	}

	accessToken = os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("Access token is required.")
	}
}

func main() {
	mgr, err := shards.New("Bot " + accessToken)
	if err != nil {
		log.Fatal("Error creating manager: ", err)
	}

	mgr.AddHandler(onConnect)
	mgr.AddHandler(voiceStateUpdate)

	mgr.RegisterIntent(discordgo.IntentsGuildMessages)
	mgr.RegisterIntent(discordgo.IntentsGuildVoiceStates)
	mgr.RegisterIntent(discordgo.IntentsGuildPresences)

	err = mgr.Start()
	if err != nil {
		log.Fatal("Error starting manager: ", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Manager
	log.Println("Stopping shard manager...")
	mgr.Shutdown()
	log.Println("Shard manager stopped. Bot is shut down.")
}

func onConnect(s *discordgo.Session, evt *discordgo.Connect) {
	log.Printf("Shart #%v connected.", s.ShardID)
}

func voiceStateUpdate (s *discordgo.Session, evt *discordgo.VoiceStateUpdate) {
	log.Printf("Shard #%v received voice state update: %v", s.ShardID, evt)
}