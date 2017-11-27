package main

import (
	"fmt"

	"github.com/aelse/seppuku"
)

func main() {
	if err := seppuku.Seppuku(); err != nil {
		fmt.Printf("Failed to request self-destruction: %v\n", err.Error())
	}
}
