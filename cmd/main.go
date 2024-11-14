package main

import (
	"Goleng_bot/internal/parser"
	"Goleng_bot/internal/webclient"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Инициализация прогресс-бара
	progress, _ := pterm.DefaultProgressbar.WithTotal(100).WithTitle("Загрузка данных...").Start()

	// Канал для передачи результата парсинга
	resultChannel := make(chan string)

	// Выполняем HTTP запрос и парсим данные в отдельной горутине
	go func() {
		// Симуляция задержки запроса
		time.Sleep(2 * time.Second)

		// Выполняем HTTP запрос
		resp, err := webclient.Request()
		if err != nil {
			pterm.Error.Println("Ошибка при выполнении запроса:", err)
			return
		}

		parser.Parse(resp)

		resultChannel <- "Парсинг завершён!"
	}()

	// прогресс от 0 до 100
	for i := 0; i <= 100; i++ {
		progress.UpdateTitle("Прогресс: " + pterm.Sprintf("%d%%", i))
		progress.Increment()
		time.Sleep(30 * time.Millisecond)
	}

	// Ожидаем результата парсинга
	result := <-resultChannel
	progress.Stop()

	// Выводим результат парсинга
	pterm.Success.Println(result) // Выводим сообщение о завершении
}
