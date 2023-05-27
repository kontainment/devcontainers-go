package devcontainers

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	ProtocolHttp  = "http"
	ProtocolHttps = "https"

	// Shows a notification when a port is automatically forwarded
	OnAutoForwardNotify = "notify"
	// Opens the browser when a port is automatically forwarded.
	// Depending on your settings this could open an embedded browser.
	OnAutoForwardOpenBrowser = "openBrowser"
	// Opens the browser when the port is automatically forwarded, but
	// only the first time the port is forward during a session.
	// Depending on your settings, this could open an embedded browser.
	OnAutoForwardOpenBrowserOnce = "openBrowserOnce"
	// Opens a preview in the same window when the port is automatically forwarded.
	OnAutoForwardOpenPreview = "openPreview"
	// Shows no notification and takes no action when this port is automatically forwarded.
	OnAutoForwardSilent = "silent"
	// This port will not be automatically forwarded.
	OnAutoForwardIgnore = "ignore"

	WaitForInitializeCommand    = "initializeCommand"
	WaitForOnCreateCommand      = "onCreateCommand"
	WaitForUpdateContentCommand = "updateContentCommand"
	WaitForPostCreateCommand    = "postCreateCommand"
	WaitForPostStartCommand     = "postStartCommand"

	UserEnvProbeNone                  = "none"
	UserEnvProbeLoginShell            = "loginShell"
	UserEnvProbeLoginShellInteractive = "loginShellInteractive"
	UserEnvProbeShellInteractive      = "shellInteractive"

	ShutdownActionNone          = "none"
	ShutdownActionStopContainer = "stopContainer"

	MountTypeBind   = "bind"
	MountTypeVolume = "volume"
)

type DevContainerCommon struct {
	// A name for the dev container which can be displayed to the user.
	Name string

	// Features to add to the dev container
	Features *InterfaceMap

	// Array consisting of the Feature id (without the semantic version) of
	// Features in the order the user wants them to be installed.
	OverrideFeatureInstallOrder []string

	// Ports that are forwarded from the container to the local machine.
	// Can be an integer port number, or a string of the format "host:port_number"
	// NOTE: For now only strings will be accepted
	// TODO(everettraven): Investigate using Go generics for allowing integers and strings
	ForwardPorts []string

	// Set default properties that are applied when a specific port number is forwarded.
	// For example:
	// ```
	// "3000": {"label": "Application"},
	// "40000-55000": {"onAutoForward": "ignore"},
	// ".+\\\\/server.js": {"onAutoForward": "openPreview"}
	// ```
	PortAttributes map[string]PortAttribute

	// Set default properties that are applied to all ports that don't get properties
	// from the setting `remote.portsAttributes`. For example:
	// ```
	// {"onAutoForward": "ignore"}
	// ```
	OtherPortsAttributes PortAttribute

	// Controls whether on Linux the container's user should be updated with the local
	// user's UID and GID. On by default when opening from a local folder.
	UpdateRemoteUserUID bool

	// Container environment variables
	ContainerEnv *InterfaceMap

	// The user the container will be started with. The default is the user on the Docker image.
	ContainerUser string

	// Mount points to set up when creating the container.
	// See Docker's documentation for the --mount option for the supported syntax.
	Mounts []Mount

	// Passes the --init flag when creating the dev container.
	Init bool

	// Passes the --privileged flag when creating the dev container.
	Privileged bool

	// Passes docker capabilities to include when creating the dev container.
	CapAdd []string

	// Passes docker security options to include when creating the dev container.
	SecurityOpt []string

	// Remote environment variables to set for processes spawned in the container
	// including lifecycle scripts and any remote editor/IDE server process.
	RemoteEnv *InterfaceMap

	// The username to use for spawning processes in the container including
	// lifecycle scripts and any remote editor/IDE server process.
	// The default is the same user as the container.
	RemoteUser string

	// A command to run locally before anything else.
	// This command is run before \"onCreateCommand\".
	// If this is a single string, it will be run in a shell.
	// If this is an array of strings, it will be run as a single command without shell.
	// NOTE: right now this is only accepting an array
	// TODO(everettraven): Investigate using generics for accepting string and []string
	InitializeCommand Command

	// A command to run when creating the container. This command is run after initializeCommand
	// and before updateContentCommand. If this is a single string, it will be run in a shell.
	// If this is an array of strings, it will be run as a single command without shell.
	// TODO(everettraven): Investigate using generics for accepting string and []string and object
	OnCreateCommand Command

	// A command to run when creating the container and rerun when the workspace content was
	// updated while creating the container. This command is run after onCreateCommand and
	// before postCreateCommand. If this is a single string, it will be run in a shell.
	// If this is an array of strings, it will be run as a single command without shell.
	// TODO(everettraven): Investigate using generics for accepting string and []string and object
	UpdateContentCommand Command

	// A command to run after creating the container. This command is run after updateContentCommand
	// and before postStartCommand. If this is a single string, it will be run in a shell.
	// If this is an array of strings, it will be run as a single command without shell.
	// TODO(everettraven): Investigate using generics for accepting string and []string and object
	PostCreateCommand Command

	// A command to run after starting the container. This command is run after postCreateCommand
	// and before postAttachCommand. If this is a single string, it will be run in a shell.
	// If this is an array of strings, it will be run as a single command without shell.
	// TODO(everettraven): Investigate using generics for accepting string and []string and object
	PostStartCommand Command

	// A command to run when attaching to the container.
	// This command is run after postStartCommand.
	// If this is a single string, it will be run in a shell.
	// If this is an array of strings, it will be run as a single command without shell.
	// TODO(everettraven): Investigate using generics for accepting string and []string and object
	PostAttachCommand Command

	// The user command to wait for before continuing execution in the background
	// while the UI is starting up. The default is updateContentCommand.
	// TODO(everettraven): Consider using a custom type that implements the
	// unmarshaller interface so we can set default values.
	WaitFor string

	// User environment probe to run. The default is loginInteractiveShell.
	// TODO(everettraven): Consider using a custom type that implements the
	// unmarshaller interface so we can set default values.
	UserEnvProbe string

	// Host hardware requirements
	HostRequirements HostRequirement

	// Tool-specific configuration. Each tool should use a JSON object
	// subproperty with a unique name to group its customizations.
	Customizations interface{}

	// Any additional custom properties
	AdditionalProperties interface{}
}

