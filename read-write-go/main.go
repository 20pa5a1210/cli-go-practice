package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	println("Enter new data to file:\n")
	scanner.Scan()
	str := scanner.Text()
	content, _ := ioutil.ReadFile("file.txt")
	bytes := []byte(str)
	content = append(content, bytes...)
	err := ioutil.WriteFile("file.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	content1, _ := ioutil.ReadFile("file.txt")
	for index, data := range content1 {
		fmt.Println(index, " ", string(data))
	}
}

