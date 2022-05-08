package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println("YA")
	fmt.Println("Time : ", time.Now())
	fmt.Println(quote.Go())
}
