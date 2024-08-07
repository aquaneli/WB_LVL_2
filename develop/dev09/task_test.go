package main

import (
	"testing"
)

func Test(t *testing.T) {
	baseUrl := "https://tech.wildberries.ru/"
	node, siteName, err := initDownloadFromHtml(baseUrl)
	if err != nil {
		t.Error("Incorrect parsing")
	}

	err = downloadSite(node, baseUrl, siteName)
	if err != nil {
		t.Error("incorrect processing of download links")
	}

}

func TestProcessNode(t *testing.T) {
	url, baseUrl, err := initDownloadFromHtml("https://exampleqwe.com")
	if url != nil || baseUrl != "" || err == nil {
		t.Error("Incorrect parsing")
	}
}
