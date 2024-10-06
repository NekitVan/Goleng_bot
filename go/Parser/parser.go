package Parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type WeatherSelectors struct {
	Name        string `json:"name"`
	Temperature string `json:"temperature"`
	Update      string `json:"update"`
}
type WeatherData struct {
	Tu_day WeatherSelectors `json:"tu_day"`
}

func Parse(resp *http.Response) {

	file, err := os.Open("json/HTML.json")
	if err != nil {
		log.Fatal("Ошибка при открытии файла:", err)
	}
	defer file.Close()

	// Декодируем JSON в структуру
	var weatherData WeatherData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&weatherData)
	if err != nil {
		log.Fatal("Ошибка при декодировании JSON:", err)
	}

	// Закрываем тело ответа после обработки
	defer resp.Body.Close()

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	cityElem := doc.Find(weatherData.Tu_day.Name)
	if cityElem.Length() > 0 {
		fmt.Printf("%s\n", cityElem.Text())
	} else {
		fmt.Println("Не удалось найти название района.")
	}

	temperatureElem := doc.Find(weatherData.Tu_day.Temperature)
	if temperatureElem.Length() > 0 {
		fmt.Printf("Температура: %s \n", temperatureElem.Text())
	} else {
		fmt.Println("Не удалось найти температуру.")
	}

	lastUpdateElem := doc.Find(weatherData.Tu_day.Update)
	if lastUpdateElem.Length() > 0 {
		fmt.Printf("Последнее обновление: %s\n", lastUpdateElem.Text())
	} else {
		fmt.Println("Не удалось найти время обновления.")
	}
}
