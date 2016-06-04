package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ruiclarateixeira/csv2json"
	"github.com/segmentio/go-prompt"
	"gopkg.in/mgo.v2"
	"log"
)

func main() {
	fileName := flag.String("file", "", "Input File Path")
	url := flag.String("url", "", "Mongo URL")
	database := flag.String("db", "", "Target Mongo Db Name")
	collection := flag.String("coll", "", "Collection Name")
	username := flag.String("user", "", "Mongo Username (Optional)")
	password := flag.String("pwd", "", "Mongo Password (Optional). If username is provided and password is not the user will be prompted for it")

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

	if *username != "" && *password == "" {
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
