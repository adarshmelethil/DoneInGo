package alternate

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
  "strconv"
  "bytes"

  "github.com/pkg/errors"
  // log "github.com/sirupsen/logrus"
)

type Puzzle struct {
  cells [9][9]Cell
}

type Cell struct {
  value uint8
  x, y int

  possibleValues [9]bool

  inConns []chan uint8
  outConns []chan<- uint8

  connVal []uint8
}

func NewCell(x, y int) Cell {
  c := Cell{
    x: x,
    y: y,
    value: 0,
  }
  for i := 0; i < len(c.possibleValues); i++ {
    c.possibleValues[i] = false
  }

  return c
}

func NewPuzzle() *Puzzle {
  p := Puzzle{}

  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      p.cells[i][j] = NewCell(i, j)
    }
  }

  // Connections
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      var conns [20]chan uint8
      for i,_ := range conns {
        conns[i] = make(chan uint8, 20)
      }
      p.setupConnections(i, j, conns[:])
    }
  }

  return &p
}

func (c *Cell) appendIncomingConnection(incoming chan uint8) {
  if len(c.inConns) >= 20 {
    panic("Adding too many incoming connections")
  }

  c.inConns = append(c.inConns, incoming)
  c.connVal = append(c.connVal, 0)
}

func (c *Cell) appendOutgoingConnection(outgoing chan<- uint8) {
  if len(c.inConns) >= 20 {
    panic("Adding too many outgoing connections")
  }

  c.outConns = append(c.outConns, outgoing)
}

func (p *Puzzle) setupConnections(x,y int, conns []chan uint8) {
  for i := 0; i < len(conns); i++ {
    p.cells[x][y].appendOutgoingConnection(conns[i])
  }

  index := 0
  // zone
  zonex := (x/3)*3
  zoney := (y/3)*3
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if !(zonex+i == x && zoney+j == y) {
        p.cells[x][y].appendIncomingConnection(conns[index])
        index++
      }
    }
  }

  // Horizontal
  for i := 0; i < 9; i++ {
    if !(i >= zoney && i <= zoney+2) {
      p.cells[x][i].appendIncomingConnection(conns[index])
      index++
    }
  }

  // Vertical
  for i := 0; i < 9; i++ {
    if !(i >= zonex && i <= zonex+2) {
      p.cells[i][y].appendIncomingConnection(conns[index])
      index++
    }
  }
}

func (p *Puzzle) SetValue(x, y int, val uint8) {
  p.cells[x][y].value = val  
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

        p.SetValue(i, j, uint8(n))
      }
    }
  }

  return p, nil
}

func (p *Puzzle) String() string {
  var buf bytes.Buffer

  for i, row := range p.cells {
    for j, cell := range row {
      buf.Write([]byte(fmt.Sprintf("%d", cell.value)))
      if (j+1)%3==0 {
        buf.Write([]byte(" "))
      }
    }

    if i != 8 {
      buf.Write([]byte("\n"))
      if (i+1)%3 == 0 {
        buf.Write([]byte("\n"))
      }
    }
  }

  return buf.String()
}

