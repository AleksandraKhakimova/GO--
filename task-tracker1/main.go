package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

func main() {
	reader := bufio.NewReader(os.Stdin) // ← добавили создание reader

	clearScreen()
	fmt.Println(colorPurple + "╔════════════════════════╗" + colorReset)
	fmt.Println(colorPurple + "║     ТРЕКЕР ЗАДАЧ      ║" + colorReset)
	fmt.Println(colorPurple + "╚════════════════════════╝" + colorReset)
	fmt.Println()
	fmt.Println(colorCyan + "Как запустить?" + colorReset)
	fmt.Println(colorWhite + "1. " + colorReset + "Консольная версия")
	fmt.Println(colorWhite + "2. " + colorReset + "Веб-версия (в браузере)")
	fmt.Print("Выберите: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var choice int
	fmt.Sscan(input, &choice)

	if choice == 2 {
		fmt.Println(colorGreen + "\nЗапускаю веб-сервер..." + colorReset)
		fmt.Println(colorCyan + "Открой браузер и перейди на:" + colorReset + colorWhite + " http://localhost:8080" + colorReset)
		fmt.Println(colorYellow + "Для остановки нажми Ctrl+C." + colorReset)
		startWebServer()
	} else {
		runConsole()
	}
}

func Menu() int {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(colorCyan + "\n--Меню--" + colorReset)
	fmt.Println(colorBlue + "1." + colorReset + "Добавить задачу")
	fmt.Println(colorBlue + "2." + colorReset + "Показать все задачи")
	fmt.Println(colorBlue + "3." + colorReset + "Отметить как выполненную")
	fmt.Println(colorBlue + "4." + colorReset + "Удалить задачу")
	fmt.Println(colorBlue + "5." + colorReset + "Выход")
	fmt.Print(colorGreen + "Выберите действие: " + colorReset)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	var choice int
	fmt.Sscan(input, &choice)
	return choice
}

func addTask(tasks *[]Task) {
	reader := bufio.NewReader(os.Stdin) // создаём читатель из консоли

	fmt.Print("Название задачи: ")
	title, _ := reader.ReadString('\n') // читаем до Enter
	title = strings.TrimSpace(title)

	fmt.Print("Описание задачи: ")
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSpace(desc)

	newTask := Task{
		Title:       title,
		Description: desc,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	*tasks = append(*tasks, newTask)
	fmt.Println(colorPurple + "Задача добавлена!" + colorReset)
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}
	printTask()
	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf(colorGreen+"%d. %s %s\n", i+1, status, task.Title+colorReset)
		fmt.Printf("   %s\n", task.Description)
	}
}

func completeTask(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println(colorRed + "Нет задач для выполнения." + colorReset)
		return
	}
	showTasks(*tasks)
	var num int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите номер задачи для выполнения: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Sscan(input, &num)

	index := num - 1
	if index < 0 || index >= len(*tasks) {
		fmt.Println(colorRed + "Неверный номер задачи!" + colorReset) //надо дать возможность ввести задачу еще раз
		return
	}
	(*tasks)[index].Done = true
	fmt.Printf("Задача '%s' выполнена!\n", (*tasks)[index].Title)
}
func deleteTask(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println(colorRed + "Нет задач для удаления." + colorReset)
		return
	}

	showTasks(*tasks)
	reader := bufio.NewReader(os.Stdin)
	var num int
	fmt.Print("Введите номер задачи для удаления: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Sscan(input, &num)

	index := num - 1

	if index < 0 || index >= len(*tasks) {
		fmt.Println(colorRed + "Неверный номер задачи!" + colorReset) //надо дать возможность ввести задачу еще раз
		return
	}
	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
}
func waitEnter() {
	fmt.Print("\nНажмите Enter для продолжения...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
func runConsole() {
	var tasks []Task

	for {
		printHeader()
		choice := Menu()

		switch choice {
		case 1:
			addTask(&tasks)
			waitEnter()
		case 2:
			showTasks(tasks)
			waitEnter()
		case 3:
			completeTask(&tasks)
			waitEnter()
		case 4:
			deleteTask(&tasks)
			waitEnter()
		case 5:
			printPoka()
			return
		default:
			fmt.Println(colorRed + "Неверный пункт меню!" + colorReset)
			waitEnter()
		}
	}
}
