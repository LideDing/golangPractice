package study

import (
	"fmt"
	"testing"
)

func TestGetMax(t *testing.T) {
	arr := []int{951, 2, 3, 4, 5, 6, 7, 1120, 8, 9, 10}
	maxNum := getMax(arr)
	fmt.Println(maxNum)

}

func TestSmallNum(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}
	fmt.Println(smallSum(arr))
}

func TestPreDutchFlag(t *testing.T) {
	arr := []int{1, 3, 4, 3,9,10,6,2, 5}
	preDutchFlag(arr, 5)
}

func TestDutchFlag(t *testing.T) {
	arr := []int{3,5,0,3,14,5,2,6,9,-6,5}
	dutchFlag(arr, 5)
}