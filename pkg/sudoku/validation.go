package sudoku

import ()

func Solved(p *Puzzle) bool {
  if len(p.RemainingCells()) > 0 {
    return false
  }
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      if p.cells[i][j].value == 0 {
        return false
      }
    }
  }
  return Validate(p)
}

func Validate(p *Puzzle) bool {
  for i := 0; i < 9; i++ {
    if !CheckHorizontal(p, i) || !CheckVertical(p, i) || !CheckZone(p, i) {
      return false
    }
  }
  return true
}

func CheckHorizontal(p *Puzzle, row int) bool {
  for i := 0; i < 8; i++ {
    for j := i+1; j < 9; j++ {
      if p.cells[row][i].value != 0 && p.cells[row][i].value == p.cells[row][j].value {
        return false
      }
    }
  }
  return true
}
func CheckVertical(p *Puzzle, col int) bool {
  for i := 0; i < 8; i++ {
    for j := i+1; j < 9; j++ {
      if p.cells[i][col].value != 0 && p.cells[i][col].value == p.cells[j][col].value {
        return false
      }
    }
  }
  return true
}
func CheckZone(p *Puzzle, zone int) bool {
  x := ((zone%3)*3)
  y := ((zone/3)*3)
  for i := 0; i < 8; i++ {
    for j := i+1; j < 9; j++ {
      x1 := (i%3)
      y1 := (i/3)
      x2 := (j%3)
      y2 := (j/3)

      if p.cells[x+x1][y+y1].value != 0 && p.cells[x+x1][y+y1].value == p.cells[x+x2][y+y2].value {
        return false
      }
    }
  }

  return true
}
