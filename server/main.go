package main

import (
	"../server/controllers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/client/", controllers.StaticContent)

	http.HandleFunc("/console/list", controllers.ListConsoles)
	http.HandleFunc("/game/list/", controllers.ListGames)

	/*/console/list
	/game/list/{id}
	/accessory/list/{id}
	/admin/login
	/admin/check
	/console/get/{id}
	/console/save/{id}
	/console/delete/{id}
	/game/get/{id}
	/game/save/{id}
	/game/delete/{id}
	/accessory/get/{id}
	/accessory/save/{id}
	/accessory/delete/{id}*/

	err := http.ListenAndServe(":8081", http.DefaultServeMux)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
