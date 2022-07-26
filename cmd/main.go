package main

import (
	"fmt"
	config "blog/pkg/config/db"
	"gorm.io/gorm"
)

type Server struct {
	Db *gorm.DB
 }

func main() {
	/*Just used to test. Shouldn't be called in main function*/
	var server = Server{}
	server.Db = config.InitDb();

    // if no error. Ping is successful
	fmt.Println("Ping to database successful, connection is still alive")	
}