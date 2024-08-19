package main

import (
	"fmt"
	"go-git/pkg/commands"
	"go-git/pkg/commands/git"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: go-git <command> [<args>...]\n")
		os.Exit(1)
	}

	baseCommand := os.Args[1]

	switch baseCommand {
	case "init":
		commands.Execute(git.NewInitCommand())
	case "cat-file":
		commands.Execute(git.NewCatFileCommand())
	default:
		fmt.Fprintf(os.Stderr, "go-git: '%s' is not a go-git command. See 'go-git --help'.\n", baseCommand)
		os.Exit(1)
	}
}
