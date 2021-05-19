package main

import (
	"log"
	"prometheus_devops/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
