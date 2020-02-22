package main

import (
	"fmt"
	"net/http"
	"runtime"
	"search-package/controllers"
	"strconv"
)

var port = 3000

func main() {
	runtime.GOMAXPROCS(4)
	controllers.RegisterSearchController()
	fmt.Printf("%s %d \n", "Server running on port", port)
	http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
}
