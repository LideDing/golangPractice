package study

/*
在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。求一个数组的小和。
例：[1,3,4,2,5],1左边比1小的数没有，3左边比3小的数为1，4左边比4小的数为1、3，2左边比2小的数，
5左边比5小的数1、3、4、2，所以小和为1+1+3+1+1+3+4+2=16

深度改写归并排序。可以得到O(NlogN)的时间复杂度
*/

func smallSum(arr []int) int { //获取数组中的最大值
	if len(arr) < 2 {
		return 0
	}
	return processSmallSum(arr, 0, len(arr)-1) //返回数组arr中[L...R]范围上的最大值
}

func processSmallSum(arr []int, L, R int) int {
	if L == R {
		return 0 //arr[L...R]范围上只有一个数，直接返回
	}
	mid := L + ((R - L) >> 1) //中点，防止溢出
	return processSmallSum(arr, L, mid) + processSmallSum(arr, mid+1, R) + merge(arr, L, mid, R)
}

func merge(arr []int, L, mid, R int) int { //将arr[L...mid]和arr[mid+1...R]两部分进行排序
	help := make([]int, R-L+1) //辅助数组
	i := 0 				   //辅助数组的索引
	p1 := L 			   //左部分的第一个数的索引
	p2 := mid + 1 		   //右部分的第一个数的索引
	res := 0 			   //小和
	for p1 <= mid && p2 <= R { //左右两部分都没有越界
		if arr[p1] < arr[p2] { //左部分的数小于右部分的数
			res += arr[p1] * (R - p2 + 1) //计算小和
			help[i] = arr[p1] 			  //将左部分的数放入辅助数组中
			i++ 						  //辅助数组的索引+1
			p1++ 						  //左部分的索引+1
		} else { //左部分的数大于等于右部分的数
			help[i] = arr[p2] //将右部分的数放入辅助数组中
			i++ 			  //辅助数组的索引+1
			p2++ 			  //右部分的索引+1
		}
	}
	for p1 <= mid { //左部分没有越界
		help[i] = arr[p1] //将左部分剩余的数放入辅助数组中
		i++ 			  //辅助数组的索引+1
		p1++ 			  //左部分的索引+1
	}
	for p2 <= R { //右部分没有越界
		help[i] = arr[p2] //将右部分剩余的数放入辅助数组中
		i++ 			  //辅助数组的索引+1
		p2++ 			  //右部分的索引+1
	}
	// for i := 0; i < len(help); i++ {
	// 	arr[L+i] = help[i]
	// }
	copy(arr[L:R+1], help) //将help数组中的值拷贝到arr[L...R]中
	return res
}