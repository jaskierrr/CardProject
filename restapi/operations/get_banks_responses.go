// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"card-project/models"
)

// GetBanksOKCode is the HTTP code returned for type GetBanksOK
const GetBanksOKCode int = 200

/*
GetBanksOK Ok

swagger:response getBanksOK
*/
type GetBanksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Bank `json:"body,omitempty"`
}

// NewGetBanksOK creates GetBanksOK with default headers values
func NewGetBanksOK() *GetBanksOK {

	return &GetBanksOK{}
}

// WithPayload adds the payload to the get banks o k response
func (o *GetBanksOK) WithPayload(payload []*models.Bank) *GetBanksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get banks o k response
func (o *GetBanksOK) SetPayload(payload []*models.Bank) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBanksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Bank, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetBanksDefault Общая ошибка

swagger:response getBanksDefault
*/
type GetBanksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetBanksDefault creates GetBanksDefault with default headers values
func NewGetBanksDefault(code int) *GetBanksDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBanksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get banks default response
func (o *GetBanksDefault) WithStatusCode(code int) *GetBanksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get banks default response
func (o *GetBanksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get banks default response
func (o *GetBanksDefault) WithPayload(payload *models.ErrorResponse) *GetBanksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get banks default response
func (o *GetBanksDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBanksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
