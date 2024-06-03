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

// DescribeBackupReader is a Reader for the DescribeBackup structure.
type DescribeBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeBackupOK creates a DescribeBackupOK with default headers values
func NewDescribeBackupOK() *DescribeBackupOK {
	return &DescribeBackupOK{}
}

/*
DescribeBackupOK describes a response with status code 200, with default header values.

Response object for Describe Backup command.
*/
type DescribeBackupOK struct {
	Payload *models.DescribeBackupResponse
}

// IsSuccess returns true when this describe backup o k response has a 2xx status code
func (o *DescribeBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe backup o k response has a 3xx status code
func (o *DescribeBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe backup o k response has a 4xx status code
func (o *DescribeBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe backup o k response has a 5xx status code
func (o *DescribeBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe backup o k response a status code equal to that given
func (o *DescribeBackupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe backup o k response
func (o *DescribeBackupOK) Code() int {
	return 200
}

func (o *DescribeBackupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/describeBackup][%d] describeBackupOK %s", 200, payload)
}

func (o *DescribeBackupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/describeBackup][%d] describeBackupOK %s", 200, payload)
}

func (o *DescribeBackupOK) GetPayload() *models.DescribeBackupResponse {
	return o.Payload
}

func (o *DescribeBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeBackupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeBackupDefault creates a DescribeBackupDefault with default headers values
func NewDescribeBackupDefault(code int) *DescribeBackupDefault {
	return &DescribeBackupDefault{
		_statusCode: code,
	}
}

/*
DescribeBackupDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeBackupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe backup default response has a 2xx status code
func (o *DescribeBackupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe backup default response has a 3xx status code
func (o *DescribeBackupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe backup default response has a 4xx status code
func (o *DescribeBackupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe backup default response has a 5xx status code
func (o *DescribeBackupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe backup default response a status code equal to that given
func (o *DescribeBackupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe backup default response
func (o *DescribeBackupDefault) Code() int {
	return o._statusCode
}

func (o *DescribeBackupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/describeBackup][%d] describeBackup default %s", o._statusCode, payload)
}

func (o *DescribeBackupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/describeBackup][%d] describeBackup default %s", o._statusCode, payload)
}

func (o *DescribeBackupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
