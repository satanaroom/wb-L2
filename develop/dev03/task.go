package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	argumentErrorNoFlags = "error argument occurred: please, use: go run task.go fileName"
	argumentError        = "error argument occurred: please, use: go run task.go -flag fileName"
	flagError            = "error flag occurred"
	kFlag                = "-k"
	nFlag                = "-n"
	rFlag                = "-r"
	uFlag                = "-u"
	MFlag                = "-M"
	bFlag                = "-b"
	cFlag                = "-c"
	hFlag                = "-h"
)

// +++ Функции обработки флагов
func parseFlag(flag string) bool {
	runes := []rune(flag)
	if len(runes) == 2 {
		if runes[0] == '-' {
			symb := runes[1]
			if symb == 'k' || symb == 'n' || symb == 'r' || symb == 'u' ||
				symb == 'M' || symb == 'b' || symb == 'c' || symb == 'h' {
				return true
			}
		}
	}
	fmt.Fprintln(os.Stderr, flagError)
	fmt.Println(`the following flags are supported:
	-k — указание колонки для сортировки
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки
	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учётом суффиксов`)
	return false
}

func chooseAFlag(flag string, content []byte) {
	switch flag {
	case kFlag:
		parseK(content)
	case nFlag:
		parseN(content)
	case rFlag:
		parseR(content)
	case uFlag:
		parseU(content)
	case MFlag:
		parseM(content)
	case bFlag:
		parseB(content)
	case cFlag:
		parseC(content)
	case hFlag:
		parseH(content)
	}
}

func parseK(content []byte) {
	fmt.Println(content)
}

func parseN(content []byte) {
	fmt.Println(content)
}

func parseR(content []byte) {
	fmt.Println(content)
}

func parseU(content []byte) {
	fmt.Println(content)
}

func parseM(content []byte) {
	fmt.Println(content)
}

func parseB(content []byte) {
	fmt.Println(content)
}

func parseC(content []byte) {
	fmt.Println(content)
}

func parseH(content []byte) {
	fmt.Println(content)
}

// ---

// ++= Функции обработки без флагов
func parseNoFlags(content []byte) {
	str := string(content)
	words := strings.Split(str, "\n")
	words = words[:len(words)-1]
	sort.Strings(words)
	for _, val := range words {
		fmt.Println(val)
	}
}

func checkNoFlag(fileName string) ([]byte, bool) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return nil, false
	}
	return f, true
}

// ---

func main() {
	if len(os.Args) == 2 {
		fileName := os.Args[1]
		content, ok := checkNoFlag(fileName)
		if !ok {
			os.Exit(1)
		}
		parseNoFlags(content)
	} else if len(os.Args) == 3 {
		flag := os.Args[1]
		ok := parseFlag(flag)
		if !ok {
			os.Exit(1)
		}
		file := os.Args[2]
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		chooseAFlag(flag, content)
	}
	os.Exit(0)
}
