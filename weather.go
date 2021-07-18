package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
    app := cli.NewApp()
	app.Name = "Weather CLI Tool" // Assign the name of the CLI app
    app.Usage = "Prompt the current weather condition for a city" // Assign the usage description for the CLI 
    app.Version = "1.0.0" //Assign the version of the CLI app

	app.Run(os.Args)
}