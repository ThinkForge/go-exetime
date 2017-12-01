package exectime

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func exectime(folder string) (int64, error) {

	// Parse the filepath into slices and then get the files.
	arr := strings.Split(folder, "/")
	pathCut, errCut := exec.LookPath(folder + arr[len(arr)-2] + ".cut.out")
	pathEx, errEx := exec.LookPath(folder + arr[len(arr)-2] + ".ex.out")

	// If either of the files could not be found, return the error.
	if errCut != nil {
		return 0, errCut
	} else if errEx != nil {
		return 0, errEx
	}

	var ratios [5]int64
	// For each of the input files we evaluate it using both executables.
	for i := 0; i < 5; i++ {
		cmdCut := exec.Command(pathCut, folder+"input/"+strconv.Itoa(i)+".txt")
		startCut := start()
		cmdCut.Run()
		cutTime := end(startCut)

		cmdEx := exec.Command(pathEx, folder+"input/"+strconv.Itoa(i)+".txt")
		startEx := start()
		cmdEx.Run()
		exTime := end(startEx)
		ratios[i] = (cutTime.Nanoseconds() / exTime.Nanoseconds())
		fmt.Printf("Ex: %d CUT: %d\n", exTime.Nanoseconds(), cutTime.Nanoseconds())
	}

	var total int64 = 0
	// Get the average ratio
	for j := 0; j < 5; j++ {
		total += ratios[j]
	}

	fmt.Printf("Total: %d Average: %d\n", total, total/5)

	return total / 5, nil
}

func start() time.Time {
	return time.Now()
}

func end(startTime time.Time) time.Duration {
	return time.Since(startTime)
}
