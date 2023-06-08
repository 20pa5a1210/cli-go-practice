package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func reversestr(str string) bool {
    str = strings.ToLower(strings.ReplaceAll(str, " ", ""))

    // Get the length of the string
    length := len(str)

    // Check characters from the beginning and end until the middle is reached
    for i := 0; i < length/2; i++ {
        if str[i] != str[length-i-1] {
            return false
        }
    }

    return true
}
func unkown() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "help" {
			println("---- help menu ")
			println("---- help menu ")
			println("---- help menu ")
			println("---- help menu ")
			println("---- help menu ")
		}
		if text == "clear" {
			for {
				println("Enter a string")
				scanner.Scan()
				str := scanner.Text()
				res := reversestr(str)
				if res {
					fmt.Println("yes", str)
				} else {
					fmt.Println("no")
				}
				if str == "exit" {
					break
				}
			}
		}
		if text == "clear" {
			print("clear")
		}
		if text == "exit" {
			os.Exit(0)
		}
	}
}

