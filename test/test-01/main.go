package test_02

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	h, v, err := parse(raw)
	if err != nil {
		os.Exit(1)
	}

	// print solution
	fmt.Println(getRectangleArea(h, v))
}

const CmdSpace = " "

func getRectangleArea(h, v int) int {
	return h * v
}

func parse(raw []byte) (int, int, error) {
	runes := bytes.Runes(raw)

	h, err := strconv.Atoi(string(runes[0]))
	if err != nil {
		return 0, 0, err
	}

	v, err := strconv.Atoi(string(runes[2]))
	if err != nil {
		return 0, 0, err
	}

	if string(runes[1]) != " " {
		return 0, 0, errors.New("Got invalid separator")
	}

	if h < 0 || v < 0 {
		return 0, 0, errors.New("Got invalid input")
	}

	return h, v, nil
}
