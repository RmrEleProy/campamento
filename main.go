package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/RmrEleProy/campamento/db"
	"github.com/RmrEleProy/campamento/server"
)

func main() {
	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")

	serv := server.ServerNew(":8080")

	go serv.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	serv.Close()
}
