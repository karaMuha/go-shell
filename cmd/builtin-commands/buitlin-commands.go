package builtincommands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var BuiltinCommands = make(map[string]cmdFn)

type cmdFn func(args []string)

func InitCommandFunctions() {
	BuiltinCommands["echo"] = EchoFn
	BuiltinCommands["exit"] = ExitFn
	BuiltinCommands["type"] = TypeFn
	BuiltinCommands["pwd"] = PwdFn
	BuiltinCommands["cd"] = CdFn
}

func EchoFn(args []string) {
	output := strings.Join(args, " ")
	fmt.Println(output)
}

func ExitFn(args []string) {
	if len(args) != 1 {
		fmt.Printf("Invalid amount of arguments for exit command: expected 1 but got %d\n", len(args))
		os.Exit(1)
	}

	exitCode, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf("Exit command has invalid argument: %v", err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}

func TypeFn(args []string) {
	if len(args) == 0 {
		fmt.Println("Need at least one argument to execute command")
		return
	}

	if _, ok := BuiltinCommands[args[0]]; ok {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}

	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		fp := filepath.Join(path, args[0])
		if _, err := os.Stat(fp); err == nil {
			fmt.Println(fp)
			return
		}
	}

	fmt.Printf("%s: not found\n", args[0])
}

func PwdFn(args []string) {
	if len(args) != 0 {
		fmt.Printf("Expected 0 arguments but got %d\n", len(args))
		return
	}

	exe, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(exe)
}

func CdFn(args []string) {
	if len(args) == 0 {
		fmt.Println("Need at least one argument to execute command")
		return
	}

	dir := args[0]

	if dir == "~" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			fmt.Println(err)
			return
		}

		dir = homeDir
	}

	err := os.Chdir(dir)

	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", args[0])
	}
}
