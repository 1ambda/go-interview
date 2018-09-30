package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const CharSpace = " "

func main() {
	//test := []byte("2\n2\n1\n2\n-1\3")
	test := []byte("2\n1\n2\n-1\n3")
	buf := strings.NewReader(string(test))
	scanner := bufio.NewScanner(buf)

	//scanner := bufio.NewScanner(os.Stdin)

	first := true
	streamWindow := stream{}

	// read stream from STDIO
	for scanner.Scan() {
		line := scanner.Text()

		// parse window size (only once)
		if first {
			window, err := parseWindowSize(line)
			if err != nil {
				os.Exit(1)
			}

			streamWindow.size = window
			first = false
			continue
		}

		// handle the remaining stream
		num, err := parseStreamNumber(line)
		if err != nil {
			os.Exit(1)
		}

		// since stream could be extremely large, (0 < W <= 3,000,000,000)
		// it's not a good idea to persist all min values in memory.
		// we will print as soon as we find a single min value.
		streamWindow.Process(num)
	}

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}

type stream struct {
	size           int64
	processedCount int64

	window []int64
}

func (s *stream) Process(number int64) {
	// move window
	(*s).window = append((*s).window, number)
	if int64(len(s.window)) > s.size {
		(*s).window = (*s).window[1:]
	}

	s.processedCount = s.processedCount + 1

	if int64(len(s.window)) == s.size {
		fmt.Println(s.currentMax())
	}
}

func (s *stream) currentMax() int64 {
	max := s.window[0]
	for i := 1; i < len(s.window); i++ {
		if max < s.window[i] {
			max = s.window[i]
		}
	}
	return max
}

func parseStreamNumber(line string) (int64, error) {
	values, err := parseBigIntegers(line)
	if err != nil {
		return 0, err
	}

	if len(values) != 1 {
		return 0, errors.New("Got invalid stream number format")
	}

	number := values[0]
	return number, err
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
