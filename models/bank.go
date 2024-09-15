// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Bank bank
//
// swagger:model Bank
type Bank struct {

	// name
	Name string `json:"Name,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this bank
func (m *Bank) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bank based on context it is used
func (m *Bank) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Bank) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bank) UnmarshalBinary(b []byte) error {
	var res Bank
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
