package main

import (
	"fmt"
	"log"

	"github.com/anthanh17/go-react-blog/pkg/config"
	"github.com/anthanh17/go-react-blog/pkg/db"
	"github.com/anthanh17/go-react-blog/pkg/db/repository"
)

func main() {
	config := config.LoadConfig()

	db.SetupConfig(config)

	err := repository.SetupDb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ping to database successful, connection is still alive")
}
