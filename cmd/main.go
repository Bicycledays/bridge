package main

import (
	"fmt"
	"github.com/bicycledays/bridge/src/handler"
	"github.com/bicycledays/bridge/src/server"
	"github.com/bicycledays/bridge/src/service"
	"github.com/gookit/color"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	label()
	s := service.NewService()
	h := handler.NewHandler(s)
	srv := new(server.Server)

	err := srv.Run(port(), h.InitRoutes())

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func label() {
	label := "\n\n\n ███████╗ █████╗ ██████╗ ████████╗ ██████╗ ██████╗ ██╗██╗   ██╗███████╗  \n ██╔════╝██╔══██╗██╔══██╗╚══██╔══╝██╔═══██╗██╔══██╗██║██║   ██║██╔════╝  \n ███████╗███████║██████╔╝   ██║   ██║   ██║██████╔╝██║██║   ██║███████╗  \n ╚════██║██╔══██║██╔══██╗   ██║   ██║   ██║██╔══██╗██║██║   ██║╚════██║  \n ███████║██║  ██║██║  ██║   ██║   ╚██████╔╝██║  ██║██║╚██████╔╝███████║  \n ███████║██║  ██║██║  ██║   ██║    ██████║ ██║  ██║██║ ██████║ ███████║  \n ╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝ ╚═════╝ ╚══════╝\n\n"
	yellow := color.FgYellow.Render
	fmt.Printf("%s\n", yellow(label))
}

func port() string {
	_ = godotenv.Load("config")
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "48700"
	}
	log.Println("network port", port)
	return port
}
