package main

import (
	"fmt"
	"os"
)

func maxn() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}
