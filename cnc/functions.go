package main

import (
	"crypto/rand"
	"encoding/base32"
	"log"
	"math/big"
	"regexp"
	"strings"
)

func fillSpace(input string, length int, filler string) string {

	var output = input
	if len(input) > length {
		output = output[:length]
	}

	for {
		if len(output) >= length {
			break
		}

		output = output + filler
	}

	return output
}

//IsIPv4 returns true if the input matches a IPv4 address
func IsIPv4(ip string) bool {
	var re = regexp.MustCompile(`(?m)^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)

	if re.MatchString(ip) == false {
		return false
	}

	return true
}

//IsDomain returns true if the input matches a domain
func IsDomain(ip string) bool {
	var re = regexp.MustCompile(`[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\.[a-zA-Z]{2,}`)

	if re.MatchString(ip) == false {
		return false
	}

	return true
}

func GenTOTPSecret() string {

	data, err := cryptoRand(32)
	if err != nil {
		log.Println(err)
		return strings.ReplaceAll(base32.StdEncoding.EncodeToString([]byte("0A9SF870A9SDUF09SDF234")), "=", "D")
	}

	return strings.ReplaceAll(base32.StdEncoding.EncodeToString([]byte(data)), "=", "D")
}

func cryptoRand(length int) (string, error) {
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}

		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}

		n := num.Int64()

		if n > 32 && n < 127 {
			result += string(n)
		}
	}
}
func formatBool(input bool) string {

	if input == false {
		return "\x1b[91mfalse\x1b[0m"
	}

	return "\x1b[92mtrue\x1b[0m"
}

func formatAdminBool(input bool) string {

	if input == false {
		return "\033[0mN/A\033[0m"
	} else {

		return "\033[102;30;140m A \033[0m"
	}
}

func formatSellerBool(input bool) string {

	if input == false {
		return "\033[0mN/A\033[0m"
	} else {

		return "\033[103;30;140m M \033[0m"
	}
}

func formatPremiumBool(input bool) string {

	if input == false {
		return "\033[0mN/A\033[0m"
	} else {

		return "\033[106;30;140m P \033[0m"
	}
}

func format2faBool(input bool) string {

	if input == false {
		return "\033[0mN/A\033[0m"
	} else {

		return "\033[103;30;140m T \033[0m"
	}
}

func censorString(input string, censor string) string {

	cut := float32(len(input)) * 0.65

	section := input[:int(cut)]

	for {
		if len(section) >= len(input) {
			break
		}

		section += censor
	}

	return section
}
