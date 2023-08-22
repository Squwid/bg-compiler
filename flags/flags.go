// Flags is a package defined to control global variables
// that are effect the control of the program.
package flags

import (
	"runtime"
	"time"
)

// WorkerCount returns the number of workers to use for
// processing jobs.
func WorkerCount() int {
	// We want to use all available threads minus 2 for pre-processing.
	return max(runtime.GOMAXPROCS(0)-2, 2)
}

// ContainerMaxDuration is the maximum amount of time that a container can run
// before being killed.
func ContainerMaxDuration() time.Duration {
	return 30 * time.Second
}

// JobChannelLength returns the length of the job channel.
// This number needs to be larger than amount of
func JobChannelLength() int {
	return 2000
}

// The amount of bytes that we can have fetched from PubSub before we stop
// accepting new payloads from the queue.
func PubSubMaxOutstandingBytes() int {
	return -1 // -1 == unlimited.
}

// ContainerMaxMemory is the maximum amount of memory that a container can use
// for a single run.
func ContainerMaxMemory() int64 {
	return 512 * 1024 * 1024 // 512MB
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
	return 1024 * 30
}
