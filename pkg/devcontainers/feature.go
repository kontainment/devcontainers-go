package devcontainers

// Struct to implement the reading of devcontainer-feature.json files
type DevContainerFeature struct {
	// ID of the feature/definition. The id should be unique in the context
	// of the repository/published package where the feature exists and
	// must match the name of the directory where the devcontainer-feature.json resides.
	Id string

	// The semantic version of the Feature.
	Version string

	// Name of the feature/definition.
	Name string

	// Description of the feature/definition.
	Description string

	// Url that points to the documentation of the Feature.
	DocumentationUrl string

	// Url that points to the license of the Feature.
	LicenseUrl string

	// List of strings relevant to a user that would search for this definition/Feature.
	Keywords []string

	// A map of options that will be passed as environment variables to the execution of the script.
	Options map[string]Option

	// A set of name value pairs that sets or overrides environment variables.
	ContainerEnv map[string]interface{}

	// Sets privileged mode for the container (required by things like docker-in-docker) when the feature is used.
	Privileged bool

	// Adds the tiny init process to the container (--init) when the Feature is used.
	Init bool

	// Adds container capabilities when the Feature is used
	CapAdd []string

	// Sets container security options like updating the seccomp profile when the Feature is used.
	SecurityOpt []string

	// Set if the feature requires an “entrypoint” script that should fire at container start up.
	Entrypoint string

	// Product specific properties, each namespace under customizations is treated as a separate
	// set of properties. For each of this sets the object is parsed, values are replaced while
	// arrays are set as a union.
	Customizations interface{}

	// Array of ID’s of Features (omitting a version tag) that should execute before this one.
	// Allows control for Feature authors on soft dependencies between different Features.
	InstallsAfter []string

	//Array of old IDs used to publish this Feature. The property is useful for renaming
	// a currently published Feature within a single namespace.
	LegacyIds []string

	// Indicates that the Feature is deprecated, and will not receive any further updates/support.
	// This property is intended to be used by the supporting tools for highlighting Feature deprecation.
	Deprecated bool

	// Defaults to unset. Cross-orchestrator way to add additional mounts to a container.
	// Each value is an object that accepts the same values as the Docker CLI --mount flag.
	// The Pre-defined devcontainerId variable may be referenced in the value. For example:
	// "mounts": [{ "source": "dind-var-lib-docker", "target": "/var/lib/docker", "type": "volume" }]
	Mounts Mount
}

type Option struct {
	// Type of the option. Valid types are currently: boolean, string
	Type string

	// A list of suggested string values. Free-form values are allowed. Omit when using optionId.enum
	Proposals []string

	// A strict list of allowed string values. Free-form values are not allowed. Omit when using optionId.proposals.
	Enum []string

	// TODO(everettraven): This can be a string or a boolean - Consider using generics?
	// Default value for the option
	Default string

	// Description for the option.
	Description string
}
