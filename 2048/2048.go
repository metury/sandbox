package main

import (
	"fmt"
	"math/rand"
	//"os"
	//"strconv"
)

const SIZE = 4;

// BOARD:
// [ 0, 1, 2, 3
//   4, 5, 6, 7
//   8, 9,10,11
//  12,13,14,15]

func printBoard(board *[SIZE*SIZE]int){
	fmt.Print("BOARD:")
	for i := 0; i < len(*board); i++{
		if i % SIZE == 0{
			fmt.Println()
		}
		fmt.Print((*board)[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

func isEnd(board *[SIZE*SIZE]int) bool{
	for i := 0; i < len(*board); i++{
		if (*board)[i] == 0{
			return false;
		}
	}
	return true;
}

func addNumber(board *[SIZE*SIZE]int){
	var free []int
	for i := 0; i < len(*board); i++{
		if (*board)[i] == 0{
			free = append(free, i)
		}
	}
	if len(free) == 0{
		return
	}
	rand.Shuffle(len(free), func(i, j int) {
		free[i], free[j] = free[j], free[i]
	})
	(*board)[free[0]] = 2
}

func step(board *[SIZE*SIZE]int, indexes ... int){
	if len(indexes) < 2{
		return
	}
	if (*board)[indexes[0]] == (*board)[indexes[1]] && (*board)[indexes[0]] != 0{
		(*board)[indexes[0]] = (*board)[indexes[0]] + (*board)[indexes[1]]
		(*board)[indexes[1]] = 0
		step(board, indexes[1:]...)
	} else if (*board)[indexes[0]] == 0{
		(*board)[indexes[0]] = (*board)[indexes[1]]
		(*board)[indexes[1]] = 0
		step(board, indexes[1:]...)
		if (*board)[indexes[0]] != 0{
			step(board, indexes[:]...)
		}
	} else{
		step(board, indexes[1:]...)
	}
}

func compareAndSum(board *[SIZE*SIZE]int, first int, second int) bool{
	if (*board)[second] == (*board)[first] && (*board)[second] != 0{
		(*board)[first] = (*board)[first] * 2
		(*board)[second] = 0
		return true
	} else if (*board)[first] == 0 && (*board)[second] != 0{
		(*board)[first] = (*board)[second]
		(*board)[second] = 0
		return true
	}
	return false
}

func moveUp(board *[SIZE*SIZE]int){
	for i := 0; i < SIZE; i++{
		step(board, i, i + SIZE, i + 2*SIZE, i + 3*SIZE)
	}
	/*cont := true
	for cont{
		cont = false
		for i := SIZE; i < len(*board); i++{
			cont = cont || compareAndSum(board, i - SIZE, i)
		}
	}*/
}

func moveDown(board *[SIZE*SIZE]int){
	cont := true
	for cont{
		cont = false
		for i := len(*board) - 1; i >= SIZE; i--{
			cont = cont || compareAndSum(board, i, i - SIZE)
		}
	}
}

func moveLeft(board *[SIZE*SIZE]int){
	cont := true
	for cont{
		cont = false
		for j := 0; j < SIZE - 1; j++{
			for i := j; i < len(*board) - 1; i += 4{
				cont = cont || compareAndSum(board, i, i + 1)
			}
		}
	}
}

func main(){
	board := [SIZE*SIZE]int{}
	printBoard(&board)
	addNumber(&board)
	printBoard(&board)
	addNumber(&board)
	printBoard(&board)
	addNumber(&board)
	printBoard(&board)
	addNumber(&board)
	printBoard(&board)
	addNumber(&board)
	printBoard(&board)
	addNumber(&board)
	printBoard(&board)
	addNumber(&board)
	printBoard(&board)
	moveUp(&board)
	printBoard(&board)
}
