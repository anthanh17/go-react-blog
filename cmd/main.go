package main

import (
	"log"

	"github.com/anthanh17/go-react-blog/pkg/config"
	"github.com/anthanh17/go-react-blog/pkg/db/repository"
	"github.com/anthanh17/go-react-blog/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.SetupConfig()

	_, err := repository.SetupDb()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")

	// telegram.SendMessage(fmt.Sprintf("[Ping] database successful, connection is still alive"), 1)
	// time.Sleep(2 * time.Second)

	// mail.SendNews("final merage")
	// time.Sleep(2 * time.Second)

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	// <-interrupt
}
