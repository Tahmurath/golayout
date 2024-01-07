package main

import (
	"fmt"
	"mod1/mod1pack1"
	"mod2/mod2pack1"
)

func main() {
	fmt.Println("Module 2 -> Pacage Main ->  main()")
	mod2pack1.M2p1app()
	mod1pack1.M1p1app()
}
