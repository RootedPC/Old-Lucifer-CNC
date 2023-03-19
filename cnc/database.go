package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

type AccountInfo struct {
	ID          int
	username    string
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
	flagged     int
}

type User struct {
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
	flagged     int
}

type Attackv2 struct {
	id       int
	method   string
	target   string
	port     int
	duration int

	username string

	end     int64
	created int64
}

func NewDatabase(dbAddr string, dbUser string, dbPassword string, dbName string) *Database {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbAddr, dbName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if error := db.Ping(); error != nil {
		fmt.Println("ERROR: " + error.Error())
		os.Exit(1)
	}

	fmt.Printf("\033[0m-> \033[0m\033[0mmysql connection: [\033[102;30;140m OK \033[0;0m]\r\n\033[0m-> \033[0m")
	//	sendScreenAlert()
	return &Database{db}
}

func (this *Database) Auth(user, pass string) bool {
	Row, error := this.db.Query("SELECT `password` FROM `users` WHERE `username` = ? AND `password` = ?", user, pass)
	if error != nil {
		return false
	}

	if !Row.Next() {
		return false
	}

	return true
}

func (this *Database) TryLogin(username string, password string, ip net.Addr) (bool, AccountInfo) {
	rows, err := this.db.Query("SELECT id, username, admin, expiry, ban, vip, mfasecret, concurrents, cooldown, hometime, bypasstime, premium, home, seller, flagged FROM users WHERE username = ? AND password = ? ", username, password)
	strRemoteAddr := ip.String()
	host, _, err := net.SplitHostPort(strRemoteAddr)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("\033[101;30;140m %s \033[0;0m has failed on: \033[97m%s\033[0m\r\n\033[0m-> \033[0m", username, host)
		this.db.Exec("INSERT INTO logins (username, action, ip) VALUES (?, ?, ?)", username, "Fail", host)
		return false, AccountInfo{}
	}

	defer rows.Close()
	if !rows.Next() {
		fmt.Printf("\033[101;30;140m %s \033[0;0m has failed on: \033[97m%s\033[0m\r\n\033[0m-> \033[0m", username, host)
		this.db.Exec("INSERT INTO logins (username, action, host) VALUES (?, ?, ?)", username, "Fail", host)
		return false, AccountInfo{}
	}

	var accInfo AccountInfo
	rows.Scan(
		&accInfo.ID,
		&accInfo.username,
		&accInfo.admin,
		&accInfo.expiry,
		&accInfo.ban,
		&accInfo.vip,
		&accInfo.mfasecret,
		&accInfo.concurrents,
		&accInfo.cooldown,
		&accInfo.hometime,
		&accInfo.bypasstime,
		&accInfo.premium,
		&accInfo.home,
		&accInfo.seller,
		&accInfo.flagged,
	)

	fmt.Printf("\033[102;30;140m %s \033[0;0m has logged in on: \033[97m%s\033[0m\r\n\033[0m-> \033[0m", username, host)
	this.db.Exec("INSERT INTO logins (username, action, ip) VALUES (?, ?, ?)", accInfo.username, "Login", host)
	return true, accInfo
}

func (this *Database) CreateUser(username string, expiry int) bool {
	rows, err := this.db.Query("SELECT username FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if rows.Next() {
		return false
	}
	this.db.Exec("INSERT INTO `users` (`id`, `username`, `password`, `admin`, `expiry`, `ban`, `vip`, `mfaSecret`, `concurrents`, `cooldown`, `hometime`, `bypasstime`, `premium`, `home`, `seller`, `flagged`) VALUES (NULL, ?, '"+username+"432@-0-0', 0, ?, 0, 0, 0, 1, 80, 300, 120, 0, 1, 0, 0);", username, time.Now().Add((time.Hour*24)*time.Duration(expiry)).Unix())
	return true
}

func (this *Database) Exists(user string) error {
	Row, error := this.db.Query("SELECT `username` FROM `users` WHERE `username` = ?", user)
	if error != nil {
		return error
	}

	if !Row.Next() {
		return error
	}

	return error
}

