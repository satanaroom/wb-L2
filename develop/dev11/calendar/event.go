package calendar

import "time"

/*
	HTTP-запросы
		|
	Handler
		|
	Service (бизнес-логика)
		|
	Repository (работа с БД)
*/

type Event struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Date        string    `json:"date"`
	ParsedDate  time.Time `db:"date"`
	MinDate     time.Time
	MaxDate     time.Time
	Done        bool `json:"done" db:"done"`
}

type ResultEvent struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Date        string `json:"date"`
	Done        bool   `json:"done" db:"done"`
}
