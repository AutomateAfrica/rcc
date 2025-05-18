package conda_test

import (
	"testing"

	"github.com/automateafrica/rcc/conda"
	"github.com/automateafrica/rcc/hamlet"
)

func TestHasDownloadLinkAvailable(t *testing.T) {
	must_be, _ := hamlet.Specifications(t)

	must_be.True(len(conda.MicromambaLink()) > 10)
}
