package l0974_subarray_sums_divisible_by_k

import (
	"fmt"
	"testing"
)

func TestSubarraysDivByK(t *testing.T) {
	fmt.Println(2 % 5)          // 2
	fmt.Println(-2 % 5)         // -2
	fmt.Println((-2%5 + 5) % 5) // 3
	// (sum % K + K) % K其实就是将负余数转为正余数，但并不是简单的去掉负号，而是增加一个k的步长，
	//比如-7%4，它的负余数为-3，加上步长4之后，就转为正余数1了；
	//即对负数有两种计算余数的方式，第一是-7=-1*4+（-3）此时余数为-3，第二是-7=-2*4+1此时余数为1。
	//而并不是说-3装为3，而是-3转为1；这样针对类似于数组[-7,8]，k=4的情况时，就有-7%4为1，（-7+8）%4为1，两者余数相等，所以（-7+8）-（-7）为8可以整除4，这样不会漏掉8这个数了；

	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) // 7
	fmt.Println(subarraysDivByK([]int{2, -2, 2, -4}, 6))       // 2
}
