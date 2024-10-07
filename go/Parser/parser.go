package Parser

import (
	"encoding/json"

	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/pterm/pterm"
)

type WeatherSelectors struct {
	Name        string `json:"name"`
	Temperature string `json:"temperature"`
	Update      string `json:"update"`
	Uf_index    string `json:"uf_index"`
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

	temperatureElem := doc.Find(weatherData.Tu_day.Temperature)

	lastUpdateElem := doc.Find(weatherData.Tu_day.Update)

	uf_index_city := doc.Find(weatherData.Tu_day.Uf_index)

	text := pterm.LightRed(cityElem.Text())
	box1 := pterm.DefaultBox.Sprint(temperatureElem.Text())
	box2 := pterm.DefaultBox.Sprint(lastUpdateElem.Text())
	box3 := pterm.DefaultBox.Sprint(uf_index_city.Text())
	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: box1}, {Data: box2}, {Data: box3}},
	}).Srender()
	pterm.DefaultBox.WithTitle(text).WithLeftPadding(4).WithRightPadding(4).WithBottomPadding(4).Println(panels)
}