func (this *Database) addcons(cons int, username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `concurrents` = ? WHERE `username` = ?", cons, username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) setalluserscons(cons int) bool {
	_, err := this.db.Query("UPDATE `users` SET `concurrents` = ?", cons)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) setalluserscooldown(cooldown int) bool {
	_, err := this.db.Query("UPDATE `users` SET `cooldown` = ?", cooldown)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) updatecooldown(username string, cooldown int) bool {
	_, err := this.db.Query("UPDATE `users` SET `cooldown` = ? WHERE `username` = ?", cooldown, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (this *Database) setallusershometime(time int) bool {
	_, err := this.db.Query("UPDATE `users` SET `hometime` = ?", time)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

//func (this *Database) setallusersdays(days int) bool {
//	rows, err := this.db.Query("UPDATE `users` SET expiry = ?", days*86400)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	if rows.Next() {
//		return false
//	}
//	this.db.Exec("UPDATE `users` SET expiry = ?", days*86400)
//	return true
//}

func (this *Database) EditHometime(username string, time int) bool {
	_, err := this.db.Query("UPDATE `users` SET `hometime` = ? WHERE `username` = ?", time, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (this *Database) EditBypasstime(username string, time int) bool {
	_, err := this.db.Query("UPDATE `users` SET `bypasstime` = ? WHERE `username` = ?", time, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (this *Database) GetUser(username string) (User, error) {

	var user User

	err := this.db.QueryRow("SELECT `id`, `username`, `password`, `admin`, `expiry`, `ban`, `vip`, `mfasecret`, `concurrents`, `cooldown`, `hometime`, `bypasstime`, `premium`, `home`, `seller`, `flagged` FROM `users` WHERE `username` = ? LIMIT 1", username).Scan(
		&user.ID,
		&user.username,
		&user.password,
		&user.admin,
		&user.expiry,
		&user.ban,
		&user.vip,
		&user.mfasecret,
		&user.concurrents,
		&user.cooldown,
		&user.hometime,
		&user.bypasstime,
		&user.premium,
		&user.home,
		&user.seller,
		&user.flagged,
	)
	return user, err
}

func (this *Database) GetUsers() ([]User, error) {

	var list []User

	rows, err := this.db.Query("SELECT `id`, `username`, `admin`, `expiry`, `ban`, `vip`, `mfasecret`, `concurrents`, `cooldown`, `hometime`, `bypasstime`, `premium`, `home`, `seller`, `flagged` from `users`")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.username,
			&user.admin,
			&user.expiry,
			&user.ban,
			&user.vip,
			&user.mfasecret,
			&user.concurrents,
			&user.cooldown,
			&user.hometime,
			&user.bypasstime,
			&user.premium,
			&user.home,
			&user.seller,
			&user.flagged,
		)

		if err != nil {
			log.Println(err)
			continue
		}
		list = append(list, user)
	}

	err = rows.Err()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (this *Database) RemoveUser(username string) bool {
	rows, err := this.db.Query("DELETE FROM `users` WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("DELETE FROM `users` WHERE username = ?", username)
	return true
}

func (this *Database) UserToggleMfa(username string, secret string) bool {
	_, err := this.db.Query("UPDATE `users` SET `mfaSecret` = ? WHERE `username` = ?", secret, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
func (this *Database) ChangeUsersPassword(username string, password string) bool {
	rows, err := this.db.Query("UPDATE `users` SET `password` = ? WHERE `username` = ?", password, username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("UPDATE `users` SET `password` = ? WHERE `username` = ?", password, username)
	return true
}
func (this *Database) EditDays(username string, days int) bool {
	rows, err := this.db.Query("UPDATE `users` SET expiry = ? WHERE username = ?", time.Now().Add((time.Hour*24)*time.Duration(days)).Unix(), username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("UPDATE `users` SET expiry = ? WHERE username = ?", time.Now().Add((time.Hour*24)*time.Duration(days)).Unix(), username)
	return true
}
func (this *Database) MakeAdmin(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `admin` = 1 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) MakeSeller(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `seller` = 1 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) RemoveSeller(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `seller` = 0 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) RemoveAdmin(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `admin` = 0 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) MakeVip(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `vip` = 1 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (this *Database) RemoveVip(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `vip` = 0 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (this *Database) MakePremium(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `premium` = 1 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) RemovePremium(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `premium` = 0 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) MakeHome(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `home` = 1 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func (this *Database) RemoveHome(username string) bool {
	_, err := this.db.Query("UPDATE `users` SET `home` = 0 WHERE `username` = ?", username)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

type running struct {
	running int
}

func (this *Database) GetRunningUser(User string) (int, error) {
	var Active running
	error := this.db.QueryRow("SELECT COUNT(*) FROM `attacksv2` WHERE `username` = ? AND `end` > ?", User, time.Now().Unix()).Scan(&Active.running)
	if error != nil {
		return 0, error
	}

	return Active.running, nil
}

func (this *Database) ListOngoing() int {
	var Active running
	error := this.db.QueryRow("SELECT COUNT(*) FROM `attacksv2` WHERE `end` > ?", time.Now().Unix()).Scan(&Active.running)
	if error != nil {
		return 0
	}

	return Active.running
}

func (this *Database) LogAttack(attacksv2 *Attackv2) (bool, error) {
	_, err := this.db.Query("INSERT INTO `attacksv2` (`id`, `username`, `target`, `method`, `port`, `duration`, `end`, `created`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?)",
		attacksv2.username,
		attacksv2.target,
		attacksv2.method,
		attacksv2.port,
		attacksv2.duration,
		attacksv2.end,
		attacksv2.created,
	)
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, nil
}

func (this *Database) AlreadyUnderAttack(User string, Target string) (*Attackv2, error) {
	var RunningDetails Attackv2
	error := this.db.QueryRow("SELECT `id`, `username`, `target`, `method`, `port`, `duration`, `end`, `created` FROM `attacksv2` WHERE `target` = ? AND `end` > ?", Target, time.Now().Unix()).Scan(
		&RunningDetails.id,
		&RunningDetails.username,
		&RunningDetails.target,
		&RunningDetails.method,
		&RunningDetails.port,
		&RunningDetails.duration,
		&RunningDetails.end,
		&RunningDetails.created,
	)
	if error != nil {
		return nil, nil
	}

	return &RunningDetails, nil
}

func (this *Database) MySent(user string) int {
	var Active running
	error := this.db.QueryRow("SELECT COUNT(*) FROM `attacksv2` WHERE `username` = ?", user).Scan(&Active.running)
	if error != nil {
		return 0
	}

	return Active.running
}

func (this *Database) Ongoing(User string) ([]*Attackv2, error) {

	var AttackRunning []*Attackv2

	var rows *sql.Rows
	var err error

	rows, err = this.db.Query("SELECT `id`, `username`, `target`, `method`, `port`, `duration`, `end`, `created` FROM `attacksv2` WHERE `end` > ?", time.Now().Unix())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return AttackRunning, err
	}

	defer rows.Close()

	for rows.Next() {
		Attacks := &Attackv2{}
		if err := scanOngoing(rows, Attacks); err != nil {
			continue
		}

		AttackRunning = append(AttackRunning, Attacks)
	}

	return AttackRunning, nil
}

func (this *Database) MyAttacking(User string) ([]*Attackv2, error) {

	var AttackRunning []*Attackv2

	var rows *sql.Rows
	var err error

	rows, err = this.db.Query("SELECT `id`, `username`, `target`, `method`, `port`, `duration`, `end`, `created` FROM `attacksv2` WHERE `end` > ? AND `username` = ?", time.Now().Unix(), User)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return AttackRunning, err
	}

	defer rows.Close()

	for rows.Next() {
		Attacks := &Attackv2{}
		if err := scanOngoing(rows, Attacks); err != nil {
			continue
		}

		AttackRunning = append(AttackRunning, Attacks)
	}

	return AttackRunning, nil
}

func scanOngoing(row *sql.Rows, Attacking *Attackv2) error {

	return row.Scan(
		&Attacking.id,
		&Attacking.username,
		&Attacking.target,
		&Attacking.method,
		&Attacking.port,
		&Attacking.duration,
		&Attacking.end,
		&Attacking.created,
	)
}

func (this *Database) CheckSessionAdmin(userz string) bool {
	var totalAttacks bool

	this.db.QueryRow("SELECT admin FROM users WHERE username = ?", userz).Scan(
		&totalAttacks,
	)

	return totalAttacks
}
func (d *Database) CheckSessionMFA(userz string) string {
	var totalAttacks string

	d.db.QueryRow("SELECT mfaSecret FROM users WHERE username = ?", userz).Scan(
		&totalAttacks,
	)

	return totalAttacks
}
func (this *Database) UserTempBan(username string, expire int64) bool {
	_, err := this.db.Query("UPDATE `users` SET `ban` = ? WHERE `username` = ?", expire, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (this *Database) flagingSystem(username string, flagged int) bool {
	_, err := this.db.Query("UPDATE `users` SET `flagged` = ? WHERE `username` = ?", flagged, username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
