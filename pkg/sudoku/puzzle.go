package sudoku

import (
  "bytes"
  "strings"
  "strconv"
  "fmt"
  "os"
  "io/ioutil"

  "github.com/pkg/errors"
  log "github.com/sirupsen/logrus"
  "github.com/fatih/color"
)

type Cell struct {
  possibleValues [9]bool
  value uint8
  connections [20]*Cell
  x,y int
}

type Puzzle struct {
  cells [9][9]Cell
  cellsLeft []*Cell
}

func (p *Puzzle) Copy() *Puzzle {
  newP := NewPuzzle()

  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      newP.SetValue(i,j, int(p.cells[i][j].value))
    }
  }
  return newP
}

func (p *Puzzle) removeFromCellsLeft(cell *Cell) {
  for i, c := range p.cellsLeft {
    if c == cell {
      p.cellsLeft[i] = p.cellsLeft[len(p.cellsLeft)-1]
      p.cellsLeft = p.cellsLeft[:len(p.cellsLeft)-1]
      return
    }
  }
}

func (p *Puzzle) SetValue(x, y, value int) error {
  if p.cells[x][y].value != 0 {
    return errors.New("Setting value for cell with value")
  }
  if value > 9 || value <= 0 {
    return errors.New("Invalid value")
  }
  if !p.cells[x][y].possibleValues[value-1] {
    return errors.New("Value would make sudoku invalid")
  }

  p.cells[x][y].value = uint8(value)
  p.removeFromCellsLeft(&p.cells[x][y])
  for _, cell := range p.cells[x][y].connections {
    cell.possibleValues[value-1] = false
  }

  return nil
}

func (p Puzzle) RemainingCells() []*Cell {
  return p.cellsLeft
}
func (c Cell) PossibleValues() [9]bool {
  return c.possibleValues
}

func (p Puzzle) String() string {
  var buf bytes.Buffer
  buf.Write([]byte(strings.Repeat("*", 37)))
  buf.Write([]byte("\n"))

  for i := 0; i < 9; i++ {
    var cells [][]string
    for j := 0; j < 9; j++ {
      cells = append(cells, strings.Split(p.cells[i][j].String(), "\n"))
    }

    for j := 0; j < 3; j++ {
      line := "‖"
      for k := 0; k < 9; k++ {
        line += cells[k][j]
        if (k+1)%3 == 0 {
          line += "‖"
        } else {
          line += "|"
        }
      }
      line += "\n"
      buf.Write([]byte(line))
    }
    if (i+1)%3 == 0 {
      buf.Write([]byte(strings.Repeat("*", 37)))
    } else {
      buf.Write([]byte(strings.Repeat("-", 37)))
    }
    if i != 8 {
      buf.Write([]byte("\n"))  
    }
  }

  return buf.String()
}

func (c Cell) String() string {
  pvColor := color.New(color.FgCyan).SprintFunc()
  valColor := color.New(color.FgRed).SprintFunc()
  emptyColor := color.New(color.FgWhite).SprintFunc()

  var buf bytes.Buffer
  if c.value == 0 {
    for i := 0; i < 9; i++ {
      if c.possibleValues[i] {
        buf.Write([]byte(pvColor(strconv.Itoa(i+1))))
      } else {
        buf.Write([]byte(emptyColor(" ")))
      }
      if (i+1)%3 == 0 && i < 8 {
        buf.Write([]byte("\n"))
      }
    }
  } else {
    buf.Write([]byte(emptyColor("   \n")))
    buf.Write([]byte(fmt.Sprintf("%s%s%s\n", emptyColor(" "), valColor(c.value), emptyColor(" "))))
    buf.Write([]byte(emptyColor("   ")))
  }
  

  return buf.String()
}

func (c Cell) Coor() (int, int) {
  return c.x, c.y
}

func (p *Puzzle) PrintCell(x, y int) {
  fmt.Println(p.cells[x][y])
}

func (p *Puzzle) PrintVals() {
  fmt.Println(strings.Repeat("*", 13))
  for i := 0; i < 9; i++ {
    fmt.Print("|")
    for j := 0; j < 9; j++ {
      fmt.Print(p.cells[i][j].value)

      if (j+1)%3 == 0{
        fmt.Print("|")
      }
    }
    fmt.Println()
    if (i+1)%3 == 0 {
      fmt.Println(strings.Repeat("*", 13))
    }
  }
}
func (p *Puzzle) PrintConnections(x, y int) {
  c := color.New(color.FgCyan)
  r := color.New(color.FgRed)
  w := color.New(color.FgWhite)
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      wrote := false
      for _, a := range p.cells[x][y].connections {
        if a.x == i && a.y == j {
          wrote = true
          c.Print("x")
        }
      }
      if x == i && y == j {
        r.Print("O")
        wrote = true
      }

      if !wrote {
        w.Print(" ")
      }

      if (j+1)%3 == 0 {
        fmt.Print("|")
      }
    }
    fmt.Println()
    if (i+1)%3 == 0 {
      fmt.Println(strings.Repeat("-", 12))
    }
  }
}


func (p *Puzzle) appendToConnections(x, y int, cell *Cell) {
  if cell.x == x && cell.y == y {
    return
  }

  a := 0
  for _, c := range p.cells[x][y].connections {
    if c == nil {
      break
    }
    if (c.x == cell.x && c.y == cell.y) {
      return
    }
    a += 1
  }
  if a >= 20 {
    panic("Trying to add more than 20 connections")
  }
  p.cells[x][y].connections[a] = cell

  return
}

func (p *Puzzle) connectCells(x, y int) {
  // zone
  zonex := (x/3)*3
  zoney := (y/3)*3
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if zonex+i != x && zoney+j != y {
        p.appendToConnections(x, y, &p.cells[zonex+i][zoney+j])
      }
    }
  }

  // Horizontal
  for i := 0; i < 9; i++ {
    p.appendToConnections(x, y, &p.cells[x][i])
  }

  // Vertical
  for i := 0; i < 9; i++ {
    p.appendToConnections(x, y, &p.cells[i][y])
  }

  return
}

func NewCell(x, y int) Cell {
  c := Cell{
    x: x,
    y: y,
  }

  for k := 0; k < 9; k++ {
    c.possibleValues[k] = true
  }

  return c
}

func NewPuzzle() *Puzzle {
  p := Puzzle{}

  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      p.cells[i][j] = NewCell(i, j)
      p.cellsLeft = append(p.cellsLeft, &p.cells[i][j])
    }
  }
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      p.connectCells(i, j)
      for k, a := range p.cells[i][j].connections {
        if a == nil {
          log.Errorf("%dx%d - %d \n", i, j, k)
          panic("nil connection")
        }
      }
    }
  }

  return &p
}

func FromFile(filename string) (*Puzzle, error) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, err
  }
  sudokuBytes, err := ioutil.ReadAll(file)
  if err != nil {
    return nil, err
  }
  file.Close()

  p := NewPuzzle()
  sudokuLines := strings.Split(string(sudokuBytes), "\n")
  for i, line := range sudokuLines {
    nums := strings.Split(line, "")
    for j, num := range nums {
      if num != "_" {
        n, err := strconv.Atoi(num)
        if err != nil {
          return nil, errors.New(fmt.Sprintf("Failed to convert number for %dx%d, line: %s", i, j, num))
        }

        err = p.SetValue(i, j, n)
        if err != nil {
          return nil, err
        }
      }
    }
  }

  return p, nil
}
