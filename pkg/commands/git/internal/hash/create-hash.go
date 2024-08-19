package hash

import (
	"crypto/sha1"
	"fmt"
)

func CreateHash(content string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(content)))
}
