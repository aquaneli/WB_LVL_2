package main

import (
	"testing"
)

func Test(t *testing.T) {
	baseURL := "https://tech.wildberries.ru/"
	node, siteName, err := initDownloadFromHTML(baseURL)
	if err != nil {
		t.Error("Incorrect parsing")
	}

	err = downloadSite(node, baseURL, siteName)
	if err != nil {
		t.Error("incorrect processing of download links")
	}

}

func TestProcessNode(t *testing.T) {
	url, baseURL, err := initDownloadFromHTML("https://exampleqwe.com")
	if url != nil || baseURL != "" || err == nil {
		t.Error("Incorrect parsing")
	}
}
