package main

import (
	"sort"
	"strings"
	"unicode/utf8"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func SearchAnagram(arr []string) map[string][]string {
	m := make(map[string][]string)

	for _, item := range arr {
		var anagrams []string
		item = strings.ToLower(strings.Trim(item, " "))

		for _, val := range arr {
			val = strings.ToLower(strings.Trim(val, " "))

			if utf8.RuneCountInString(item) == utf8.RuneCountInString(val) {
				hash := make(map[string]int)

				for _, r := range item {
					j := hash[string(r)]

					if j == 0 {
						hash[string(r)] = 1
					} else {
						hash[string(r)] = j + 1
					}
				}

				for _, r := range val {
					j := hash[string(r)]

					if j == 0 {
						hash[string(r)] = 1
					} else {
						hash[string(r)] = j + 1
					}
				}

				var isAnagram bool = true
				for _, value := range hash {
					if value%2 != 0 {
						isAnagram = false
					}
				}

				if isAnagram {
					anagrams = append(anagrams, val)
					if len(anagrams) >= 2 {
						m[anagrams[0]] = anagrams[1:]
					}
				}
			}
		}
	}
	for val := range m {
		sort.Strings(m[val])
	}

	return m
}
