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

// Get1AppliancesReader is a Reader for the Get1Appliances structure.
type Get1AppliancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Get1AppliancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGet1AppliancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGet1AppliancesOK creates a Get1AppliancesOK with default headers values
func NewGet1AppliancesOK() *Get1AppliancesOK {
	return &Get1AppliancesOK{}
}

/* Get1AppliancesOK describes a response with status code 200, with default header values.

List of appliances
*/
type Get1AppliancesOK struct {
	Payload []*models.Appliance
}

// IsSuccess returns true when this get1 appliances o k response has a 2xx status code
func (o *Get1AppliancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get1 appliances o k response has a 3xx status code
func (o *Get1AppliancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get1 appliances o k response has a 4xx status code
func (o *Get1AppliancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get1 appliances o k response has a 5xx status code
func (o *Get1AppliancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get1 appliances o k response a status code equal to that given
func (o *Get1AppliancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *Get1AppliancesOK) Error() string {
	return fmt.Sprintf("[GET /1/appliances][%d] get1AppliancesOK  %+v", 200, o.Payload)
}

func (o *Get1AppliancesOK) String() string {
	return fmt.Sprintf("[GET /1/appliances][%d] get1AppliancesOK  %+v", 200, o.Payload)
}

func (o *Get1AppliancesOK) GetPayload() []*models.Appliance {
	return o.Payload
}

func (o *Get1AppliancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
