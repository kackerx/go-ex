package search

import (
	"fmt"
	"testing"
)

var arr = []int{0, 1, 4, 5, 7, 8, 9, 12, 17, 19}

var a = []int{1, 1, 3, 3, 5, 5, 7}

func TestBinarySearch(t *testing.T) {
	fmt.Println(BinarySearch(arr, 19))

	fmt.Println(BinarySearchRecursion(arr, 0, len(arr)-1, 19))
}

func TestCeil(t *testing.T) {
	fmt.Println(Ceil(a, 5))
	fmt.Println(Ceil(a, 6))
	fmt.Println(Ceil(a, 1))
}

func TestLower(t *testing.T) {
	fmt.Println(Lower(a, 2))
	fmt.Println(Lower(a, 4))
	fmt.Println(Lower(a, 5))
	fmt.Println(Lower(a, 1))
}
