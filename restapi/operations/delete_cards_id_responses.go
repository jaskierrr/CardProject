// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"card-project/models"
)

// DeleteCardsIDNoContentCode is the HTTP code returned for type DeleteCardsIDNoContent
const DeleteCardsIDNoContentCode int = 204

/*
DeleteCardsIDNoContent Card deleted

swagger:response deleteCardsIdNoContent
*/
type DeleteCardsIDNoContent struct {
}

// NewDeleteCardsIDNoContent creates DeleteCardsIDNoContent with default headers values
func NewDeleteCardsIDNoContent() *DeleteCardsIDNoContent {

	return &DeleteCardsIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteCardsIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
DeleteCardsIDDefault Общая ошибка

swagger:response deleteCardsIdDefault
*/
type DeleteCardsIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteCardsIDDefault creates DeleteCardsIDDefault with default headers values
func NewDeleteCardsIDDefault(code int) *DeleteCardsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteCardsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete cards ID default response
func (o *DeleteCardsIDDefault) WithStatusCode(code int) *DeleteCardsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete cards ID default response
func (o *DeleteCardsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete cards ID default response
func (o *DeleteCardsIDDefault) WithPayload(payload *models.ErrorResponse) *DeleteCardsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cards ID default response
func (o *DeleteCardsIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCardsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
