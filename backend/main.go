package main

import (
	"apiserver/apiserver"
	"fmt"
)

func main() {
	err := apiserver.Start()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("Программа закончила своё выполнение")
}
