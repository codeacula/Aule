// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/codeacula/Aule/discord-bot/models"
)

// GetSummaryOKCode is the HTTP code returned for type GetSummaryOK
const GetSummaryOKCode int = 200

/*
GetSummaryOK OK message.

swagger:response getSummaryOK
*/
type GetSummaryOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Node `json:"body,omitempty"`
}

// NewGetSummaryOK creates GetSummaryOK with default headers values
func NewGetSummaryOK() *GetSummaryOK {

	return &GetSummaryOK{}
}

// WithPayload adds the payload to the get summary o k response
func (o *GetSummaryOK) WithPayload(payload []*models.Node) *GetSummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get summary o k response
func (o *GetSummaryOK) SetPayload(payload []*models.Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Node, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetSummaryDefault Unexpected error

swagger:response getSummaryDefault
*/
type GetSummaryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSummaryDefault creates GetSummaryDefault with default headers values
func NewGetSummaryDefault(code int) *GetSummaryDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSummaryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get summary default response
func (o *GetSummaryDefault) WithStatusCode(code int) *GetSummaryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get summary default response
func (o *GetSummaryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get summary default response
func (o *GetSummaryDefault) WithPayload(payload *models.Error) *GetSummaryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get summary default response
func (o *GetSummaryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSummaryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
