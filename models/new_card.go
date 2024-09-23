// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewCard new card
//
// swagger:model NewCard
type NewCard struct {

	// bank ID
	BankID int64 `json:"BankID,omitempty"`

	// create date
	// Format: date-time
	CreateDate strfmt.DateTime `json:"CreateDate,omitempty"`

	// number
	Number int64 `json:"Number,omitempty"`

	// user ID
	UserID int64 `json:"UserID,omitempty"`
}

// Validate validates this new card
func (m *NewCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewCard) validateCreateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CreateDate", "body", "date-time", m.CreateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this new card based on context it is used
func (m *NewCard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewCard) UnmarshalBinary(b []byte) error {
	var res NewCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
