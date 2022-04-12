Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:

```bash
fatal error: all goroutines are asleep - deadlock!
```
Так как в канал никто не пишет, а мы ожидаем значение из него, чтобы напечатать, произойдет deadlock. Для решения проблемы необходимо закрыть канал после цикла:
```go
go func() {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}()
```
