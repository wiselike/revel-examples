package tests

import (
	"github.com/wiselike/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}
