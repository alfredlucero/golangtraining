package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Go programs need to spawn other non-Go processes

func main() {
	// exec.Command helper creates an object to represent this external process
	dateCmd := exec.Command("date")

	// Runs command, waits for it to finish and collects it output
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	// If no errors, dateOut will hold bytes with date info
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Pipe data to external process on its stdin and collect results from stdout
	grepCmd := exec.Command("grep", "hello")

	// Grab input/output pipes, start process, write some input to it, read resulting output, wait for process to exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// Collect StdoutPipe results to display
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// When spawning commands we need to provide an explicit delineated command and argument array
	// If you want to spawn a full command with a string, you can use bash's -c option
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

	// Spawned programs return output that is same as if we had run them directly from command-line
}
