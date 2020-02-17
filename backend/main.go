package main

import (
	"fmt"
	"net/http"
	"search-package/controllers"
	"strconv"
)

var port = 3000

func main() {
	controllers.RegisterSearchController()
	fmt.Printf("%s %d \n", "Server running on port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
