package main

import (
	"fmt"
	"os"
)

func main() {
	version := os.Getenv("LATEST_VERSION")

	fmt.Printf("The release version is: %s\n", version)
}
