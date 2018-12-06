package models

import "fmt"

type Game struct {
	GameId    int    `json:"id"`
	ConsoleId int    `json:"consoleId"`
	Name      string `json:"name"`
	Sales     string `json:"sales"`
	Year      string `json:"year"`
	ImageURL  string `json:"imageURL"`
}

func SelectAllGames() []Game {
	games := make([]Game, 0)
	rows, err := database.Query("select GameId, ConsoleId, Name, Sales, Year, ImageURL from Games")
	if err != nil {
		fmt.Println("Database select all error:", err)
		return games
	}
	defer rows.Close()
	for rows.Next() {
		var g Game
		err = rows.Scan(&g.GameId, &g.ConsoleId, &g.Name, &g.Sales, &g.Year, &g.ImageURL)
		if err != nil {
			fmt.Println("Database select all error:", err)
			break
		}
		games = append(games, g)
	}
	return games
}

func SelectGame(gameId int) Game {
	var g Game
	rows, err := database.Query("select GameId, ConsoleId, Name, Sales, Year, ImageURL from Games where GameId = ?", gameId)
	if err != nil {
		fmt.Println("Database select error:", err)
		return g
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&g.GameId, &g.ConsoleId, &g.Name, &g.Sales, &g.Year, &g.ImageURL)
	if err != nil {
		fmt.Println("Database select error:", err)
	}
	return g
}

func InsertGame(g Game) {
	statement, _ := database.Prepare("insert into Games (GameId, ConsoleId, Name, Sales, Year, ImageURL) values (?, ?, ?, ?, ?, ?)")
	_, err := statement.Exec(g.GameId, g.ConsoleId, g.Name, g.Sales, g.Year, g.ImageURL)
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateGame(g Game) {
	statement, _ := database.Prepare("update Games set ConsoleId = ?, Name = ?, Sales = ?, Year = ?, ImageURL = ? where GameId = ?")
	_, err := statement.Exec(g.ConsoleId, g.Name, g.Sales, g.Year, g.ImageURL, g.GameId)
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteGame(gameId int) {
	statement, _ := database.Prepare("delete from Messages where GameId = ?")
	_, err := statement.Exec(gameId)
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
