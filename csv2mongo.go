package main

import (
	"fmt"
	"log"
	"flag"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"github.com/segmentio/go-prompt"
	"github.com/ruiclarateixeira/csv2json"
)

func main() {
	fileName := flag.String("file", "", "Input File Path")
	url := flag.String("url", "", "Mongo URL")
	database := flag.String("db", "", "Target Mongo Db Name")
	collection := flag.String("coll", "", "Collection Name")
	username := flag.String("user", "", "Mongo Username")
	password := flag.String("pwd", "", "Mongo Password")

	flag.Parse()

	if *fileName == "" || *url == "" || *database == "" || *collection == "" {
		flag.Usage()
		fmt.Println("Bad Inputs:", flag.Args())
		return
	}

	result, Err := csv2json.ReadFile(*fileName)

	if Err != nil {
		log.Fatal(Err)
	}

	session, Err := mgo.Dial(*url)

	if Err != nil {
		log.Fatal(Err)
	}

	jsonResult := make([]map[string]string, 0)
	Err = json.Unmarshal(result, &jsonResult)

	if Err != nil {
		log.Fatal(Err)
	}

	db := session.DB(*database)

	if *username != ""  && *password == "" {
		input := prompt.Password("Please enter the password for " + *username)
		password = &input
	}
	
	db.Login(*username, *password)
	
	for _, document := range jsonResult {
		Err = db.C(*collection).Insert(document)

		if Err != nil {
			log.Fatal(Err)
		}
	}
}
