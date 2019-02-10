package s6d

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/shiny/screen"
)

func Main(f func(screen.Screen)) {
	fmt.Println("\x1b[?1006h")       // request mouse tracking (SGR_EXT_MODE_MOUSE)
	defer fmt.Println("\x1b[?1006l") // stop mouse tracking
	r := bufio.NewReader(os.Stdin)

	state := def
	for {
		b, err := r.ReadByte()
		if err != nil {
			fmt.Println(err)
			return
		}

		switch state {
		case esc:
			state = oth
			if b == '[' {
				state = csi
			}
		case oth:
			state = def
			fmt.Println("ignoring 0x1b 0x%x\n", b)
		default:
		}

	}
}

const (
	def int = iota
	esc
	csi
	oth
)
