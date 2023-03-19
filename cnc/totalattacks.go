package main

import (
	"time"
)

func (d *Database) getTotalAttacks() int {
	var totalAttacks int

	d.db.QueryRow("SELECT COUNT(*) FROM `attacksv2`").Scan(
		&totalAttacks,
	)

	return totalAttacks
}

func (d *Database) getTotalAttacksRunning() int {
	var totalAttacks int

	d.db.QueryRow("SELECT COUNT(*) FROM `history` WHERE `time_sent` + `duration` > ?", time.Now().Unix()).Scan(
		&totalAttacks,
	)

	return totalAttacks
}

func (d *Database) GrabCons(id int) int {
	var intcons int

	d.db.QueryRow("SELECT COUNT(*) FROM `history` WHERE `time_sent` + `duration` > ? && user_id = ?", time.Now().Unix(), id).Scan(
		&intcons,
	)

	return intcons
}

func (this *Database) CheckIDFromUser(username string) string {
	var userID string

	this.db.QueryRow("SELECT id FROM `users` WHERE username = ?", username).Scan(
		&userID,
	)

	return userID
}
