package controllers_test

import (
	"testing"

	"github.com/wiselike/revel-examples/booking/app/tmp/run"
	"github.com/wiselike/revel-modules/server-engine/gohttptest/testsuite"
)

//  go test -coverprofile=coverage.out github.com/wiselike/revel-examples/booking/app/controllers/  -args -revel.importPath=github.com/wiselike/revel-examples/booking
func TestMain(m *testing.M) {
	testsuite.RevelTestHelper(m, "dev", run.Run)
}

func TestIndex(t *testing.T) {
	tester := testsuite.NewTestSuite(t)
	tester.Get("/").AssertOk()
}
