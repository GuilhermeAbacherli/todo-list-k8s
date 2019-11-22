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

// PressAnyKeyToContinue wait for the user to press any key
func PressAnyKeyToContinue(reader *bufio.Reader) {
	key := Input(reader, "\n\nPressione enter para continuar... ")
	if key != "" {
		return
	}
}
