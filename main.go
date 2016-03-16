package main

import (
	"net/http"
	"fmt"
	"github.com/zacharyjacobcollins/ShiftAPI/api"
)

func main() {
	var port string =  ":3000"
	fmt.Sprintf("Server Starting on port %s", port)
	http.ListenAndServe(port, api.Handlers())
}