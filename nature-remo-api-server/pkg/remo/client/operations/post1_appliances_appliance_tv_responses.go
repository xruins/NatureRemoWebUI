// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Post1AppliancesApplianceTvReader is a Reader for the Post1AppliancesApplianceTv structure.
type Post1AppliancesApplianceTvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Post1AppliancesApplianceTvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPost1AppliancesApplianceTvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPost1AppliancesApplianceTvOK creates a Post1AppliancesApplianceTvOK with default headers values
func NewPost1AppliancesApplianceTvOK() *Post1AppliancesApplianceTvOK {
	return &Post1AppliancesApplianceTvOK{}
}

/* Post1AppliancesApplianceTvOK describes a response with status code 200, with default header values.

Updated tv state
*/
type Post1AppliancesApplianceTvOK struct {
}

// IsSuccess returns true when this post1 appliances appliance tv o k response has a 2xx status code
func (o *Post1AppliancesApplianceTvOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post1 appliances appliance tv o k response has a 3xx status code
func (o *Post1AppliancesApplianceTvOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post1 appliances appliance tv o k response has a 4xx status code
func (o *Post1AppliancesApplianceTvOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post1 appliances appliance tv o k response has a 5xx status code
func (o *Post1AppliancesApplianceTvOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post1 appliances appliance tv o k response a status code equal to that given
func (o *Post1AppliancesApplianceTvOK) IsCode(code int) bool {
	return code == 200
}

func (o *Post1AppliancesApplianceTvOK) Error() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}/tv][%d] post1AppliancesApplianceTvOK ", 200)
}

func (o *Post1AppliancesApplianceTvOK) String() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}/tv][%d] post1AppliancesApplianceTvOK ", 200)
}

func (o *Post1AppliancesApplianceTvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
