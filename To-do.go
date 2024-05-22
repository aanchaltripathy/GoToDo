package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID   int
	Text string
	Done bool
}
type List struct {
	todo []Task
}

func NewToDoList() *List {
	return &List{}
}
func (l *List) AddTask(text string) {
	task := Task{
		ID:   len(l.todo) + 1,
		Text: text,
		Done: false,
	}
	l.todo = append(l.todo, task)
	fmt.Println("task added successfully")
}
func (l *List) DeleteTask(id int) {
	for i, task := range l.todo {
		if task.ID == id {
			l.todo = append(l.todo[:i], l.todo[i+1])
			fmt.Println("task deleted successfully")
		}
	}
	fmt.Println("TASK DOES NOT EXIST")
}
func (l *List) MarkAsDone(id int) {
	for i, task := range l.todo {
		if task.ID == id {
			l.todo[i].Done = true
			fmt.Println("task completed")
			return
		}
	}
	fmt.Println("task not found")
}
func (l *List) Display() {
	for _, task := range l.todo {
		status := " "
		if task.Done == true {
			status = "✓"
		}
		fmt.Printf("[%d]%s %s\n", task.ID, status, task.Text)
	}
}
func main() {
	list := NewToDoList()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("/n--------------------TO-DO LIST---------------------")
		fmt.Println("1. add task")
		fmt.Println("2. delete task")
		fmt.Println("3. mark as done")
		fmt.Println("4. display tasks")
		fmt.Println("5. exit")

		var inp int
		fmt.Println("enter a choice")
		fmt.Scanln(&inp)

		switch inp {
		case 1:
			{
				//var text string
				fmt.Println("Add the task")
				text, _ := reader.ReadString('\n')

				text = strings.TrimSpace(text)
				list.AddTask(text)
			}
		case 2:
			{
				var ind int
				fmt.Println("Enter id of the task to be deleted")
				fmt.Scanln(&ind)
				list.DeleteTask(ind)
			}
		case 3:
			{
				var ind int
				fmt.Println("Enter id to mark as done")
				fmt.Scanln(&ind)
				list.MarkAsDone(ind)
			}
		case 4:
			{
				list.Display()
			}
		case 5:
			{
				fmt.Println("Exiting the program...")
				return
			}
		default:
			{
				fmt.Println("Enter a valid choice")
			}
		}
	}
}
