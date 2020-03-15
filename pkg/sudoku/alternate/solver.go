package alternate

import (
  "sync"

  log "github.com/sirupsen/logrus"
)

func (p *Puzzle) Solve() {
  var waitgroup sync.WaitGroup

  log.Debug("Starting cells")
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      log.Infof("Starting %dx%d", i, j)
      waitgroup.Add(1)
      go p.cells[i][j].runCell(&waitgroup)
    }
  }

  log.Info("Sending first value")
  p.cells[0][0].outConns[0] <- uint8(8)

  waitgroup.Wait()
}

func (c *Cell) runCell(waitgroup *sync.WaitGroup) {
  for {
    for i, inConn := range c.inConns {
      // log.Debugf("%dx%d: %d", c.x, c.y, i)
      select {
        case val := <-inConn:
          log.Debugf("%dx%d: Recieved: %d -> %d\n", c.x, c.y, i, val)

          c.connVal[i] = val
          c.closeAll()
          c.broadcast(val)
          log.Debugf("%dx%d: Finished broadcasting", c.x, c.y)
          waitgroup.Done()
          return
      }
    }
  }
}
func (c *Cell) closeAll() {
  for _, inConn := range c.inConns {
    close(inConn)
  }
}
func (c *Cell) broadcast(val uint8) {
  for _, outConn := range c.outConns {
    sendToChan(outConn, val)
  }
}

func sendToChan(c chan<-uint8, v uint8) {
  defer func() {
    recover()
  }()

  c <- v
}

