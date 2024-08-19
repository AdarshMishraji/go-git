package commands

import (
	"go-git/pkg/commands/git"
	commandUtil "go-git/pkg/commands/internal"
	"go-git/pkg/utils"
	"os"
)

func ExecuteCommand() {
	if len(os.Args) < 2 {
		utils.ErrorLogger("usage: go-git <command> [<args>...]")
		os.Exit(1)
	}

	baseCommand := os.Args[1]

	switch baseCommand {
	case "init":
		commandUtil.Execute(git.NewInitCommand())
	case "cat-file":
		commandUtil.Execute(git.NewCatFileCommand())
	case "hash-object":
		commandUtil.Execute(git.NewHashObjectCommand())
	case "ls-tree":
		commandUtil.Execute(git.NewLsTreeCommand())
	case "write-tree":
		commandUtil.Execute(git.NewWriteTreeCommand())
	default:
		utils.ErrorLoggerF("go-git: '%s' is not a go-git command. See 'go-git --help'.\n", baseCommand)
		os.Exit(1)
	}
}
