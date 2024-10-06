package main

import (
	"Goleng_bot/go/Parser"
	"Goleng_bot/go/Request_HTTP" // Путь к Request_HTTP
)

func main() {
	// Выполняем HTTP запрос и парсим данные
	resp := Request_HTTP.Request() // Получаем ответ
	Parser.Parse(resp)             // Передаем ответ для парсинга
}
