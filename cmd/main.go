package main

import (
	"fmt"
	"github.com/rousquille/poc-go-svelte-websockets/internal/api"
	"github.com/rousquille/poc-go-svelte-websockets/internal/cli"
	"log"
	"os"
)

func main() {
	cli.ParseArgs()

	if *cli.Version {
		fmt.Println("poc-go-svelte-websockets version :", cli.GlobalVersion)
		os.Exit(0)
	}

	err := api.RunWebServer(*cli.Port)
	if err != nil {
		log.Fatal(err)
	}

}
