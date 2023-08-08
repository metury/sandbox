package main

import (
	"fmt"
	"math/rand"
	"os"
)

const SIZE = 4;

// BOARD:
// [ 0, 1, 2, 3
//   4, 5, 6, 7
//   8, 9,10,11
//  12,13,14,15]


func printBoard(board *[SIZE*SIZE]int){	
	sum := 0
	for i := 0; i < len(*board); i++{
		if i % SIZE == 0{
			fmt.Println()
			fmt.Println("-------------------------------------")
			fmt.Print("| ")
		}
		fmt.Printf("% 6d", (*board)[i])
		sum += (*board)[i]
		fmt.Print(" | ")
	}
	fmt.Println()
	fmt.Println("-------------------------------------")
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

func freeSpace(board *[SIZE*SIZE]int) int{
	sum := 0
	for i := 0; i < len(*board); i++{
		if (*board)[i] == 0{
			sum++
		}
	}
	return sum
}

func isEnd(board *[SIZE*SIZE]int) bool{
	for i := 0; i < len(*board); i++{
		if (*board)[i] == 0{
			return false;
		}
	}
	for i := 0; i < len(*board); i++{
		var neighbours []int
		if i % SIZE == 0 || i % SIZE == SIZE - 1{
			neighbours = []int{i - 4, i + 4}
		} else{
			neighbours = []int{i - 4, i - 1, i + 1, i + 4}
		}
		for j := 0; j < len(neighbours); j++{
			if neighbours[j] >= 0 && neighbours[j] < len(*board){
				if (*board)[i] == (*board)[neighbours[j]]{
					return false
				}
			}
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

func copy(board *[SIZE*SIZE]int) [SIZE*SIZE]int{
	var ret [SIZE*SIZE]int
	for i := 0; i < len(ret); i++{
		ret[i] = (*board)[i]
	}
	return ret
}

func cycle() int{
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
	return sum(&board)
}

func leftUp() int{
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
	return sum(&board)
}

func best() int{
	board := [SIZE*SIZE]int{}
	for !isEnd(&board) {
		addNumber(&board)
		printBoard(&board)
		boards := [4][SIZE*SIZE]int{copy(&board), copy(&board), copy(&board), copy(&board)}
		moveUp(&boards[0])
		moveDown(&boards[1])
		moveLeft(&boards[2])
		moveRight(&boards[3])
		results := [4]int{}
		for i := 0; i < len(boards); i++{
			results[i] = freeSpace(&boards[i])
		}
		max := 0
		index := 0
		for i := 0; i < len(boards); i++{
			if results[i] > max{
				index = i
				max = results[i]
			}
		}
		board = copy(&boards[index])
	}
	return sum(&board)
}

func random() int{
	board := [SIZE*SIZE]int{}
	for !isEnd(&board) {
		addNumber(&board)
		printBoard(&board)
		boards := [4][SIZE*SIZE]int{copy(&board), copy(&board), copy(&board), copy(&board)}
		moveUp(&boards[0])
		moveDown(&boards[1])
		moveLeft(&boards[2])
		moveRight(&boards[3])
		index := rand.Int() % len(boards)
		board = copy(&boards[index])
	}
	return sum(&board)
}

func main(){
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	result := 0
	switch args[0]{
		case "-c":
			result = cycle()
			break
		case "-lu":
			result = leftUp()
			break
		case "-b":
			result = best()
			break
		case "-r":
			result = random()
			break
		default:
			fmt.Println("Cycle (-c), leftup (-lu), best (-b), random (-r).")
	}
	if result > 0{
		fmt.Println(result)
	}
}
