package controllers

import (
	"../../server/models"
	"../../server/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ListGames(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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
			if g.ImageURL == "" {
				g.ImageURL = "/client/img/none.png"
			}
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

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := utils.PathTail(r.URL.Path)

	if id == -1 {
		return
	}

	fmt.Println("/game/get/", id)

	game := models.SelectGame(id)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(game)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}

func SaveGame(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var game models.Game

	game.GameId = utils.PathTail(r.URL.Path)

	if game.GameId != -1 {
		game = models.SelectGame(game.GameId)
	}

	game.ConsoleId, _ = strconv.Atoi(r.FormValue("consoleId"))
	game.Name = r.FormValue("name")
	game.Sales = r.FormValue("sales")
	game.Year = r.FormValue("year")
	game.ImageURL = r.FormValue("imageURL")

	fmt.Println("/game/save/", game)

	if game.GameId == -1 {

		fmt.Println("Saving new game...")

		lastId := 0
		allGames := models.SelectAllGames()
		for _, g := range allGames {
			if g.GameId > lastId {
				lastId = g.GameId
			}
		}
		game.GameId = lastId + 1

		models.InsertGame(game)
		fmt.Fprint(w, "OK")

	} else {

		fmt.Println("Saving existing game...")

		models.UpdateGame(game)
		fmt.Fprint(w, "OK")

	}

}

func DeleteGame(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := utils.PathTail(r.URL.Path)

	if id == -1 {
		return
	}

	fmt.Println("/game/delete/", id)

	allGames := models.SelectAllGames()
	for _, g := range allGames {
		if g.GameId == id {
			models.DeleteGame(id)
			fmt.Fprint(w, "OK")
			return
		}
	}

	fmt.Fprint(w, "Error: Can't find game with id", id)

}
