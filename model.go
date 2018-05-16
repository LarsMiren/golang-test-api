package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Room struct {
	Number int      `json:"number"`
	Guests []string `json:"guests"`
}

var db *sql.DB

func initializeDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite.db")
	return err
}

func getGuests(room int) ([]string, error) {
	stmt, err := db.Prepare("SELECT name FROM guest WHERE room_number=?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(room)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []string
	for rows.Next() {
		var g string
		if err := rows.Scan(&g); err != nil {
			return nil, err
		}
		guests = append(guests, g)
	}
	return guests, nil
}

func getRooms() ([]Room, error) {
	stmt, err := db.Prepare("SELECT number FROM room")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var nums []int
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	var rooms []Room
	for _, n := range nums {
		g, err := getGuests(n)
		if err != nil {
			g = []string{}
		}
		r := Room{Number: n, Guests: g}
		rooms = append(rooms, r)
	}

	return rooms, nil
}

func addRoom(r Room) error {
	stmt, err := db.Prepare("INSERT INTO room(number) VALUES (?)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(r.Number); err != nil {
		return err
	}

	for _, g := range r.Guests {
		if err := addGuest(r.Number, g); err != nil {

		}
	}
	return nil
}

func addGuest(room int, name string) error {
	stmt, err := db.Prepare("INSERT INTO guest(room_number, name) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(room, name)
	return err
}

func deleteGuests(room int) error {
	stmt, err := db.Prepare("DELETE FROM guest WHERE room_number=?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(room); err != nil {
		return err
	}

	return nil
}

func deleteRoom(numb int) error {
	if err := deleteGuests(numb); err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM room WHERE number=?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(numb); err != nil {
		return err
	}

	return nil
}
