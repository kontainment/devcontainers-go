package devcontainers

// TODO(everettraven): Consider using Go generics to implement the different potential types:
// 1) "dockerfile"
// 2) "dockerFile"
type DevContainerDockerfileContainer struct {
	// Docker build-related options.
	Build DockerfileBuildOpts
}

// TODO(everettraven): Consider using Go generics or custom unmarshaller to implement the different potential types:
// 1) "dockerfile"
// 2) "dockerFile"
type DockerfileBuildOpts struct {
	// The location of the Dockerfile that defines the contents of the container.
	// The path is relative to the folder containing the `devcontainer.json` file.
	Dockerfile string

	// The location of the context folder for building the Docker image.
	// The path is relative to the folder containing the `devcontainer.json` file.
	Context string
}

type BuildOptions struct {
	// Target stage in a multi-stage build.
	Target string

	// Build arguments.
	Args []string

	// The image to consider as a cache. Use an array to specify multiple images.
	CacheFrom []string
}

type ImageContainer struct {
	// The docker image that will be used to create the container.
	Image string
}

type ComposeContainer struct {
	// The name of the docker-compose file(s) used to start the services.
	DockerComposeFile []string

	// The service you want to work on. This is considered the primary
	// container for your dev environment which your editor will connect to.
	Service string

	// An array of services that should be started and stopped.
	RunServices []string

	// The path of the workspace folder inside the container.
	// This is typically the target path of a volume mount in the docker-compose.yml.
	WorkspaceFolder string

	// Action to take when the user disconnects from the primary container in their
	// editor. The default is to stop all of the compose containers.
	ShutdownAction string

	// Whether to overwrite the command specified in the image. The default is false.
	OverrideCommand bool
}
