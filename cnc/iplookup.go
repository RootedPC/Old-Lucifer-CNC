package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"unicode/utf8"
)

type API_Resp struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

var (
	ErrHTTP = errors.New("\033[91mError connecting to Yukari iplookup.sys\033[0m")

	ErrJsonUnmarshal = errors.New("\033[91mUnkown error has occured\033[0m")
)

func Reach(Target string, Token string) (*API_Resp, error) {

	Resp, error := http.Get("http://ipinfo.io/" + url.QueryEscape(Target) + "?token=" + url.QueryEscape(Token))
	if error != nil {
		return nil, error
	}

	if Resp.StatusCode != 200 {
		return nil, ErrHTTP
	}

	Read, error := ioutil.ReadAll(Resp.Body)
	if error != nil {
		return nil, ErrJsonUnmarshal
	}

	var NewRes API_Resp
	error = json.Unmarshal(Read, &NewRes)
	if error != nil {
		return nil, ErrJsonUnmarshal
	}

	return &NewRes, nil
}

func Lookup(ip string, this *Admin) {
	l, error := Reach(ip, "b338a9aeaca3dc")
	if error != nil || l == nil {
		this.conn.Write([]byte("\033[91mError\033[97m:\033[0m failed to lookup `" + ip + "`\r\n"))
		return
	}
	if l.IP == " " {
		this.conn.Write([]byte("\033[91mError\033[97m:\033[0m failed to lookup `" + ip + "`\r\n"))
		return
	} else if l.IP == "" {
		this.conn.Write([]byte("\033[91mError\033[97m:\033[0m failed to lookup `" + ip + "`\r\n"))
		return
	} else {
		this.conn.Write([]byte("\033[96m═════════════════\033[106;30;140m yukari.v1 \033[0;96m════════════════╗\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mIPv4\033[96m:\033[0m " + FillSpace(l.IP, 38) + "\033[96m║\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mCity\033[96m:\033[0m " + FillSpace(l.City, 38) + "\033[96mL\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mRegion\033[96m:\033[0m " + FillSpace(l.Region, 36) + "\033[96mO\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mCountry\033[96m:\033[0m " + FillSpace(l.Country, 35) + "\033[96mO\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mLocation\033[96m:\033[0m " + FillSpace(l.Loc, 34) + "\033[96mK\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mPostal\033[96m:\033[0m " + FillSpace(l.Postal, 36) + "\033[96mU\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mOrganization\033[96m:\033[0m " + FillSpace(l.Org, 30) + "\033[96mP\033[0m\r\n"))
		this.conn.Write([]byte("\033[37mTimezone\033[96m:\033[0m " + FillSpace(l.Timezone, 34) + "\033[96m║\033[0m\r\n"))
		this.conn.Write([]byte("\033[96m════════════════════════════════════════════╝\033[0m\r\n"))
		return
	}
}

func FillSpace(Object string, LenNeeded int) string {

	if utf8.RuneCountInString(Object) == LenNeeded {
		return Object
	}

	var Complete string = Object

	for I := utf8.RuneCountInString(Object); I < LenNeeded; I++ {
		Complete += " "
	}

	return Complete
}
