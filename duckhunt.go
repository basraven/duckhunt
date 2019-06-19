package main

import (
	"duckhunt/internal/gamemasters/photohunt"
	"log"
)

func main() {
	defaultChannels := []string{"telegram", "messenger"}

	log.Printf("Starting duckhunt\n")
	game := &photohunt.Photohunt{Channels: defaultChannels, StateStore: "postgress"}
	game.Init()
	err := game.JoinParty(100, "party123")
	if err != nil {
		log.Fatal(err)
	}
	// for i := 0; i < len(defaultChannels); i++ {
	// 	log.Println(defaultChannels[i])
	// }
}
