package main

import (
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

// User for saving users
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthdate time.Time `json:"birthdate"`
}

func createUser(w http.ResponseWriter, req *http.Request) {

}
