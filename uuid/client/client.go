package main

import (
	"fmt"

	"github.com/tawakhal/util/uuid"
)

func main() {
	data, _ := uuid.New()
	fmt.Printf("MSB : %d\nLSB : %d\n\n", data.MSB, data.LSB)

	asdas, _ := uuid.FromString(data.String())
	fmt.Println(asdas.MSB, asdas.LSB)
}
