package internal

import "go-git/pkg/interfaces"

func Execute(c interfaces.Command) {
	c.Execute()
}
