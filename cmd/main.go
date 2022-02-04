package main

/*
	1200 бит/с
	паритет нечётный
	7 бит
	1 стоповый
*/

import (
	"github.com/bicycledays/bridge/src/handler"
	"github.com/bicycledays/bridge/src/server"
	"github.com/bicycledays/bridge/src/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	envInit()
	port, found := os.LookupEnv("LOCAL_PORT")
	if !found {
		port = "8080"
	}
	s := service.NewService()
	h := handler.NewHandler(s)
	srv := new(server.Server)
	if err := srv.Run(port, h.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}

func envInit() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
