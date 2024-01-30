package main

import (
	"fmt"

	version "github.com/ynhummel/github-env-test/internal"
)

func main() {

	version := version.Version
	if version == "" {
		version = "versionNotFound"
	}

	fmt.Printf("The release version is: %s\n", version)
}
