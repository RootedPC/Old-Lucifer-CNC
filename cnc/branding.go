package main

import (
	"io/ioutil"
	"strings"
	"sync"
)

var (
	branding = make(map[string][]string)
	mux      sync.Mutex
)

func LoadBranding(dir string) (error, int) {

	for key, _ := range branding {
		delete(branding, key)
	}

	File, err := ioutil.ReadDir(dir)
	if err != nil {
		return err, 0
	}

	for _, i := range File {

		if strings.Split(i.Name(), ".")[len(strings.Split(i.Name(), "."))-1] != "tfx" {
			continue
		}

		content, err := ioutil.ReadFile(dir + "/" + i.Name())
		if err != nil {
			continue
		}

		mux.Lock()
		branding[i.Name()] = strings.SplitAfter(string(content), "\n")
		mux.Unlock()
	}

	return nil, len(branding)
}

func GetItem(file string) []string {
	return branding[file]
}
