package study

import "fmt"

//荷兰国旗问题
/*
1. 给定一个数组arr，和一个数num，请把小于等于num的数放在数组的左边，
大于num的数放在数组的右边。要求额外空间复杂度O（1），时间复杂度O（N）
*/

func preDutchFlag(arr []int, num int) {
	less:=-1 //小于等于区域的右边界
	for i:=0;i<len(arr);i++{ //遍历数组
		if arr[i]<=num{ //当前数小于等于num
			arr[i],arr[less+1]=arr[less+1],arr[i] //将当前数与小于等于区域的下一个数交换
			less++ //小于等于区域的右边界+1
		}
	}
	fmt.Println(arr)
}

/*
2. 荷兰国旗问题，给定一个数组arr，和一个数num，请把小于num的数放在数组的左边，
等于num的数放在数组中间，大于num的数放在数组的右边。要求额外空间复杂度O（1），时间复杂度O(N)
*/

func dutchFlag(arr []int, num int) {
	less:=-1 //小于区域的右边界
	more:=len(arr) //大于区域的左边界
	for i:=0;i<more;{ //遍历数组，终止遍历条件为i==more，即当前和大于区重合
		if arr[i]<num{ //当前数小于num
			arr[i],arr[less+1]=arr[less+1],arr[i] //将当前数与小于区域的下一个数交换
			less++ //小于区域的右边界+1
			i++
		}else if arr[i]>num{ //当前数大于num
			arr[i],arr[more-1]=arr[more-1],arr[i] //将当前数与大于区域的前一个数交换
			more-- //大于区域的左边界-1
		}else{
			i++ //当前数等于num，不交换，继续遍历
		}
	}
	fmt.Println(arr)
} 

