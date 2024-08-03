package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	baseUrl := "https://www.google.com/"

	//подключаемся к серверу
	client := http.Client{}
	resp, err := client.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//создаем папку куда будем сохранять файлы
	err = os.Mkdir("site", 0777)
	if err != nil {
		log.Fatal(err)
	}

	//считываемвесь html файл или изображения и сохраняем их
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(data))
	os.WriteFile("site/index.html", data, 0777)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//нужно проходиться по всем тегам рекурсивно и выводить из каждого тега внутренние теги
	for itr := doc.FirstChild; itr != nil; itr = itr.NextSibling {
		fmt.Println(itr.Type)
		if itr.Type == html.ElementNode {
			fmt.Println(itr.Attr)
		}
	}

}
