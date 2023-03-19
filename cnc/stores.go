package main

import (
	"net"
	"sync"
)

var (
	TypeJson *Json
	JsonFile string = "json/configure.json"

	Clients map[string]*Client = make(map[string]*Client)
	mutex2  sync.Mutex
)

type Client struct {
	Conn net.Conn
}

type Json struct {
	ListenerStat string `json:"ListenerStat"`
	ProxyServer  struct {
		DialTarget   string `json:"DialTarget"`
		DialProtocol string `json:"DialProtocol"`
		DialTimeout  int    `json:"DialTimeout"`
	} `json:"ProxyServer"`
	AppName string   `json:"AppName"`
	Owners  []string `json:"Owners"`
	Sec     struct {
		ConnectionsPerIP int `json:"ConnectionsPerIP"`
		Whitelists       struct {
			Status     bool     `json:"status"`
			Whitelists []string `json:"hosts"`
		} `json:"whitelists"`
		Blacklists struct {
			Status     bool     `json:"status"`
			Whitelists []string `json:"hosts"`
		} `json:"blacklists"`
	} `json:"sec"`
}
