package devcontainers

type DevContainerNonComposeBase struct {
	// Application ports that are exposed by the container.
	// This can be a single port or an array of ports.
	// Each port can be a number or a string. A number is
	// mapped to the same port on the host. A string is
	// passed to Docker unchanged and can be used to map
	// ports differently, e.g. \"8000:8010\".
	// TODO(everettraven): Investigate Go generics to allow this to be of type integer, string, or []string
	AppPort []string

	// The arguments required when starting in the container.
	RunArgs []string

	// Action to take when the user disconnects from the container
	// in their editor. The default is to stop the container.
	// TODO(everettraven): Consider using a custom type for this and implementing the Unmarshaller interface to set a default value
	ShutdownAction string

	// Whether to overwrite the command specified in the image. The default is true.
	// TODO(everettraven): How to default a bool to true?
	OverrideCommand bool

	// The path of the workspace folder inside the container.
	WorkspaceFolder string

	// The --mount parameter for docker run. The default is to mount the project folder at /workspaces/$project.
	WorkspaceMount string
}
