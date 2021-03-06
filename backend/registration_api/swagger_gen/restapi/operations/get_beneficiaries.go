// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/registration-api/swagger_gen/models"
)

// GetBeneficiariesHandlerFunc turns a function with the right signature into a get beneficiaries handler
type GetBeneficiariesHandlerFunc func(GetBeneficiariesParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBeneficiariesHandlerFunc) Handle(params GetBeneficiariesParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetBeneficiariesHandler interface for that can handle valid get beneficiaries params
type GetBeneficiariesHandler interface {
	Handle(GetBeneficiariesParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetBeneficiaries creates a new http.Handler for the get beneficiaries operation
func NewGetBeneficiaries(ctx *middleware.Context, handler GetBeneficiariesHandler) *GetBeneficiaries {
	return &GetBeneficiaries{Context: ctx, Handler: handler}
}

/*GetBeneficiaries swagger:route GET /beneficiaries/search getBeneficiaries

Get all beneficiaries

*/
type GetBeneficiaries struct {
	Context *middleware.Context
	Handler GetBeneficiariesHandler
}

func (o *GetBeneficiaries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBeneficiariesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
