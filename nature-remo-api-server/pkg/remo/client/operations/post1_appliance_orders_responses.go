// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Post1ApplianceOrdersReader is a Reader for the Post1ApplianceOrders structure.
type Post1ApplianceOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Post1ApplianceOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPost1ApplianceOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPost1ApplianceOrdersOK creates a Post1ApplianceOrdersOK with default headers values
func NewPost1ApplianceOrdersOK() *Post1ApplianceOrdersOK {
	return &Post1ApplianceOrdersOK{}
}

/* Post1ApplianceOrdersOK describes a response with status code 200, with default header values.

Reordered appliances
*/
type Post1ApplianceOrdersOK struct {
}

// IsSuccess returns true when this post1 appliance orders o k response has a 2xx status code
func (o *Post1ApplianceOrdersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post1 appliance orders o k response has a 3xx status code
func (o *Post1ApplianceOrdersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post1 appliance orders o k response has a 4xx status code
func (o *Post1ApplianceOrdersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post1 appliance orders o k response has a 5xx status code
func (o *Post1ApplianceOrdersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post1 appliance orders o k response a status code equal to that given
func (o *Post1ApplianceOrdersOK) IsCode(code int) bool {
	return code == 200
}

func (o *Post1ApplianceOrdersOK) Error() string {
	return fmt.Sprintf("[POST /1/appliance_orders][%d] post1ApplianceOrdersOK ", 200)
}

func (o *Post1ApplianceOrdersOK) String() string {
	return fmt.Sprintf("[POST /1/appliance_orders][%d] post1ApplianceOrdersOK ", 200)
}

func (o *Post1ApplianceOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}