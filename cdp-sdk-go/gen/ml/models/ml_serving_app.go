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

// MlServingApp The Cloudera AI Inference Service instance.
//
// swagger:model MlServingApp
type MlServingApp struct {

	// The CRN of the Cloudera AI Inference Service instance.
	AppCrn string `json:"appCrn,omitempty"`

	// The name of the Cloudera AI Inference Service instance.
	AppName string `json:"appName,omitempty"`

	// The cloud platform of the environment that was used to create this instance.
	// Required: true
	CloudPlatform *string `json:"cloudPlatform"`

	// The Kubernetes cluster information.
	Cluster *KubernetesCluster `json:"cluster,omitempty"`

	// Creation date of Cloudera AI Inference Service instance.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The CRN of the environment.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The name of the environment.
	EnvironmentName string `json:"environmentName,omitempty"`

	// Indicates if HTTPs communication was enabled on this application when it was provisioned.
	HTTPSEnabled bool `json:"httpsEnabled,omitempty"`

	// Is this service installed on a private cluster.
	IsPrivateCluster bool `json:"isPrivateCluster,omitempty"`

	// The Cloudera AI Inference Service version running on this instance.
	MlServingVersion string `json:"mlServingVersion,omitempty"`

	// The namespace used for this service.
	Namespace string `json:"namespace,omitempty"`

	// The email of the user who created this service.
	OwnerEmail string `json:"ownerEmail,omitempty"`

	// The status of the Cloudera AI Inference Service instance.
	Status string `json:"status,omitempty"`

	// The list of subnets used for the load balancer.
	SubnetsForLoadBalancers []string `json:"subnetsForLoadBalancers"`

	// Indicates if this Cloudera AI Inference Service instance uses a public load balancer.
	UsePublicLoadBalancer bool `json:"usePublicLoadBalancer,omitempty"`
}

// Validate validates this ml serving app
func (m *MlServingApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MlServingApp) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.Required("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *MlServingApp) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *MlServingApp) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ml serving app based on the context it is used
func (m *MlServingApp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MlServingApp) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {

		if swag.IsZero(m.Cluster) { // not required
			return nil
		}

		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MlServingApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MlServingApp) UnmarshalBinary(b []byte) error {
	var res MlServingApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
