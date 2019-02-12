package main

import (
	"log"
	"os"

	"github.com/ldmtam/serverless-todo/src/utils"

	"github.com/apex/gateway"
	"github.com/ldmtam/serverless-todo/src/database"
)

func main() {
	mode := os.Getenv("MODE")

	mysql, err := database.NewMySQL()
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	router := initRouter(mode, mysql)

	if mode == utils.Production {
		log.Fatal(gateway.ListenAndServe(":3000", router))
	} else {
		log.Fatal(router.Run(":8080"))
	}
}
