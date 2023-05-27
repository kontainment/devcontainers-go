package devcontainers

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO(everettraven): Add JSON tags to all fields

// TODO(everettraven): Look into custom JSON parsing+validation
// logic to match the devcontainers JSON schema

// Go type to represent a devcontainer
type DevContainer struct {
	DevContainerBase
}

// Go type to represent the devcontainer.base schema
type DevContainerBase struct {
	DevContainerCommon
	DevContainerNonComposeBase
	DevContainerDockerfileContainer
	BuildOptions
	ImageContainer
	ComposeContainer
	Mount
}

// DevContainerFromPath reads a devcontainer.json file from the specified path
// Inputs: A string file path. This must include the file name.
// Outputs: DevContainer object if successful, error if one is encountered
func DevContainerFromPath(path string) (*DevContainer, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	devcontainer := &DevContainer{}

	err = json.Unmarshal(contents, devcontainer)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling devcontainer: %w", err)
	}

	return devcontainer, nil
}
