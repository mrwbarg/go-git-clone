package utils

import (
	"fmt"
	"os"
)

func ErrorAndExit(message string) {
	fmt.Fprint(os.Stderr, fmt.Sprintln(message))
	os.Exit(1)
}
