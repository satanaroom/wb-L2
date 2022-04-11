package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	argumentError = "error argument occurred: please, use: go-grep [option] template [fileName]"
	flagError     = `error flag occurred: the following flags are supported:
	-A — печатать +N строк после совпадения
	-B — печатать +N строк до совпадения
	-C — печатать ±N строк вокруг совпадения
	-c — количество строк
	-i — игнорировать регистр
	-v — вместо совпадения, исключать
	-F — точное совпадение со строкой
	-n — напечатать номер строки`
	AFlag = 'A'
	BFlag = 'B'
	CFlag = 'C'
	cFlag = 'c'
	iFlag = 'i'
	vFlag = 'v'
	FFlag = 'F'
	nFlag = 'n'
)

func executeNoFlag(grep *Grep) {
	var keys []int
	for k := range grep.content {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if strings.Contains(grep.content[k], grep.template) {
			fmt.Println(grep.content[k])
		}
	}
}

func executeAFlag(grep *Grep) {
	wasPrinted := make(map[int]bool)
	var keys []int
	for k := range grep.content {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if strings.Contains(grep.content[k], grep.template) {
			lines := 0
			for lines < grep.aNum+1 {
				if !wasPrinted[k+lines] {
					if grep.content[k+lines] != "" {
						fmt.Println(grep.content[k+lines])
					}
				}
				wasPrinted[k+lines] = true
				lines++
			}
		}
	}
}

func executeBFlag(grep *Grep) {
	wasPrinted := make(map[int]bool)
	var keys []int
	for k := range grep.content {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if strings.Contains(grep.content[k], grep.template) {
			lines := grep.bNum
			for lines != -1 {
				if !wasPrinted[k-lines] {
					if grep.content[k-lines] != "" {
						fmt.Println(grep.content[k-lines])
					}
				}
				wasPrinted[k-lines] = true
				lines--
			}
		}
	}
}

func executeCFlag(grep *Grep) {
	wasPrinted := make(map[int]bool)
	var keys []int
	for k := range grep.content {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if strings.Contains(grep.content[k], grep.template) {
			linesBefore := grep.cNum
			for linesBefore != -1 {
				if !wasPrinted[k-linesBefore] {
					if grep.content[k-linesBefore] != "" {
						fmt.Println(grep.content[k-linesBefore])
					}
				}
				wasPrinted[k-linesBefore] = true
				linesBefore--
			}
			linesAfter := 0
			for linesAfter < grep.cNum+1 {
				if !wasPrinted[k+linesAfter] {
					if grep.content[k+linesAfter] != "" {
						fmt.Println(grep.content[k+linesAfter])
					}
				}
				wasPrinted[k+linesAfter] = true
				linesAfter++
			}
		}
	}
}

func executecFlag(grep *Grep) {
	contains := 0
	for _, val := range grep.content {
		if strings.Contains(val, grep.template) {
			contains++
		}
	}
	fmt.Println(contains)
}

func executeiFlag(grep *Grep) {
	var keys []int
	for k := range grep.content {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if strings.Contains(strings.ToLower(grep.content[k]), grep.template) {
			fmt.Println(grep.content[k])
		}
	}
}

func executevFlag(grep *Grep) {
	var keys []int
	for k := range grep.content {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if !strings.Contains(grep.content[k], grep.template) {
			fmt.Println(grep.content[k])
		}
	}
}

func executeFFlag(grep *Grep) {
	executeNoFlag(grep)
}

func executenFlag(grep *Grep) {
	for i, val := range grep.content {
		if strings.Contains(val, grep.template) {
			fmt.Printf("%d:%s\n", i+1, val)
		}
	}
}

func executeCommands(grep *Grep) {
	var flag rune
	if len(grep.flags) != 0 {
		flag = grep.flags[1]
	}
	switch flag {
	default:
		executeNoFlag(grep)
	case AFlag:
		executeAFlag(grep)
	case BFlag:
		executeBFlag(grep)
	case CFlag:
		executeCFlag(grep)
	case cFlag:
		executecFlag(grep)
	case iFlag:
		executeiFlag(grep)
	case vFlag:
		executevFlag(grep)
	case FFlag:
		executeFFlag(grep)
	case nFlag:
		executenFlag(grep)
	}
}

func createContent(grep *Grep) error {
	file, err := os.Open(grep.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	grep.lineCount = 0
	for sc.Scan() {
		grep.content[grep.lineCount] = sc.Text()
		grep.lineCount++
	}
	return nil
}

func setNumbers(grep *Grep) {
	for i, val := range grep.flags {
		if val >= 48 && val <= 57 {
			prev := grep.flags[i-1]
			if prev == 'A' {
				grep.aNum = int(val - '0')
			} else if prev == 'B' {
				grep.bNum = int(val - '0')
			} else if prev == 'C' {
				grep.cNum = int(val - '0')
			}
		}
	}
}

func parseFlags(flags []rune) bool {
	if len(flags) == 0 {
		return true
	}
	if len(flags) >= 2 {
		if flags[0] == '-' {
			for i := 1; i < len(flags); i++ {
				symb := flags[i]
				if symb != 'A' && symb != 'B' && symb != 'C' && symb != 'c' &&
					symb != 'i' && symb != 'v' && symb != 'F' && symb != 'n' &&
					!unicode.IsDigit(symb) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func parseArgs(grep *Grep) bool {
	if len(os.Args) == 3 {
		grep.template = os.Args[1]
		grep.fileName = os.Args[2]
	} else if len(os.Args) == 4 {
		grep.flags = []rune(os.Args[1])
		grep.template = os.Args[2]
		grep.fileName = os.Args[3]
	} else {
		return false
	}
	return true
}

type Grep struct {
	flags     []rune
	template  string
	fileName  string
	aNum      int
	bNum      int
	cNum      int
	content   map[int]string
	lineCount int
}

func main() {
	grep := &Grep{
		content: map[int]string{},
	}
	ok := parseArgs(grep)
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", argumentError)
		os.Exit(1)
	}
	ok = parseFlags(grep.flags)
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", flagError)
		os.Exit(1)
	}
	setNumbers(grep)
	err := createContent(grep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	executeCommands(grep)
}
