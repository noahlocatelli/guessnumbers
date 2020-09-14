package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())

	var number = rand.Intn(100)
	for  {

		var try int
		fmt.Println("choisir un nombre")
		fmt.Scan(&try)

		if try<number {
			fmt.Println("plus")

		} else if try>number {fmt.Println("moins")} else if try==number{fmt.Print("égale")}
break
	}
}