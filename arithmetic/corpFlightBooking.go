package arithmetic

/**
* @ Description:这里有 n 个航班，它们分别从 1 到 n 进行编号。
我们这儿有一份航班预订表，表中第 i 条预订记录 bookings[i] = [i, j, k] 意味着我们在从 i 到 j 的每个航班上预订了 k 个座位。
请你返回一个长度为 n 的数组 answer，按航班编号顺序返回每个航班上预订的座位数。
示例：
输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
提示：
1 <= bookings.length <= 20000
1 <= bookings[i][0] <= bookings[i][1] <= n <= 20000
1 <= bookings[i][2] <= 10000
* @Author:
* @Date: 2020/2/17 23:55
*/

func CorpFlightBooking(booking [][]int, n int) []int { // 针对一些需要往上增加的数
	dif := make([]int, n+1) // 初始化所有差分为0
	for _, book := range booking {
		dif[book[0]-1] += book[2] // 区间内为加
		dif[book[1]] -= book[2]   // 区间外为减
	}
	for i := 1; i <= n; i++ { // 因为是累加所以只需要上面的循环在区间的第一个数加//第0个为基准值
		dif[i] += dif[i-1]
	}
	return dif[:n]
}

//关注点：差分序列