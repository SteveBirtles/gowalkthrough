package controllers

import (
	"../../server/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListConsoles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("/console/list")
	consoles := models.SelectAllConsoles()

	type ConsoleWithManufacturer struct {
		models.Console
		Manufacturer string `json:"manufacturer"`
	}

	consolesWithManufacturers := make([]ConsoleWithManufacturer, 0)
	for _, c := range consoles {
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

}

func SaveConsole(w http.ResponseWriter, r *http.Request) {

}

func DeleteConsole(w http.ResponseWriter, r *http.Request) {

}
