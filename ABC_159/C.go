package main

import ("fmt")

func main()  {
	var l int
	fmt.Scanf("%d", &l)

	var a float64 = float64(l)/3
	v := a * a * a

	fmt.Printf("%v", v)
}