package Parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Parse(resp *http.Response) {
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	cityElem := doc.Find(".hdr__inner")
	if cityElem.Length() > 0 {
		fmt.Printf("%s\n", cityElem.Text())
	} else {
		fmt.Println("Не удалось найти название района.")
	}

	temperatureElem := doc.Find(".p-forecast__temperature-value")
	if temperatureElem.Length() > 0 {
		fmt.Printf("Температура: %s \n", temperatureElem.Text())
	} else {
		fmt.Println("Не удалось найти температуру.")
	}

	lastUpdateElem := doc.Find(".p-forecast__title")
	if lastUpdateElem.Length() > 0 {
		fmt.Printf("Последнее обновление: %s\n", lastUpdateElem.Text())
	} else {
		fmt.Println("Не удалось найти время обновления.")
	}
}
