package cmd

import (
	"os"

	"github.com/automateafrica/rcc/common"
	"github.com/automateafrica/rcc/pathlib"
	"github.com/automateafrica/rcc/pretty"
)

func osSpecificHolotreeSharing(enable bool) {
	if !enable {
		return
	}
	pathlib.ForceShared()
	err := os.WriteFile(common.SharedMarkerLocation(), []byte(common.Version), 0644)
	pretty.Guard(err == nil, 3, "Could not write %q, reason: %v", common.SharedMarkerLocation(), err)
}
