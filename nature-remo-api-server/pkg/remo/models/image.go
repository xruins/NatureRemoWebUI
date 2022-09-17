// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Image Basename of the image file included in the app. Ex: "ico_ac_1"
//
//
// swagger:model Image
type Image string

// Validate validates this image
func (m Image) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image based on context it is used
func (m Image) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}