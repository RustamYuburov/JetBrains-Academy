package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	handleInput()
}

func handleInput() {
	scanner := bufio.NewScanner(os.Stdin)

	var records []string
	maxNotes := getMaxLength()

	for {
		input := getInput(*scanner)

		switch input[0] {
		case "exit":
			exitApp()
			break

		case "clear":
			clearRecords(&records)

		case "list":
			listRecords(&records)

		case "create":
			createRecord(input, &records, &maxNotes)

		case "update":
			updateRecord(input, &records, &maxNotes)

		case "delete":
			deleteRecords(input, &records, &maxNotes)

		default:
			handleUnknownCommand()
		}
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

func createRecord(input []string, recordsPointer *[]string, maxNotes *int) {
	if len(input) < 2 {
		fmt.Println("[Error] Missing note argument")
		return
	}

	if len(*recordsPointer) >= *maxNotes {
		fmt.Println("[Error] Notepad is full")
		return
	}

	line := input[1]
	*recordsPointer = append(*recordsPointer, line)
	fmt.Println("[OK] The note was successfully created")
}

func updateRecord(input []string, recordsPointer *[]string, maxNotes *int) {
	if len(input) <= 1 {
		fmt.Println("[Error] Missing position argument")
		return
	}

	data := strings.SplitN(input[1], " ", 2)

	if len(data) < 2 {
		fmt.Println("[Error] Missing note argument")
		return
	}

	position, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", data[0])
		return
	}

	if position > *maxNotes {
		fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, *maxNotes)
		return
	} else if position > len(*recordsPointer) {
		fmt.Println("[Error] There is nothing to update")
		return
	}

	(*recordsPointer)[position-1] = data[1]
	fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
}

func deleteRecords(input []string, recordsPointer *[]string, maxNotes *int) {
	if len(input) <= 1 {
		fmt.Println("[Error] Missing position argument")
		return
	}

	position, err := strconv.Atoi(input[1])
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", input[1])
		return
	}

	if position > *maxNotes {
		fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, *maxNotes)
		return
	} else if position > len(*recordsPointer) {
		fmt.Println("[Error] There is nothing to delete")
		return
	}

	*recordsPointer = append((*recordsPointer)[:position-1], (*recordsPointer)[position:]...)
	fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
}
