package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Паттерн Строитель является пораждающим паттерном проектирования. Он позволяет создавать сложные объекты
// используя шаги. На каждом шаге производится какая-то часть общего объекта. Тем самым, выполняя все шаги по очереди,
// формируется некий объект, представляющий сложную структуру. Строитель позволяет использовать один и тот же код
// строительства объекта для получения разных представлений этого объекта.

// Плюсы:
// 1. Позволяет пошагово создавать общий объект, который зависит от составляющих частей;
// 2. Позволяет использовать один и тот же код для создания различных объектов;
// 3. Изолирует сложный код при сборке объекта и его бизнес-логики.

// Минусы:
// 1. Усложняет код программы из-за введения дополнительных классов (структур, интерфейсов).
// 2. Привязка клиента к конкретному объекту строителя, т.к. в интерфейсе может не быть какого-то метода,
// поэтому будет необходимо его добавить.

// +++ Инициализация сборщика
const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}

// ---

// +++ Инициализация структуры компьютера
type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (pc *Computer) Print() {
	fmt.Printf("%s Core:[%d] Memory:[%d] GraphicCard:[%d] Monitor:[%d]\n",
		pc.Brand, pc.Core, pc.Memory, pc.GraphicCard, pc.Monitor)
}

// ---

// +++ Инициализация сборщика одного бренда
type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

// ---

// +++ Инициализация сборщика другого бренда
type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 4
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

// ---

// +++ Инициализация завода по производству компьтеров
type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

// Выполнение последовательных шагов сборки
func (factory Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()
	return factory.Collector.GetComputer()
}

// ---

/*
func main() {
	// Определение возможных комплектаций
	asusCollector := GetCollector("asus")
	hpCollector := GetCollector("hp")

	// Создание базового завода с комплектацией по умолчанию
	factory := NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	// Смена комплектации
	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	factory.SetCollector(asusCollector)
	pc := factory.CreateComputer()
	pc.Print()
}
*/
