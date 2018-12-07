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
	db := myDB.db
	var data string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("People"))
		v := b.Get([]byte(strId))
		data = string(v)
		return nil
	})

	return data

}

func (myDB *MyDB) QueryFilm(strId string) string {
	db := myDB.db
	var data string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Film"))
		v := b.Get([]byte(strId))
		data = string(v)
		return nil
	})

	return data

}

func (myDB *MyDB) QueryPlanet(strId string) string {
	db := myDB.db
	var data string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Planet"))
		v := b.Get([]byte(strId))
		data = string(v)
		return nil
	})

	return data

}
