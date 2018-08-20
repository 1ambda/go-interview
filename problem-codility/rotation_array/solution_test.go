package solution

import (
	"testing"
	"reflect"
	"log"
)

func TestSolution(t *testing.T) {
	arr1 := []int{3, 8, 9, 7, 6}
	exp1 := []int{9, 7, 6, 3, 8}
	ret1 := Solution(arr1, 3)
	if !reflect.DeepEqual(arr1, ret1) {
		log.Fatalf("%v is not equal to %v", ret1, exp1)
	}

	var arr2 []int
	var exp2 []int
	ret2 := Solution(arr2, 3)
	if !reflect.DeepEqual(arr2, ret2) {
		log.Fatalf("%v is not equal to %v", ret2, exp2)
	}
}