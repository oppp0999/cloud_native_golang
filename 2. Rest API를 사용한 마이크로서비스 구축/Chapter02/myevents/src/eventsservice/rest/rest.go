package rest

import (
	"net/http"

	"github.com/PacktPublishing/Cloud-Native-programming-with-Golang/chapter02/myevents/src/lib/persistence"

	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler) //업무 1
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler) //업무 2
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler) //업무 3
	return http.ListenAndServe(endpoint, r)
}
