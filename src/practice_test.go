package src

import (
	"fmt"
	"golangPractice/src/sort"
	"testing"
)

func TestQucikSort(t *testing.T) {
	arr:=[]int{9,0,3,6,-7,12}
	src.QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
	
}

func TestMergeSort(t *testing.T) {
	arr:=[]int{9,0,3,6,-7,12}
	res:=src.MergeSort(arr)
	fmt.Println(res)
	
}