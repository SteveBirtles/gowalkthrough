package controllers

import (
	"../../server/models"
	"../../server/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type AccessoriesWithConsole struct {
	ConsoleName string             `json:"consoleName"`
	Accessories []models.Accessory `json:"accessoriesList"`
}

type AccessoryWithCategory struct {
	models.Accessory
	CategoryName string             `json:"category"`
}

func ListAccessories(w http.ResponseWriter, r *http.Request) {

	fmt.Println("/accessory/list")

	consoleId := utils.PathTail(r.URL.Path)

	allAccessories := models.SelectAllAccessorys()
	relevantAccessories := make([]models.Accessory, 0)
	for _, g := range allAccessories {
		if g.ConsoleId == consoleId {
			relevantAccessories = append(relevantAccessories, g)
		}
	}

	accessoriesWithConsole := AccessoriesWithConsole{models.SelectConsole(consoleId).Name, relevantAccessories}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accessoriesWithConsole)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}

func GetAccessory(w http.ResponseWriter, r *http.Request) {

	id := utils.PathTail(r.URL.Path)

	fmt.Println("/accessory/get/", id)

	accessory := models.SelectAccessory(id)
	category := models.SelectCategory(accessory.CategoryId)
	accessoryWithCategory := AccessoryWithCategory{accessory, category.Name}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accessoryWithCategory)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}

}

func SaveAccessory(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var accessory models.Accessory

	accessory.AccessoryId = utils.PathTail(r.URL.Path)

	if accessory.AccessoryId != -1 {
		accessory = models.SelectAccessory(accessory.AccessoryId)
	}

	fmt.Print("consoleId:", r.FormValue("consoleId"))
	accessory.ConsoleId, _ = strconv.Atoi(r.FormValue("consoleId"))

	accessory.Description = r.FormValue("description")
	accessory.Quantity = r.FormValue("quantity")
	accessory.ThirdParty = r.FormValue("thirdParty") != ""
	accessory.ImageURL = r.FormValue("imageURL")
	category := r.FormValue("category")
	accessory.CategoryId = -1

	allCategories := models.SelectAllCategorys()

	lastCategoryId := 0
	for _, c := range allCategories {
		if c.CategoryId > lastCategoryId {
			lastCategoryId = c.CategoryId
		}
		if c.Name == category {
			accessory.CategoryId = c.CategoryId
		}
	}

	if accessory.CategoryId == -1 {

		fmt.Println("Saving new category...")

		newCategory := models.Category{CategoryId: lastCategoryId + 1, Name: category}
		models.InsertCategory(newCategory)
		accessory.CategoryId = newCategory.CategoryId

	}

	fmt.Println("/accessory/save/", accessory)

	if accessory.AccessoryId == -1 {

		fmt.Println("Saving new accessory...")

		lastId := 0
		allAccessories := models.SelectAllAccessorys()
		for _, a := range allAccessories {
			if a.AccessoryId > lastId {
				lastId = a.AccessoryId
			}
		}
		accessory.AccessoryId = lastId + 1

		models.InsertAccessory(accessory)
		fmt.Fprint(w, "OK")

	} else {

		fmt.Println("Saving existing accessory...")

		models.UpdateAccessory(accessory)
		fmt.Fprint(w, "OK")

	}

}

func DeleteAccessory(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := utils.PathTail(r.URL.Path)

	if id == -1 {
		return
	}

	fmt.Println("/accessory/delete/", id)

	accessory := models.SelectAccessory(id)
	if accessory.AccessoryId == 0 {
		fmt.Fprint(w, "Error: Can't find accessory with id", id)
		return
	}

	allAccessories := models.SelectAllAccessorys()
	count := 0
	for _, c := range allAccessories {
		if c.AccessoryId != id && c.CategoryId == accessory.CategoryId {
			count++
		}
	}
	if count == 0 {
		models.DeleteManufacturer(accessory.CategoryId)
	}

	for _, a := range allAccessories {
		if a.AccessoryId == id {
			models.DeleteConsole(id)
			fmt.Fprint(w, "OK")
			return
		}
	}

}

