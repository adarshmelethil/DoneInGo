package sudoku

import (
  "github.com/spf13/cobra"

  "github.com/adarshmelethil/GoGo/cmd/sudoku/alternate"
)

var (
  alternateCmd = &cobra.Command{
    Use: "alternate",
    Aliases: []string{"a"},
    Short: "sudoku puzzle with alternate",
    Run: DelegateSubcommands,
  }
)

func setAlternateCmdFlags() {
}

func init() {

  alternate.AddSubCommands(alternateCmd)
}
