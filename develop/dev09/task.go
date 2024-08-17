package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	baseURL, err := parseURL()
	if err != nil {
		log.Fatal(err)
	}

	node, siteName, err := initDownloadFromHTML(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	err = downloadSite(node, baseURL, siteName)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
}

// Парсим адрес сайта
func parseURL() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New("usage: ./task [URL]")
	}

	return os.Args[1], nil
}

// Конкатенация нескольких строк
func concatStrings(str ...string) string {
	sb := strings.Builder{}
	for _, val := range str {
		sb.WriteString(val)
	}
	return sb.String()
}

// Cамый первый html файл который мы будем парсить
func initDownloadFromHTML(baseURL string) (*html.Node, string, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	err = os.MkdirAll(resp.TLS.ServerName, 0777)
	if err != nil {
		return nil, "", err
	}

	buf := bytes.Buffer{}
	tee := io.TeeReader(resp.Body, &buf)

	data, err := io.ReadAll(tee)
	if err != nil {
		return nil, "", err
	}
	err = os.WriteFile(concatStrings(resp.TLS.ServerName, "/index.html"), data, 0777)
	if err != nil {
		return nil, "", err
	}

	bufReader := bytes.NewReader(buf.Bytes())
	node, err := html.Parse(bufReader)
	if err != nil {
		return nil, "", err
	}

	return node, resp.TLS.ServerName, nil
}

// Обработка всех узлов html страницы рекурсивно
func downloadSite(node *html.Node, baseURL, siteName string) error {

	if node.Type == html.ElementNode && node.Data == "style" {
		if node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
			err := processStyle(node.FirstChild, baseURL, siteName)
			if err != nil {
				return err
			}
		}
	}

	if node.Type == html.ElementNode {
		err := processNode(node, baseURL, siteName)
		if err != nil {
			return err
		}
	}

	for itr := node.FirstChild; itr != nil; itr = itr.NextSibling {
		err := downloadSite(itr, baseURL, siteName)
		if err != nil {
			return err
		}
	}
	return nil
}

// Обработка style ноды
func processStyle(node *html.Node, baseURL, siteName string) error {
	attr := strings.Split(node.Data, " ")
	for _, v := range attr {
		if len(v) > 3 && v[:3] == "url" {
			url := strings.Split(v[3:], "\"")
			if len(url) > 1 {
				err := downloadMaterial(siteName, baseURL, url[1])
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Если попался тег с ссылкой то обрабатываем там атрибуты из получаем ссылку из атрибута
func processNode(node *html.Node, baseURL, siteName string) error {
	if node.Data == "link" {
		err := processAttr(node, baseURL, siteName, "href")
		if err != nil {
			return err
		}
	} else if node.Data == "script" || node.Data == "source" || node.Data == "img" {
		err := processAttr(node, baseURL, siteName, "src")
		if err != nil {
			return err
		}
	}

	return nil
}

// Проходим по всем атрибуам ноды
func processAttr(node *html.Node, baseURL, siteName, key string) error {
	for _, v := range node.Attr {
		if v.Key == key {
			val, _ := url.Parse(v.Val)
			if len(val.Host) == 0 || len(val.Scheme) == 0 {
				err := downloadMaterial(siteName, baseURL, val.Path)
				if err != nil {
					return err
				}
			} else {
				fullURL := concatStrings(val.Scheme, "://", val.Host)
				err := downloadMaterial(siteName, fullURL, val.Path)
				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}

// Создаем директории и сохраняем туда данные
func downloadMaterial(siteName, baseURL, val string) error {
	resp, err := http.Get(concatStrings(baseURL, val))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = os.MkdirAll(getPathDirToFile(siteName, val), 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "не удалось создать каталог")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(concatStrings(siteName, val), data, 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "не удалось сохранить файл")
	}

	return nil
}

// Получить все директории до файла
func getPathDirToFile(siteName, val string) string {
	dir := strings.Split(val, "/")
	return concatStrings(siteName, strings.Join(dir[:len(dir)-1], "/"), "/")
}
