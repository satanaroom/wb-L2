package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	commandError = `error command occurred: the following commands are supported:
	$ cd   — смена директории
	$ pwd  — вывод текущей директории
	$ echo — вывод строки текста
	$ kill — убить процесс
	$ ps   — вывод списка процессов`
	ShellToUse    = "bash"
	MyShellStart  = `======================Welcome to MyShell 1.0======================`
	MyShellFinish = `==================================================================`
)

func Shellout(command string, cd *Cd) (error, string) {
	var stdout bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Dir = cd.dirName
	cmd.Stdout = &stdout
	err := cmd.Run()
	return err, stdout.String()
}

func parseArgs(cmd string, cd *Cd) {
	var wasCd bool
	var res []string
	res = strings.Split(cmd, " ")
	for _, val := range res {
		if val != "cd" && !wasCd {
			continue
		}
		if !wasCd {
			wasCd = true
			continue
		}
		fileInfo, _ := os.Stat(val)
		if fileInfo.IsDir() {
			cd.dirName = val
		}
		break
	}
}

type Cd struct {
	dirName string
}

func main() {
	cd := &Cd{}
	sc := bufio.NewScanner(os.Stdin)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	fmt.Println(MyShellStart)
	go func() {
		for {
			if sc.Scan() {
				cmd := sc.Text()
				parseArgs(cmd, cd)
				err, out := Shellout(cmd, cd)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}
				fmt.Printf("%s", out)
			}
		}
	}()
	select {
	case <-sigChan:
		fmt.Println(MyShellFinish)
		os.Exit(0)
	}
}
