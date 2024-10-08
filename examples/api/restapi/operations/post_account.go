// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostAccountHandlerFunc turns a function with the right signature into a post account handler
type PostAccountHandlerFunc func(PostAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAccountHandlerFunc) Handle(params PostAccountParams) middleware.Responder {
	return fn(params)
}

// PostAccountHandler interface for that can handle valid post account params
type PostAccountHandler interface {
	Handle(PostAccountParams) middleware.Responder
}

// NewPostAccount creates a new http.Handler for the post account operation
func NewPostAccount(ctx *middleware.Context, handler PostAccountHandler) *PostAccount {
	return &PostAccount{Context: ctx, Handler: handler}
}

/*
	PostAccount swagger:route POST /account postAccount

# Create a new account service한글도 되긴함

Create a new account with...
*/
type PostAccount struct {
	Context *middleware.Context
	Handler PostAccountHandler
}

func (o *PostAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAccountParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
