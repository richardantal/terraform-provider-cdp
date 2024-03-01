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

// DisengageAutoAdminRequest A request to disengage the autonomous nature of the database.
//
// swagger:model DisengageAutoAdminRequest
type DisengageAutoAdminRequest struct {

	// Database name.
	// Required: true
	DatabaseName *string `json:"databaseName"`

	// Environment name.
	// Required: true
	EnvironmentName *string `json:"environmentName"`
}

// Validate validates this disengage auto admin request
func (m *DisengageAutoAdminRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisengageAutoAdminRequest) validateDatabaseName(formats strfmt.Registry) error {

	if err := validate.Required("databaseName", "body", m.DatabaseName); err != nil {
		return err
	}

	return nil
}

func (m *DisengageAutoAdminRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this disengage auto admin request based on context it is used
func (m *DisengageAutoAdminRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DisengageAutoAdminRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisengageAutoAdminRequest) UnmarshalBinary(b []byte) error {
	var res DisengageAutoAdminRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
