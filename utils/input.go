package utils

import (
	"bufio"
	"fmt"
	"strings"
)

// Input let the user type something and get it's value
func Input(reader *bufio.Reader, question string) (typedValue string) {
	fmt.Printf("%s", question)
	typedValue, _ = reader.ReadString('\n')
	typedValue = strings.TrimSpace(typedValue)
	return
}
