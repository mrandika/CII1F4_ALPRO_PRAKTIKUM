package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5 int
	var s1, s2, s3, s4, s5 int
	var is_straight, is_flush bool

	fmt.Scanf("%d%c %d%c %d%c %d%c %d%c", &a1, &s1, &a2, &s2, &a3, &s3, &a4, &s4, &a5, &s5)

	is_straight = (a1+1 == a2 && a2+1 == a3 && a3+1 == a4 && a4+1 == a5 && a1 != 1) || (a1 == 10 && a2 == 11 && a3 == 12 && a3 == 13 && a4 == 1)
	is_flush = (s1 == s2 && s1 == s3 && s1 == s4 && s1 == s5)

	fmt.Println(is_straight)
	fmt.Println(is_flush)
}