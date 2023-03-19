package main

//import (
//	"encoding/json"
//	"io/ioutil"
//	"log"
//	"os"
//	"strings"
//	"unicode/utf8"
//)
//
////loads the json tech
//func jsonAuthencation() error {
//	File, error := os.Open(JsonFile)
//	if error != nil {
//		return error
//	}
//
//	defer File.Close()
//
//	ByteValue, error := ioutil.ReadAll(File)
//	if error != nil {
//		return error
//	}
//
//	var Params Json
//	error = json.Unmarshal(ByteValue, &Params)
//	if error != nil {
//		return error
//	}
//
//	TypeJson = &Params
//	return nil
//}
//
//func VerboseLog(msg ...interface{}) {
//	if VerBose {
//		log.Println(msg)
//	}
//	return
//}
//
//func NeedleHayStack(array []string, value string) bool {
//	for _, val := range array {
//		if strings.Split(val, ":")[0] == strings.Split(value, ":")[0] {
//			return true
//		}
//	}
//	return false
//}
//
//func FillSpace(Object string, LenNeeded int) string {
//
//	if len(Object) == LenNeeded {
//		return Object
//	}
//
//	var Complete string = Object
//
//	for I := len(Object); I < LenNeeded; I++ {
//		Complete += " "
//	}
//
//	return Complete
//}
//
//func CentreObject(Length int, Text string) string {
//	var Line string
//
//	var Middle = Length / 2
//
//	for i := 0; i < Length; i++ {
//
//		if i == Middle {
//			Line += Text
//			i = utf8.RuneCountInString(Text)
//		} else {
//			Line += " "
//			i++
//		}
//	}
//	return Line
//}
//
//func GetOpenSessions(addr string) int {
//	var OPN int = 0
//	for _, value := range Clients {
//
//		if strings.Split(value.Conn.RemoteAddr().String(), ":")[0] == strings.Split(addr, ":")[0] {
//			OPN++
//			continue
//		}
//		continue
//	}
//	return OPN
//}
//
