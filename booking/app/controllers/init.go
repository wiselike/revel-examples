package controllers

import "github.com/wiselike/revel"

func init() {
	revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
}
