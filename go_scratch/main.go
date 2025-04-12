package main

import (
	"fmt"
)

// func main() {
//
// 	// fmt.Println("=======================Variables & Constants=======================")
// 	// var a int = 1
// 	// var b string = "test"
// 	// c := 1
// 	// const test = 1
// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	// fmt.Println(c)
// 	// fmt.Println(test)
// 	//
// 	// fmt.Println("=======================For loops --> Only loop construct in GO=======================")
// 	// for i := 0; i < 3; i++ {
// 	// 	fmt.Println(i)
// 	// }
// 	//
// 	// for i := range 2 {
// 	// 	fmt.Println(i)
// 	// }
// 	//
// 	// i := 0
// 	// for {
// 	// 	fmt.Println("infinite loop to be break after some condition")
// 	// 	if i > 1 {
// 	// 		break
// 	// 	}
// 	// 	i += 1
// 	// }
// 	//
// 	// if n := 1; n > 0 {
// 	// 	fmt.Println("if test")
// 	// } else {
// 	// 	fmt.Println("else test")
// 	// }
// 	//
// 	// fmt.Println("=======================Switch=======================")
// 	// weekday := "test"
// 	// switch weekday {
// 	// case "test", "tt":
// 	// 	fmt.Println("case 1")
// 	// default:
// 	// 	fmt.Println("case 2")
// 	// }
// 	//
// 	// fmt.Println("=======================Arrays --> Fixed-size=======================")
// 	// arr := [3]int{1, 2}
// 	// // arr = append(arr, 6)
// 	// fmt.Println("Arr lenght -->", len(arr))
// 	// fmt.Println("Arr  -->", arr)
// 	//
// 	// for a := range arr {
// 	// 	fmt.Println(arr[a])
// 	// }
// 	//
// 	// // [...] --> infers size of array from the elements present in the array
// 	// // 5: 500 --> 5 is the index of the array where 500 will be placed
// 	// z := [...]int{100, 400, 5: 500}
// 	// fmt.Println("idx:", z)
// 	//
// 	// var twoD [2][3]int
// 	// for i := 0; i < 2; i++ {
// 	// 	for j := 0; j < 3; j++ {
// 	// 		twoD[i][j] = i + j
// 	// 	}
// 	// }
// 	// fmt.Println("2d: ", twoD)
// 	// twoD = [2][3]int{
// 	// 	{1, 2, 3},
// 	// 	{1, 2, 3},
// 	// }
// 	// fmt.Println("2d: ", twoD)
// 	//
// 	// fmt.Println("=======================Slices --> Variable-sized Sequences=======================")
// 	// // Slices
// 	// s := make([]string, 3)
// 	// fmt.Println("emp:", s)
// 	// s[0] = "a"
// 	// s[1] = "b"
// 	// s[2] = "c"
// 	// fmt.Println("set:", s)
// 	//
// 	// s = append(s, "d", "e")
// 	// fmt.Println("append:", s)
// 	//
// 	// f := make([]string, len(s))
// 	// copy(f, s)
// 	// fmt.Println("copy:", f)
// 	// f = append(f, "f")
// 	// fmt.Println("copy:", f, s)
// 	//
// 	// fmt.Println("are slices equal?? ", slices.Equal(f, s))
// 	//
// 	// fmt.Println("=======================Maps --> Unordered Collection of Key-Value Pairs=======================")
// 	// _map := make(map[string]int)
// 	// _map["k1"] = 7
// 	// _map["k2"] = 13
// 	// fmt.Println("map:", _map)
// 	// fmt.Println("key not exists:", _map["k3"])
// 	//
// 	// delete(_map, "k2")
// 	// fmt.Println("after deletion of k2 key ", _map)
// 	//
// 	// _, prs := _map["k2"]
// 	// fmt.Println("prs:", prs)
// 	//
// 	// for key, value := range map[string]string{"1": "1", "2": "2"} {
// 	// 	fmt.Println("key:", key, "value:", value)
// 	// }
// 	//
// 	// _sum := sum(1, 2)
// 	// fmt.Println("sum:", _sum)
// 	//
// 	// // multiple multiple_return
// 	// _, r, _ := multiple_return(1, 2)
// 	// fmt.Println("multiple_return:", r)
// 	// variadic(1, 1, 1, 342, 324234)
// 	//
// 	// arr = [3]int{1, 2, 3}
// 	// for idx, num := range arr {
// 	// 	fmt.Println(idx, num)
// 	// }
// 	//
// 	// fmt.Println("=======================Range over string and runes=======================")
// 	// for i, c := range "abc" {
// 	// 	fmt.Println(i, c)
// 	// }
// 	//
// 	// testSlice := []int{}
// 	// for i := 0; i < 100; i++ {
// 	// 	testSlice = append(testSlice, i)
// 	// 	fmt.Println("Capacity:", cap(testSlice), "Length:", len(testSlice))
// 	// }
// 	// slice := make([]int, 10, 15)
// 	// fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
//
// 	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
// 	fmt.Println(sample)
// 	fmt.Printf("%q\n", sample)
// 	fmt.Printf("%+q\n", sample)
//
// 	const rawString = "สวัสดี"
// 	for i := 0; i < len(rawString); i++ {
// 		// Prints the Rune (unicode codepoint) value of the character
// 		fmt.Println(rawString[i])
// 	}
//
// 	fmt.Println("Rune count:", utf8.RuneCountInString(rawString))
// 	for idx, runeValue := range rawString {
// 		fmt.Printf("%#U starts at %d\n", runeValue, idx)
// 	}
// }

func sum(a, b int) int {
	return a + b
}

func multiple_return(a, b int) (int, int, int) {
	return a, b, 3
}

func variadic(nums ...int) int {
	fmt.Println(nums)
	return 1
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
