package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/torreswoo/blog/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
