package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"sync"

	"github.com/vaughan0/go-ini"
)

var (
	BetaMapHandler = make(map[string]*CommandText)
	Handle         sync.Mutex
)

type CommandText struct {
	CommandName        string
	CommandAdmin       bool
	CommandReseller    bool
	CommandVip         bool
	CommandContains    string
	CommandDescription string
}

var CommandFile string = "branding/"

func PolyLoader() bool {

	for i, _ := range BetaMapHandler {
		delete(BetaMapHandler, i)
	}

	Files, err := ioutil.ReadDir(CommandFile)
	if err != nil {
		return false
	}
	loaded := 0

	for _, f := range Files {
		_, _ = ini.LoadFile(CommandFile + f.Name())
		loaded++
	}

	log.Println("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded " + strconv.Itoa(len(BetaMapHandler)) + " Commands Correctly")

	log.Println(" [RELOADED] [Reloaded Commands Correctly] [" + strconv.Itoa(len(BetaMapHandler)) + "]")

	return true
}
