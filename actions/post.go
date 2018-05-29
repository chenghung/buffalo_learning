package actions

import (
	"strconv"

	"github.com/buffalo_learning/models"
	"github.com/gobuffalo/buffalo"
)

// POSTIndex default implementation.
func PostIndex(c buffalo.Context) error {
	page, _ := strconv.Atoi(c.Param("page"))
	size, _ := strconv.Atoi(c.Param("size"))
	posts := models.Posts{}
	tx := models.DB
	c.Logger().Debug("find posts by user_id=", c.Param("user_id"))
	c.Logger().Debug("page=", page, ", size=", size)
	tx.
		Where("user_id = ?", c.Param("user_id")).
		Paginate(page, size).
		All(&posts)

	return c.Render(200, r.JSON(posts))
}
