package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/notes/Note"
)

func main() {
	var noteDatatitle, noteDataContent = getNoteData()

	userNote, err := note.New(noteDatatitle, noteDataContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Saving of file was failed!")
		return
	}

	fmt.Println("Saving of file was successful")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note Content: ")
	return title, content
}
