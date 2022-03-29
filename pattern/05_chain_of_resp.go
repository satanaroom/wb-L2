package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Паттерн Цепочка вызовов является поведенческим паттерном. Он позволяет передавать выполнение
// запросов последовательно по цепочке. Вызывает обработчик после того, как вызовет этот процесс
// предыдущий обработчик. В цепочке может быть много звеньев, которую выполняют определенную задачу.
// Пользователь -> Запрос в сервис (одно звено) -> Авторизация (другое звено) + сложная бизнес логика

// Плюсы:
// 1. Уменьшает зависимость между клиентом и обработчиками. При этом, изменять эту логику можно
// независимо от того, является ли обработка общей;
// 2. Реализует принцип single responsibility;
// 3. Реализует принцип open–closed (программные сущности (классы, модули, функции и т. п.) должны быть открыты для расширения, но закрыты для изменения).

// Минусы:
// 1. Запрос может остаться неотработанным в случае нарушения логики.

// +++ Инициализация структуры цепочки вызовов
type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	// Выполнился ли прием данных
	GetSource bool
	// Выполнилась ли обработка данных
	UpdateSource bool
}

// ---

// +++ Инициализация сервиса получения данных
type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data from device [%s] has already been received.\n", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s].\n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *Device) SetNext(svc Service) {
	device.Next = svc
}

// ---

// +++ Инициализация сервиса обработки данных
type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data from device [%s] has already been updated.\n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from device [%s].\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc
}

// ---

// +++ Инициализация сервиса сохранения обработанных данных
type DataService struct {
	Next Service
}

func (ds *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("Data not updated")
		return
	}
	fmt.Println("Data saved.")
}

func (ds *DataService) SetNext(svc Service) {
	ds.Next = svc
}

// ---

/*
func main() {
	// Инициализируем сервисы
	device := &Device{Name: "Device-1"}
	updateSvc := &UpdateDataService{Name: "Update-1"}
	dataSvc := &DataService{}
	// Инициализируем цепочку вызовов
	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)
	data := &Data{}
	device.Execute(data)
}
*/
