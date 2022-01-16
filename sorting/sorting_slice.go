package main

import (
	"fmt"
	"sort"
)



func main()  {
	//s := []int{4, 2, 3, 1}
	s := []string{"Ramos", "Cristiano", "Shubho"}
	sort.Strings(s)
	fmt.Println(s)


	family := []struct{
		club 	string
		ucl 	int
	}{

		{"RealMadrid", 13},
		{"Barcelona", 5},
		{"Liverpool", 5},
		{"Bayarn", 6},
	}
	//sort.SliceStable(family, func(i, j int) bool {
	//	return family[i].ucl > family[j].ucl
	//})
	sort.Slice(family, func(i, j int) bool {
		return family[i].ucl > family[j].ucl
	})
	fmt.Println(family)


}