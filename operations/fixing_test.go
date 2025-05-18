package operations_test

import (
	"testing"

	"github.com/automateafrica/rcc/hamlet"
	"github.com/automateafrica/rcc/operations"
)

func TestCanRemoveWindowsNewlines(t *testing.T) {
	must, _ := hamlet.Specifications(t)

	must.Equal("Hello", string(operations.ToUnix([]byte("Hello"))))
	must.Equal("A\nB\n", string(operations.ToUnix([]byte("A\nB\n"))))
	must.Equal([]byte("A\n\nB\n"), operations.ToUnix([]byte("A\r\n\r\nB\r\n")))
}
