package cmd

import (
	"github.com/automateafrica/rcc/common"
	"github.com/automateafrica/rcc/conda"
	"github.com/automateafrica/rcc/pretty"
	"github.com/spf13/cobra"
)

var (
	enableLongpaths bool
)

var longpathsCmd = &cobra.Command{
	Use:   "longpaths",
	Short: "Check and enable Windows longpath support",
	Long:  "Check and enable Windows longpath support",
	Run: func(cmd *cobra.Command, args []string) {
		if common.OverrideSystemRequirements() {
			pretty.Exit(100, "This operation is prevented, because ROBOCORP_OVERRIDE_SYSTEM_REQUIREMENTS is effective!")
		}
		var err error
		if enableLongpaths {
			err = conda.EnforceLongpathSupport()
		}
		if err != nil {
			pretty.Exit(1, "Failure to modify registry: %v", err)
		}
		if !conda.HasLongPathSupport() {
			pretty.Exit(2, "Long paths do not work!")
		}
		pretty.Ok()
	},
}

func init() {
	configureCmd.AddCommand(longpathsCmd)
	longpathsCmd.Flags().BoolVarP(&enableLongpaths, "enable", "e", false, "Change registry settings and enable longpath support")
}
