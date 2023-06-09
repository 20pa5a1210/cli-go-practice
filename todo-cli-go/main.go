package main

import (
	"bufio"
	"fmt"
	"os"

	"nerdypunk.me/commands"
)

type Todo struct {
	Task string
	Done bool
}

type TodoList struct {
	Tasks []Todo
}

func main() {
	todoList := TodoList{}
	repl(&todoList)

}

func repl(config *TodoList) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(" >")
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
			fmt.Println("Invalid Command!")
			continue
		}
		err := commandName.CallBack(args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}
