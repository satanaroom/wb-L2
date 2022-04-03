package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
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

func CreateLines(content []byte) []string {
	str := string(content)
	lines := strings.Split(str, "\n")
	lines = lines[:len(lines)-1]
	return lines
}

// +++ Функции обработки флагов
func parseFlag(flag string) bool {
	runes := []rune(flag)
	if len(runes) >= 2 {
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
		sortK(content)
	case nFlag:
		sortN(content)
	case rFlag:
		sortR(content)
	case uFlag:
		sortU(content)
	case MFlag:
		sortM(content)
	case bFlag:
		sortB(content)
	case cFlag:
		sortC(content)
	case hFlag:
		sortH(content)
	}
}

// +++ Сортировка по колонке
type KSort [][]string

var ColumnIndex int

func (data KSort) Less(i, j int) bool {
	columnIndex := ColumnIndex
	word1 := data[i][columnIndex]
	word2 := data[j][columnIndex]

	return word1 < word2
}

func (data KSort) Len() int {
	return len(data)
}
func (data KSort) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func sortK(content []byte) {
	var lines [][]string
	parts := CreateLines(content)
	for _, val := range parts {
		lines = append(lines, strings.Split(val, " "))
	}
	sort.Sort(KSort(lines))
	for _, val := range lines {
		fmt.Println(strings.Join(val, " "))
	}
}

// ---

// +++ Сотрировка по числовому значению
type NSort [][]string

func (data NSort) Less(i, j int) bool {
	columnIndex := ColumnIndex
	num1 := data[i][columnIndex]
	num2 := data[j][columnIndex]
	convNum1, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	convNum2, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	return convNum1 < convNum2
}

func (data NSort) Len() int {
	return len(data)
}
func (data NSort) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}
func sortN(content []byte) {
	var lines [][]string
	parts := CreateLines(content)
	for _, val := range parts {
		lines = append(lines, strings.Split(val, " "))
	}
	sort.Sort(NSort(lines))
	for _, val := range lines {
		fmt.Println(strings.Join(val, " "))
	}
}

// ---

// +++ Сотрировка в обратном порядке
func sortR(content []byte) {
	lines := CreateLines(content)
	sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	for _, val := range lines {
		fmt.Println(val)
	}
}

// ---

func RemoveDuplicates(str []string) []string {
	// Объявление результируещего множества
	var unique []string
	// Создание мапы для определения уникальности значений в результирующем множестве
	attend := make(map[string]bool)
	for _, val := range str {
		// Если в мапе нет элемента, добавляем
		if !attend[val] {
			unique = append(unique, val)
			// Устанавливаем ключ уникальности
			attend[val] = true
		}
	}
	return unique
}

func sortU(content []byte) {
	lines := CreateLines(content)
	lines = RemoveDuplicates(lines)
	for _, val := range lines {
		fmt.Println(val)
	}
}

func sortM(content []byte) {
	fmt.Println(content)
}

func sortB(content []byte) {
	fmt.Println(content)
}

func sortC(content []byte) {
	fmt.Println(content)
}

func sortH(content []byte) {
	fmt.Println(content)
}

// ---

// ++= Функции обработки без флагов
func parseNoFlags(content []byte) {
	words := CreateLines(content)
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
		flags := os.Args[1]
		ok := parseFlag(flags)
		if !ok {
			os.Exit(1)
		}
		file := os.Args[2]
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		indicator := flags[:2]
		number := flags[2:]
		if number != "" {
			num, err := strconv.Atoi(number)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
				os.Exit(1)
			}
			if num == 0 {
				fmt.Fprintf(os.Stderr, "%s\n", "number of column should not zero")
				os.Exit(1)
			}
			ColumnIndex = num
		}
		chooseAFlag(indicator, content)
	}
	os.Exit(0)
}
