package models

import "fmt"

type Console struct {
	ConsoleId int `json:"consoleId"`
	Name string `json:"name"`
	ManufacturerId int `json:"manufacturerId"`
	MediaType string `json:"mediaType"`
	Year string `json:"year"`
	Sales string `json:"sales"`
	Handheld string `json:"handheld"`
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
	rows, err := database.Query(fmt.Sprintf("select ConsoleId, Name, ManufacturerId, MediaType, Year, Sales, Handheld, ImageURL, Notes from Consoles where ConsoleId = %d", consoleId))
	if err != nil {
		fmt.Println("Database select all error:", err)
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
	_, err := database.Exec(fmt.Sprintf("insert into Consoles (ConsoleId, Name, ManufacturerId, MediaType, Year, Sales, Handheld, ImageURL, Notes) values (%d, '%s', '%d', '%s', '%s', '%s', '%s', '%s', '%s')",
		c.ConsoleId, c.Name, c.ManufacturerId, c.MediaType, c.Year, c.Sales, c.Handheld, c.ImageURL, c.Notes))
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateConsole(c Console) {
	_, err := database.Exec(fmt.Sprintf("update Consoles set Name = '%s', ManufacturerId = '%d', MediaType = '%s', Year = '%s', Sales = '%s', Handheld = '%s', ImageURL = '%s', Notes = '%s' where ConsoleId = %d",
		c.Name, c.ManufacturerId, c.MediaType, c.Year, c.Sales, c.Handheld, c.ImageURL, c.Notes, c.ConsoleId))
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteConsole(consoleId int) {
	_, err := database.Exec(fmt.Sprintf("delete from Messages where ConsoleId = %d", consoleId))
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
