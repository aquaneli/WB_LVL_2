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
	baseUrl := "https://www.google.ru/"
	DownloadSite(baseUrl)
}

func DownloadSite(baseUrl string) {
	//отправляем get запрос
	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//создаем папку куда будем сохранять файлы
	err = os.Mkdir(resp.TLS.ServerName, 0777)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.Buffer{}
	tee := io.TeeReader(resp.Body, &buf)

	//считали все символы из ридера чтобы сохранить
	data, err := io.ReadAll(tee)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(resp.TLS.ServerName+"/index.html", data, 0777)

	//распарсили главный файл html
	bufReader := bytes.NewReader(buf.Bytes())
	node, err := html.Parse(bufReader)
	if err != nil {
		log.Fatal(err)
	}

	//начинаем рекурсивно проходиться по каждому элементу  html файла
	processElem(resp.TLS.ServerName, baseUrl, node)

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
	if node.Data == "link" {
		donwnloadMaterials(path, baseUrl, "href", node)
	} else if node.Data == "script" || node.Data == "source" || node.Data == "div" || node.Data == "img" {
		donwnloadMaterials(path, baseUrl, "src", node)
	}
}

func donwnloadMaterials(path, baseUrl, key string, node *html.Node) {
	sb := strings.Builder{}
	for _, val := range node.Attr {
		if val.Key == key {
			sb.WriteString(baseUrl)
			sb.WriteString(val.Val)
			resp, err := http.Get(sb.String())
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

			sb.Reset()
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			sb.WriteString(pathAll)
			sb.WriteString("/")
			sb.WriteString(file)

			os.WriteFile(sb.String(), data, 0777)

		} else if val.Key == "style" {

			url := strings.Split(val.Val, ":")
			if len(url) > 1 && (url[0] == "--mask-url" || url[0] == "background-image") {
				materialPath := url[1][4 : len(url[1])-1]
				sb.WriteString(baseUrl)
				sb.WriteString(materialPath)
				resp, err := http.Get(sb.String())
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
				sb.Reset()

				dir := strings.Split(materialPath, "/")
				file := dir[len(dir)-1]
				sb.WriteString(path)
				sb.WriteString(strings.Join(dir[:len(dir)-1], "/"))
				

				err = os.MkdirAll(sb.String(), 0777)
				if err != nil {
					log.Fatal(err)
				}

				data, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				sb.WriteString("/")
				sb.WriteString(file)
				os.WriteFile(sb.String(), data, 0777)
				sb.Reset()
			}
		}

	}
}
