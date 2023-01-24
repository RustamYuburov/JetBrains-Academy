package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maxNotes int

func main() {
	var records []string

	handleInput(records)
}

func handleInput(records []string) {
	scanner := bufio.NewScanner(os.Stdin)

	maxNotes = getMaxLength()

	for {
		input := getInput(*scanner)
		resolveCommand(input, &records)
	}
}

func getMaxLength() int {
	var maxLength int

	fmt.Print("Enter the maximum number of notes:")
	fmt.Scan(&maxLength)

	return maxLength
}

func getInput(scanner bufio.Scanner) []string {
	fmt.Print("Enter a command and data: ")
	scanner.Scan()

	input := strings.SplitN(scanner.Text(), " ", 2)

	return input
}

func resolveCommand(input []string, recordsPointer *[]string) {
	command := input[0]

	switch command {
	case "exit":
		exitApp()
		break

	case "clear":
		clearRecords(recordsPointer)

	case "list":
		listRecords(recordsPointer)

	case "create":
		createRecord(input, recordsPointer)

	default:
		handleUnknownCommand()
	}
}

func exitApp() {
	fmt.Println("[Info] Bye!")
}

func handleUnknownCommand() {
	fmt.Println("[Error] Unknown command")
}

func clearRecords(recordsPointer *[]string) {
	*recordsPointer = nil
	fmt.Println("[OK] All notes were successfully deleted")
}

func listRecords(recordsPointer *[]string) {
	if len(*recordsPointer) == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}

	for i, el := range *recordsPointer {
		fmt.Printf("[Info] %d: %s\n", i+1, el)
	}
}

func createRecord(input []string, recordsPointer *[]string) {
	if len(input) < 2 {
		fmt.Println("[Error] Missing note argument")
		return
	}

	if len(*recordsPointer) >= maxNotes {
		fmt.Println("[Error] Notepad is full")
		return
	}

	line := input[1]
	*recordsPointer = append(*recordsPointer, line)
	fmt.Println("[OK] The note was successfully created")
}


// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	var records []string
// 	var maxNotes int
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Print("Enter the maximum number of notes:")
// 	fmt.Scan(&maxNotes)

// 	for {
// 		fmt.Print("Enter a command and data: ")
// 		scanner.Scan()

// 		input := strings.SplitN(scanner.Text(), " ", 2)
// 		command := input[0]

// 		switch command {
// 		case "exit":
// 			fmt.Println("[Info] Bye!")
// 			break

// 		case "clear":
// 			records = nil
// 			fmt.Println("[OK] All notes were successfully deleted")

// 		case "list":
// 			if len(records) == 0 {
// 				fmt.Println("[Info] Notepad is empty")
// 				continue
// 			}

// 			for i, el := range records {
// 				fmt.Printf("[Info] %d: %s\n", i+1, el)
// 			}

// 		case "create":
// 			if len(input) < 2 {
// 				fmt.Println("[Error] Missing note argument")
// 				continue
// 			}

// 			if len(records) >= maxNotes {
// 				fmt.Println("[Error] Notepad is full")
// 				continue
// 			}

// 			line := input[1]
// 			records = append(records, line)
// 			fmt.Println("[OK] The note was successfully created")

// 		default:
// 			fmt.Println("[Error] Unknown command")
// 		}
// 	}
// }
