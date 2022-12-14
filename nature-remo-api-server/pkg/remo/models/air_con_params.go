// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AirConParams air con params
//
// swagger:model AirConParams
type AirConParams struct {

	// button
	Button ACButton `json:"button,omitempty"`

	// dir
	Dir AirDirection `json:"dir,omitempty"`

	// mode
	Mode OperationMode `json:"mode,omitempty"`

	// temp
	Temp Temperature `json:"temp,omitempty"`

	// vol
	Vol AirVolume `json:"vol,omitempty"`
}

// Validate validates this air con params
func (m *AirConParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateButton(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AirConParams) validateButton(formats strfmt.Registry) error {
	if swag.IsZero(m.Button) { // not required
		return nil
	}

	if err := m.Button.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("button")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("button")
		}
		return err
	}

	return nil
}

func (m *AirConParams) validateDir(formats strfmt.Registry) error {
	if swag.IsZero(m.Dir) { // not required
		return nil
	}

	if err := m.Dir.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dir")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dir")
		}
		return err
	}

	return nil
}

func (m *AirConParams) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *AirConParams) validateTemp(formats strfmt.Registry) error {
	if swag.IsZero(m.Temp) { // not required
		return nil
	}

	if err := m.Temp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("temp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("temp")
		}
		return err
	}

	return nil
}

func (m *AirConParams) validateVol(formats strfmt.Registry) error {
	if swag.IsZero(m.Vol) { // not required
		return nil
	}

	if err := m.Vol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vol")
		}
		return err
	}

	return nil
}

// ContextValidate validate this air con params based on the context it is used
func (m *AirConParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateButton(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDir(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AirConParams) contextValidateButton(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Button.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("button")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("button")
		}
		return err
	}

	return nil
}

func (m *AirConParams) contextValidateDir(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Dir.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dir")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dir")
		}
		return err
	}

	return nil
}

func (m *AirConParams) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Mode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *AirConParams) contextValidateTemp(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Temp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("temp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("temp")
		}
		return err
	}

	return nil
}

func (m *AirConParams) contextValidateVol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Vol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vol")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AirConParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AirConParams) UnmarshalBinary(b []byte) error {
	var res AirConParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
