package pattern

import (
	"errors"
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Паттерн Фасад является структурным паттерном проектирования. Представляет собой простой интерфейс к сложной системе.
// Если имеется много разных подсистем, которые используют свои интерфейсы и реализуют свой функционал поведения,
// следуюет применить паттерн Фасад, чтобы создать простой интерфейс для максимально простого взаимодействия с подсистемами.

// Плюсы:
// 1. Изолирует клиентов от поведения сложной системы

// Минусы:
// 1. Интерфейс Фасада может стать супер-объектом (супер-классом).
// Другими словами, все последующие функции будут проходить через этот объект.

// Описание товара
type Product struct {
	Name  string
	Price float64
}

// +===+ Фасад над сложной бизнес логикой покупки и оплате товаров по безналичному рассчету
// Описание магазина - основного объекта Фасада
type Shop struct {
	Name     string
	Products []Product
}

// Метод продажи магазином товара
func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка: может ли пользователь %s купить товар\n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] У пользователя недостаточно средств для покупки товара")
		}
		fmt.Printf("[Магазин] Товар %s куплен\n", prod.Name)
	}
	return nil
}
// -===-

// Описание платежной системы
type Bank struct {
	Name  string
	Cards []Card
}

// Метод проверки баланса карты банком
func (bank Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("[Банк] Получение остатка по карте %s\n", cardNumber)
	time.Sleep(time.Millisecond * 300)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		}
	}
	fmt.Println("[Банк] Остаток положительный")
	return nil
}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

// Метод проверки баланса карты
func (card Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return card.Bank.CheckBalance(card.Name)
}

// Описание пользователя (покупателя)
type User struct {
	Name string
	Card *Card
}

// Метод получения баланса пользователя
func (user User) GetBalance() float64 {
	return user.Card.Balance
}

/*
var (
	bank = Bank{
		Name:  "Sber",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "CARD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CARD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = User{
		Name: "Ivan",
		Card: &card1,
	}
	user2 = User{
		Name: "Igor",
		Card: &card2,
	}
	prod = Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = Shop{
		Name: "Vkusvill",
		Products: []Product{
			prod,
		},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s] Приходит в магазин\n", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[%s] Приходит в магазин\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
*/
