package cmd

// import (
//   "fmt"
//   "math/rand"

//   "github.com/adarshmelethil/GoGo/pkg/sudoku"
// )

// func runQuick() {
//   puzzle, _ := sudoku.NewPuzzle()
//   fmt.Print(puzzle)
//   fmt.Println(sudoku.Validate(puzzle))

//   type move struct {
//     move int
//     p1, p2 int
//   }
//   var moves [100]move
//   for i := 0; i < 100; i++ {
//     curMove := rand.Intn(2)
//     p1 := rand.Intn(9)
//     p2 := rand.Intn(9)

//     moves[i] = move{
//       move: curMove,
//       p1: p1,
//       p2: p2,
//     }

//     if curMove == 0 {
//       puzzle.SSwapH(p1, p2)
//     } else {
//       puzzle.SSwapV(p1, p2)
//     }
//     if !sudoku.Validate(puzzle) {
//       fmt.Println(i, curMove, p1, p2, "Invalid")
//       break
//     }
//   }

//   fmt.Print(puzzle)

//   // fmt.Println(puzzle.SwapHorizontal(0,1))
//   // fmt.Println(sudoku.Validate(puzzle))

//   // for i := 0; i < 9; i++ {
//   //   sudoku.CheckZone(puzzle, i)
//   // }
//   // sudoku.CheckZone(puzzle, 0)


//   // var mailboxes [20*9*9]chan int
//   // for i, _ := range mailboxes {
//   //   mailboxes[i] = make(chan int)
//   // }
//   // fmt.Println(len(mailboxes))
  
//   fmt.Println("End")
// }


// func cell(id, mynum int, home chan<- int, mailbox []<-chan int, connections []chan<- int) {
//   // fixed := false
//   // if mynum != 0 {
//   //   fixed = true
//   // }


// }
