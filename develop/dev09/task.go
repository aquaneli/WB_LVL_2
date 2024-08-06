package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	baseUrl := "https://tech.wildberries.ru/"
	DownloadSite(baseUrl)
}

func DownloadSite(baseUrl string) {
	//отправляем get запрос
	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//базовый путь к файлу куда сохранять
	path := resp.TLS.ServerName + "1"

	//создаем папку куда будем сохранять файлы
	// err = os.Mkdir(path, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	buf := bytes.Buffer{}
	tee := io.TeeReader(resp.Body, &buf)

	//считали все символы из файла чтобы сохранить
	_, err = io.ReadAll(tee)
	if err != nil {
		log.Fatal(err)
	}
	// os.WriteFile(path+"/index.html", data, 0777)

	//распарсили главный файл html
	bufReader := bytes.NewReader(buf.Bytes())
	node, err := html.Parse(bufReader)
	if err != nil {
		log.Fatal(err)
	}

	//начинаем рекурсивно проходиться по каждому элементу  html файла
	processElem(path, baseUrl, node)

}

func processElem(path string, baseUrl string, node *html.Node) {
	if node.Type == html.ElementNode {
		processNode(path, baseUrl, node)
	}

	//нужно проходиться по всем тегам рекурсивно и выводить из каждого тега внутренние теги
	for itr := node.FirstChild; itr != nil; itr = itr.NextSibling {
		processElem(path, baseUrl, itr)
	}
}

func processNode(path string, baseUrl string, node *html.Node) {
	switch node.Data {
	case "link":
		for _, val := range node.Attr {
			if val.Key == "href" {

				resp, err := http.Get(baseUrl + val.Val)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()

				dir := strings.Split(val.Val, "/")
				file := dir[len(dir)-1]
				pathAll := path + strings.Join(dir[:len(dir)-1], "/")

				err = os.MkdirAll(pathAll, 0777)
				if err != nil {
					log.Fatal(err)
				}

				//считали все символы из файла чтобы сохранить
				data, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				os.WriteFile(pathAll+"/"+file, data, 0777)

				// fmt.Println(pathAll+val.Val, node.Data, val.Key, val.Val)

			}

		}
	case "script":
		for _, val := range node.Attr {
			if val.Key == "src" {

				resp, err := http.Get(baseUrl + val.Val)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()

				dir := strings.Split(val.Val, "/")
				file := dir[len(dir)-1]
				pathAll := path + strings.Join(dir[:len(dir)-1], "/")

				err = os.MkdirAll(pathAll, 0777)
				if err != nil {
					log.Fatal(err)
				}

				//считали все символы из файла чтобы сохранить
				data, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				os.WriteFile(pathAll+"/"+file, data, 0777)

			}

		}

	}
}
