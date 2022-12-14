// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Post1AppliancesApplianceLightReader is a Reader for the Post1AppliancesApplianceLight structure.
type Post1AppliancesApplianceLightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Post1AppliancesApplianceLightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPost1AppliancesApplianceLightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPost1AppliancesApplianceLightOK creates a Post1AppliancesApplianceLightOK with default headers values
func NewPost1AppliancesApplianceLightOK() *Post1AppliancesApplianceLightOK {
	return &Post1AppliancesApplianceLightOK{}
}

/* Post1AppliancesApplianceLightOK describes a response with status code 200, with default header values.

Updated light state
*/
type Post1AppliancesApplianceLightOK struct {
}

// IsSuccess returns true when this post1 appliances appliance light o k response has a 2xx status code
func (o *Post1AppliancesApplianceLightOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post1 appliances appliance light o k response has a 3xx status code
func (o *Post1AppliancesApplianceLightOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post1 appliances appliance light o k response has a 4xx status code
func (o *Post1AppliancesApplianceLightOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post1 appliances appliance light o k response has a 5xx status code
func (o *Post1AppliancesApplianceLightOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post1 appliances appliance light o k response a status code equal to that given
func (o *Post1AppliancesApplianceLightOK) IsCode(code int) bool {
	return code == 200
}

func (o *Post1AppliancesApplianceLightOK) Error() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}/light][%d] post1AppliancesApplianceLightOK ", 200)
}

func (o *Post1AppliancesApplianceLightOK) String() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}/light][%d] post1AppliancesApplianceLightOK ", 200)
}

func (o *Post1AppliancesApplianceLightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
