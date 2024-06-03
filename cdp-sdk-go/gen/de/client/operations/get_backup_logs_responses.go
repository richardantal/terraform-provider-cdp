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

// GetBackupLogsReader is a Reader for the GetBackupLogs structure.
type GetBackupLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackupLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackupLogsOK creates a GetBackupLogsOK with default headers values
func NewGetBackupLogsOK() *GetBackupLogsOK {
	return &GetBackupLogsOK{}
}

/*
GetBackupLogsOK describes a response with status code 200, with default header values.

Response object for Get Backup Logs command.
*/
type GetBackupLogsOK struct {
	Payload *models.GetBackupLogsResponse
}

// IsSuccess returns true when this get backup logs o k response has a 2xx status code
func (o *GetBackupLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup logs o k response has a 3xx status code
func (o *GetBackupLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup logs o k response has a 4xx status code
func (o *GetBackupLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup logs o k response has a 5xx status code
func (o *GetBackupLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup logs o k response a status code equal to that given
func (o *GetBackupLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get backup logs o k response
func (o *GetBackupLogsOK) Code() int {
	return 200
}

func (o *GetBackupLogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getBackupLogs][%d] getBackupLogsOK %s", 200, payload)
}

func (o *GetBackupLogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getBackupLogs][%d] getBackupLogsOK %s", 200, payload)
}

func (o *GetBackupLogsOK) GetPayload() *models.GetBackupLogsResponse {
	return o.Payload
}

func (o *GetBackupLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetBackupLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupLogsDefault creates a GetBackupLogsDefault with default headers values
func NewGetBackupLogsDefault(code int) *GetBackupLogsDefault {
	return &GetBackupLogsDefault{
		_statusCode: code,
	}
}

/*
GetBackupLogsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetBackupLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get backup logs default response has a 2xx status code
func (o *GetBackupLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get backup logs default response has a 3xx status code
func (o *GetBackupLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get backup logs default response has a 4xx status code
func (o *GetBackupLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get backup logs default response has a 5xx status code
func (o *GetBackupLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get backup logs default response a status code equal to that given
func (o *GetBackupLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get backup logs default response
func (o *GetBackupLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetBackupLogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getBackupLogs][%d] getBackupLogs default %s", o._statusCode, payload)
}

func (o *GetBackupLogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getBackupLogs][%d] getBackupLogs default %s", o._statusCode, payload)
}

func (o *GetBackupLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
