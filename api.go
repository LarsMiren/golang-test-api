package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initializeRoutes() {
	router = mux.NewRouter()

	router.HandleFunc("/room/{number:[0-9]+}", getRoomHandler).Methods("GET")
	router.HandleFunc("/room", addRoomHandler).Methods("POST")
	router.HandleFunc("/room/{number:[0-9]+}", changeRoomHandler).Methods("PUT")
	router.HandleFunc("/room/{number:[0-9]+}", deleteRoomHandler).Methods("DELETE")
	router.HandleFunc("/rooms", getRoomsHandler).Methods("GET")
}

func getRoomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["number"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	g, err := getGuests(n)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	room := Room{Number: n, Guests: g}
	respondWithJSON(w, http.StatusOK, room)
}

func getRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rooms, err := getRooms()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, rooms)
}

func addRoomHandler(w http.ResponseWriter, r *http.Request) {
	var room Room
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&room); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	if err := addRoom(room); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func deleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["number"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := deleteRoom(n); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func changeRoomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["number"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := deleteRoom(n); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var room Room
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&room); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	if err := addRoom(room); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
