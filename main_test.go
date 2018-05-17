package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDBconnection(t *testing.T) {
	if err := initializeDB(); err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func clearTables() error {
	if _, err := db.Exec("DELETE FROM guest"); err != nil {
		return err
	}
	if _, err := db.Exec("DELETE FROM room"); err != nil {
		return err
	}
	return nil
}

func TestGetAllRoomsWithEmptyDB(t *testing.T) {
	if err := initializeDB(); err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err := clearTables(); err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(getRoomsHandler)

	r, err := http.NewRequest("GET", "/rooms", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Error("Wrong status. Expected OK")
	}

	if w.Body.String() != "[]" {
		t.Error("Wrong response. Expected empty array, got " + w.Body.String())
	}
}

func TestCreateRoom(t *testing.T) {
	if err := initializeDB(); err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err := clearTables(); err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(addRoomHandler)

	room := Room{
		Number: 3,
		Guests: []string{"John Doe", "Jane Doe"}}

	jr, _ := json.Marshal(room)

	r := httptest.NewRequest("POST", "/rooms", bytes.NewReader(jr))

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Error("Wrong status. Expected OK, got " + string(w.Code))
	}
}
