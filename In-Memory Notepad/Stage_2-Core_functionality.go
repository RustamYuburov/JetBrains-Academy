package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
	var records []string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command and data: ")
		scanner.Scan()

		input := strings.SplitN(scanner.Text(), " ", 2)
		command := input[0]

		switch command {
		case "exit":
			fmt.Println("[Info] Bye!")
			break
		case "clear":
			records = nil
			fmt.Println("[OK] All notes were successfully deleted")
		case "list":
			for i, el := range records {
				fmt.Printf("[Info] %d: %s\n", i+1, el)
			}
		case "create":
			if len(records) >= 5 {
				fmt.Println("[Error] Notepad is full")
				continue
			}
			line := input[1]
			records = append(records, line)
			fmt.Println("[OK] The note was successfully created")
		}
	}
}
