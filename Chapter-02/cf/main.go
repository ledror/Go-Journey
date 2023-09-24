package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			handleInput(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			handleInput(input.Text())
		}
	}
}

func handleInput(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celcius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}
