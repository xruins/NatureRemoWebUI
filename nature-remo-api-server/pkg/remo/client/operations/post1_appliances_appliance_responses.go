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

// Post1AppliancesApplianceReader is a Reader for the Post1AppliancesAppliance structure.
type Post1AppliancesApplianceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Post1AppliancesApplianceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPost1AppliancesApplianceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPost1AppliancesApplianceOK creates a Post1AppliancesApplianceOK with default headers values
func NewPost1AppliancesApplianceOK() *Post1AppliancesApplianceOK {
	return &Post1AppliancesApplianceOK{}
}

/* Post1AppliancesApplianceOK describes a response with status code 200, with default header values.

Updated an appliance
*/
type Post1AppliancesApplianceOK struct {
	Payload *models.Appliance
}

// IsSuccess returns true when this post1 appliances appliance o k response has a 2xx status code
func (o *Post1AppliancesApplianceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post1 appliances appliance o k response has a 3xx status code
func (o *Post1AppliancesApplianceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post1 appliances appliance o k response has a 4xx status code
func (o *Post1AppliancesApplianceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post1 appliances appliance o k response has a 5xx status code
func (o *Post1AppliancesApplianceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post1 appliances appliance o k response a status code equal to that given
func (o *Post1AppliancesApplianceOK) IsCode(code int) bool {
	return code == 200
}

func (o *Post1AppliancesApplianceOK) Error() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}][%d] post1AppliancesApplianceOK  %+v", 200, o.Payload)
}

func (o *Post1AppliancesApplianceOK) String() string {
	return fmt.Sprintf("[POST /1/appliances/{appliance}][%d] post1AppliancesApplianceOK  %+v", 200, o.Payload)
}

func (o *Post1AppliancesApplianceOK) GetPayload() *models.Appliance {
	return o.Payload
}

func (o *Post1AppliancesApplianceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Appliance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
