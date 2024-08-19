package git

import (
	"fmt"
	"go-git/pkg/commands/git/internal/blob"
	"go-git/pkg/commands/git/internal/hash"
	"go-git/pkg/constants"
	"go-git/pkg/utils"
	"os"
	"path"
	"strconv"
	"strings"
)

type WriteTreeCommand struct {
	ignoreList *utils.Set
}

func NewWriteTreeCommand() *WriteTreeCommand {
	return &WriteTreeCommand{
		ignoreList: utils.ParseGitIgnore(),
	}
}

func (c *WriteTreeCommand) recursiveWriteTree(basePath string) string {
	currentDirectorContents, error := os.ReadDir(basePath)
	if error != nil {
		utils.ErrorLoggerF("go-git write-tree: %s: no such file or directory\n", basePath)
		os.Exit(1)
	}

	if len(currentDirectorContents) == 0 {
		return ""
	}

	treeContents := []map[string]interface{}{}

	for _, content := range currentDirectorContents {
		currentPath := path.Join(basePath, content.Name())

		if content.Name() == ".git" || c.ignoreList.Has(content.Name()) {
			continue
		}

		if content.IsDir() {
			sha1Hash := c.recursiveWriteTree(currentPath)
			if sha1Hash != "" {
				treeContents = append(treeContents, map[string]interface{}{
					"mode": constants.TreeDirMode,
					"hash": sha1Hash,
					"path": content.Name(),
				})
			}
		} else {
			formattedContent := blob.CreateBlob(currentPath)
			sha1Hash := hash.CreateHash(formattedContent)
			blob.WriteBlob(formattedContent, sha1Hash, currentPath)
			treeContents = append(treeContents, map[string]interface{}{
				"mode": constants.TreeFileMode,
				"hash": sha1Hash,
				"path": content.Name(),
			})
		}
	}

	if len(treeContents) == 0 {
		return ""
	}

	var treeContent strings.Builder

	treeContent.WriteString("tree")
	treeContent.WriteString(strconv.Itoa(len(treeContents)))
	treeContent.WriteString("\x00")

	for _, content := range treeContents {
		treeContent.WriteString(content["mode"].(string))
		treeContent.WriteString(" ")
		treeContent.WriteString(content["path"].(string))
		treeContent.WriteString("\x00")
		treeContent.WriteString(content["hash"].(string))
	}

	formattedContent := treeContent.String()
	sha1Hash := hash.CreateHash(formattedContent)

	blob.WriteBlob(formattedContent, sha1Hash, basePath)

	return sha1Hash
}

func (c *WriteTreeCommand) Execute() {
	sha1Hash := c.recursiveWriteTree(".")
	fmt.Println(sha1Hash)
}
