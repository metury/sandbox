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

func printStack(stack *[]int){
	fmt.Print("Kolo: ")
	for i := 0; i < len(*stack); i++{
		fmt.Print((*stack)[i])
	}
	fmt.Println()
}

func singleStep(playingFirst bool, firstPack *[]int, secondPack *[]int) bool{
	var stack []int
	var current int
	for ; current != 1; {
		if playingFirst{
			current = (*firstPack)[0]
			(*firstPack) = (*firstPack)[1:]
		} else {
			current = (*secondPack)[0]
			(*secondPack) = (*secondPack)[1:]
		}
		stack = append(stack, current)
		printStack(&stack)
		playingFirst = !playingFirst
	}
	if !playingFirst{
		(*secondPack) = append((*secondPack), stack[:]...)
	} else{
		(*firstPack) = append((*firstPack), stack[:]...)
	}
	return !playingFirst
}

func printPacks(firstPack *[]int, secondPack *[]int){
	fmt.Print("První balíček: ")
	for i := 0; i < len((*firstPack)); i++{
		fmt.Print((*firstPack)[i])
	}
	fmt.Println()
	fmt.Print("Druhý balíček: ")
	for i := 0; i < len((*secondPack)); i++{
		fmt.Print((*secondPack)[i])
	}
	fmt.Println()
}

func main() {
	var packSize int
	args := os.Args[1:]
	if len(args) != 0{
		i, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		packSize = i
	} else {
		packSize = 60
	}
	round := 0
	pack := genPacks(packSize)
	packOne := pack[0 : (packSize / 2) + 1]
	packTwo := pack[(packSize / 2) - 1 : packSize]
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

