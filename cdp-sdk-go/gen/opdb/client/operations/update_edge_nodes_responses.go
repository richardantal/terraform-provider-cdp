// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// UpdateEdgeNodesReader is a Reader for the UpdateEdgeNodes structure.
type UpdateEdgeNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEdgeNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEdgeNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateEdgeNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateEdgeNodesOK creates a UpdateEdgeNodesOK with default headers values
func NewUpdateEdgeNodesOK() *UpdateEdgeNodesOK {
	return &UpdateEdgeNodesOK{}
}

/*
UpdateEdgeNodesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateEdgeNodesOK struct {
	Payload *models.UpdateEdgeNodesResponse
}

// IsSuccess returns true when this update edge nodes o k response has a 2xx status code
func (o *UpdateEdgeNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update edge nodes o k response has a 3xx status code
func (o *UpdateEdgeNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update edge nodes o k response has a 4xx status code
func (o *UpdateEdgeNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update edge nodes o k response has a 5xx status code
func (o *UpdateEdgeNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update edge nodes o k response a status code equal to that given
func (o *UpdateEdgeNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update edge nodes o k response
func (o *UpdateEdgeNodesOK) Code() int {
	return 200
}

func (o *UpdateEdgeNodesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateEdgeNodes][%d] updateEdgeNodesOK  %+v", 200, o.Payload)
}

func (o *UpdateEdgeNodesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateEdgeNodes][%d] updateEdgeNodesOK  %+v", 200, o.Payload)
}

func (o *UpdateEdgeNodesOK) GetPayload() *models.UpdateEdgeNodesResponse {
	return o.Payload
}

func (o *UpdateEdgeNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateEdgeNodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeNodesDefault creates a UpdateEdgeNodesDefault with default headers values
func NewUpdateEdgeNodesDefault(code int) *UpdateEdgeNodesDefault {
	return &UpdateEdgeNodesDefault{
		_statusCode: code,
	}
}

/*
UpdateEdgeNodesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateEdgeNodesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update edge nodes default response has a 2xx status code
func (o *UpdateEdgeNodesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update edge nodes default response has a 3xx status code
func (o *UpdateEdgeNodesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update edge nodes default response has a 4xx status code
func (o *UpdateEdgeNodesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update edge nodes default response has a 5xx status code
func (o *UpdateEdgeNodesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update edge nodes default response a status code equal to that given
func (o *UpdateEdgeNodesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update edge nodes default response
func (o *UpdateEdgeNodesDefault) Code() int {
	return o._statusCode
}

func (o *UpdateEdgeNodesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateEdgeNodes][%d] updateEdgeNodes default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateEdgeNodesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateEdgeNodes][%d] updateEdgeNodes default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateEdgeNodesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateEdgeNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
