package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

//func genCode(r, g, b int) int {
//	return 16 + 36*r + 6*g + b
//}

// https://github.com/busyloop/lolcat/blob/master/lib/lolcat/lol.rb
func ColorScreen(i int, freq float64) (int, int, int) {
	if freq == 0 {
		freq = 0.1
	}
	return int(math.Sin(freq*float64(i)+0)*127 + 128),
		int(math.Sin(freq*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(freq*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune) {
	for i, value := range output {
		r, g, b := ColorScreen(i, 0.1)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, value)
		//fmt.Printf("\033[38;5;%dm%c\033[0m", genCode(r, g, b), value)
	}
	fmt.Println()
}

func main() {
	var output []rune

	reader := bufio.NewReader(os.Stdin)

	for {
		val, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		output = append(output, val)
	}

	print(output)
}
