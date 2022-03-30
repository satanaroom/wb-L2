package pattern

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Паттерн Комманда является поведенческим паттерном уровня объекта. Он позволяет
// представить запрос в виде объекта. Из этого следует, что команда - это объект.
// Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
// Command - запрос в виде объекта на выполнение;
// Receiver - объект-получатель запроса, который будет обрабатывать нашу команду;
// Invoker* - объект-инициатор запроса.
// *Invoker умеет складывать команды в стопку и инициировать их выполнение по какому-то событию.
//  Обратившись к Invoker можно отменить команду, пока та не выполнена.
// Комманда отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
// Единственное, что должен знать инициатор, это как отправить команду.
// Паттерн следует применять в следующих случаях:
// 1. Кнопки пользовательского интерфейса и пункты меню;
// 2. Запись макросов;
// 3. Многоуровневая отмена операций (Undo);
// 4. Индикаторы выполнения;
// 5. Транзакции.

// Инициализация Команды
type Command interface {
	Execute() string
}

// +++ Реализация Команды ToggleOnCommand
type ToggleOnCommand struct {
	receiver *Receiver
}

func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ---

// +++ Реализация Команды ToggleOffCommand
type ToggleOffCommand struct {
	receiver *Receiver
}

func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// +++

// +++ Реализация отправителя
type Receiver struct {
}

func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// ---

// +++ Реализация инвокера (вызывающего)
type Invoker struct {
	commands []Command
}

// Команда сохранения
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// Команда удаления
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Выполнение всех команд
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

/*
func main() {
	expect := "Toggle On\n" +
		"Toggle Off\n"

	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})

	result := invoker.Execute()

	if result != expect {
		fmt.Printf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
*/
