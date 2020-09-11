package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
		fmt.Print("moins")
		rand.Seed(time.Now().UnixNano())
		var number = rand.Intn(100)
		var try int
		fmt.Print("choisir un nombre")
		fmt.Scan(&try)

		for try<number{
	}
}