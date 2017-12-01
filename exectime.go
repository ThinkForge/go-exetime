package exectime

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func exectime(folder string) (timeresult, error) {

	// Store the time for each run.
	var cutTimeArr, exTimeArr [5]int64

	// Parse the filepath into slices and then get the files.
	arr := strings.Split(folder, "/")
	pathCut, errCut := exec.LookPath(folder + arr[len(arr)-2] + ".cut.out")
	pathEx, errEx := exec.LookPath(folder + arr[len(arr)-2] + ".ex.out")

	// If either of the files could not be found, return the error.
	if errCut != nil {
		return timeresult{ratio: nil}, errCut
	} else if errEx != nil {
		return timeresult{ration: nil}, errEx
	}

	// For each of the input files we evaluate it using both executables.
	for i := 0; i < 5; i++ {
		cmd := exec.Command(pathCut, folder+strconv.Itoa(i)+".txt")
		started := start()
		cmd.Run()
		time := end(started)

		result := timeresult{cutTime: time, exTime: time}
	}
	return result, nil
}

func start() time.Time {
	return time.Now()
}

func end(startTime time.Time) int64 {
	endTime := time.Now()
	return (endTime.Sub(startTime)).Nanoseconds()
}

type timeresult struct {
	ratio [5]int64
}
