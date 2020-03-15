package cmd

import (
  "os"
  "fmt"
  "strings"
  
  "github.com/spf13/cobra"
)

func DelegateSubcommands(cmd *cobra.Command, args []string) {
  EnsureCommand(cmd, args)
  cmd.Help()
  os.Exit(0)
}

func EnsureCommand(cmd *cobra.Command, args []string) {
  if len(args) > 0 {
    fmt.Println("Unknown command", strings.Join(args, " "))
  }
}

