package controllers

import (
	"github.com/wiselike/revel-examples/validation/app/models"
	"github.com/wiselike/revel"
)

type Sample4 struct {
	*revel.Controller
}

func (c Sample4) Index() revel.Result {
	return c.Render()
}

func (c Sample4) HandleSubmit(user *models.User) revel.Result {
	user.Validate(c.Validation)

	// Handle errors
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()

		return c.Redirect(Sample4.Index)
	}

	// Ok, display the created user
	return c.Render(user)
}
