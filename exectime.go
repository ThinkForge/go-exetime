package exectime

import (
	"os/exec"
	"strings"
	"time"
)

func exectime(folder string) (timeresult, error) {
	arr := strings.Split(folder, "/")
	path, err := exec.LookPath(folder + arr[len(arr)-2] + ".cut.out")

	if err != nil {
		return timeresult{ratio: nil}, err
	}

	cmd := exec.Command(path)

	started := start()

	cmd.Run()

	time := end(started)

	result := timeresult{cutTime: time, exTime: time}

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
	ratio int64
}
