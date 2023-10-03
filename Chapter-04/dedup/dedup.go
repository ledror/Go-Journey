package dedup

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Dedup(in io.Reader, out io.Writer) {
	seen := make(map[string]bool)
	input := bufio.NewScanner(in)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Fprintf(out, line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
