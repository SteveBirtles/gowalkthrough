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

	var game models.Game

	id := utils.PathTail(r.URL.Path)

	if id != -1 {
		game = models.SelectGame(id)
	}

	consoleId, _ := strconv.Atoi(r.FormValue("consoleId"))
	name := r.FormValue("name")
	sales := r.FormValue("sales")
	year := r.FormValue("year")
	imageURL := r.FormValue("imageURL")

	fmt.Println("/game/save/", id, consoleId, "(", r.FormValue("consoleId"), ")", name, sales, year, imageURL)

	game.ConsoleId = consoleId
	game.Name = name
	game.Sales = sales
	game.Year = year
	game.ImageURL = imageURL

	if game.GameId == 0 {

		lastId := 0
		allGames := models.SelectAllGames()
		for _, g := range allGames {
			if g.GameId > lastId {
				lastId = g.GameId
			}
		}
		game.GameId = lastId + 1

		models.InsertGame(game)
		_, _ = fmt.Fprint(w, "OK")

	} else {

		models.UpdateGame(game)
		_, _ = fmt.Fprint(w, "OK")

	}

}

func DeleteGame(w http.ResponseWriter, r *http.Request) {

}
