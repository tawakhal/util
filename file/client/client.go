package main

import (
	"fmt"
	"olgi/util/file"
)

func main() {
	_, err := file.CreateFileLog("")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Berasil !! ")
}
