package sudoku

import (
  "fmt"
)

func (p *Puzzle) SwapC(x1,y1,x2,y2 int) *Puzzle {
  temp := p.cells[x1][y1].value
  p.cells[x1][y1].value = p.cells[x2][y2].value
  p.cells[x2][y2].value = temp
  return p
}

func (p *Puzzle) SwapH(h1,h2 int) *Puzzle {
  for i := 0; i < 9; i++ {
    p.SwapC(h1, i, h2, i) 
  }
  return p
}

func (p *Puzzle) SwapV(v1,v2 int) *Puzzle {
  for i := 0; i < 9; i++ {
    p.SwapC(i, v1, i, v2) 
  }
  return p
}

func (p *Puzzle) SSwapH(v1, v2 int) *Puzzle {
  if v1/3 != v2/3 {
    panic(fmt.Sprintf("SSwapH called with %d<->%d", v1, v2))
  }
  return p.SwapH(v1, v2)
}

func (p *Puzzle) SSwapV(h1, h2 int) *Puzzle {
  if h1/3 != h2/3 {
    panic(fmt.Sprintf("SSwapH called with %d<->%d", h1, h2))
  }
  return p.SwapV(h1, h2)
}
