package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", CreateEvent)
	mux.HandleFunc("/update_event", UpdateEvent)
	mux.HandleFunc("/delete_event", DeleteEvent)
	mux.HandleFunc("/events_for_day", EventsForDay)
	mux.HandleFunc("/events_for_week", EventsForWeek)
	mux.HandleFunc("/events_for_month", EventsForMonth)

	return mux
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func main() {
	srv := new(Server)

	go func() {
		log.Fatalln(srv.Run("8080", InitRoutes()))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func CreateEvent(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}

func UpdateEvent(res http.ResponseWriter, req *http.Request) {

}

func DeleteEvent(res http.ResponseWriter, req *http.Request) {

}

func EventsForDay(res http.ResponseWriter, req *http.Request) {

}

func EventsForWeek(res http.ResponseWriter, req *http.Request) {

}

func EventsForMonth(res http.ResponseWriter, req *http.Request) {

}
