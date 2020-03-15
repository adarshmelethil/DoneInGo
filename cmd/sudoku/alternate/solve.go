package alternate

import (
  "fmt"
  // "time"

  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"

  alt "github.com/adarshmelethil/GoGo/pkg/sudoku/alternate"
)

var (
  solveCmd = &cobra.Command{
    Use: "solve",
    Aliases: []string{"s"},
    Short: "sudoku puzzle alternate solved",
    Run: runSolveCmd,
  }
)

func setSolveCmdFlags() {
  solveCmd.Flags().StringP("puzzle", "p", "", "Sudoku file")
}

func runSolveCmd(cmd *cobra.Command, args []string) {
  filename := viper.GetString("puzzle")

  var puzzle *alt.Puzzle
  var err error
  if filename == "" {
    fmt.Println("random")
  } else {
    puzzle, err = alt.FromFile(filename)
    if err != nil {
      log.Fatal("Failed to read sudoku:", err)
    }
  }

  fmt.Println(puzzle)
  puzzle.Solve()
  fmt.Println(puzzle)
  // start := time.Now()
  // solved := solver.BruteForceFork(puzzle)
  // elapsed := time.Since(start)
  // fmt.Println(solved)
  // log.Debugf("Elapsed: %dhh %dmm %dss %dns", int(elapsed.Hours()), int(elapsed.Minutes()), int(elapsed.Seconds()), elapsed.Nanoseconds)
}
