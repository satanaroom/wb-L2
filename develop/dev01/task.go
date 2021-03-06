package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	// Для получения текущего времени, а также некоторых дополнительных
	// метаданных, используется функция запроса.
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	// Response выполняет дополнительные проверки работоспособности,
	// чтобы определить, подходит ли ответ для целей синхронизации времени.
	err = response.Validate()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	// ClockOffset - предполагаемое смещение локальных системных часов относительно часов сервера.
	// Используется для более точного считывания времени.
	time := time.Now().Add(response.ClockOffset)
	fmt.Printf("%s", time)
}
