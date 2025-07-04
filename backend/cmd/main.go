package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Benson003/nuntius/api"
	"github.com/Benson003/nuntius/database"
	router "github.com/Benson003/nuntius/routes"
	"github.com/Benson003/nuntius/tools/files"
)

func main() {
	if err := files.InitalizeFolder(); err != nil {
		fmt.Printf("error initalising folder:%e", err)
		return
	}
	db := database.InitDataBase("nuntius.db")
	if db.Error != nil {
		log.Fatalf("couldn't initialize database: %v", db.Error)
		return
	}

	h := api.Handler{
		DBObject: db,
	} // 🔹 create your handler with DB

	r := router.NewRouter(&h) // 🔹 pass handler into router

	fmt.Println("🚀 Server started on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
