// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IsDeviceDeregistrableHandlerFunc turns a function with the right signature into a is device deregistrable handler
type IsDeviceDeregistrableHandlerFunc func(IsDeviceDeregistrableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IsDeviceDeregistrableHandlerFunc) Handle(params IsDeviceDeregistrableParams) middleware.Responder {
	return fn(params)
}

// IsDeviceDeregistrableHandler interface for that can handle valid is device deregistrable params
type IsDeviceDeregistrableHandler interface {
	Handle(IsDeviceDeregistrableParams) middleware.Responder
}

// NewIsDeviceDeregistrable creates a new http.Handler for the is device deregistrable operation
func NewIsDeviceDeregistrable(ctx *middleware.Context, handler IsDeviceDeregistrableHandler) *IsDeviceDeregistrable {
	return &IsDeviceDeregistrable{Context: ctx, Handler: handler}
}

/*IsDeviceDeregistrable swagger:route GET /namespaces/{namespace}/devices/{device-id}/deregister isDeviceDeregistrable

Confirms if a device can be deregistered

*/
type IsDeviceDeregistrable struct {
	Context *middleware.Context
	Handler IsDeviceDeregistrableHandler
}

func (o *IsDeviceDeregistrable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIsDeviceDeregistrableParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
