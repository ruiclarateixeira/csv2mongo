package main

import (
	"log"
	"flag"
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"github.com/ruiclarateixeira/csv2json"
)

func main() {
	fileName := flag.String("file", "", "Input File Path")
	url := flag.String("url", "", "Mongo URL")
	database := flag.String("db", "", "Target Mongo Db Name")
	collection := flag.String("coll", "", "Collection Name")

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

	c := session.DB(*database).C(*collection)

	for _, document := range jsonResult {
		Err = c.Insert(document)

		if Err != nil {
			log.Fatal(Err)
		}
	}
}