package grifts

import (
	"github.com/chenghung/buffalo_learning/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
