// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"swaggertest/models"
)

// GetAccountAllHandlerFunc turns a function with the right signature into a get account all handler
type GetAccountAllHandlerFunc func(GetAccountAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountAllHandlerFunc) Handle(params GetAccountAllParams) middleware.Responder {
	return fn(params)
}

// GetAccountAllHandler interface for that can handle valid get account all params
type GetAccountAllHandler interface {
	Handle(GetAccountAllParams) middleware.Responder
}

// NewGetAccountAll creates a new http.Handler for the get account all operation
func NewGetAccountAll(ctx *middleware.Context, handler GetAccountAllHandler) *GetAccountAll {
	return &GetAccountAll{Context: ctx, Handler: handler}
}

/*
	GetAccountAll swagger:route GET /account/all getAccountAll

# Get all of user

Get all of the user infomation.
*/
type GetAccountAll struct {
	Context *middleware.Context
	Handler GetAccountAllHandler
}

func (o *GetAccountAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAccountAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetAccountAllOKBody get account all o k body
//
// swagger:model GetAccountAllOKBody
type GetAccountAllOKBody struct {

	// attr
	Attr []*models.AccountEntry `json:"Attr"`
}

// Validate validates this get account all o k body
func (o *GetAccountAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAccountAllOKBody) validateAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.Attr) { // not required
		return nil
	}

	for i := 0; i < len(o.Attr); i++ {
		if swag.IsZero(o.Attr[i]) { // not required
			continue
		}

		if o.Attr[i] != nil {
			if err := o.Attr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAccountAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getAccountAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get account all o k body based on the context it is used
func (o *GetAccountAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAccountAllOKBody) contextValidateAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Attr); i++ {

		if o.Attr[i] != nil {
			if err := o.Attr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAccountAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getAccountAllOK" + "." + "Attr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAccountAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAccountAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetAccountAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
