package main

import "fmt"
import "time"

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	//	10億回ループで約1.2秒
	for i := 0; i < 1000000000; i++ {
		if i == 5000 {
			fmt.Print(i)
			fmt.Println("");
		}
	}
	t2 := time.Now()
	fmt.Println(t2)
}
