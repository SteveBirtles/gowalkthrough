package controllers

import (
	"../../server/models"
	"../../server/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListGames(w http.ResponseWriter, r *http.Request) {

	type GamesWithConsole struct {
		ConsoleName string        `json:"consoleName"`
		Games       []models.Game `json:"gamesList"`
	}

	fmt.Println("/game/list")

	consoleId := utils.PathTail(r.URL.Path)

	allGames := models.SelectAllGames()
	relevantGames := make([]models.Game, 0)
	for _, g := range allGames {
		if g.ConsoleId == consoleId {
			relevantGames = append(relevantGames, g)
		}
	}

	gamesWithConsole := GamesWithConsole{models.SelectConsole(consoleId).Name, relevantGames}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(gamesWithConsole)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}

func GetGame(w http.ResponseWriter, r *http.Request) {

}

func SaveGame(w http.ResponseWriter, r *http.Request) {

}

func DeleteGame(w http.ResponseWriter, r *http.Request) {

}
