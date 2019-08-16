package main

import (
	"fmt"
)

func main() {
	cnt := 0

	permMePlease := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//permMePlease := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	permCount := factorial(len(permMePlease))

	permutations := make([][]int, permCount)
	for i := 0; i < permCount; i++ {
		permutations[i] = make([]int, len(permMePlease))
	}

	for perm := range GeneratePermutations(permMePlease) {
		permutations[cnt] = perm
		//fmt.Println(perm)
		cnt++
	}

	fmt.Println(permutations)
	fmt.Println(cnt)

}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func GeneratePermutations(data []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}
func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}
