// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xruins/NatureRemoWebUI/nature-remo-api-server/pkg/remo/models"
)

// Get1UsersMeReader is a Reader for the Get1UsersMe structure.
type Get1UsersMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Get1UsersMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGet1UsersMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGet1UsersMeOK creates a Get1UsersMeOK with default headers values
func NewGet1UsersMeOK() *Get1UsersMeOK {
	return &Get1UsersMeOK{}
}

/* Get1UsersMeOK describes a response with status code 200, with default header values.

User information
*/
type Get1UsersMeOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get1 users me o k response has a 2xx status code
func (o *Get1UsersMeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get1 users me o k response has a 3xx status code
func (o *Get1UsersMeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get1 users me o k response has a 4xx status code
func (o *Get1UsersMeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get1 users me o k response has a 5xx status code
func (o *Get1UsersMeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get1 users me o k response a status code equal to that given
func (o *Get1UsersMeOK) IsCode(code int) bool {
	return code == 200
}

func (o *Get1UsersMeOK) Error() string {
	return fmt.Sprintf("[GET /1/users/me][%d] get1UsersMeOK  %+v", 200, o.Payload)
}

func (o *Get1UsersMeOK) String() string {
	return fmt.Sprintf("[GET /1/users/me][%d] get1UsersMeOK  %+v", 200, o.Payload)
}

func (o *Get1UsersMeOK) GetPayload() *models.User {
	return o.Payload
}

func (o *Get1UsersMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
