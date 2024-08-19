package git

import (
	"fmt"
	"go-git/pkg/constants"
	"go-git/pkg/utils"
	"os"
)

type InitCommand struct {
}

func NewInitCommand() *InitCommand {
	return &InitCommand{}
}

func (c *InitCommand) Execute() {
	if _, error := os.Stat(constants.GitRootDir); error == nil {
		utils.ErrorLoggerF("go-git init: %s: repository already exists\n", constants.GitRootDir)
		os.Exit(1)
	}

	dirs := []string{constants.GitRootDir, constants.GitRootDir + "/objects", constants.GitRootDir + "/refs"}

	for _, dir := range dirs {
		error := os.MkdirAll(dir, 0755) // 0755 is the permission (rwxr-xr-x)
		if error != nil {
			utils.ErrorLoggerF("go-git init: %s: unable to create directory\n", dir)
			os.Exit(1)
		}
	}

	headFileContent := "ref: refs/heads/main\n"

	error := os.WriteFile(constants.GitRootDir+"/HEAD", []byte(headFileContent), 0644) // 0644 is the permission (rw-r--r--)
	if error != nil {
		utils.ErrorLoggerF("go-git init: %s: unable to create file\n", ".git/HEAD")
		os.Exit(1)
	}

	fmt.Println("Initialized empty go-git repository in .git/")
}
