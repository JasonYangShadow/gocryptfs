//go:build race
// +build race

package pkg

func init() {
	// adds " -race" to the output of "gocryptfs -version"
	raceDetector = true
}
