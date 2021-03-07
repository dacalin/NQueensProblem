package main

import (
    "fmt"
    "time"
    //"sync"
)

var numSolutions uint = 0

func find(board Board)  {
    if board.queens_count == board.size {
        //fmt.Println(" Sol ", numSolutions)
        numSolutions ++
        //board.Print()
    } else {
        var row = board.queens_count

        for col:=0; col< board.size; col++ {
            if board.IsValidPos(row, col) {
                board.SetQueen(row, col)
                find(board)
                board.RemoveQueen(row, col)
            }
        }
    }
}

func main() {
    var boardSize = 0
    var time1 time.Time
    var time2 time.Time
    //var wg sync.WaitGroup

    fmt.Println("Board Size: ")
    fmt.Scanln(&boardSize)
    fmt.Println("Board Size: ", boardSize, "x", boardSize)

    time1 = time.Now()

    var board Board
    board.Init(boardSize)
    find(board)

    time2 = time.Now()

    fmt.Println("Num solutions: ", numSolutions)
    fmt.Println("seconds", time2.Unix() - time1.Unix())
}