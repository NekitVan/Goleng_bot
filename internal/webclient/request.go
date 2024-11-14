package webclient

import (
	"fmt"
	"log"
	"net/http"
)

var (
	request_url  = "https://pogoda.mail.ru/prognoz/magnitogorsk/24hours/"
	request_type = "GET"
	user_agent   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36"
)

// Request делает HTTP запрос и возвращает ответ и ошибку
func Request() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(request_type, request_url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", user_agent)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err // Возвращаем ошибку
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Ошибка при запросе страницы: %d %s", resp.StatusCode, resp.Status)
		return nil, fmt.Errorf("ошибка: статус %d %s", resp.StatusCode, resp.Status) // Возвращаем ошибку
	}

	return resp, nil // Возвращаем ответ и nil как ошибку
}
