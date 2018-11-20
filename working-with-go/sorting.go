package main

import (
	"fmt"
	"sort"
)

func main() {
	abc := []string{"gih", "def", "zxy", "abc"}
	nums := []int{4, 2, 12, 5, 1, 3}
	fmt.Println("Raw abc: ", abc)
	fmt.Println("Raw nums: ", nums)

	//string array排序
	sort.Strings(abc)
	fmt.Println("Sorted abc: ", abc)
	//int array排序
	sort.Ints(nums)
	fmt.Println("Sorted nums: ", nums)

	//逆序
	sort.Sort(sort.Reverse(sort.StringSlice(abc)))
	fmt.Println("Reverse abc: ", abc)

	hash := map[string]int{
		"c": 3,
		"a": 1,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	for k, v := range hash {
		fmt.Printf("%s => %v\n", k, v)
	}

	//把mapn keys放到数组里面排序, map是无序的
	keys := make([]string, 0, len(hash))
	for k := range hash {
		fmt.Println("The value of map is: ", k)
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		fmt.Printf("%s => %v\n", keys[i], hash[keys[i]])
	}

}
