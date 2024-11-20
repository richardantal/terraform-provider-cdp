// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateDatabaseResponse A response to database update request
//
// swagger:model UpdateDatabaseResponse
type UpdateDatabaseResponse struct {

	// Status of whether the update was successful or not.
	Status string `json:"status,omitempty"`
}

// Validate validates this update database response
func (m *UpdateDatabaseResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update database response based on context it is used
func (m *UpdateDatabaseResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDatabaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDatabaseResponse) UnmarshalBinary(b []byte) error {
	var res UpdateDatabaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
