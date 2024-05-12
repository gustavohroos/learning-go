package main

import (
	"fmt"

	"github.com/gustavohroos/learning-go/packages/pkg2"
)

func main() {
	fmt.Println("This is the main function of the package 1!")
	Other()
	pkg2.Pkg2()
}
