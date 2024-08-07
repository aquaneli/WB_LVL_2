package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	baseUrl := "https://dev.to/dave3130/golang-html-tokenizer-5fh7"
	node, siteName, err := initDownloadFromHtml(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = downloadSite(node, baseUrl, siteName)
	if err != nil {
		log.Fatal(err)
	}
}

// конкатенация нескольких строк
func concatStrings(str ...string) string {
	sb := strings.Builder{}
	for _, val := range str {
		sb.WriteString(val)
	}
	return sb.String()
}

func initDownloadFromHtml(baseUrl string) (*html.Node, string, error) {
	resp, err := http.Get(baseUrl)
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
	os.WriteFile(concatStrings(resp.TLS.ServerName, "/index.html"), data, 0777)

	bufReader := bytes.NewReader(buf.Bytes())
	node, err := html.Parse(bufReader)
	if err != nil {
		return nil, "", err
	}

	return node, resp.TLS.ServerName, nil
}

// обработка всех узлов html страницы рекурсивно
func downloadSite(node *html.Node, baseUrl, siteName string) error {

	// if node.Type == html.ElementNode && node.Data == "style" {
	// 	if node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
	// 		processStyle(node.FirstChild, baseUrl, siteName)
	// 	}
	// }

	if node.Type == html.ElementNode {
		err := processNode(node, baseUrl, siteName)
		if err != nil {
			return err
		}
	}

	for itr := node.FirstChild; itr != nil; itr = itr.NextSibling {
		err := downloadSite(itr, baseUrl, siteName)
		if err != nil {
			return err
		}
	}
	return nil
}

func processStyle(node *html.Node, baseUrl, siteName string) error {
	attr := strings.Split(node.Data, " ")
	fmt.Println(attr)
	for _, v := range attr {
		if len(v) > 3 && v[:3] == "url" {
			url := strings.Split(v[3:], "\"")
			if len(url) > 1 {
				err := downloadMaterial(siteName, baseUrl, url[1])
				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}

// если попался тег с ссылкой то обрабатываем там атрибуты из получаем ссылку из атрибута
func processNode(node *html.Node, baseUrl, siteName string) error {
	if node.Data == "link" {
		err := processAttr(node, baseUrl, siteName, "href")
		if err != nil {
			return err
		}
	} else if node.Data == "script" || node.Data == "source" || node.Data == "img" {
		err := processAttr(node, baseUrl, siteName, "src")
		if err != nil {
			return err
		}
	} else if node.Data == "div" {
		for _, v := range node.Attr {
			if v.Key == "style" {
				res := strings.Split(v.Val, ":")
				if len(res) > 1 {
					if res[1][:3] == "url" {
						processStyle(&html.Node{Data: res[1]}, baseUrl, siteName)
					}

				}
			}
		}
	}
	return nil
}

func processAttr(node *html.Node, baseUrl, siteName, key string) error {
	for _, v := range node.Attr {
		if v.Key == key {
			err := downloadMaterial(siteName, baseUrl, v.Val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// создаем директории и сохраняем туда данные
func downloadMaterial(siteName, baseUrl, val string) error {
	resp, err := http.Get(concatStrings(baseUrl, val[1:]))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = os.MkdirAll(getPathDirToFile(siteName, val), 0777)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	os.WriteFile(concatStrings(siteName, val), data, 0777)

	return nil
}

// Получить все директории до файла
func getPathDirToFile(siteName, val string) string {
	dir := strings.Split(val, "/")
	return concatStrings(siteName, strings.Join(dir[:len(dir)-1], "/"), "/")
}
