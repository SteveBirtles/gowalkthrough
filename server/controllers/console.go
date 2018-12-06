package controllers

import (
	"../../server/models"
	"../../server/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ConsoleWithManufacturer struct {
	models.Console
	Manufacturer string `json:"manufacturer"`
}

func ListConsoles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("/console/list")
	consoles := models.SelectAllConsoles()

	consolesWithManufacturers := make([]ConsoleWithManufacturer, 0)
	for _, c := range consoles {
		if c.ImageURL == "" {
			c.ImageURL = "/client/img/none.png"
		}
		manufacturer := models.SelectManufacturer(c.ManufacturerId)
		consolesWithManufacturers = append(consolesWithManufacturers, ConsoleWithManufacturer{c, manufacturer.Name})
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(consolesWithManufacturers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}

func GetConsole(w http.ResponseWriter, r *http.Request) {

	id := utils.PathTail(r.URL.Path)

	fmt.Println("/console/get/", id)

	console := models.SelectConsole(id)
	manufacturer := models.SelectManufacturer(console.ManufacturerId)
	consolesWithManufacturer := ConsoleWithManufacturer{console, manufacturer.Name}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(consolesWithManufacturer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}


func SaveConsole(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var console models.Console

	console.ConsoleId = utils.PathTail(r.URL.Path)

	if console.ConsoleId != -1 {
		console = models.SelectConsole(console.ConsoleId)
	}

	console.Name = r.FormValue("name")
	console.MediaType = r.FormValue("mediaType")
	console.Sales = r.FormValue("sales")
	console.Year = r.FormValue("year")
	console.Handheld = r.FormValue("handheld") != ""
	console.ImageURL = r.FormValue("imageURL")
	console.Notes = r.FormValue("notes")
	manufacturer := r.FormValue("manufacturer")

	allManufacturers := models.SelectAllManufacturers()
	console.ManufacturerId = -1
	lastManufacturerId := 0
	for _, m := range allManufacturers {
		if m.ManufacturerId > lastManufacturerId {
			lastManufacturerId = m.ManufacturerId
		}
		if m.Name == manufacturer {
			console.ManufacturerId = m.ManufacturerId
		}
	}

	if console.ManufacturerId == -1 {

		fmt.Println("Saving new manufacturer...")

		newManufacturer := models.Manufacturer{ManufacturerId: lastManufacturerId + 1, Name: manufacturer}
		models.InsertManufacturer(newManufacturer)
		console.ManufacturerId = newManufacturer.ManufacturerId

	}

	fmt.Println("/console/save/", console)

	if console.ConsoleId == -1 {

		fmt.Println("Saving new console...")

		lastId := 0
		allConsoles := models.SelectAllConsoles()
		for _, c := range allConsoles {
			if c.ConsoleId > lastId {
				lastId = c.ConsoleId
			}
		}
		console.ConsoleId = lastId + 1

		models.InsertConsole(console)
		fmt.Fprint(w, "OK")

	} else {

		fmt.Println("Saving existing console...")

		models.UpdateConsole(console)
		fmt.Fprint(w, "OK")

	}

}

func DeleteConsole(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := utils.PathTail(r.URL.Path)

	if id == -1 {
		return
	}

	fmt.Println("/game/delete/", id)

	console := models.SelectConsole(id)
	if console.ConsoleId == 0 {
		fmt.Fprint(w, "Error: Can't find console with id", id)
		return
	}

	allConsoles := models.SelectAllConsoles()
	count := 0
	for _, c := range allConsoles {
		if c.ConsoleId != id && c.ManufacturerId == console.ManufacturerId {
			count++
		}
	}
	if count == 0 {
		models.DeleteManufacturer(console.ManufacturerId)
	}

	for _, c := range allConsoles {
		if c.ConsoleId == id {
			models.DeleteConsole(id)
			fmt.Fprint(w, "OK")
			return
		}
	}

}
