package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/scottshotgg/toyota-test/restapi/operations"
)

// GetSymbol handles retrieving data to a single endpoint
func GetSymbol(params operations.GetCurrencySymbolParams) middleware.Responder {
	fmt.Println("serving bonus")

	return operations.NewGetCurrencySymbolOK().WithPayload(nil)
}

// GetAll handles retrieving data for every currency
func GetAll(params operations.GetCurrencyAllParams) middleware.Responder {
	fmt.Println("serving bonus")

	return operations.NewGetCurrencyAllOK().WithPayload(nil)
}
