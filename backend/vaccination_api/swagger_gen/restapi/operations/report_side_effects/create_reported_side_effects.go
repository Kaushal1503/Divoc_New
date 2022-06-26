// Code generated by go-swagger; DO NOT EDIT.

package report_side_effects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/divoc/api/swagger_gen/models"
)

// CreateReportedSideEffectsHandlerFunc turns a function with the right signature into a create reported side effects handler
type CreateReportedSideEffectsHandlerFunc func(CreateReportedSideEffectsParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateReportedSideEffectsHandlerFunc) Handle(params CreateReportedSideEffectsParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// CreateReportedSideEffectsHandler interface for that can handle valid create reported side effects params
type CreateReportedSideEffectsHandler interface {
	Handle(CreateReportedSideEffectsParams, *models.JWTClaimBody) middleware.Responder
}

// NewCreateReportedSideEffects creates a new http.Handler for the create reported side effects operation
func NewCreateReportedSideEffects(ctx *middleware.Context, handler CreateReportedSideEffectsHandler) *CreateReportedSideEffects {
	return &CreateReportedSideEffects{Context: ctx, Handler: handler}
}

/*CreateReportedSideEffects swagger:route POST /v1/report-side-effects reportSideEffects createReportedSideEffects

Create reported side effects

*/
type CreateReportedSideEffects struct {
	Context *middleware.Context
	Handler CreateReportedSideEffectsHandler
}

func (o *CreateReportedSideEffects) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateReportedSideEffectsParams()

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

// CreateReportedSideEffectsBody create reported side effects body
//
// swagger:model CreateReportedSideEffectsBody
type CreateReportedSideEffectsBody struct {

	// certificate Id
	CertificateID string `json:"certificateId,omitempty"`

	// side effects response
	SideEffectsResponse []*models.SideEffectsResponse `json:"sideEffectsResponse"`
}

// Validate validates this create reported side effects body
func (o *CreateReportedSideEffectsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSideEffectsResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateReportedSideEffectsBody) validateSideEffectsResponse(formats strfmt.Registry) error {

	if swag.IsZero(o.SideEffectsResponse) { // not required
		return nil
	}

	for i := 0; i < len(o.SideEffectsResponse); i++ {
		if swag.IsZero(o.SideEffectsResponse[i]) { // not required
			continue
		}

		if o.SideEffectsResponse[i] != nil {
			if err := o.SideEffectsResponse[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "sideEffectsResponse" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateReportedSideEffectsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateReportedSideEffectsBody) UnmarshalBinary(b []byte) error {
	var res CreateReportedSideEffectsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}