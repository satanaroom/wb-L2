Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:

```bash
error
```
Тип переменной err после возврата из функции - *main.customError, поэтому при её сравнении с nil будет ложь.

Можно проверить, при помощи спецификаторов:
- %v	значение в формате по умолчанию
- %#v	Go-синтаксическое представление значения
- %T	Go-синтаксическое представление типа значения
