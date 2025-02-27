package interfaces

import (
	"bufio"
	"os"
	"fmt"
)

func GetInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	scanner.Scan()
	return scanner.Text()
}
