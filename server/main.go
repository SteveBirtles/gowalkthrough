package main

import (
	"../server/controllers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/client/", controllers.StaticContent)

	http.HandleFunc("/admin/login", controllers.AdminLogin)
	http.HandleFunc("/admin/check", controllers.AdminCheck)

	http.HandleFunc("/console/list", controllers.ListConsoles)
	http.HandleFunc("/console/get/", controllers.GetConsole)
	http.HandleFunc("/console/save/", controllers.SaveConsole)
	http.HandleFunc("/console/delete/", controllers.DeleteConsole)

	http.HandleFunc("/game/list/", controllers.ListGames)
	http.HandleFunc("/game/get/", controllers.GetGame)
	http.HandleFunc("/game/save/", controllers.SaveGame)
	http.HandleFunc("/game/delete/", controllers.DeleteGame)

	http.HandleFunc("/accessory/list/", controllers.ListAccessories)
	http.HandleFunc("/accessory/get/", controllers.GetAccessory)
	http.HandleFunc("/accessory/save/", controllers.SaveAccessory)
	http.HandleFunc("/accessory/delete/", controllers.DeleteAccessory)

	err := http.ListenAndServe(":8081", http.DefaultServeMux)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
