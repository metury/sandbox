package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func genPacks(packSize int) []int{
	pack := rand.Perm(packSize)
	for i := 0; i < packSize; i++ {
		pack[i] = (pack[i] % 2)
	}
	return pack
}

func printCard(card int){
	if card == 1{
		fmt.Print("R")
	} else {
		fmt.Print("B")
	}
}

func printStack(stack *[]int){
	fmt.Print("Kolo: ")
	for i := 0; i < len(*stack); i++{
		printCard((*stack)[i])
	}
	fmt.Println()
}

func singleStep(playingFirst bool, firstPack *[]int, secondPack *[]int) bool{
	var stack []int
	var current int
	for current != 1 {
		if playingFirst && len(*firstPack) > 0{
			current = (*firstPack)[0]
			(*firstPack) = (*firstPack)[1:]
		} else if len(*secondPack) > 0 {
			current = (*secondPack)[0]
			(*secondPack) = (*secondPack)[1:]
		} else{
			break
		}
		stack = append(stack, current)
		playingFirst = !playingFirst
	}
	printStack(&stack)
	if !playingFirst{
		(*secondPack) = append((*secondPack), stack[:]...)
	} else{
		(*firstPack) = append((*firstPack), stack[:]...)
	}
	return playingFirst
}

func printPacks(firstPack *[]int, secondPack *[]int){
	fmt.Print("První balíček: ")
	for i := 0; i < len((*firstPack)); i++{
		printCard((*firstPack)[i])
	}
	fmt.Println()
	fmt.Print("Druhý balíček: ")
	for i := 0; i < len((*secondPack)); i++{
		printCard((*secondPack)[i])
	}
	fmt.Println()
}

func main() {
	var packSize int
	args := os.Args[1:]
	var pack []int
	var packOne []int
	var packTwo []int
	if len(args) == 1{
		i, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		packSize = i
		pack = genPacks(packSize)
		packOne = pack[0 : (packSize / 2) + 1]
		packTwo = pack[(packSize / 2) - 1 : packSize]
	} else if len(args) == 2{
		for i := 0; i < len(args[0]); i++{
			if args[0][i] == '1'{
				packOne = append(packOne, 1)
			} else {
				packOne = append(packOne, 0)
			}
		}
		for i := 0; i < len(args[1]); i++{
			if args[1][i] == '1'{
				packTwo = append(packTwo, 1)
			} else {
				packTwo = append(packTwo, 0)
			}
		}
	} else {
		packSize = 80
		pack = genPacks(packSize)
		packOne = pack[0 : (packSize / 2) + 1]
		packTwo = pack[(packSize / 2) - 1 : packSize]
	}
	round := 0
	fmt.Print("Balíček: ")
	for i := 0; i < packSize; i++{
		fmt.Print(pack[i])
	}
	fmt.Println()
	printPacks(&packOne, &packTwo)
	playingFirst := true
	for len(packOne) != 0 && len(packTwo) != 0{
		playingFirst = singleStep(playingFirst, &packOne, &packTwo)
		printPacks(&packOne, &packTwo)
		round ++
	}
	fmt.Println("Number of round: ", round)
}

