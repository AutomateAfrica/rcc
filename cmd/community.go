package cmd

import (
	"fmt"

	"github.com/automateafrica/rcc/common"
	"github.com/spf13/cobra"
)

var communityCmd = &cobra.Command{
	Use:     "community",
	Aliases: []string{"co"},
	Short:   fmt.Sprintf("Group of commands related to `%s Community`.", common.Product.Name()),
	Long:    `This group of commands apply to community provided robots and services.`,
}

func init() {
	if common.Product.IsLegacy() {
		rootCmd.AddCommand(communityCmd)
	}
}
