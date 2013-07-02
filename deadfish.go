package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x int
	for {
		if x < 0 || x == 256 {
			x = 0
		}
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		c := input[0]
		switch c {
		case 's':
			x = x * x
		case 'i':
			x++
		case 'd':
			x--
		case 'o':
			fmt.Println(x)
		}
	}
}
