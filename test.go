package main

import "fmt"

var house = "Malibu Point 10880, 90265"

func main()  {
	i := &house
	fmt.Println(*i)
}
