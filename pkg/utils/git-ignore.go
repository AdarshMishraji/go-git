package utils

import (
	"bufio"
	"os"
)

func ParseGitIgnore() *Set {
	ignoreList := &Set{}
	ignoreList.Add(".git")

	file, err := os.Open(".gitignore")
	if err != nil {
		return ignoreList
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ignoreList = ignoreList.Add(scanner.Text())
	}

	return ignoreList
}
