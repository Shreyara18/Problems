package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	ID   uint32 `json:"id"`
	Note string `json:"note"`
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement user creation logic

	w.WriteHeader(http.StatusOK)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement login logic

	// Generate session ID
	sessionID := "session_id"

	response := struct {
		SID string `json:"sid"`
	}{
		SID: sessionID,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func listNotesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		SID string `json:"sid"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement logic to retrieve notes for the user based on the session ID

	notes := []Note{
		{ID: 1, Note: "Note 1"},
		{ID: 2, Note: "Note 2"},
		{ID: 3, Note: "Note 3"},
	}
	response := struct {
		Notes []Note `json:"notes"`
	}{
		Notes: notes,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		SID  string `json:"sid"`
		Note string `json:"note"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement logic to create a new note for the user based on the session ID

	// Generate a new note ID
	noteID := uint32(1)

	response := struct {
		ID uint32 `json:"id"`
	}{
		ID: noteID,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		SID string `json:"sid"`
		ID  uint32 `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
//////
	// TODO: Implement logic to delete the note for the user based on the session ID and note ID

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/notes", listNotesHandler)
	http.HandleFunc("/notes", createNoteHandler)
	http.HandleFunc("/notes", deleteNoteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
