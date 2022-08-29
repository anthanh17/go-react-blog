package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anthanh17/go-react-blog/pkg/config"
	"github.com/anthanh17/go-react-blog/pkg/db/repository"
	"github.com/anthanh17/go-react-blog/pkg/mail"
	"github.com/anthanh17/go-react-blog/pkg/telegram"
)

func main() {
	config.SetupConfig()

	_, err := repository.SetupDb()
	if err != nil {
		log.Fatal(err)
	}

	telegram.SendMessage(fmt.Sprintf("[Ping] database successful, connection is still alive"), 1)
	time.Sleep(2 * time.Second)

	mail.SendNews("final merage")
	time.Sleep(2 * time.Second)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
