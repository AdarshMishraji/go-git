package blob

import (
	"strings"
)

func ReadTree(hash string, nameOnly bool) (string, string, string) {
	rawContent := ReadObject(hash)
	firstNullChar := strings.Index(rawContent, "\x00")
	if firstNullChar == -1 {
		return "", "", ""
	}

	var content strings.Builder
	if nameOnly {
		splitContent := strings.Split(rawContent[firstNullChar:], "\x00")
		for _, line := range splitContent {
			spaceSplit := strings.Split(line, " ")
			if len(spaceSplit) > 1 {
				content.WriteString(spaceSplit[1])
				content.WriteString("\n")
			}
		}
	} else {
		content.WriteString(rawContent[firstNullChar:])
	}

	header := rawContent[:firstNullChar]
	hashType := strings.Split(header, " ")[0]
	contentLength := strings.Split(header, " ")[1]

	return hashType, contentLength, content.String()
}
