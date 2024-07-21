package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/harryfung22/ScrapeHive/internal/databse"
)

func handleRes(w http.ResponseWriter, r *http.Request) {
	resJson(w, 200, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	resErr(w, 400, "Something went wrong")
}

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		resErr(w, 400, fmt.Sprintf("Error parsion JSON: ", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), databse.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		resErr(w, 400, fmt.Sprintf("Couldn't create user: ", err))
		return
	}

	resJson(w, 200, dbUserToUser(user))
}
