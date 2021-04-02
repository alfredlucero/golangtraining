package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Sometimes want to replace current Go process with another
	// exec ls -> Go requires an absolute path to binary we want to execute so we need exec.LookPath (probably in /bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form
	// First argument should be program name
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec needs set of environment variables
	env := os.Environ()

	// syscall.Exec; if successful, the execution of our process will end here and be replaced by /bin/ls -a -l -h process; if there is an error we'll get a return value
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	// Go doesn't offer a Unix fork function
	// goroutines, spawning processes, execing processes cover most use cases for a fork
}
