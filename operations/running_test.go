package operations_test

import (
	"testing"

	"github.com/automateafrica/rcc/hamlet"
	"github.com/automateafrica/rcc/operations"
)

func TestTokenPeriodWorksAsExpected(t *testing.T) {
	must, wont := hamlet.Specifications(t)

	var period *operations.TokenPeriod
	must.Nil(period)
	wont.Panic(func() {
		period.Deadline()
	})
}
