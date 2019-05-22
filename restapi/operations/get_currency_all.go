// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCurrencyAllHandlerFunc turns a function with the right signature into a get currency all handler
type GetCurrencyAllHandlerFunc func(GetCurrencyAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCurrencyAllHandlerFunc) Handle(params GetCurrencyAllParams) middleware.Responder {
	return fn(params)
}

// GetCurrencyAllHandler interface for that can handle valid get currency all params
type GetCurrencyAllHandler interface {
	Handle(GetCurrencyAllParams) middleware.Responder
}

// NewGetCurrencyAll creates a new http.Handler for the get currency all operation
func NewGetCurrencyAll(ctx *middleware.Context, handler GetCurrencyAllHandler) *GetCurrencyAll {
	return &GetCurrencyAll{Context: ctx, Handler: handler}
}

/*GetCurrencyAll swagger:route GET /currency/all getCurrencyAll

Retrieve information for every currency

*/
type GetCurrencyAll struct {
	Context *middleware.Context
	Handler GetCurrencyAllHandler
}

func (o *GetCurrencyAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCurrencyAllParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}