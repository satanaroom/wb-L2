package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const argumentError = `error argument occurred: utility has two options:
	download file: go run task.go -m UrlToFile
	download site: go run task.go Url`

type Wget struct {
	Url    string
	isFile bool
}

func (w Wget) createFile(p []byte) error {
	var fileName string
	if w.isFile {
		fileName = w.Url[strings.LastIndex(w.Url, "/")+1:]
	} else {
		fileName = "index.html"
	}
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		return err
	}
	ioutil.WriteFile(f.Name(), p, 0777)
	return nil
}

func (w Wget) getPage(client *http.Client) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, w.Url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %d", resp.StatusCode)
	}
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (w Wget) checkUrl() error {
	_, err := url.ParseRequestURI(w.Url)
	if err != nil {
		return err
	}
	return nil
}

func parseArgs(wget *Wget) bool {
	if len(os.Args) == 2 {
		wget.Url = os.Args[1]
		return true
	}
	if len(os.Args) == 3 {
		if os.Args[1] == "-m" {
			wget.isFile = true
			wget.Url = os.Args[2]
			return true
		}
	}
	return false
}

func main() {
	w := &Wget{}
	ok := parseArgs(w)
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", argumentError)
		os.Exit(1)
	}
	err := w.checkUrl()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	client := &http.Client{}
	p, err := w.getPage(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	err = w.createFile(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
