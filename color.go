// Run tmux with -2 option to force using 256 colors.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("\x1b[38;5;8m2\x1b[38;5;9m5\x1b[38;5;10m6\x1b[38;5;11mc\x1b[38;5;12mo\x1b[38;5;13ml\x1b[38;5;14mo\x1b[38;5;15mr\x1b[0m")
	fmt.Println()

	for i := 0; i < 256; i++ {
		print_color(i)

		if (i % 16 == 15) {
			fmt.Println()
		}
	}
}

func print_color(n int) {
	fmt.Printf("%03d\x1b[48;5;%dm    \x1b[0m", n, n)
}
