package main

import (
	"fmt"

	op "filestore/internal/apiserver/options"
)

func main() {
	a := op.NewOptions()
	fmt.Println(a.MySQLOptions)
	fmt.Println(a.MySQLOptions.NewClient())
}
