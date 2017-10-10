// Code generated by go-swagger; DO NOT EDIT.

package distribution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/checkr/flagr/swagger_gen/models"
)

// FindDistributionsOKCode is the HTTP code returned for type FindDistributionsOK
const FindDistributionsOKCode int = 200

/*FindDistributionsOK distribution under the segment

swagger:response findDistributionsOK
*/
type FindDistributionsOK struct {

	/*
	  In: Body
	*/
	Payload models.FindDistributionsOKBody `json:"body,omitempty"`
}

// NewFindDistributionsOK creates FindDistributionsOK with default headers values
func NewFindDistributionsOK() *FindDistributionsOK {
	return &FindDistributionsOK{}
}

// WithPayload adds the payload to the find distributions o k response
func (o *FindDistributionsOK) WithPayload(payload models.FindDistributionsOKBody) *FindDistributionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find distributions o k response
func (o *FindDistributionsOK) SetPayload(payload models.FindDistributionsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindDistributionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.FindDistributionsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindDistributionsDefault generic error response

swagger:response findDistributionsDefault
*/
type FindDistributionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindDistributionsDefault creates FindDistributionsDefault with default headers values
func NewFindDistributionsDefault(code int) *FindDistributionsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindDistributionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find distributions default response
func (o *FindDistributionsDefault) WithStatusCode(code int) *FindDistributionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find distributions default response
func (o *FindDistributionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find distributions default response
func (o *FindDistributionsDefault) WithPayload(payload *models.Error) *FindDistributionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find distributions default response
func (o *FindDistributionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindDistributionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}