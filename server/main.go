package main

import (
	"../server/controllers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/client/", controllers.StaticContent)
	http.HandleFunc("/console/list", controllers.ListConsoles)

	err := http.ListenAndServe(":8081", http.DefaultServeMux)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
