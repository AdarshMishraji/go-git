package git

import (
	"fmt"
	"go-git/pkg/commands/git/internal/blob"
	"go-git/pkg/utils"
	"os"
)

type CatFileCommand struct {
	flag string
	hash string
}

func NewCatFileCommand() *CatFileCommand {
	flag := os.Args[2]
	hash := os.Args[3]

	validArgs := []string{"-p", "-t", "-s"}
	if !utils.SliceContains(validArgs, flag) {
		utils.ErrorLoggerF("go-git cat-file: invalid option %s\n", flag)
		os.Exit(1)
	}

	if len(hash) == 0 {
		utils.ErrorLoggerF("go-git cat-file: missing object\n")
		os.Exit(1)
	}

	return &CatFileCommand{
		flag: flag,
		hash: hash,
	}
}

func (c *CatFileCommand) Execute() {
	hashType, contentLength, contentString := blob.ReadBlob(c.hash)

	switch c.flag {
	case "-p":
		fmt.Println(contentString)
	case "-t":
		fmt.Println(hashType)
	case "-s":
		fmt.Println(contentLength)
	default:
		fmt.Println("Invalid flag")
	}
}
