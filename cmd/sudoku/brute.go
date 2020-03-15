package sudoku

import (
  "fmt"
  "time"

  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"

  "github.com/adarshmelethil/GoGo/pkg/sudoku"
  "github.com/adarshmelethil/GoGo/pkg/sudoku/solver"
)

var (
  bruteCmd = &cobra.Command{
    Use: "brute",
    Short: "sudoku puzzle stuff",
    Run: runBruteCmd,
  }
)

func setBruteCmdFlags() {
  bruteCmd.Flags().StringP("puzzle", "p", "", "Sudoku file")
}

func runBruteCmd(cmd *cobra.Command, args []string) {
  filename := viper.GetString("puzzle")

  var puzzle *sudoku.Puzzle
  var err error
  if filename == "" {
    fmt.Println("random")
  } else {
    puzzle, err = sudoku.FromFile(filename)
    if err != nil {
      log.Fatal("Failed to read sudoku:", err)
    }
  }

  start := time.Now()
  solved := solver.BruteForceFork(puzzle)
  elapsed := time.Since(start)
  fmt.Println(solved)
  log.Debugf("Elapsed: %dhh %dmm %dss %dns", int(elapsed.Hours()), int(elapsed.Minutes()), int(elapsed.Seconds()), elapsed.Nanoseconds)
}
