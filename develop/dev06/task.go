package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	flagError = `error flag occurred: the following flags are supported:
	-f — выбрать поля (колонки): STDIN | go-cut -f 2
	-d — использовать другой разделитель: STDIN | go-cut -d ' '
	-s — только строки с разделителем: STDIN | go-cut -s -f 3 -d ':'`
	fFlag = "-f"
	dFlag = "-d"
	sFlag = "-s"
)

func createContent(cut *Cut) {
	var tmp []string
	var res []string
	var fields []int
	for _, val := range cut.fields {
		num, _ := strconv.Atoi(val)
		fields = append(fields, num)
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		tmp = append(tmp, txt)
	}
	if !cut.needChange {
		cut.content = tmp
	} else {
		for _, val := range tmp {
			if strings.Contains(val, string(cut.delimiter)) {
				res = strings.Split(val, string(cut.delimiter))
				canTake := len(res)
				for _, val := range fields {
					if val > canTake {
						fmt.Println()
					} else {
						fmt.Println(res[val-1])
					}
				}
			} else {
				if !cut.onlyWithDelimiter {
					fmt.Println(val)
				}
			}
		}
	}
}

func parseFlags(cut *Cut) bool {
	if len(os.Args) != 1 {
		for i := 1; i < len(os.Args); i++ {
			runes := []rune(os.Args[i])
			if os.Args[i] == fFlag {
				continue
			}
			for _, r := range runes {
				if unicode.IsDigit(rune(r)) {
					if strings.Contains(os.Args[i], ",") {
						cut.fields = strings.Split(os.Args[i], ",")
						break
					} else if strings.Contains(os.Args[i], "-") {
						cut.fields = strings.Split(os.Args[i], "-")
						break
					} else {
						cut.fields = append(cut.fields, os.Args[i])
						break
					}
				}
			}
			if os.Args[i] == dFlag {
				cut.needChange = true
				continue
			}
			if os.Args[i] == sFlag {
				cut.onlyWithDelimiter = true
			}
			if len(runes) == 1 && cut.needChange {
				for _, r := range runes {
					cut.delimiter = r
				}
			} else if len(runes) != 1 && cut.needChange {
				return false
			}
		}
		if len(cut.fields) == 0 || (cut.needChange && unicode.IsDigit(cut.delimiter)) {
			return false
		}
	} else {
		return false
	}
	return true
}

type Cut struct {
	content           []string
	fields            []string
	delimiter         rune
	needChange        bool
	onlyWithDelimiter bool
}

func main() {
	cut := &Cut{}
	ok := parseFlags(cut)
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", flagError)
		os.Exit(1)
	}
	createContent(cut)
	os.Exit(0)
}