type PortAttribute struct {
	// Label that will be shown in the UI for this port
	Label Label

	// Defines the action that occurs when the port is discovered for automatic forwarding
	OnAutoForward AutoForward

	// Automatically prompt for elevation (if needed) when this port is forwarded.
	// Elevate is required if the local port is a privileged port.
	ElevateIfNeeded bool

	// When true, a modal dialog will show if the chosen local port isn't used for forwarding.
	RequireLocalPort bool

	// The protocol to use when forwarding this port
	Protocol string
}

// Host hardware requirements.
// TODO(everettraven): Consider implementing a generic
// unit size type to verify when unmarshalling based on
// a regex pattern.
type HostRequirement struct {
	// Number of required CPUs.
	Cpus int

	// Amount of required RAM in bytes. Supports units tb, gb, mb and kb.
	Memory string

	// Amount of required disk space in bytes. Supports units tb, gb, mb and kb.
	Storage string

	// TODO(everettraven): Implement a custom type that implements the
	// unmarshaller interface and can unmarshal the devcontainers specification for a GPU
	Gpu bool
}

// TODO(everettraven): implement the Unmarshaller interface on this type
type Label string

// TODO(everettraven): implement the Unmarshaller interface on this type
type AutoForward string

type InterfaceMap map[string]interface{}

// TODO(everettraven): implement this && consider generics
type Mount struct {
	// Mount type.
	Type string

	// Mount source.
	Source string

	// Mount target.
	Target string
}

type Command struct {
	Cmd  string
	Args []string
}

func (c *Command) UnmarshalJSON(data []byte) error {
	jsonData := fmt.Sprintf("{ \"data\": %s }", string(data))
	parseData := map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonData), &parseData)
	if err != nil {
		return err
	}

	val, ok := parseData["data"]
	if !ok {
		return fmt.Errorf("no data??")
	}

	switch v := val.(type) {
	case string:
		strs := strings.Split(v, " ")
		c.Cmd = strs[0]
		c.Args = strs[1:]
	// TODO(everettraven): Fix this to parse []interface{} since that is what an array is actually parsed as
	case []string:
		c.Cmd = v[0]
		c.Args = v[1:]
	default:
		return fmt.Errorf("invalid type %T", v)
	}

	return nil
}
