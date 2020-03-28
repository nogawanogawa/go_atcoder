package main

import "fmt"

func count_choice(a int) int{
	count := a * (a-1)/2
	return count
}

func main()  {
	var b, c int
	fmt.Scanf("%d %d", &b, &c)

	b_ := count_choice(b)
	c_ := count_choice(c)
	fmt.Printf("%d", b_+c_)
}