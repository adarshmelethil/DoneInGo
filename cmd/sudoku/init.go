package sudoku

import (
  "github.com/spf13/cobra"
)

func AddSubCommands(parent *cobra.Command) {
  setBruteCmdFlags()
  setAlternateCmdFlags()

  parent.AddCommand(bruteCmd)
  parent.AddCommand(alternateCmd)
}
