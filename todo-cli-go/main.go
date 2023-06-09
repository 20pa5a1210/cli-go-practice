package main

import (
	"bufio"
	"fmt"
	"os"

	"nerdypunk.me/commands"
)

func main() {
	config := commands.TodoList{}
	repl(&config)
}

func repl(config *commands.TodoList) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := commands.Cleaned(text)

		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCommands := commands.GetCommands()
		commandName, ok := availableCommands[command]
		if !ok {
    		fmt.Printf("Uncaught ReferenceError: %s is not defined",command)
			continue
		}
		err := commandName.CallBack(config, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}
