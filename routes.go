package main

import (
	"fmt"
	"net/http"

	"monthlyCoder/modules/http_controllers"
)

// start routes
func startRoutes() {
	fmt.Println("starting routes..")

	http.HandleFunc("/", http_controllers.Index)
}
