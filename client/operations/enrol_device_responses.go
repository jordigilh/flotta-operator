// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/project-flotta/flotta-operator/models"
)

// EnrolDeviceReader is a Reader for the EnrolDevice structure.
type EnrolDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnrolDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnrolDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnrolDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnrolDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnrolDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnrolDeviceOK creates a EnrolDeviceOK with default headers values
func NewEnrolDeviceOK() *EnrolDeviceOK {
	return &EnrolDeviceOK{}
}

/*EnrolDeviceOK handles this case with default header values.

Success
*/
type EnrolDeviceOK struct {
}

func (o *EnrolDeviceOK) Error() string {
	return fmt.Sprintf("[POST /namespaces/{namespace}/devices/{device-id}/enrol][%d] enrolDeviceOK ", 200)
}

func (o *EnrolDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnrolDeviceUnauthorized creates a EnrolDeviceUnauthorized with default headers values
func NewEnrolDeviceUnauthorized() *EnrolDeviceUnauthorized {
	return &EnrolDeviceUnauthorized{}
}

/*EnrolDeviceUnauthorized handles this case with default header values.

Unauthorized
*/
type EnrolDeviceUnauthorized struct {
}

func (o *EnrolDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /namespaces/{namespace}/devices/{device-id}/enrol][%d] enrolDeviceUnauthorized ", 401)
}

func (o *EnrolDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnrolDeviceForbidden creates a EnrolDeviceForbidden with default headers values
func NewEnrolDeviceForbidden() *EnrolDeviceForbidden {
	return &EnrolDeviceForbidden{}
}

/*EnrolDeviceForbidden handles this case with default header values.

Forbidden
*/
type EnrolDeviceForbidden struct {
}

func (o *EnrolDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /namespaces/{namespace}/devices/{device-id}/enrol][%d] enrolDeviceForbidden ", 403)
}

func (o *EnrolDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnrolDeviceDefault creates a EnrolDeviceDefault with default headers values
func NewEnrolDeviceDefault(code int) *EnrolDeviceDefault {
	return &EnrolDeviceDefault{
		_statusCode: code,
	}
}

/*EnrolDeviceDefault handles this case with default header values.

Error
*/
type EnrolDeviceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the enrol device default response
func (o *EnrolDeviceDefault) Code() int {
	return o._statusCode
}

func (o *EnrolDeviceDefault) Error() string {
	return fmt.Sprintf("[POST /namespaces/{namespace}/devices/{device-id}/enrol][%d] EnrolDevice default  %+v", o._statusCode, o.Payload)
}

func (o *EnrolDeviceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnrolDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
