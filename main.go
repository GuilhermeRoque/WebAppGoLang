package main

import (
	"goWebApp/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000",nil)
}
