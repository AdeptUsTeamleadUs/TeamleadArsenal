package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	checklist := []string{
		"✅ Все задачи завершены и протестированы?",
		"✅ Подготовлен сценарий демонстрации?",
		"✅ Есть ли данные, готовые для показа?",
		"✅ Демо-стенд развёрнут и проверен?",
		"✅ Команда знает порядок выступления?",
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🚀 Чеклист перед демо команды:")

	for _, item := range checklist {
		fmt.Printf("%s (y/n): ", item)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input != "y" {
			fmt.Println("❌ Демо ещё не готово, проверьте:", item)
			os.Exit(1)
		}
	}

	fmt.Println("🎉 Всё готово к демо, можно начинать!")
}