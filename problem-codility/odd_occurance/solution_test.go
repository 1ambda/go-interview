package solution

import (
	"testing"
	"reflect"
	"log"
)

func TestSolution(t *testing.T) {
	arr1 := []int{9, 3, 9, 3, 9, 7, 9}
	ret1 := Solution(arr1)
	if !reflect.DeepEqual(arr1, ret1) {
		log.Fatalf("%v is not equal to %v", ret1, 7)
	}
}