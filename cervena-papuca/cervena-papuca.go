package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen_packs(pack_size int) []int {
	// Actually Seed is deprecated.
	rand.Seed(time.Now().UnixNano())
	pack := rand.Perm(pack_size)
	for i := 0; i < pack_size; i++ {
		pack[i] = (pack[i] % 2)
	}
	return pack
}

func get_standard_pack() []int {
	pack := []int{0, 0, 0, 0, 0, 0, 1, 1,
		0, 0, 0, 0, 0, 0, 1, 1,
		0, 0, 0, 0, 0, 0, 1, 1,
		0, 0, 0, 0, 0, 0, 1, 1}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pack), func(i, j int) {
		pack[i], pack[j] = pack[j], pack[i]
	})
	return pack
}

func print_stack(stack *[]int) {
	fmt.Print("Round: ")
	for i := 0; i < len(*stack); i++ {
		print_card((*stack)[i])
	}
	fmt.Println()
}

func single_step(playing_first bool, first_pack *[]int, second_pack *[]int) bool {
	var stack []int
	var current int
	for current != 1 {
		if playing_first && len(*first_pack) > 0 {
			current = (*first_pack)[0]
			(*first_pack) = (*first_pack)[1:]
		} else if len(*second_pack) > 0 {
			current = (*second_pack)[0]
			(*second_pack) = (*second_pack)[1:]
		} else {
			break
		}
		stack = append(stack, current)
		playing_first = !playing_first
	}
	print_stack(&stack)
	if !playing_first {
		(*second_pack) = append((*second_pack), stack[:]...)
	} else {
		(*first_pack) = append((*first_pack), stack[:]...)
	}
	return playing_first
}

func print_card(card int) {
	if card == 1 {
		fmt.Print("♦️")
	} else {
		fmt.Print("♠️")
	}
}

func print_pack(text string, pack *[]int) {
	fmt.Print(text)
	for i := 0; i < len((*pack)); i++ {
		print_card((*pack)[i])
	}
	fmt.Println()
}

func print_packs(first_pack *[]int, second_pack *[]int) {
	print_pack(" *  First pack: ", first_pack)
	print_pack(" * Second pack: ", second_pack)
}

func parse_pack(text string, pack *[]int) {
	for i := 0; i < len(text); i++ {
		var value int = int(text[i] - '0')
		*pack = append(*pack, value)
	}
}

func main() {
	var pack_size int
	args := os.Args[1:]
	var pack []int
	var first_pack []int
	var second_pack []int
	if len(args) == 1 {
		size, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		pack_size = size
		pack = gen_packs(pack_size)
		first_pack = pack[0 : (pack_size/2)+1]
		second_pack = pack[(pack_size/2)-1 : pack_size]
	} else if len(args) == 2 {
		parse_pack(args[0], &first_pack)
		parse_pack(args[1], &second_pack)
	} else {
		pack = get_standard_pack()
		first_pack = pack[0 : (len(pack)/2)+1]
		second_pack = pack[(len(pack)/2)-1 : len(pack)]
	}
	round := 0
	print_packs(&first_pack, &second_pack)
	playing_first := true
	for len(first_pack) != 0 && len(second_pack) != 0 {
		playing_first = single_step(playing_first, &first_pack, &second_pack)
		print_packs(&first_pack, &second_pack)
		round++
	}
	fmt.Println("Number of rounds: ", round)
}
