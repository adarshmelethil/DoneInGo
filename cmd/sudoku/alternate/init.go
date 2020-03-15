package alternate

import (
	"github.com/spf13/cobra"
)

func AddSubCommands(parent *cobra.Command) {
  setSolveCmdFlags()

  parent.AddCommand(solveCmd)
}
