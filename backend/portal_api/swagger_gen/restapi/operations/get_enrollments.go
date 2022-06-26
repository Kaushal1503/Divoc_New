// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// GetEnrollmentsHandlerFunc turns a function with the right signature into a get enrollments handler
type GetEnrollmentsHandlerFunc func(GetEnrollmentsParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEnrollmentsHandlerFunc) Handle(params GetEnrollmentsParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetEnrollmentsHandler interface for that can handle valid get enrollments params
type GetEnrollmentsHandler interface {
	Handle(GetEnrollmentsParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetEnrollments creates a new http.Handler for the get enrollments operation
func NewGetEnrollments(ctx *middleware.Context, handler GetEnrollmentsHandler) *GetEnrollments {
	return &GetEnrollments{Context: ctx, Handler: handler}
}

/*GetEnrollments swagger:route GET /enrollments getEnrollments

get enrollments

*/
type GetEnrollments struct {
	Context *middleware.Context
	Handler GetEnrollmentsHandler
}

func (o *GetEnrollments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEnrollmentsParams()

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
