package tests

import (
	"fmt"
	"testing"
)

// go mod init tests
// go test -v

// testing for unit tests
// go test command to run tests
// testing code typically lives in the same package as the code it tests

// Usually would be in a file like intutils.go and test file in intutils_test.go
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test is created by writing a function with a name beginning with Test
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* will report test failures but continue executing the test
		// t.Fatal* will report test failures and stop the test immediately
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Idiomatic to use table-driven style unit tests where test inputs and expected outputs are listed in a table and a single loop walks over them and performs the test logic
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// Enables running "subtests", one for each table entry
		// Shown separately when executing go test -v
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
