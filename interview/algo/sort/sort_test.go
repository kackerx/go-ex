package sort

import (
	"fmt"
	"testing"
)

var arr = []int{1, 8, 7, 4, 4, 9, 10, 5, 3, 7, 9}

func TestSort(t *testing.T) {
	Maopao(arr)
	fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	partitionV3(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
	HeapSort(arr)
	fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
	SelectSort(arr)

	fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
	insertSort(arr)
	fmt.Println(arr)
}

func TestName(t *testing.T) {
	fmt.Println(0 / 2)
}
