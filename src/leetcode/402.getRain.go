package leetcode

func trap(height []int) int {
	return trap2(height)
}

// 接雨水
func trap1(height []int) int { //动态规划解法
	//=============初始化答案和长度=======================
	n := len(height) //数组长度
	if n < 3 {
		return 0
	}
	//===================初始化动态规划数组====================
	leftMax := make([]int, n)   //左边最大值
	rightMax := make([]int, n)  //右边最大值
	leftMax[0] = height[0]      //第一个元素的左边最大值为自己
	rightMax[n-1] = height[n-1] //最后一个元素的右边最大值为自己

	//====================计算动态规划数组=====================
	for i := 1; i < n; i++ { //从左到右遍历，求出每个元素的左边最大值
		leftMax[i] = max(leftMax[i-1], height[i]) //当前元素的左边最大值为前一个元素的左边最大值和当前元素的高度的最大值
	}
	for i := n - 2; i >= 0; i-- { //从右到左遍历，求出每个元素的右边最大值
		rightMax[i] = max(rightMax[i+1], height[i]) //当前元素的右边最大值为后一个元素的右边最大值和当前元素的高度的最大值
	}
	//======================计算并返回结果======================
	ans := 0                 //结果
	for i := 0; i < n; i++ { //遍历每个元素，求出每个元素的左右最大值的最小值减去当前元素的高度，就是当前元素能接到的雨水
		ans += min(leftMax[i], rightMax[i]) - height[i] //左右最大值的最小值减去当前元素的高度，就是当前元素能接到的雨水
	}
	return ans
}

// =====================工具函数=========================
func max[T int | int16 | int32 | int64 | int8 | float32 | float64](a, b T) T { //求最大值，使用了泛型
	if a > b {
		return a
	}
	return b
}

func min[T int | int16 | int32 | int64 | int8 | float32 | float64](a, b T) T { //求最小值，使用了泛型
	if a < b {
		return a
	}
	return b
}

// 双指针解法
func trap2(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	left, right := 0, n-1                       //左右指针
	leftMax, rightMax := height[0], height[n-1] //左右最大值
	ans := 0                                    //结果
	for left <= right {                         //左右指针相遇结束
		leftMax = max(leftMax, height[left])    //更新左边最大值
		rightMax = max(rightMax, height[right]) //更新右边最大值
		if leftMax < rightMax {                 //左边最大值小于右边最大值
			ans += leftMax - height[left] //左边最大值减去当前元素的高度，就是当前元素能接到的雨水
			left++                        //左指针右移
		} else { //左边最大值大于右边最大值
			ans += rightMax - height[right] //右边最大值减去当前元素的高度，就是当前元素能接到的雨水
			right--                         //右指针左移
		}
	}
	return ans
}

// 单调栈解法
func trap3(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	ans := 0                 //结果
	stack := make([]int, 0)  //单调栈
	for i := 0; i < n; i++ { //遍历每个元素
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] { //单调栈不为空且当前元素大于栈顶元素
			top := stack[len(stack)-1]   //栈顶元素
			stack = stack[:len(stack)-1] //出栈
			if len(stack) == 0 {         //栈空了
				break
			}
			left := stack[len(stack)-1]                             //左边界
			curWidth := i - left - 1                                //当前宽度
			curHeight := min(height[left], height[i]) - height[top] //当前高度
			ans += curWidth * curHeight                             //当前面积
		}
		stack = append(stack, i) //当前元素入栈
	}
	return ans
}
