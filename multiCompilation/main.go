package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Esse é programa do exercício de compilação cruzada. Foi compilado num darwin/arm64, e agora está rodando num sistema:", runtime.GOARCH, runtime.GOOS)
}
