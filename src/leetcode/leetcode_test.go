package leetcode

import(
	"fmt"
	"testing"
)

func TestGetRain(t *testing.T) {
	arr:=[]int{0,1,0,2,1,0,1,3,2,1,2,1}
	res:=trap(arr)
	fmt.Println(res)
}