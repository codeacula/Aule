// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/codeacula/Aule/discord-bot/models"
)

// GetNodeOKCode is the HTTP code returned for type GetNodeOK
const GetNodeOKCode int = 200

/*
GetNodeOK OK message.

swagger:response getNodeOK
*/
type GetNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Node `json:"body,omitempty"`
}

// NewGetNodeOK creates GetNodeOK with default headers values
func NewGetNodeOK() *GetNodeOK {

	return &GetNodeOK{}
}

// WithPayload adds the payload to the get node o k response
func (o *GetNodeOK) WithPayload(payload *models.Node) *GetNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node o k response
func (o *GetNodeOK) SetPayload(payload *models.Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetNodeDefault Unexpected error

swagger:response getNodeDefault
*/
type GetNodeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNodeDefault creates GetNodeDefault with default headers values
func NewGetNodeDefault(code int) *GetNodeDefault {
	if code <= 0 {
		code = 500
	}

	return &GetNodeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get node default response
func (o *GetNodeDefault) WithStatusCode(code int) *GetNodeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get node default response
func (o *GetNodeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get node default response
func (o *GetNodeDefault) WithPayload(payload *models.Error) *GetNodeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node default response
func (o *GetNodeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}