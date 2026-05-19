package main

import "fmt"

// Коды цветов
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m" //начало команды \033[ а дальше задает цвет, 0m сброс цвета
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// Функция очистки экрана
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Цветные сообщения
func success(msg string) {
	fmt.Println(colorGreen + "✓ " + msg + colorReset)
}

func warning(msg string) {
	fmt.Println(colorYellow + "⚠ " + msg + colorReset)
}

func errorMsg(msg string) {
	fmt.Println(colorRed + "✗ " + msg + colorReset)
}

func info(msg string) {
	fmt.Println(colorCyan + "→ " + msg + colorReset)
}

// Красивый заголовок
func printHeader() {
	clearScreen()
	fmt.Println(colorPurple + "╔════════════════════════╗" + colorReset)
	fmt.Println(colorPurple + "║     ТРЕКЕР ЗАДАЧ       ║" + colorReset)
	fmt.Println(colorPurple + "╚════════════════════════╝" + colorReset)
	fmt.Println()
}

func printTask() {
	clearScreen()
	fmt.Println(colorYellow + "╔════════════════════════╗" + colorReset)
	fmt.Println(colorYellow + "║     СПИСОК ЗАДАЧ       ║" + colorReset)
	fmt.Println(colorYellow + "╚════════════════════════╝" + colorReset)
	fmt.Println()
}

func printPoka() {
	clearScreen()
	fmt.Println(colorRed + "╔════════════════════════╗" + colorReset)
	fmt.Println(colorRed + "║     ДО СВИДАНИЯ!!!     ║" + colorReset)
	fmt.Println(colorRed + "╚════════════════════════╝" + colorReset)
	fmt.Println()
}
