package main

import (
	"fmt"
	"os"

	"github.com/everettraven/devcontainers-go/pkg/devcontainers"
)

func main() {
	devcontainer, err := devcontainers.DevContainerFromPath(".devcontainer/devcontainer.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Devcontainer:")
	fmt.Printf("\tName: %s\n\tImage: %s\n\tFeatures: %v\n\tPostCreateCommand: %v\n", devcontainer.Name, devcontainer.Image, devcontainer.Features, devcontainer.PostCreateCommand)
}
