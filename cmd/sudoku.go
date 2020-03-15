package cmd

import (
  "github.com/spf13/cobra"

  "github.com/adarshmelethil/GoGo/cmd/sudoku"
)

var sudokuCmd = &cobra.Command{
  Use: "sudoku",
  Short: "sudoku puzzle stuff",
  Run: DelegateSubcommands,
}

func init() {
  rootCmd.AddCommand(sudokuCmd)
  sudoku.AddSubCommands(sudokuCmd)
}
