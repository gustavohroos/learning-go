package main

import (
	"fmt"

	"github.com/gustavohroos/learning-go/packages/pkg2"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("01novembro2001")
	sb, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(sb))
	pkg2.Pkg2()

	fmt.Println()
	fmt.Println("Comparing right password")
	err = bcrypt.CompareHashAndPassword(sb, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Password was correct")
	fmt.Println()
	fmt.Println("Comparing wrong password")
	err = bcrypt.CompareHashAndPassword(sb, []byte("01novembro2002"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Password was correct")
}
