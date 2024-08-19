package blob

import (
	"fmt"
	"go-git/pkg/utils"
	"io"
	"os"
)

func CreateBlob(filePath string) string {
	file, error := os.Open(filePath)
	if error != nil {
		utils.ErrorLoggerF("go-git hash-object: %s: no such file or directory\n", filePath)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, error := file.Stat()
	if error != nil {
		utils.ErrorLoggerF("go-git hash-object: %s: unable to read file\n", filePath)
		os.Exit(1)
	}

	size := fileInfo.Size()
	content, error := io.ReadAll(file)
	if error != nil {
		utils.ErrorLoggerF("go-git hash-object: %s: unable to read file\n", filePath)
		os.Exit(1)
	}

	formattedContent := fmt.Sprintf("blob %d\x00%s", size, content)
	return formattedContent
}
