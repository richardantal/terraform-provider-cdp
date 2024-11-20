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

// ModifyWorkspaceLoadBalancerRequest Request object for ModifyWorkspaceLoadBalancer.
//
// swagger:model ModifyWorkspaceLoadBalancerRequest
type ModifyWorkspaceLoadBalancerRequest struct {

	// The allowlist of CIDR blocks which can access the loadbalancer.
	// Required: true
	LoadBalancerIPAllowLists []string `json:"loadBalancerIPAllowLists"`

	// The CRN of the workbench cluster to modify.
	// Required: true
	WorkspaceCrn *string `json:"workspaceCrn"`
}

// Validate validates this modify workspace load balancer request
func (m *ModifyWorkspaceLoadBalancerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadBalancerIPAllowLists(formats); err != nil {
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

func (m *ModifyWorkspaceLoadBalancerRequest) validateLoadBalancerIPAllowLists(formats strfmt.Registry) error {

	if err := validate.Required("loadBalancerIPAllowLists", "body", m.LoadBalancerIPAllowLists); err != nil {
		return err
	}

	return nil
}

func (m *ModifyWorkspaceLoadBalancerRequest) validateWorkspaceCrn(formats strfmt.Registry) error {

	if err := validate.Required("workspaceCrn", "body", m.WorkspaceCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this modify workspace load balancer request based on context it is used
func (m *ModifyWorkspaceLoadBalancerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModifyWorkspaceLoadBalancerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifyWorkspaceLoadBalancerRequest) UnmarshalBinary(b []byte) error {
	var res ModifyWorkspaceLoadBalancerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
