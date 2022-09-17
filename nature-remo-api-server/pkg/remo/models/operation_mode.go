// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OperationMode The range of OperationModes which the air conditioner accepts depends on the air conditioner model. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model.
//
// swagger:model OperationMode
type OperationMode string

func NewOperationMode(value OperationMode) *OperationMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OperationMode.
func (m OperationMode) Pointer() *OperationMode {
	return &m
}

const (

	// OperationModeEmpty captures enum value ""
	OperationModeEmpty OperationMode = ""

	// OperationModeCool captures enum value "cool"
	OperationModeCool OperationMode = "cool"

	// OperationModeWarm captures enum value "warm"
	OperationModeWarm OperationMode = "warm"

	// OperationModeDry captures enum value "dry"
	OperationModeDry OperationMode = "dry"

	// OperationModeBlow captures enum value "blow"
	OperationModeBlow OperationMode = "blow"

	// OperationModeAuto captures enum value "auto"
	OperationModeAuto OperationMode = "auto"
)

// for schema
var operationModeEnum []interface{}

func init() {
	var res []OperationMode
	if err := json.Unmarshal([]byte(`["","cool","warm","dry","blow","auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationModeEnum = append(operationModeEnum, v)
	}
}

func (m OperationMode) validateOperationModeEnum(path, location string, value OperationMode) error {
	if err := validate.EnumCase(path, location, value, operationModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this operation mode
func (m OperationMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOperationModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this operation mode based on context it is used
func (m OperationMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}