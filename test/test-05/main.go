package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const CharSpace = " "

func main() {
	//bytes := []byte("2\n2\n1\n2\n-1\n3")
	//bytes := []byte("2\n1\n2\n-1\n3")

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	// split into multiple lines
	raw := string(bytes)
	commands := strings.Split(raw, "\n")
	if len(commands) < 1 {
		// invalid STDIN
		os.Exit(1)
	}

	// parse window size
	windowSize, err := parseWindowSize(commands[0])
	if err != nil {
		os.Exit(1)
	}

	// parse a stream of numbers
	numbers, err := parseStream(commands[1:])
	if err != nil {
		os.Exit(1)
	}

	printMinValueInSlidingWindow(numbers, windowSize)
}

func printMinValueInSlidingWindow(numbers []int64, windowSize int64) {
	var i int64 = 0
	for ; i+windowSize <= int64(len(numbers)); i++ {
		min := maxOf(numbers[i : i+windowSize])
		fmt.Println(min)
	}
}

func maxOf(nums []int64) int64 {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func parseStream(lines []string) ([]int64, error) {

	stream := make([]int64, 0)

	for i := range lines {
		line := lines[i]
		values, err := parseBigIntegers(line)
		if err != nil {
			return nil, err
		}

		if len(values) != 1 {
			return nil, errors.New("Stream has invalid number format")
		}

		stream = append(stream, values[0])
	}

	return stream, nil
}

func parseWindowSize(raw string) (int64, error) {
	values, err := parseBigIntegers(raw)
	if err != nil {
		return 0, err
	}

	if len(values) != 1 {
		return 0, errors.New("Failed to parse window size. Got invalid format")
	}

	windowSize := values[0]
	if windowSize <= 0 {
		return 0, errors.New("Window size shuold be greather than 0")
	}

	return windowSize, err
}

// parseBigIntegers parses string and return the slice of int64
func parseBigIntegers(raw string) ([]int64, error) {
	// trim string to make sure about the input (STDIN)
	raw = strings.TrimSpace(raw)

	splitted := strings.Split(raw, CharSpace)
	integers := make([]int64, 0)

	for i := range splitted {
		stringified := splitted[i]
		x, err := strconv.ParseInt(stringified, 10, 64)
		if err != nil {
			return nil, err
		}

		integers = append(integers, x)
	}

	return integers, nil
}
