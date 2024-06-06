package main

import (
	"flag"
	"fmt"
)

func main() {
	arg1 := flag.String("flag", "Default value", "Description")
	flag.Parse()
	fmt.Println(*arg1)
}
