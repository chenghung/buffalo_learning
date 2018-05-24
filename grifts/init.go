package grifts

import (
	"github.com/buffalo_learning/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
