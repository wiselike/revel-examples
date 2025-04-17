package jobs

import (
	"fmt"

	"github.com/wiselike/revel-examples/booking/app/models"
	"github.com/wiselike/revel-modules/jobs/app/jobs"
	gorp "github.com/wiselike/revel-modules/orm/gorp/app"
	"github.com/wiselike/revel"
)

// Periodically count the bookings in the database.
type BookingCounter struct{}

func (c BookingCounter) Run() {
	bookings, err := gorp.Db.Map.Select(models.Booking{},
		`select * from Booking`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d bookings.\n", len(bookings))
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1m", BookingCounter{})
	})
}
