package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not connnect to the database: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSer := user.NewService(userRep)
	userHan := user.NewHandler(userSer)

	hub := ws.NewHub()
	hubHan := ws.NewHandler(hub)

	router.InitRouter(userHan, hubHan)
	router.Start("0.0.0.0:8080")
}
