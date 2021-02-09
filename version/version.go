package version

var (
	// BuildVersion is the git tag of this library.
	BuildVersion = "v0.3.0"
	// BuildCommit is the short git commit of the built binary.
	// If using this as a library, it will always be "dev".
	BuildCommit = "dev"
	// BuildName is the name of the build. If using as a library, this
	// will always be magicbell-go. But can be overridden for binary builds,
	// for example the mbctl CLI tool.
	BuildName = "magicbell-go"
)
