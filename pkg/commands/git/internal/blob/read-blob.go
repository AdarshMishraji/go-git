package blob

import "strings"

func ReadBlob(hash string) (string, string, string) {
	rawContent := ReadObject(hash)
	headerAndContent := strings.Split(rawContent, "\x00")
	header := headerAndContent[0]
	hashType := strings.Split(header, " ")[0]
	contentLength := strings.Split(header, " ")[1]

	return hashType, contentLength, headerAndContent[1]
}
