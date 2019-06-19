package pgstate

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "db"
	port     = 5432
	user     = "duckhunt"
	password = "duckhunt"
	dbname   = "duckhunt"
)

type PgState struct {
	db *sql.DB
}

func (pgState *PgState) Connect() error {
	log.Println("pgstate in connect")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	pgState.db = db
	return nil
}

func (pgState *PgState) GetPartyIds(userId int, partyName string) ([]int, error) {
	var parties []int
	// rows, err := pgState.db.Query(`SELECT pId FROM parties WHERE pId=$1`, userId)
	rows, err := pgState.db.Query(`SELECT parties.pId FROM parties INNER JOIN parties_name ON parties.pId = parties_name.pId  WHERE parties.uId=$1 AND parties_name.pName=$2`, userId, partyName)
	if err != nil {
		return parties, err
	}
	defer rows.Close()
	for rows.Next() {
		var PId int
		err = rows.Scan(&PId)
		if err != nil {
			return parties, err
		}
		parties = append(parties, PId)
	}
	err = rows.Err()
	if err != nil {
		return parties, err
	}
	return parties, nil
}

func (pgState *PgState) CreateParty(userId int, partyName string) ([]int, error) {
	var parties []int
	rows, err := pgState.db.Query(`INSERT INTO parties (uId) VALUES ($1) RETURNING pId`, userId)
	if err != nil {
		return parties, err
	}
	defer rows.Close()
	for rows.Next() {
		var PId int
		err = rows.Scan(&PId)
		if err != nil {
			return parties, err
		}
		pgState.db.Query(`INSERT INTO parties_name (pId, pName) VALUES ($1, $2)`, PId, partyName)
		parties = append(parties, PId)
	}
	err = rows.Err()
	if err != nil {
		return parties, err
	}
	return parties, nil
}
