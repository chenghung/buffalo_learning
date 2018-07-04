package actions

import (
	"github.com/chenghung/buffalo_learning/models"
	"github.com/gobuffalo/buffalo"
)

// POSTIndex default implementation.
func PostIndex(c buffalo.Context) error {
	posts := models.Posts{}
	tx := models.DB
	c.Logger().Debug("find posts by user_id=", c.Param("user_id"))
	tx.
		Where("user_id = ?", c.Param("user_id")).
		PaginateFromParams(c.Params()).
		All(&posts)

	return c.Render(200, r.JSON(posts))
}
