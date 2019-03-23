package main

import (
	"fmt"
	fdr "olgi/util/folder"
)

func main() {
	err := fdr.CreateFolders("/home/olgi/Project/Go/src/olgi/util/folder/example/1/2/3/4/5")
	fmt.Println("Error : ", err)
}
