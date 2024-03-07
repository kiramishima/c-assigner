package interfaces

import "net/http"

type CreditHandlers interface {
	CreateAssignHandler(w http.ResponseWriter, req *http.Request)
	GetStatistics(w http.ResponseWriter, req *http.Request)
}
