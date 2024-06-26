package flags

type FlagConfiguration struct {
	// WorkerCount is the number of workers to use for processing jobs.
	// Defaults to the number of threads minus 2 for pre-processing.
	WorkerCount int

	// Maximum amount of time that a container can run before being killed.
	// Defaults to 30 seconds.
	ContainerTimeoutSeconds int

	// Backlogged job count before the server starts rejecting requests.
	// Defaults to 2000.
	JobChannelLength int

	// Maximum amount of memory that a container can use for a single run.
	// Defaults to 512MB.
	ContainerMaxMemoryMB int64

	// CPU shares for a container.
	// Defaults to 1024.
	ContainerCPUShares int64

	// Number of bytes that can be read from a container output before the
	// container is killed.
	// Defaults to 30KB.
	MaxReadOutputBytesKB int

	// Server port to listen on.
	Port int

	// UseGVisor is a flag to enable the use of gVisor for container isolation.
	// runsc needs to be installed on the host machine for this to work.
	// Defaults to false.
	UseGVisor bool
}
