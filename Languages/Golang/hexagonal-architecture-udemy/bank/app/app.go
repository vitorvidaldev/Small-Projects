package app

import (
	"log"
	"net/http"

	"dev/vitorvidaldev/banking/domain"
	"dev/vitorvidaldev/banking/service"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
