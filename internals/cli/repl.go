package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Repl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter the filters for the car you want")
		reader.Scan()

		fmt.Printf("these are the inputs:%v", reader.Text())
	}
}
