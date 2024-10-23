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
	Pressure    string `json:"pressure"`
	WindSpeed   string `json:"wind_speed"`
	Humidity    string `json:"humidity"`
	UfIndex     string `json:"uf_index"`
	Update      string `json:"update"`
}

type WeatherData struct {
	TuDay WeatherSelectors `json:"tu_day"`
}

func Parse(resp *http.Response) {
	// Открываем JSON файл с селекторами
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

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Ищем элементы по селекторам из JSON
	cityElem := doc.Find(weatherData.TuDay.Name)
	temperatureElem := doc.Find(weatherData.TuDay.Temperature)
	lastUpdateElem := doc.Find(weatherData.TuDay.Update)
	pressureElem := doc.Find(weatherData.TuDay.Pressure)
	windSpeedElem := doc.Find(weatherData.TuDay.WindSpeed)
	humidityElem := doc.Find(weatherData.TuDay.Humidity)
	ufIndexElem := doc.Find(weatherData.TuDay.UfIndex)

	// Выводим данные
	text := pterm.LightRed(cityElem.Text())
	box1 := pterm.DefaultBox.Sprint(temperatureElem.Text())
	box2 := pterm.DefaultBox.Sprint(lastUpdateElem.Text())
	box3 := pterm.DefaultBox.Sprint(pressureElem.Text())
	box4 := pterm.DefaultBox.Sprint(windSpeedElem.Text())
	box5 := pterm.DefaultBox.Sprint(humidityElem.Text())
	box6 := pterm.DefaultBox.Sprint(ufIndexElem.Text())

	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: box1}, {Data: box2}, {Data: box3}, {Data: box4}, {Data: box5}, {Data: box6}},
	}).Srender()
	pterm.DefaultBox.WithTitle(text).WithLeftPadding(4).WithRightPadding(4).WithBottomPadding(4).Println(panels)
}
