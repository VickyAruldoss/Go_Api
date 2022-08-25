package main

import (
	"fmt"
	"go-api/jetpack"
)

func main() {
	jet := jetpack.NewJetpack()
	fmt.Println(jet.Print())
}
