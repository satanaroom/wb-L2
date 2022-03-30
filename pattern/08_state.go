package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Паттерн Состояние является поведенческим паттерном уровня объекта. Он позволяет
// позволяет объекту изменять свое поведение в зависимости от внутреннего состояния
// и является объектно-ориентированной реализацией конечного автомата.
// Поведение объекта изменяется настолько, что создается впечатление, будто изменился тип объекта.
// Паттерн должен применяться:
// 1. Когда поведение объекта зависит от его состояния;
// 2. Поведение объекта должно изменяться во время выполнения программы;
// 3. Состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно.

// Инициализация общего интерфейса для различных состояний
type MobileAlertStater interface {
	Alert() string
}

// +++ Реализация оповещения, в зависимости от своего состояния
type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// ---

// Инициализация конструктора оповещения
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// ---

// +++ Реализация оповещения вибрацией
type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// ---

// +++ Реализация оповещения о звуковым сигналом
type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

// ---

/*
func main() {
	expect := "Vrrr... Brrr... Vrrr..." +
		"Vrrr... Brrr... Vrrr..." +
		"Белые розы, Белые розы. Беззащитны шипы..."

	mobile := NewMobileAlert()

	result := mobile.Alert()
	result += mobile.Alert()

	mobile.SetState(&MobileAlertSong{})

	result += mobile.Alert()

	if result != expect {
		fmt.Printf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
*/
