package models

import "fmt"

type Console struct {
	ConsoleId int `json:"id"`
	Name string `json:"name"`
	ManufacturerId int `json:"manufacturerId"`
	MediaType string `json:"mediaType"`
	Year string `json:"year"`
	Sales string `json:"sales"`
	Handheld bool `json:"handheld"`
	ImageURL string `json:"imageURL"`
	Notes string `json:"notes"`
}

func SelectAllConsoles() []Console {
	consoles := make([]Console, 0)
	rows, err := database.Query("select ConsoleId, Name, ManufacturerId, MediaType, Year, Sales, Handheld, ImageURL, Notes from Consoles")
	if err != nil {
		fmt.Println("Database select all error:", err)
		return consoles
	}
	defer rows.Close()
	for rows.Next() {
		var c Console
		err = rows.Scan(&c.ConsoleId, &c.Name, &c.ManufacturerId, &c.MediaType, &c.Year, &c.Sales, &c.Handheld, &c.ImageURL, &c.Notes)
		if err != nil {
			fmt.Println("Database select all error:", err)
			break
		}
		consoles = append(consoles, c)
	}
	return consoles
}

func SelectConsole(consoleId int) Console {
	var c Console
	rows, err := database.Query("select ConsoleId, Name, ManufacturerId, MediaType, Year, Sales, Handheld, ImageURL, Notes from Consoles where ConsoleId = ?", consoleId)
	if err != nil {
		fmt.Println("Database select error:", err)
		return c
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&c.ConsoleId, &c.Name, &c.ManufacturerId, &c.MediaType, &c.Year, &c.Sales, &c.Handheld, &c.ImageURL, &c.Notes)
	if err != nil {
		fmt.Println("Database select error:", err)
	}
	return c
}

func InsertConsole(c Console) {
	statement, err := database.Prepare("insert into Consoles (ConsoleId, Name, ManufacturerId, MediaType, Year, Sales, Handheld, ImageURL, Notes) values (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err == nil { _, err = statement.Exec(c.ConsoleId, c.Name, c.ManufacturerId, c.MediaType, c.Year, c.Sales, c.Handheld, c.ImageURL, c.Notes) }
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateConsole(c Console) {
	statement, err := database.Prepare("update Consoles set Name = ?, ManufacturerId = ?, MediaType = ?, Year = ?, Sales = ?, Handheld = ?, ImageURL = ?, Notes = ? where ConsoleId = ?")
	if err == nil { _, err = statement.Exec(c.Name, c.ManufacturerId, c.MediaType, c.Year, c.Sales, c.Handheld, c.ImageURL, c.Notes, c.ConsoleId) }
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteConsole(consoleId int) {
	statement, err := database.Prepare("delete from Consoles where ConsoleId = ?")
	if err == nil { _, err = statement.Exec(consoleId) }
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
