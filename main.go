package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".version")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	version := os.Getenv("LATEST_VERSION")
	if version == "" {
		version = "versionNotFound"
	}

	fmt.Printf("The release version is: %s\n", version)
}
