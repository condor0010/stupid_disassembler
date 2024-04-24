package main

import (
	"bufio"
	"os"
	"fmt"
)

func dis (bytes string) {
	fmt.Println ( bytes )
}


func main () {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println ( "enter asm:" )
	input_bytes, _ := reader.ReadString('\n')

	dis ( input_bytes )

}
