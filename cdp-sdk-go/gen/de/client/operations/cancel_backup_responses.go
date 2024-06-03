// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/models"
)

// CancelBackupReader is a Reader for the CancelBackup structure.
type CancelBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCancelBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelBackupOK creates a CancelBackupOK with default headers values
func NewCancelBackupOK() *CancelBackupOK {
	return &CancelBackupOK{}
}

/*
CancelBackupOK describes a response with status code 200, with default header values.

Response object for Cancel Backup command.
*/
type CancelBackupOK struct {
	Payload models.CancelBackupResponse
}

// IsSuccess returns true when this cancel backup o k response has a 2xx status code
func (o *CancelBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel backup o k response has a 3xx status code
func (o *CancelBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel backup o k response has a 4xx status code
func (o *CancelBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel backup o k response has a 5xx status code
func (o *CancelBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel backup o k response a status code equal to that given
func (o *CancelBackupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancel backup o k response
func (o *CancelBackupOK) Code() int {
	return 200
}

func (o *CancelBackupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/cancelBackup][%d] cancelBackupOK %s", 200, payload)
}

func (o *CancelBackupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/cancelBackup][%d] cancelBackupOK %s", 200, payload)
}

func (o *CancelBackupOK) GetPayload() models.CancelBackupResponse {
	return o.Payload
}

func (o *CancelBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelBackupDefault creates a CancelBackupDefault with default headers values
func NewCancelBackupDefault(code int) *CancelBackupDefault {
	return &CancelBackupDefault{
		_statusCode: code,
	}
}

/*
CancelBackupDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CancelBackupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this cancel backup default response has a 2xx status code
func (o *CancelBackupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cancel backup default response has a 3xx status code
func (o *CancelBackupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cancel backup default response has a 4xx status code
func (o *CancelBackupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cancel backup default response has a 5xx status code
func (o *CancelBackupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cancel backup default response a status code equal to that given
func (o *CancelBackupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cancel backup default response
func (o *CancelBackupDefault) Code() int {
	return o._statusCode
}

func (o *CancelBackupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/cancelBackup][%d] cancelBackup default %s", o._statusCode, payload)
}

func (o *CancelBackupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/cancelBackup][%d] cancelBackup default %s", o._statusCode, payload)
}

func (o *CancelBackupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
