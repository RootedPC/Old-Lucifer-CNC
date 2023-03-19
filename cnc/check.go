package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type toyotaq struct {
	Toyota []audiq `json:"methods"`
}

type audiq struct {
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

func CheckJSONMethod(method string) bool {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotaq

	json.Unmarshal(byteValue, &toyotas)
	for i := 0; i < len(toyotas.Toyota); i++ {
		if len(method) == 0 {
			return false
		} else if method == strings.ToUpper(toyotas.Toyota[i].Name) || method == strings.ToLower(toyotas.Toyota[i].Name) {
			return true
		}
	}
	return false
}

func CheckJSONURL(method string) string {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotaq

	json.Unmarshal(byteValue, &toyotas)
	for i := 0; i < len(toyotas.Toyota); i++ {
		if method == strings.ToUpper(toyotas.Toyota[i].Name) || method == strings.ToLower(toyotas.Toyota[i].Name) {
			return toyotas.Toyota[i].APIShit.API
		}
	}
	return ""
}

func CheckJSONEnabled(method string) bool {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotaq

	json.Unmarshal(byteValue, &toyotas)
	for i := 0; i < len(toyotas.Toyota); i++ {
		if len(method) == 0 {
			return false
		} else if method == strings.ToUpper(toyotas.Toyota[i].Name) || method == strings.ToLower(toyotas.Toyota[i].Name) {
			if toyotas.Toyota[i].Enabled == true {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func CheckJSONVip(method string) bool {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotaq

	json.Unmarshal(byteValue, &toyotas)
	for i := 0; i < len(toyotas.Toyota); i++ {
		if len(method) == 0 {
			return false
		} else if method == strings.ToUpper(toyotas.Toyota[i].Name) || method == strings.ToLower(toyotas.Toyota[i].Name) {
			if toyotas.Toyota[i].Vip == true {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func CheckJSONPremium(method string) bool {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotaq

	json.Unmarshal(byteValue, &toyotas)
	for i := 0; i < len(toyotas.Toyota); i++ {
		if len(method) == 0 {
			return false
		} else if method == strings.ToUpper(toyotas.Toyota[i].Name) || method == strings.ToLower(toyotas.Toyota[i].Name) {
			if toyotas.Toyota[i].Premium == true {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func CheckJSONHome(method string) bool {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotaq

	json.Unmarshal(byteValue, &toyotas)
	for i := 0; i < len(toyotas.Toyota); i++ {
		if len(method) == 0 {
			return false
		} else if method == strings.ToUpper(toyotas.Toyota[i].Name) || method == strings.ToLower(toyotas.Toyota[i].Name) {
			if toyotas.Toyota[i].Home == true {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func CheckJSONDescription(method string) string {
	jsonFile, err := os.Open("json/apis.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyotaq

	json.Unmarshal(byteValue, &toyotas)
	for i := 0; i < len(toyotas.Toyota); i++ {
		if method == strings.ToUpper(toyotas.Toyota[i].Name) || method == strings.ToLower(toyotas.Toyota[i].Name) {
			return toyotas.Toyota[i].Description
		}
	}
	return ""
}

func BoolToString(argbool bool) string {
	if argbool == true {
		return "true"
	} else if argbool == false {
		return "false"
	}
	return ""
}
