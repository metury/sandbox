package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	fmt.Print(" BOARD:")
	sum := 0
	for i := 0; i < len(*board); i++{
		if i % SIZE == 0{
			fmt.Println()
		}
		fmt.Print((*board)[i])
		sum += (*board)[i]
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Print("SCORE: ")
	fmt.Println(sum)
}

func sum(board *[SIZE*SIZE]int) int{
	sum := 0
	for i := 0; i < len(*board); i++{
		sum += (*board)[i]
	}
	return sum
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
	number := [3]int{2,2,4}
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
	rand.Shuffle(len(number), func(i, j int) {
		number[i], number[j] = number[j], number[i]
	})
	(*board)[free[0]] = number[0]
}

func singleStep(board *[SIZE*SIZE]int, indexes [SIZE]int, flags *[SIZE]bool) bool{
	cont := false
	for i := 0; i < SIZE - 1; i++{
		if (*board)[indexes[i]] == (*board)[indexes[i+1]] && !(*flags)[i] && !(*flags)[i+1] && (*board)[indexes[i]] != 0{
			(*board)[indexes[i]] = (*board)[indexes[i]] + (*board)[indexes[i+1]]
			(*board)[indexes[i+1]] = 0
			(*flags)[i] = true
			cont = true
		} else if (*board)[indexes[i]] == 0 && (*board)[indexes[i+1]] != 0{
			(*board)[indexes[i]] = (*board)[indexes[i+1]]
			(*board)[indexes[i+1]] = 0
			(*flags)[i], (*flags)[i+1] = (*flags)[i+1],(*flags)[i]
			cont = true
		}
	}
	return cont
}

func step(board *[SIZE*SIZE]int, indexes  [SIZE]int){
	flags := [SIZE]bool{}
	for singleStep(board, indexes, &flags){}
}

func moveUp(board *[SIZE*SIZE]int){
	for i := 0; i < SIZE; i++{
		indexes := [SIZE]int {i, i + SIZE, i + 2*SIZE, i + 3*SIZE}
		step(board, indexes)
	}
}

func moveDown(board *[SIZE*SIZE]int){
	for i := 0; i < SIZE; i++{
		indexes := [SIZE]int {i + 3*SIZE, i + 2*SIZE, i + SIZE, i}
		step(board, indexes)
	}
}

func moveLeft(board *[SIZE*SIZE]int){
	for i := 0; i < len(*board); i += 4{
		indexes := [SIZE]int {i, i + 1, i + 2, i + 3}
		step(board, indexes)
	}
}

func moveRight(board *[SIZE*SIZE]int){
	for i := 0; i < len(*board); i += 4{
		indexes := [SIZE]int {i + 3, i + 2, i + 1, i}
		step(board, indexes)
	}
}

func cycle(){
	board := [SIZE*SIZE]int{}
	rounds := 0
	for !	isEnd(&board){
		addNumber(&board)
		printBoard(&board)
		switch rounds{
			case 0:
				moveUp(&board)
				break
			case 1:
				moveLeft(&board)
				break
			case 2:
				moveDown(&board)
				break
			case 3:
				moveRight(&board)
				break
			default:
				rounds = 0
				break
		}
		rounds = (rounds + 1) % 4
	}
}

func leftUp(){
	board := [SIZE*SIZE]int{}
	rounds := 0
	for !isEnd(&board){
		addNumber(&board)
		printBoard(&board)
		switch rounds{
			case 0:
				moveUp(&board)
				break
			case 1:
				moveLeft(&board)
				break
			default:
				rounds = 0
				break
		}
		rounds = (rounds + 1) % 2
	}
}

func best(){
	board := [SIZE*SIZE]int{}
	for !isEnd(&board){
		addNumber(&board)
		printBoard(&board)
		var board_1, board_2, board_3, board_4 [16]int
		for i := 0; i < len(board); i++{
			board_1[i] = board[i]
			board_2[i] = board[i]
			board_3[i] = board[i]
			board_4[i] = board[i]
		}		
		moveUp(&board_1)
		moveDown(&board_2)
		moveLeft(&board_3)
		moveRight(&board_4)
		boards := [][SIZE*SIZE]int{board_1, board_2, board_3, board_4}
		results := []int{sum(&board_1), sum(&board_2), sum(&board_3), sum(&board_4)}
		origin := results
		sort.Ints(results)
		for i := 0; i < len(origin); i++{
			if origin[i] == results[0]{
				for j := 0; j < len(board); j++{
					board[j] = boards[i][j]
				}
				printBoard(&board)
				break
			}
		}
	}
}

func main(){
	//cycle()
	//leftUp()
	best()
}
