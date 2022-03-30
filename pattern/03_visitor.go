package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Паттерн Посетитель является поведенческим паттерном уровня объекта. Он позволяет обойти
// набор элементов (объектов) с разнородными интерфейсами, а также позволяет добавить
// новый метод в тип объекта, при этом, не изменяя сам тип этого объекта.
// Решает задачу определения новой операции, не изменяя типы объектов, над которыми выполняется одна или более операций.
// Паттерн следует применять если:
// 1. Имеются различные объекты разных типов с разными интерфейсами, но над ними нужно совершать операции, зависящие от конкретных типов;
// 2. Необходимо над структурой выполнить различные, усложняющие структуру операции;
// 3. Часто добавляются новые операции над структурой.

// Плюсы:
// 1. Упрощается добавление новых операций;
// 2. Объединение родственных операции Посетителе;
// 3. Посетитель может запоминать в себе какое-то состояние по мере обхода контейнера.

// Минусы:
// 1. Затруднено добавление новых типов, поскольку нужно обновлять иерархию Посетителя и его сыновей.

// Инициализация Посетителя
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Инициализация места, которое должен посетить Посетитель
type Place interface {
	Accept(v Visitor) string
}

// +++ Реализация посетителя
type People struct {
}

// Реализация посещения суши бара
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// Реализация посещения пиццерии
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// Реализация посещения бургерной
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// ---

// Реализация коллекции мест для посещения
type City struct {
	places []Place
}

// Добавления места в коллекцию
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Реализация посещения всех мест в городе
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// +++ Реализация интерфейса места суши баром
type SushiBar struct {
}

func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// ---

// +++ Реализация интерфейса места пиццерией
type Pizzeria struct {
}

func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// ---

// +++ Реализация интерфейса места бургерной
type BurgerBar struct {
}

func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

// ---

/*
func main() {
	expect := "Buy sushi...Buy pizza...Buy burger..."

	city := new(City)

	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	result := city.Accept(&People{})

	if result != expect {
		fmt.Printf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
*/
