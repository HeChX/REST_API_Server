package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Howlyao/Server/model"

	"github.com/boltdb/bolt"
)

type Field struct {
	Fields map[string]interface{}
	Pk     int
}

var db, err = bolt.Open("../database/my.db", 0600, nil)

func main() {
	creataBucket()
	initPlanets()
	initPeople()
	initFilms()

	// db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("People"))
	// 	v := b.Get([]byte("1"))
	// 	fmt.Println(string(v))
	// 	return nil
	// })
}

func initPeople() {
	data, err := ioutil.ReadFile("resourses/people.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertPeople(&fields[i])
	}

}

func initFilms() {
	data, err := ioutil.ReadFile("resourses/films.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertFilm(&fields[i])
	}

}

func initPlanets() {
	data, err := ioutil.ReadFile("resourses/planets.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertPlanet(&fields[i])
	}

}

func creataBucket() {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("People"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Film"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Planet"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func insertPlanet(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	m["url"] = "http://localhost:8080/planets/" + strId

	data, _ := json.Marshal(&m)
	planet := model.Planet{}
	_ = json.Unmarshal(data, &planet)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Planet"))
		data, _ := json.Marshal(&planet)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func insertPeople(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	homeworld := int(m["homeworld"].(float64))
	delete(m, "homeworld")
	m["url"] = "http://localhost:8080/people/" + strId
	m["homeworld"] = "http://localhost:8080/planet/" + strconv.Itoa(homeworld)

	data, _ := json.Marshal(&m)
	people := model.People{}
	_ = json.Unmarshal(data, &people)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("People"))
		data, _ := json.Marshal(&people)
		err := b.Put([]byte(strId), data)
		return err
	})

	planetUpdate(strconv.Itoa(homeworld), m["url"].(string), true)

}

func insertFilm(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	m["url"] = "http://localhost:8080/films/" + strId
	characters := m["characters"].([]interface{})
	planets := m["planets"].([]interface{})

	var charactersURL []string
	var planetsURL []string
	for i := 0; i < len(characters); i++ {
		cid := strconv.Itoa(int(characters[i].(float64)))
		url := "http://localhost:8080/people/" + cid
		charactersURL = append(charactersURL, url)
		peopleUpdate(cid, m["url"].(string))
	}
	delete(m, "characters")
	m["characters"] = charactersURL

	for i := 0; i < len(planets); i++ {
		pid := strconv.Itoa(int(planets[i].(float64)))
		url := "http://localhost:8080/planets/" + pid
		planetsURL = append(planetsURL, url)
		planetUpdate(pid, m["url"].(string), false)
	}
	delete(m, "planets")
	m["planets"] = planetsURL

	data, _ := json.Marshal(&m)
	film := model.Film{}
	_ = json.Unmarshal(data, &film)

	//updata people film url
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Film"))
		data, _ := json.Marshal(&film)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func peopleUpdate(strId string, url string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("People"))
		v := b.Get([]byte(strId))
		people := model.People{}
		_ = json.Unmarshal(v, &people)

		people.Films = append(people.Films, url)
		data, _ := json.Marshal(&people)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func planetUpdate(strId string, url string, flag bool) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Planet"))
		v := b.Get([]byte(strId))
		planet := model.Planet{}
		_ = json.Unmarshal(v, &planet)
		if flag == true {
			planet.Residents = append(planet.Residents, url)
		} else {
			planet.Films = append(planet.Films, url)
		}
		data, _ := json.Marshal(&planet)
		err := b.Put([]byte(strId), data)
		return err
	})

}
