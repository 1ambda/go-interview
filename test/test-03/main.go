package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//bytes := []byte("1 2 3 4\n7")
	//bytes := []byte("1 2 3 4 5\n6")

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	raw := string(bytes)
	commands := strings.Split(raw, "\n")

	// invalid STDIN
	if len(commands) != 2 {
		os.Exit(1)
	}

	sticks, err := parseIntegers(commands[0])
	if err != nil {
		os.Exit(1)
	}
	target, err := parseTarget(commands[1])
	if err != nil {
		os.Exit(1)
	}

	printCombinationSums(sticks, target)
}

const CharSpace = " "

func printCombinationSums(candidates []int, target int) {
	solutions := make([][]int, 0)

	current := make([]int, 0)
	sum := 0

	// use recursive backtracking to find combination sums
	backtrack(candidates, target, sum, current, &solutions)

	// filter out invalid solutions for this problem
	// solution's length is greater than 2
	adjustedSolutions := make([][]int, 0)
	for i := range solutions {
		if len(solutions[i]) == 2 {
			adjustedSolutions = append(adjustedSolutions, solutions[i])
		}
	}

	// then, print the solution
	printSolution(adjustedSolutions)
}

func printSolution(solutions [][]int) {
	// formatting the calculated solution
	if len(solutions) == 0 {
		fmt.Println(-1)
		return
	}

	// output should be sorted and find the shortest stick pair
	var shortestStickPair []int
	for i := range solutions {
		s := solutions[i]
		sort.Ints(s[:])

		if shortestStickPair == nil || shortestStickPair[0] > s[0] {
			shortestStickPair = s
		}

	}

	if len(solutions) == 1 {
		printSlice(solutions[0])
		return
	}

	// If you find 2 or more pairs,
	// output the pair which contains the stick of the shortest length.
	printSlice(shortestStickPair)
}

func printSlice(candidates []int) {
	fmt.Printf("%d %d", candidates[0], candidates[1])
}

func backtrack(nums []int, target int, sum int, current []int, solutions *[][]int) {
	if sum > target {
		return
	}

	if sum == target {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*solutions = append(*solutions, tmp)
		return
	}

	for i := 0; i < len(nums); i ++ {
		n := nums[i]

		current = append(current, n)
		sum += n

		backtrack(nums[i+1:], target, sum, current, solutions)

		current = current[:len(current)-1]
		sum -= n
	}
}

func parseIntegers(raw string) ([]int, error) {
	// trim string to make sure about the input (STDIN)
	raw = strings.TrimSpace(raw)

	splitted := strings.Split(raw, CharSpace)
	integers := make([]int, 0)

	for i := range splitted {
		stringified := splitted[i]
		x, err := strconv.Atoi(stringified)
		if err != nil {
			return nil, err
		}

		integers = append(integers, x)
	}

	return integers, nil
}

func parseTarget(raw string) (int, error) {
	// trim string to make sure about the input (STDIN)
	raw = strings.TrimSpace(raw)

	target, err := strconv.Atoi(raw)
	if err != nil {
		return -1, err
	}

	return target, nil
}
