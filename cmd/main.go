package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	builtincommands "github.com/karaMuha/go-shell/builtin-commands"
)

func main() {
	builtincommands.InitCommandFunctions()

	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		// replace os specific delimiter to run on different operating systems
		input = strings.ReplaceAll(input, "\r\n", "\n")
		input = strings.Trim(input, "\n")

		if err != nil {
			fmt.Println("Error while reading input: ", err)
			os.Exit(1)
		}

		inputArray := strings.Split(input, " ")
		cmd := inputArray[0]
		args := inputArray[1:]

		cmdFn, ok := builtincommands.BuiltinCommands[cmd]

		if ok {
			cmdFn(args)
		} else {
			runProgramm(cmd, args)
		}
	}
}

func runProgramm(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmd)
	}
}
