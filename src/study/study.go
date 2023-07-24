package study

func getMax(arr []int)int{ //获取数组中的最大值
    return processGetMax(arr,0,len(arr)-1) //返回数组arr中[L...R]范围上的最大值
}

func processGetMax(arr []int,L,R int)int{
    if L==R{
        return arr[L] //arr[L...R]范围上只有一个数，直接返回
    }
    mid:=L+((R-L)>>1) //中点，防止溢出
    leftMax:=processGetMax(arr,L,mid) //左部分获得的最大值
    rightMax:=processGetMax(arr,mid+1,R) //右部分获得的最大值
    return max(leftMax,rightMax) //返回左右部分获得的最大值
}

func max[T int|int8|int16|int32|int64|float32|float64](a,b T)T{ //泛型
    if a>b{
        return a
    }else{
        return b
    }
}