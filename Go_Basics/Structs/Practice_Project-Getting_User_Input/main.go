package main

import (
	"bufio"
	"fmt"
	// "go/types"
	"os"
	"strings"

	"example.com/getting_user_input/note"
	"example.com/getting_user_input/todo"
)

type saver interface {
	Save() error
}

// type displayer interface{
// 	Display()
// }

// type outputtable interface{
// 	Save() error
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}



func main() {

	// printSomething(1)
	// printSomething(2.5)
	// printSomething("App")

	title, content := getNoteData()
	todoText := GetUserInput("Note title: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
}

func printSomething(value any){
	intVal,ok := value.(int)
	if ok {
		fmt.Println("Integer: ",intVal)
		return 
	}

	floatVal,ok := value.(float64)
	if ok {
		fmt.Println("Float: ",floatVal)
		return 
	}

	stringVal,ok := value.(string)
	if ok {
		fmt.Println("String: ",stringVal)
		return 
	}


	// switch value.(type){
	// case int:
	// 	fmt.Println("Integer: ",value)
	// case float64:
	// 	fmt.Println("Integer: ",value)
	// case string:
	// 	fmt.Println("Integer: ",value)
	// default:
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note successful")
	return nil
}

func getNoteData() (string, string) {
	title := GetUserInput("Note title:")
	content := GetUserInput("Note content:")

	return title, content
}

func GetUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}


