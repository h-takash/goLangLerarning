package main

import (
	"fmt"
	"./zoo/animals"
)

func main()  {

	fmt.Println("Hello World!")

	fmt.Println(Appname())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}

