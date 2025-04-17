package controllers

import (
	"github.com/wiselike/revel-examples/orm/gorm/app/models"
	gorm "github.com/wiselike/revel-modules/orm/gorm/app"
	"github.com/wiselike/revel"
)

func initializeDB() {
	gorm.DB.AutoMigrate(&models.User{})
	firstUser := models.User{Name: "Demo", Email: "demo@demo.com"}
	firstUser.SetNewPassword("demo")
	firstUser.Active = true
	gorm.DB.Create(&firstUser)
}

func init() {
	revel.OnAppStart(initializeDB)
}
