// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Temperature The temperature in string format. The unit is described in Aircon object. The range of Temperatures which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
//
// swagger:model Temperature
type Temperature string

// Validate validates this temperature
func (m Temperature) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this temperature based on context it is used
func (m Temperature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
