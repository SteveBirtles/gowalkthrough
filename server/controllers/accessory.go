package controllers

import (
	"../../server/models"
	"../../server/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListAccessories(w http.ResponseWriter, r *http.Request) {

	type AccessoryWithConsole struct {
		ConsoleName string             `json:"consoleName"`
		Accessories []models.Accessory `json:"accessoriesList"`
	}

	fmt.Println("/accessory/list")

	consoleId := utils.PathTail(r.URL.Path)

	allAccessories := models.SelectAllAccessories()
	relevantAccessories := make([]models.Accessory, 0)
	for _, g := range allAccessories {
		if g.ConsoleId == consoleId {
			relevantAccessories = append(relevantAccessories, g)
		}
	}

	accessoryWithConsole := AccessoryWithConsole{models.SelectConsole(consoleId).Name, relevantAccessories}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accessoryWithConsole)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}

func GetAccessory(w http.ResponseWriter, r *http.Request) {

	id := utils.PathTail(r.URL.Path)

	fmt.Println("/accessory/get/", id)

	accessory := models.SelectAccessory(id)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accessory)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}

func SaveAccessory(w http.ResponseWriter, r *http.Request) {

}

func DeleteAccessory(w http.ResponseWriter, r *http.Request) {

}
