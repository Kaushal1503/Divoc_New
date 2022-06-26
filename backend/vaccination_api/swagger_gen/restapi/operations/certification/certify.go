// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/api/swagger_gen/models"
)

// CertifyHandlerFunc turns a function with the right signature into a certify handler
type CertifyHandlerFunc func(CertifyParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn CertifyHandlerFunc) Handle(params CertifyParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// CertifyHandler interface for that can handle valid certify params
type CertifyHandler interface {
	Handle(CertifyParams, *models.JWTClaimBody) middleware.Responder
}

// NewCertify creates a new http.Handler for the certify operation
func NewCertify(ctx *middleware.Context, handler CertifyHandler) *Certify {
	return &Certify{Context: ctx, Handler: handler}
}

/*Certify swagger:route POST /v1/certify certification certify

Certify the one or more vaccination

Certification happens asynchronously, this requires vaccinator authorization and vaccinator should be trained for the vaccination that is being certified.

*/
type Certify struct {
	Context *middleware.Context
	Handler CertifyHandler
}

func (o *Certify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCertifyParams()

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