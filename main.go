package main

import (
	"BE_datnd/config"
	"BE_datnd/route"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization

func main() {
	err := gotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn := config.InitMysql()
	defer config.CloseConnectDB(conn)


	r := route.SetupRouter(conn)

	appPort := os.Getenv("APP_PORT")
	r.Run(appPort)
}
