package leetcode

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
*/
func convert(s string, numRows int) string {
	if numRows <= 1 { //特殊情况
		return s
	}
	res := make([]string, numRows) //创建一个字符串数组
	i := 0                         //行数标志
	flag := -1                     //往上走还是往下走的标志
	for _, b := range s {          //遍历字符串
		res[i] += string(b)           //将字符加入对应行
		if i == 0 || i == numRows-1 { //行首行尾变换
			flag = -flag //变换标志
		}
		i += flag //行数变换
	}
	ans := ""               //结果
	for _, v := range res { //遍历字符串数组
		ans += v
	}
	return ans
}
