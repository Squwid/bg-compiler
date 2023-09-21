// Flags is a package defined to configure global variables
// that are effect the control of the program.
package flags

import (
	"runtime"
	"time"
)

// Overrides for FlagConfig based on flag values from the command line.
var FlagConfig *FlagConfiguration = &FlagConfiguration{
	// Use all available threads minus 2 for pre/post-processing.
	WorkerCount:             max(runtime.GOMAXPROCS(0)-2, 2),
	ContainerTimeoutSeconds: 30,
	JobChannelLength:        2000,
	ContainerMaxMemoryMB:    512, // 512MB
	MaxReadOutputBytesKB:    30,  // 30KB
}

// WorkerCount returns the number of workers to use for
// processing jobs.
func WorkerCount() int {
	return FlagConfig.WorkerCount
}

// ContainerMaxDuration is the maximum amount of time that a container can run
// before being killed.
func ContainerMaxDuration() time.Duration {
	return time.Duration(FlagConfig.ContainerTimeoutSeconds) * time.Second
}

// JobChannelLength returns the length of the job channel.
// This number needs to be larger than amount of
func JobChannelLength() int {
	return FlagConfig.JobChannelLength
}

// ContainerMaxMemory is the maximum amount of memory that a container can use
// for a single run.
func ContainerMaxMemory() int64 {
	return FlagConfig.ContainerMaxMemoryMB * 1024 * 1024
}

// ContainerMaxCPU is the maximum amount of CPU that a container can use for a
// single run.
func ContainerMaxCPU() int64 {
	return int64((1 / WorkerCount()) * WorkerCount() * 1024)
}

// MaxReadOutputBytes is the maximum number of bytes that can be read from a
// container output before the container is killed.
//
// The largest API call output should be approximately this * MaxBatchCompiles.
func MaxReadOutputBytes() int {
	return FlagConfig.MaxReadOutputBytesKB * 1024
}

// TODO: List of additional flags that should be added at some point:
//     - ShouldKillWhenOutputFull
//     - EnforceTimeout
