package database

import (
	"github.com/boltdb/bolt"
)

type MyDB struct {
	db *bolt.DB
}

var myDB *MyDB

func GetDB() *MyDB {
	if myDB == nil {
		// fmt.Println("test")
		myDB = &MyDB{}
		myDB.db, _ = bolt.Open("./database/my.db", 0600, nil)
	}
	return myDB
}

func (myDB *MyDB) QueryPeople(strId string) string {
	return queryString(myDB, strId, "People")

}

func (myDB *MyDB) QueryFilm(strId string) string {
	return queryString(myDB, strId, "Film")

}

func (myDB *MyDB) QueryPlanet(strId string) string {
	return queryString(myDB, strId, "Planet")
}

func (myDB *MyDB) QuerySpecies(strId string) string {
	return queryString(myDB, strId, "Species")
}

func (myDB *MyDB) QueryStarship(strId string) string {
	return queryString(myDB, strId, "Starship")
}

func (myDB *MyDB) QueryVehicle(strId string) string {
	return queryString(myDB, strId, "Vehicle")
}
func queryString(myDB *MyDB, strId string, model string) string {
	db := myDB.db
	var data string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(model))
		v := b.Get([]byte(strId))
		data = string(v)
		return nil
	})

	return data
}
