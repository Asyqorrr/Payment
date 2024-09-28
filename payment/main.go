package main

import (
	"log"
	"os"
	"payment/config"
	"payment/server"
)

func main() {
	log.Println("Starting payment API")

	log.Println("Initialiazing Configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)

	// log.Println(dbHandler)

	log.Println("Initializing HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env != "" {
		return "payment-" + env
	}

	return "payment"
}
