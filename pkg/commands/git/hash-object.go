package git

import (
	"fmt"
	"go-git/pkg/commands/git/internal/blob"
	"go-git/pkg/commands/git/internal/hash"
	"go-git/pkg/utils"
	"os"
)

type HashObjectCommand struct {
	flag     string
	filePath string
}

func NewHashObjectCommand() *HashObjectCommand {
	flag := os.Args[2]
	var filePath string

	if len(os.Args) == 3 {
		filePath = flag
		flag = ""
	} else {
		filePath = os.Args[3]
	}

	validArgs := []string{"-w", ""}
	if !utils.SliceContains(validArgs, flag) {
		utils.ErrorLoggerF("go-git hash-object: invalid option %s\n", flag)
		os.Exit(1)
	}

	if len(filePath) == 0 {
		utils.ErrorLogger("go-git hash-object: missing file\n")
		os.Exit(1)
	}

	return &HashObjectCommand{
		flag:     flag,
		filePath: filePath,
	}
}

func (c *HashObjectCommand) Execute() {
	formattedContent := blob.CreateBlob(c.filePath)
	sha1Hash := hash.CreateHash(formattedContent)

	if c.flag == "-w" {
		blob.WriteBlob(formattedContent, sha1Hash, c.filePath)
	}

	fmt.Println(sha1Hash)
}
