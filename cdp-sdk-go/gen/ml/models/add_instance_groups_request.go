// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddInstanceGroupsRequest Request object for AddInstanceGroups.
//
// swagger:model AddInstanceGroupsRequest
type AddInstanceGroupsRequest struct {

	// The instance groups that we want to add to the workbench.
	// Required: true
	InstanceGroups []*InstanceGroup `json:"instanceGroups"`

	// The CRN of the workbench to which the instance groups are to be added.
	// Required: true
	WorkspaceCrn *string `json:"workspaceCrn"`
}

// Validate validates this add instance groups request
func (m *AddInstanceGroupsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaceCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddInstanceGroupsRequest) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddInstanceGroupsRequest) validateWorkspaceCrn(formats strfmt.Registry) error {

	if err := validate.Required("workspaceCrn", "body", m.WorkspaceCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this add instance groups request based on the context it is used
func (m *AddInstanceGroupsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddInstanceGroupsRequest) contextValidateInstanceGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceGroups); i++ {

		if m.InstanceGroups[i] != nil {

			if swag.IsZero(m.InstanceGroups[i]) { // not required
				return nil
			}

			if err := m.InstanceGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddInstanceGroupsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddInstanceGroupsRequest) UnmarshalBinary(b []byte) error {
	var res AddInstanceGroupsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
