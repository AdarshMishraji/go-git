package git

import (
	"fmt"
	"go-git/pkg/commands/git/internal/blob"
	"go-git/pkg/utils"
	"os"
)

type LsTreeCommand struct {
	flag string
	hash string
}

func NewLsTreeCommand() *LsTreeCommand {
	flag := os.Args[2]
	var hash string

	if len(os.Args) == 3 {
		hash = flag
		flag = ""
	} else {
		hash = os.Args[3]
	}

	validArgs := []string{"-u", ""}
	if !utils.SliceContains(validArgs, flag) {
		utils.ErrorLoggerF("go-git read-tree: invalid option %s\n", flag)
		os.Exit(1)
	}

	if len(hash) == 0 {
		utils.ErrorLoggerF("go-git read-tree: missing tree\n")
		os.Exit(1)
	}

	return &LsTreeCommand{
		flag: flag,
		hash: hash,
	}
}

func (c *LsTreeCommand) Execute() {
	// _, _, content := blob.ReadTree(c.hash, c.flag == "-name-only")
	_, _, content := blob.ReadTree(c.hash, true)

	fmt.Println(content)
}
