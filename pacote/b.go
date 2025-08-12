package pacote

import (
	"fmt"
	"myFirstGoProject/pacote/internal/foo"
)

var Bar string = "Hello, World!"

func PrintMinha() {
	fmt.Println(foo.Minha)
}
