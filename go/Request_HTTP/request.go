package Request_HTTP // название пакета должно совпадать с именем папки

import (
	"log"
	"net/http"
)

func Request() *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pogoda.mail.ru/prognoz/magnitogorsk/24hours/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Ошибка при запросе страницы: %d %s", resp.StatusCode, resp.Status)
	}
	return resp // Верните ответ
}
