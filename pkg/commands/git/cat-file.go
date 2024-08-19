package git

import (
	"compress/zlib"
	"fmt"
	"go-git/pkg/utils"
	"io"
	"os"
	"path"
)

type CatFileCommand struct {
	flag string
	hash string
}

var validArgs = []string{"-p", "-t", "-s", ""}

func NewCatFileCommand() *CatFileCommand {
	flag := os.Args[3]
	hash := os.Args[4]

	if !utils.ArrContains(validArgs, flag) {
		fmt.Fprintf(os.Stderr, "go-git cat-file: invalid option %s\n", flag)
		os.Exit(1)
	}

	if len(hash) == 0 {
		fmt.Fprintf(os.Stderr, "go-git cat-file: missing object\n")
		os.Exit(1)
	}

	return &CatFileCommand{
		flag: flag,
		hash: hash,
	}
}

func (c *CatFileCommand) Execute() {
	pathString := path.Join(".git/objects", c.hash[:2], c.hash[2:])

	file, error := os.Open(pathString)
	if error != nil {
		fmt.Fprintf(os.Stderr, "go-git cat-file: %s: no such file or directory\n", c.hash)
		os.Exit(1)
	}
	defer file.Close()

	reader, error := zlib.NewReader(file)
	if error != nil {
		fmt.Fprintf(os.Stderr, "go-git cat-file: %s: unable to read object\n", c.hash)
		os.Exit(1)
	}
	defer reader.Close()

	content, error := io.ReadAll(reader)
	if error != nil {
		fmt.Fprintf(os.Stderr, "go-git cat-file: %s: unable to read object\n", c.hash)
		os.Exit(1)
	}

	switch c.flag {
	case "-p":
		fmt.Println(string(content))
	case "-t":
		fmt.Println("blob")
	case "-s":
		fmt.Println(len(content))
	default:
		fmt.Println(string(content))
	}
}
