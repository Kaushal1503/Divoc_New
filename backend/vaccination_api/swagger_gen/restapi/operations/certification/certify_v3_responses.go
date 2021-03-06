// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/api/swagger_gen/models"
)

// CertifyV3OKCode is the HTTP code returned for type CertifyV3OK
const CertifyV3OKCode int = 200

/*CertifyV3OK OK

swagger:response certifyV3OK
*/
type CertifyV3OK struct {
}

// NewCertifyV3OK creates CertifyV3OK with default headers values
func NewCertifyV3OK() *CertifyV3OK {

	return &CertifyV3OK{}
}

// WriteResponse to the client
func (o *CertifyV3OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CertifyV3BadRequestCode is the HTTP code returned for type CertifyV3BadRequest
const CertifyV3BadRequestCode int = 400

/*CertifyV3BadRequest Invalid input

swagger:response certifyV3BadRequest
*/
type CertifyV3BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCertifyV3BadRequest creates CertifyV3BadRequest with default headers values
func NewCertifyV3BadRequest() *CertifyV3BadRequest {

	return &CertifyV3BadRequest{}
}

// WithPayload adds the payload to the certify v3 bad request response
func (o *CertifyV3BadRequest) WithPayload(payload *models.Error) *CertifyV3BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the certify v3 bad request response
func (o *CertifyV3BadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CertifyV3BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
