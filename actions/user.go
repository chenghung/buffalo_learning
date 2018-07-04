package actions

import (
	"strconv"

	"github.com/chenghung/buffalo_learning/models"
	"github.com/gobuffalo/buffalo"
)

type ApiError struct {
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
}

// UserIndex default implementation.
func UserIndex(c buffalo.Context) error {
	users := models.Users{}
	models.DB.All(&users)

	return c.Render(200, r.JSON(users))
}

// UserShow default implementation.
func UserShow(c buffalo.Context) error {
	user := models.User{}
	err := models.DB.Find(&user, c.Param("id"))

	if err != nil {
		c.Logger().Debug(err)
		var message string
		var statusCode int

		if err.Error() == "mysql select one: sql: no rows in result set" {
			message = "user not found"
			statusCode = 404
		} else {
			message = err.Error()
			statusCode = 500
		}

		apiErr := ApiError{
			Message: message,
		}
		return c.Render(statusCode, r.JSON(apiErr))
	}

	return c.Render(200, r.JSON(user))
}

// UserCreate default implementation.
func UserCreate(c buffalo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := models.DB.Save(&user); err != nil {
		return err
	}

	return c.Render(200, r.JSON(user))
}

// UserUpdate default implementation.
func UserUpdate(c buffalo.Context) error {
	user := models.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := c.Bind(&user); err != nil {
		return err
	}

	user.ID = id
	c.Logger().Debug("user", user)

	if err := models.DB.Save(&user); err != nil {
		return err
	}

	return c.Render(200, r.JSON(user))
}

// UserDestroy default implementation.
func UserDestroy(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user := models.User{
		ID: id,
	}

	if err := models.DB.Destroy(&user); err != nil {
		return err
	}

	return c.Render(204, r.JSON(""))
}
