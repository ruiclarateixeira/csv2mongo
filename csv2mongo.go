package main

import (
	"encoding/json"
	"github.com/ruiclarateixeira/csv2json"
	"gopkg.in/mgo.v2"
	"log"
)

func main() {
	fileName := "/Users/ruijorgeclarateixeira/Development/CodingProjects.csv"
	url := "127.0.0.1"
	database := "myapps"
	collection := "apps"

	result, Err := csv2json.ReadFile(fileName)

	if Err != nil {
		log.Fatal(Err)
	}

	session, Err := mgo.Dial(url)

	if Err != nil {
		log.Fatal(Err)
	}

	jsonResult := make([]map[string]string, 0)
	Err = json.Unmarshal(result, &jsonResult)

	if Err != nil {
		log.Fatal(Err)
	}

	c := session.DB(database).C(collection)

	for _, document := range jsonResult {
		Err = c.Insert(document)

		if Err != nil {
			log.Fatal(Err)
		}
	}
}
