package main

import "fmt"

func preFix() {
	var n int
	fmt.Scanf("%d", &n)
	var arr, preFix []int = make([]int,n), make([]int, n+1)
	for i := 0; i < n; i++ {
	    fmt.Scanf("%d" , &arr[i])
	}
	for i := 1; i <= n; i++ {
		preFix[i] = preFix[i-1] + arr[i-1]
	}
	fmt.Println(arr)
	fmt.Println(preFix)
}

func main() {
	preFix()

}
