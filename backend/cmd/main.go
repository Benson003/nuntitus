package main

import (
	"net/http"

	"github.com/Benson003/nuntius/tools/router"
)

func main() {
	router := router.New()
	http.ListenAndServe(":3000", router)
}
