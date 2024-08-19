package blob

import (
	"compress/zlib"
	"go-git/pkg/constants"
	"go-git/pkg/utils"
	"io"
	"os"
	"path"
	"strings"
)

func ReadObject(sha1Hash string) string {
	pathString := path.Join(constants.GitRootDir, "objects", sha1Hash[:2], sha1Hash[2:])

	file, error := os.Open(pathString)
	if error != nil {
		utils.ErrorLoggerF("go-git cat-file: %s: no such file or directory\n", pathString)
		os.Exit(1)
	}
	defer file.Close()

	reader, error := zlib.NewReader(file)
	if error != nil {
		utils.ErrorLoggerF("go-git cat-file: %s: unable to read object\n", pathString)
		os.Exit(1)
	}
	defer reader.Close()

	content, error := io.ReadAll(reader)
	if error != nil {
		utils.ErrorLoggerF("go-git cat-file: %s: unable to read object\n", pathString)
		os.Exit(1)
	}

	rawContent := strings.TrimSpace(string(content))
	return rawContent
}
