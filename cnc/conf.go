package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var Methodz map[string]*Attackz = make(map[string]*Attackz)

type toyotav struct {
	Toyota []audiv `json:"methods"`
}

type audiv struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Vip         bool   `json:"vip"`
	Premium     bool   `json:"premium"`
	Home        bool   `json:"home"`
	APIShit     struct {
		API string `json:"url"`
	} `json:"Links"`
}

func Configure() {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotav

	json.Unmarshal(byteValue, &toyotas)

	for i := 0; i < len(toyotas.Toyota); i++ {
		log.Println(" [Method Found] | Name \"" + toyotas.Toyota[i].Name + "\"")
		Methodz[toyotas.Toyota[i].Name] = &Attackz{
			Name:        toyotas.Toyota[i].Name,
			Description: toyotas.Toyota[i].Description,
			Enabled:     toyotas.Toyota[i].Enabled,
			Vip:         toyotas.Toyota[i].Vip,
			Premium:     toyotas.Toyota[i].Premium,
			Home:        toyotas.Toyota[i].Home,
			API:         toyotas.Toyota[i].APIShit.API,
		}
	}
}
