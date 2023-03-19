package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	sessions     = make(map[int64]*Session)
	sessionMutex sync.Mutex
	mutex        sync.RWMutex

	queueSize = 60
)

type Session struct {
	ID          int64
	Username    string
	Conn        net.Conn
	Chat        bool
	Created     time.Time
	LastCommand time.Time
	Msg         string
	User        string
	Cmdhis      string
	admin       bool
	vip         bool
	premium     bool
	seller      bool
	Bhat        bool
	Api         bool
	Loggedlol   bool
	Listen      bool
	Queue       chan []byte
}

func (s *Session) Broadcast(payload []byte) {
	mutex.RLock()
	defer mutex.RUnlock()

	for _, session := range sessions {
		if len(session.Queue) >= queueSize-1 {
			continue
		}

		session.Queue <- payload
	}
}

func (s *Session) Remove() {
	setsessions := len(sessions)           // Setting the real sessions
	removedupesession := (setsessions - 1) // Removes the duped session
	fmt.Printf("\033[103;30;140m %s's \033[0;0m session closed, there is now : \033[97m%d\033[0m online\r\n\033[0m-> \033[0m", s.Username, removedupesession)
	sessionMutex.Lock()
	delete(sessions, s.ID)
	sessionMutex.Unlock()
}

var selected string

func (S *Session) AutoComplete(line string, pos int, key rune) (newLine string, newPos int, ok bool) {

	if key != '\t' || pos != len(line) {
		return
	}
	lastWord := regexp.MustCompile(`.+\W(\w+)$`)
	if !strings.Contains(line, " ") {
		var name string
		return name, len(name), true
	}

	if strings.HasSuffix(line, " ") {
		return line, pos, true
	}
	m := lastWord.FindStringSubmatch(line)
	if m == nil {
		return line, len(line), true
	}
	soFar := m[1]
	var match []string

	if len(match) == 0 {
		return
	}
	if len(match) > 1 {
		return line, pos, true
	}
	newLine = line[:len(line)-len(soFar)] + match[0]
	return newLine, len(newLine), true
}

func (s *Session) Open(username string) int {
	Ammount := 0
	// loops throughout the session list checking the ammount of sessions open by that person
	for _, s := range sessions {
		if s.Username == username {
			Ammount++
		}
	}
	return Ammount
}

func usersSessions(username string) []*Session {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	var list []*Session
	for i := range sessions {
		if sessions[i].Username == username {
			list = append(list, sessions[i])
		}
	}

	return list
}

type support struct {
	ID       int64
	Username string
	Password string
	Admin    bool
	Conn     net.Conn
	Msg      string
	User     string
	Cmdhis   string

	Chat bool
}

var (
	cmds         = make(map[int64]*history)
	historyMutex sync.Mutex
)

type history struct {
	ID       int64
	Username string
	Password string
	Admin    bool
	Conn     net.Conn
	Msg      string
	User     string
	Cmdhis   string

	Chat bool
}

var (
	cmdsss      = make(map[int64]*cmdatt)
	cmdattMutex sync.Mutex
)

type cmdatt struct {
	ID       int64
	Username string
	Password string
	Admin    bool
	Conn     net.Conn
	Msg      string
	User     string
	Cmdhis   string
	Atks     string

	Chat bool
}
