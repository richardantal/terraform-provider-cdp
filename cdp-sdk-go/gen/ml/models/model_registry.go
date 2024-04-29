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

// ModelRegistry Model registry object
//
// swagger:model ModelRegistry
type ModelRegistry struct {

	// The creation time of model registry.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// The creator of model registry.
	Creator string `json:"creator,omitempty"`

	// ModelRegistry CRN.
	Crn string `json:"crn,omitempty"`

	// The domain of the model registry
	Domain string `json:"domain,omitempty"`

	// The environment CRN of model registry.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// Environment name.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The id of model registry.
	ID int64 `json:"id,omitempty"`

	// The machine user CRN of the model registry.
	MachineUserCrn string `json:"machineUserCrn,omitempty"`

	// The namespace of model registry.
	Namespace string `json:"namespace,omitempty"`

	// The s3 bucket of model registry.
	S3Bucket string `json:"s3Bucket,omitempty"`

	// The Ozone endpoint of model registry.
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// The service name of model registry.
	ServiceName string `json:"serviceName,omitempty"`

	// The status of model registry.
	Status string `json:"status,omitempty"`

	// Deprecated. WorkSpace CRN. This field is no longer used.
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`

	// Deprecated, please refer to serviceName. Workspace name.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// Validate validates this model registry
func (m *ModelRegistry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelRegistry) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model registry based on context it is used
func (m *ModelRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelRegistry) UnmarshalBinary(b []byte) error {
	var res ModelRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
