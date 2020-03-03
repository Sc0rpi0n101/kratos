// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// SelfServiceBrowserVerifyReader is a Reader for the SelfServiceBrowserVerify structure.
type SelfServiceBrowserVerifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SelfServiceBrowserVerifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewSelfServiceBrowserVerifyFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSelfServiceBrowserVerifyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSelfServiceBrowserVerifyFound creates a SelfServiceBrowserVerifyFound with default headers values
func NewSelfServiceBrowserVerifyFound() *SelfServiceBrowserVerifyFound {
	return &SelfServiceBrowserVerifyFound{}
}

/*SelfServiceBrowserVerifyFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type SelfServiceBrowserVerifyFound struct {
}

func (o *SelfServiceBrowserVerifyFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/verification/{via}/confirm/{code}][%d] selfServiceBrowserVerifyFound ", 302)
}

func (o *SelfServiceBrowserVerifyFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSelfServiceBrowserVerifyInternalServerError creates a SelfServiceBrowserVerifyInternalServerError with default headers values
func NewSelfServiceBrowserVerifyInternalServerError() *SelfServiceBrowserVerifyInternalServerError {
	return &SelfServiceBrowserVerifyInternalServerError{}
}

/*SelfServiceBrowserVerifyInternalServerError handles this case with default header values.

genericError
*/
type SelfServiceBrowserVerifyInternalServerError struct {
	Payload *models.GenericError
}

func (o *SelfServiceBrowserVerifyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/verification/{via}/confirm/{code}][%d] selfServiceBrowserVerifyInternalServerError  %+v", 500, o.Payload)
}

func (o *SelfServiceBrowserVerifyInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *SelfServiceBrowserVerifyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}