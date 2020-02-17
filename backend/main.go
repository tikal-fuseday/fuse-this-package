package main

import (
	"net/http"
	"search-package/controllers"
)

func main() {
	controllers.RegisterSearchController()
	http.ListenAndServe(":3000", nil)
}
