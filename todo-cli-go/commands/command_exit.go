package commands

import (
	"fmt"
	"os"
)

func Exit_command(cfg *TodoList, args ...string) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
