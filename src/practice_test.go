package src

import (
	"fmt"
	src "golangPractice/src/sort"

	// src "golangPractice/src/mutiply"
	"testing"
)

// func TestQucikSort(t *testing.T) {
// 	arr:=[]int{9,0,3,6,-7,12}
// 	src.QuickSort(arr,0,len(arr)-1)
// 	fmt.Println(arr)

// }

func TestMergeSort(t *testing.T) {
	arr := []int{9, 0, 3, 6, -7, 12}
	res := src.MergeSort(arr)
	fmt.Println(res)

}

func TestInsertSort(t *testing.T) {
	arr := []int{9, 0, 3, 6, -7, 12}
	res := src.InsertSort(arr)
	fmt.Println(res)
}

func TestSelectSort(t *testing.T) {
	arr := []int{9, 0, 3, 6, -7, 12,-69}
	res := src.SelectSort(arr)
	fmt.Println(res)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 0, 3, 6, -7, 12}
	res := src.BubbleSort(arr)
	fmt.Println(res)
}
// func TestFuncGoroutine(t *testing.T) {
// 	src.FuncGoroutine()
// }

// func TestGoroutineControl(t *testing.T) {
// 	src.GoroutineControl(10, 2)
// }
