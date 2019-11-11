// learn from Mr.PanMaolin's "main.go"

package main

import (
	"os"

	"github.com/Passenger0/ServiceComputing/cloudgo/service"
	// or "./service"
	flag "github.com/spf13/pflag"
)

const (
	//PORT default:8080
	PORT string = "8080"
)

func main() {
	//set the port,if not set ,user default value 8080
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	// set the port for httpd listening
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	// setup the server
	server := service.NewServer()
	//run the server
	server.Run(":" + port)
}
