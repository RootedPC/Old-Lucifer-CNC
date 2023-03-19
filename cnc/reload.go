package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	BrandingWorker = make(map[int64]*Branding)
	Commands       sync.Mutex
)

type Branding struct {
	ID              int64
	CommandName     string
	CommandFile     string
	CommandContains string
}

var BrandingFile = "branding/"
var BrandingFile2 = "json/"
var BrandingFile3 = "titleshit/"

func CompleteLoad() error {
	for _, f := range BrandingWorker {
		delete(BrandingWorker, f.ID)
	}

	files, err := ioutil.ReadDir(BrandingFile)
	if err != nil {
		log.Println(err)
		return err
	}

	loaded := 0

	for _, f := range files {
		CommandVisual, err := ioutil.ReadFile(BrandingFile + f.Name())
		if err != nil {
			log.Printf(" [Failed To Grab Command %s]", BrandingFile+f.Name())
		}

		var Working = &Branding{
			ID:              time.Now().UnixNano(),
			CommandName:     strings.Replace(f.Name(), ".tfx", "", -1),
			CommandFile:     f.Name(),
			CommandContains: string(CommandVisual),
		}

		Commands.Lock()
		BrandingWorker[Working.ID] = Working
		Commands.Unlock()
		loaded++
		continue
	}
	return nil
}

func CompleteLoad2() error {
	for _, f := range BrandingWorker {
		delete(BrandingWorker, f.ID)
	}

	files2, err := ioutil.ReadDir(BrandingFile2)
	if err != nil {
		log.Println(err)
		return err
	}

	loaded2 := 0

	for _, f2 := range files2 {
		CommandVisual, err := ioutil.ReadFile(BrandingFile2 + f2.Name())
		if err != nil {
			log.Printf(" [Failed To Grab Json %s]", BrandingFile2+f2.Name())
		}

		var Working = &Branding{
			ID:              time.Now().UnixNano(),
			CommandName:     strings.Replace(f2.Name(), ".json", "", -1),
			CommandFile:     f2.Name(),
			CommandContains: string(CommandVisual),
		}

		Commands.Lock()
		BrandingWorker[Working.ID] = Working
		Commands.Unlock()
		loaded2++
		continue
	}
	return nil
}

func CompleteLoad3() error {
	for _, f := range BrandingWorker {
		delete(BrandingWorker, f.ID)
	}

	files3, err := ioutil.ReadDir(BrandingFile3)
	if err != nil {
		log.Println(err)
		return err
	}

	loaded3 := 0

	for _, f3 := range files3 {
		CommandVisual, err := ioutil.ReadFile(BrandingFile3 + f3.Name())
		if err != nil {
			log.Printf(" [Failed To Grab Title %s]", BrandingFile3+f3.Name())
		}

		var Working = &Branding{
			ID:              time.Now().UnixNano(),
			CommandName:     strings.Replace(f3.Name(), ".tfx", "", -1),
			CommandFile:     f3.Name(),
			CommandContains: string(CommandVisual),
		}

		Commands.Lock()
		BrandingWorker[Working.ID] = Working
		Commands.Unlock()
		loaded3++
		continue
	}
	return nil
}
