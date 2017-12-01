package exectime

import (
	"fmt"
	"testing"
)

func TestExectime(t *testing.T) {

	fmt.Println("Running exectime tests.")
	res, err := exectime("./test/")
	if err != nil {
		t.Fail()
		t.Logf("Error thrown: " + err.Error())
	} else {
		t.Logf("Average Runtime Ratio: %d", res)
	}
}
