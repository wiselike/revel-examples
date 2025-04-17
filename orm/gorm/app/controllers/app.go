package controllers

import (
	"github.com/wiselike/revel-examples/orm/gorm/app/models"
	gormc "github.com/wiselike/revel-modules/orm/gorm/app/controllers"
	"github.com/wiselike/revel"
)

type App struct {
	gormc.TxnController
}

func (c App) Index() revel.Result {
	users := []models.User{}
	c.Txn.Find(&users)
	return c.RenderJSON(users)
}
