package leetcode

//接雨水
func trap(height []int) int { //双指针
    n:=len(height) //数组长度
	if n<3{
		return 0
	}
	leftMax:=make([]int,n) //左边最大值，动态规划思路
	rightMax:=make([]int,n) //右边最大值，动态规划思路
	leftMax[0]=height[0] //第一个元素的左边最大值为自己
	rightMax[n-1]=height[n-1] //最后一个元素的右边最大值为自己
	for i:=1;i<n;i++{ //从左到右遍历，求出每个元素的左边最大值
		leftMax[i]=max(leftMax[i-1],height[i]) //当前元素的左边最大值为前一个元素的左边最大值和当前元素的高度的最大值
	}
	for i:=n-2;i>=0;i--{ //从右到左遍历，求出每个元素的右边最大值
		rightMax[i]=max(rightMax[i+1],height[i]) //当前元素的右边最大值为后一个元素的右边最大值和当前元素的高度的最大值
	}
	ans:=0 //结果
	for i:=0;i<n;i++{ //遍历每个元素，求出每个元素的左右最大值的最小值减去当前元素的高度，就是当前元素能接到的雨水
		ans+=min(leftMax[i],rightMax[i])-height[i] //左右最大值的最小值减去当前元素的高度，就是当前元素能接到的雨水
	}
	return ans
}

func max[T int|int16|int32|int64|int8|float32|float64](a,b T)T{ //求最大值
	if a>b{
		return a
	}
	return b
}

func min[T int|int16|int32|int64|int8|float32|float64](a,b T)T{ //求最小值
	if a<b{
		return a
	}
	return b
}

/*
        ans = 0  # 结果
        n = len(height)
        if n < 3:
            return 0
        left_max_arr = [0] * n
        right_max_arr = [0] * n
        left_max_arr[0] = height[0]
        right_max_arr[n - 1] = height[n - 1]
        for i in range(1, n):
            left_max_arr[i] = max(left_max_arr[i - 1], height[i])
        for i in range(n - 2, -1, -1):
            right_max_arr[i] = max(right_max_arr[i + 1], height[i])

        for i in range(n):
            ans += min(left_max_arr[i], right_max_arr[i]) - height[i]

        return ans
*/