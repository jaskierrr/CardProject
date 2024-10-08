// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchBanksHandlerFunc turns a function with the right signature into a patch banks handler
type PatchBanksHandlerFunc func(PatchBanksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchBanksHandlerFunc) Handle(params PatchBanksParams) middleware.Responder {
	return fn(params)
}

// PatchBanksHandler interface for that can handle valid patch banks params
type PatchBanksHandler interface {
	Handle(PatchBanksParams) middleware.Responder
}

// NewPatchBanks creates a new http.Handler for the patch banks operation
func NewPatchBanks(ctx *middleware.Context, handler PatchBanksHandler) *PatchBanks {
	return &PatchBanks{Context: ctx, Handler: handler}
}

/*
	PatchBanks swagger:route PATCH /banks patchBanks

Patch bank
*/
type PatchBanks struct {
	Context *middleware.Context
	Handler PatchBanksHandler
}

func (o *PatchBanks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchBanksParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
