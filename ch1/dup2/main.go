package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 : %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for fileName, duplicateLines := range counts {
		fmt.Println(fileName)
		for duplicateLine, dupCounter := range duplicateLines{  
			if dupCounter > 1 {
				fmt.Printf("%d\t%s\n", dupCounter, duplicateLine)
			}
		}

	}

}

func countLines(file *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
			value, exists := counts[file.Name()]
			if exists {
				value[input.Text()]++
			}else{
				counts[file.Name()] = make(map[string]int)
			}
	}
}
