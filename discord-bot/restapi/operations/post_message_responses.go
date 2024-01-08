// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/codeacula/Aule/discord-bot/models"
)

// PostMessageOKCode is the HTTP code returned for type PostMessageOK
const PostMessageOKCode int = 200

/*
PostMessageOK OK message.

swagger:response postMessageOK
*/
type PostMessageOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostMessageOK creates PostMessageOK with default headers values
func NewPostMessageOK() *PostMessageOK {

	return &PostMessageOK{}
}

// WithPayload adds the payload to the post message o k response
func (o *PostMessageOK) WithPayload(payload string) *PostMessageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post message o k response
func (o *PostMessageOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMessageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
PostMessageDefault Unexpected error

swagger:response postMessageDefault
*/
type PostMessageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMessageDefault creates PostMessageDefault with default headers values
func NewPostMessageDefault(code int) *PostMessageDefault {
	if code <= 0 {
		code = 500
	}

	return &PostMessageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post message default response
func (o *PostMessageDefault) WithStatusCode(code int) *PostMessageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post message default response
func (o *PostMessageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post message default response
func (o *PostMessageDefault) WithPayload(payload *models.Error) *PostMessageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post message default response
func (o *PostMessageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMessageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
