package project

import (
	"bufio"
	"fmt"
	"go-learning/ch7/project/note"
	"os"
	"strings"
)

func RunNotesApp() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error Occured.! Please try again", err)
		return
	}
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error Creating Object", err)
		return
	}
	note.Display()
	err_save := note.Save()
	if err_save != nil {
		fmt.Println("Result: Saving the note failed")
	} else {
		fmt.Println("Result: Saving the note succeeded")
	}

}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter Note Title")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	content, err := getUserInput("Enter Note Content")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Printf("%v: ", prompt)
	// var value string
	// fmt.Scan(&value)
	// if value == "" {
	// 	return "", errors.New("field is empty! Please try again")
	// }
	// return value, nil
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // For Windows
	return text, nil

}
