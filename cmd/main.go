package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anthanh17/go-react-blog/pkg/config"
	"github.com/anthanh17/go-react-blog/pkg/db/repository"
	"github.com/anthanh17/go-react-blog/pkg/telegram"
)

func main() {
	config.SetupConfig()

	err := repository.SetupDb()
	if err != nil {
		log.Fatal(err)
	}
	for {
		telegram.SendMessage(fmt.Sprintf("[Ping] database successful, connection is still alive"), 1)
		time.Sleep(2 * time.Second)
	}
	//defer telegram.SendMessage(fmt.Sprintf("[Ping] database successful, connection is still alive"), 1)
	fmt.Println("Ping to database successful, connection is still alive")
}
