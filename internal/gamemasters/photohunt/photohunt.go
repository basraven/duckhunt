package photohunt

import (
	"duckhunt/internal/storage/state/pgstate"
	"fmt"
	"log"
)

type Photohunt struct {
	Channels   []string
	StateStore string
	Store      *pgstate.PgState
}

func (photohunt *Photohunt) JoinParty(userId int, partyName string) error {
	log.Println("Joining Party" + partyName)
	var parties []int
	var err error
	parties, err = photohunt.Store.GetPartyIds(userId, partyName)
	if err != nil {
		return err
	}
	if len(parties) < 1 {
		parties, err = photohunt.Store.CreateParty(userId, partyName)
	}
	if err != nil {
		return err
	}
	fmt.Println(parties)
	return nil
}

func (photohunt *Photohunt) Init() error {
	// FIXME: this can be nicer
	if photohunt.StateStore == "postgress" {
		log.Println("Gamemaster postgress")
		photohunt.Store = &pgstate.PgState{}
		photohunt.Store.Connect()
	}
	return nil
}
