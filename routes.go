package main

import (
	"encoding/json"
	"fmt"
	"github/Himesh-Kundal/PropHunt/db"

	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, struct{}{})	
}


func (apiCfg *apiConfig)handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}
	params := parameter{}
	err := json.NewDecoder(r.Body).Decode(&params)
	params.PassWord, err = HashPassword(params.PassWord)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if params.UserName == "" || params.PassWord == "" {
		responseWithError(w, http.StatusBadRequest, "Something is empty")
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		Username: params.UserName,
		HashedPassword: params.PassWord,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(w, http.StatusCreated, user)
}

func (apiCfg *apiConfig)handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, users)
}

func (apiCfg *apiConfig) handleGetJwt(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}
	params := parameter{
		UserName: r.URL.Query().Get("username"),
		PassWord: r.URL.Query().Get("password"),
	}

	if params.UserName == "" || params.PassWord == "" {
		responseWithError(w, http.StatusBadRequest, "Something is empty")
		return
	}
	user, err := apiCfg.DB.GetUserByUsername(r.Context(), params.UserName)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if !CheckPasswordHash(params.PassWord, user.HashedPassword) {
		responseWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}
	token, err := GenerateJWT(apiCfg.jwtSecretKey, user.Username, user.HashedPassword)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
}

func (apiCfg *apiConfig) handleUserData(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	user, err := apiCfg.DB.GetUserByUsername(r.Context(),username)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, user)
}

func (apiCfg *apiConfig) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Kills int `json:"kills"`
		Deaths int `json:"deaths"`
		Wins int `json:"wins"`
		Losses int `json:"losses"`
		Draws int `json:"draws"`
		TimeAlive int `json:"time_alive"`
	}
	params := parameter{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println("ye wala")
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if (params.Kills == 0 && params.Deaths == 0 && params.Wins == 0 && params.Losses == 0 && params.Draws == 0 && params.TimeAlive == 0) || 
		params.Kills < 0 || params.Deaths < 0 || params.Wins < 0 || params.Losses < 0 || params.Draws < 0 || params.TimeAlive < 0 {
		fmt.Println("ya fir ye wala")
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user,err := apiCfg.DB.UpdateUserStats(r.Context(), db.UpdateUserStatsParams{
		Username: r.Header.Get("username"),
		Kills: int32(params.Kills),
		Deaths: int32(params.Deaths),
		Wins: int32(params.Wins),
		Losses: int32(params.Losses),
		Draws: int32(params.Draws),
		TimeAlive: int32(params.TimeAlive),
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, user)
}
