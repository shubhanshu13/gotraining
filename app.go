package main

func main() {
	// var array1 = []int{2, 1, 3, 5, 3, 2}
	// fmt.Println(array1)

	// var array2 int = firstDuplicate(array1)
	// fmt.Println(array2)
	// checkPalindrome("az"
}

// func reverseArray(ar []string) int {
// 	var Length int = len(ar) - 1
// 	fmt.Println(Length)

// 	var rev = make([]string, Length+1)
// 	for i := Length; i >= 0; i-- {
// 		rev[Length-i-1] = ar[i]

// 	}
// 	return len(rev)
// }
// func firstDuplicate(a []int) int {
// 	var length int = len(a)
// 	for i := 0; i < length; i++ {
// 		for j := i + 1; j < length; j++ {
// 			if a[i] == a[j] {
// 				fmt.Println(a[i])
// 				break
// 			}
// 		}
// 	}
// 	return -1
// }

//
//
func shapeArea(n int) int {
	var area int
	j := 0
	for i := 1; i <= n; i++ {
		area = i*j + 1
		j += 2
	}
	return area
}
