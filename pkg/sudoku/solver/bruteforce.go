package solver

import (
  "fmt"

  "github.com/pkg/errors"

  "github.com/adarshmelethil/GoGo/pkg/sudoku"
)

func BruteForceFork(p *sudoku.Puzzle) *sudoku.Puzzle {
  ans := make(chan *sudoku.Puzzle)
  newP := p.Copy()

  go runBruteForceFork(newP, ans)

  return <-ans
}

func runBruteForceFork(p *sudoku.Puzzle, ans chan<- *sudoku.Puzzle) {
  if sudoku.Solved(p) {
    ans <- p
    return
  }

  if err := writeOnlyOneAns(p); err != nil {
    return
  }

  forkMultipleValues(p, ans)
  return
}

func forkMultipleValues(p *sudoku.Puzzle, ans chan<- *sudoku.Puzzle) {
  cell := p.RemainingCells()[0]

  for i, pv := range cell.PossibleValues() {
    if pv {
      val := i+1 
      newP := p.Copy()
      x, y := cell.Coor()
      newP.SetValue(x, y, val)
      go runBruteForceFork(newP, ans)
    }
  }
}
func writeOnlyOneAns(p *sudoku.Puzzle) error {
  repeat := false
  for repeat {
    repeat = false
    for _, cell := range p.RemainingCells() {
      numPossibilities := 0
      val := 0
      for i, pv := range cell.PossibleValues() {
        if pv {
          numPossibilities++
          val = i+1
        }
      }
      if numPossibilities == 0 {
        x, y := cell.Coor()
        return errors.New(fmt.Sprintf("Mistake num possible values is zero for %dx%d", x, y))
      }
      if numPossibilities == 1 {
        x, y := cell.Coor()
        p.SetValue(x, y, val)
        repeat = true
        break
      }
    }
  }
  return nil
}

