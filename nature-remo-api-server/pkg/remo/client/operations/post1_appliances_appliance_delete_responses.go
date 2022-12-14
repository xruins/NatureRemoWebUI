// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Post1AppliancesApplianceDeleteReader is a Reader for the Post1AppliancesApplianceDelete structure.
type Post1AppliancesApplianceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Post1AppliancesApplianceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPost1AppliancesApplianceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPost1AppliancesApplianceDeleteOK creates a Post1AppliancesApplianceDeleteOK with default headers values
func NewPost1AppliancesApplianceDeleteOK() *Post1AppliancesApplianceDeleteOK {
	return &Post1AppliancesApplianceDeleteOK{}
}

/* Post1AppliancesApplianceDeleteOK describes a response with status code 200, with default header values.

Deleted an appliance
*/
type Post1AppliancesApplianceDeleteOK struct {
}

// IsSuccess returns true when this post1 appliances appliance delete o k response has a 2xx status code
func (o *Post1AppliancesApplianceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post1 appliances appliance delete o k response has a 3xx status code
func (o *Post1AppliancesApplianceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post1 appliances appliance delete o k response has a 4xx status code
func (o *Post1AppliancesApplianceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post1 appliances appliance delete o k response has a 5xx status code
func (o *Post1AppliancesApplianceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post1 appliances appliance delete o k response a status code equal to that given
func (o *Post1AppliancesApplianceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *Post1AppliancesApplianceDeleteOK) Error() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}/delete][%d] post1AppliancesApplianceDeleteOK ", 200)
}

func (o *Post1AppliancesApplianceDeleteOK) String() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}/delete][%d] post1AppliancesApplianceDeleteOK ", 200)
}

func (o *Post1AppliancesApplianceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
