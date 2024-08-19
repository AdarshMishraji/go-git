package git

import (
	"fmt"
	"os"
)

type InitCommand struct {
}

func NewInitCommand() *InitCommand {
	return &InitCommand{}
}

func (c *InitCommand) Execute() {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		error := os.MkdirAll(dir, 0755) // 0755 is the permission (rwxr-xr-x)
		if error != nil {
			fmt.Fprintf(os.Stderr, "go-git init: %s: unable to create directory\n", dir)
			os.Exit(1)
		}
	}

	headFileContent := "ref: refs/heads/main\n"

	error := os.WriteFile(".git/HEAD", []byte(headFileContent), 0644) // 0644 is the permission (rw-r--r--)
	if error != nil {
		fmt.Fprintf(os.Stderr, "go-git init: %s: unable to create file\n", ".git/HEAD")
		os.Exit(1)
	}

	fmt.Println("Initialized empty go-git repository in .git/")
}
