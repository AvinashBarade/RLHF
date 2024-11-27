package main

import (
	"fmt"

	"github.com/myusername/mygopackage"
)

func main() {
	result := mygopackage.DoSomething()
	fmt.Println(result)

	input := "Hello, World!"
	result2 := mygopackage.DoSomethingElse(input)
	fmt.Println(result2)
}
