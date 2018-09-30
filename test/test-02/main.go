package solution

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	raw := string(bytes)
	commands := strings.Split(raw, "\n")

	// invalid commands structure
	if len(commands) < 2 {
		os.Exit(1)
	}

	// parse init command and create queue
	commandsNum, maxQueueSize, err := parseInitCommand(commands[0])
	if err != nil {
		os.Exit(1)
	}
	queue := NewFIFOQueue(maxQueueSize)

	// handle the following operation commands
	if err = handleOperation(queue, commands[1:], commandsNum); err != nil {
		os.Exit(1)
	}
}

const CmdTake = "TAKE"
const CmdOffer = "OFFER"
const CmdSize = "SIZE"
const CharSpace = " "

func parseInitCommand(initCommand string) (totalCommandNumber, maxQueueSize int, err error) {
	if len(initCommand) != 3 {
		return 0, 0, errors.New("Invalid init command")
	}

	if space := string(initCommand[1]); space != " " {
		return 0, 0, errors.New("Invalid init command")
	}

	totalCommandNumber, err = strconv.Atoi(string(initCommand[0]))
	if err != nil {
		return 0, 0, err
	}

	maxQueueSize, err = strconv.Atoi(string(initCommand[2]))
	if err != nil {
		return 0, 0, err
	}

	return totalCommandNumber, maxQueueSize, nil
}

func handleOperation(queue FIFOQueue, operations []string, totalCommandNum int) (error) {
	for i := 0; i < totalCommandNum; i++ {
		operation := operations[i]
		operation = strings.TrimSpace(operation)
		spliited := strings.Split(operation, CharSpace)

		switch spliited[0] {
		case CmdSize:
			fmt.Println(queue.Size())
		case CmdTake:
			elem, err := queue.Take()
			if err != nil {
				return err
			}
			fmt.Println(elem)
		case CmdOffer:
			if len(spliited) == 1 {
				return errors.New("Invalid OFFER command.")
			}
			elem := spliited[1]
			offered := queue.Offer(elem)
			fmt.Println(offered)
		default:
			return errors.New("Unknown command")
		}

	}

	return nil
}

type FIFOQueue interface {
	Size() int
	Offer(x string) bool
	Take() (string, error)
}

func NewFIFOQueue(maxSize int) FIFOQueue {
	return &fifoQueue{
		maxSize: maxSize,
		vector:  make([]string, 0),
	}
}

type fifoQueue struct {
	maxSize int
	vector  []string
}

func (q fifoQueue) Size() int {
	return len(q.vector)
}

func (q *fifoQueue) Offer(x string) bool {
	if q.Size() == q.maxSize {
		return false
	}

	(*q).vector = append((*q).vector, x)
	return true
}

func (q *fifoQueue) Take() (string, error) {
	// problem doesn't describe this situation, but handle invalid case here.
	if q.Size() == 0 {
		return "", errors.New("Can't take element from an empty queue")
	}

	elem := (*q).vector[0]
	(*q).vector = (*q).vector[1:] // remove first element from the internal vector

	return elem, nil
}
