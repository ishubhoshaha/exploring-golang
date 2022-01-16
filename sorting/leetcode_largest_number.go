package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	numbers := make([]string, 0)

	for i := 0; i < len(nums); i++ {
		numbers = append(numbers, strconv.Itoa(nums[i]))
		//fmt.Println(len(numbers))
	}

	sort.Slice(numbers, func(i, j int) bool {
		//fmt.Println(numbers[i] + numbers[j], numbers[j] + numbers[i])
		return numbers[i] + numbers[j] > numbers[j] + numbers[i]
	})
	//fmt.Println(numbers)
	ans := ""

	for _, v := range numbers {
		ans += v
	}
	//fmt.Println(ans)
	if ans[0] == '0' {
		return "0"
	}
	fmt.Println(ans)
	return ans
}

func main()  {
	//nums := []int{3,30,34,5,9}
	nums := []int{1,2,3,0}
	largestNumber(nums)
}