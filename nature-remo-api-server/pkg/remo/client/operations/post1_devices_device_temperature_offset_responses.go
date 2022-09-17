// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Post1DevicesDeviceTemperatureOffsetReader is a Reader for the Post1DevicesDeviceTemperatureOffset structure.
type Post1DevicesDeviceTemperatureOffsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Post1DevicesDeviceTemperatureOffsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPost1DevicesDeviceTemperatureOffsetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPost1DevicesDeviceTemperatureOffsetOK creates a Post1DevicesDeviceTemperatureOffsetOK with default headers values
func NewPost1DevicesDeviceTemperatureOffsetOK() *Post1DevicesDeviceTemperatureOffsetOK {
	return &Post1DevicesDeviceTemperatureOffsetOK{}
}

/* Post1DevicesDeviceTemperatureOffsetOK describes a response with status code 200, with default header values.

Updated
*/
type Post1DevicesDeviceTemperatureOffsetOK struct {
}

// IsSuccess returns true when this post1 devices device temperature offset o k response has a 2xx status code
func (o *Post1DevicesDeviceTemperatureOffsetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post1 devices device temperature offset o k response has a 3xx status code
func (o *Post1DevicesDeviceTemperatureOffsetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post1 devices device temperature offset o k response has a 4xx status code
func (o *Post1DevicesDeviceTemperatureOffsetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post1 devices device temperature offset o k response has a 5xx status code
func (o *Post1DevicesDeviceTemperatureOffsetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post1 devices device temperature offset o k response a status code equal to that given
func (o *Post1DevicesDeviceTemperatureOffsetOK) IsCode(code int) bool {
	return code == 200
}

func (o *Post1DevicesDeviceTemperatureOffsetOK) Error() string {
	return fmt.Sprintf("[POST /1/devices/{device}/temperature_offset][%d] post1DevicesDeviceTemperatureOffsetOK ", 200)
}

func (o *Post1DevicesDeviceTemperatureOffsetOK) String() string {
	return fmt.Sprintf("[POST /1/devices/{device}/temperature_offset][%d] post1DevicesDeviceTemperatureOffsetOK ", 200)
}

func (o *Post1DevicesDeviceTemperatureOffsetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}