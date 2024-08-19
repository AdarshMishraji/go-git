package blob

import (
	"bytes"
	"compress/zlib"
	"go-git/pkg/constants"
	"go-git/pkg/utils"
	"os"
	"path"
)

func WriteBlob(content string, sha1Hash string, filePath string) {
	var buffer bytes.Buffer
	compressedContentWriter := zlib.NewWriter(&buffer)
	compressedContentWriter.Write([]byte(content))
	compressedContentWriter.Close()

	os.MkdirAll(path.Join(constants.GitRootDir, "objects", sha1Hash[:2]), 0755) // 0755 is the permission (rwxr-xr-x)

	pathString := path.Join(constants.GitRootDir, "objects", sha1Hash[:2], sha1Hash[2:])
	error := os.WriteFile(pathString, buffer.Bytes(), 0644) // 0644 is the permission (rw-r--r--)
	if error != nil {
		utils.ErrorLoggerF("go-git hash-object: %s: unable to write object\n", filePath)
		os.Exit(1)
	}
}
