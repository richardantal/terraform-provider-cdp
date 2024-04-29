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

// UpgradeDatabaseRequest Asynchronous request to upgrade the CDP Runtime for a database.
//
// swagger:model UpgradeDatabaseRequest
type UpgradeDatabaseRequest struct {

	// The name or CRN of the database.
	Database string `json:"database,omitempty"`

	// Upgrade strategy for edge nodes.
	EdgeUpgradeStrategy EdgeUpgradeStrategy `json:"edgeUpgradeStrategy,omitempty"`

	// The name or CRN of the environment.
	Environment string `json:"environment,omitempty"`

	// ID of a database image to upgrade to. If this is empty, the default image is used.
	ImageID string `json:"imageId,omitempty"`

	// Batch size of non-worker nodes for restart.
	NonWorkerBatchSize int32 `json:"nonWorkerBatchSize,omitempty"`

	// Only perform an Operating System upgrade.
	OsUpgradeOnly bool `json:"osUpgradeOnly,omitempty"`

	// Perform upgrade in a rolling fashion.
	RollingUpgrade bool `json:"rollingUpgrade,omitempty"`

	// The CDP Runtime version to upgrade to.
	Runtime string `json:"runtime,omitempty"`

	// Batch size of worker nodes for restart.
	WorkerBatchSize int32 `json:"workerBatchSize,omitempty"`
}

// Validate validates this upgrade database request
func (m *UpgradeDatabaseRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeUpgradeStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeDatabaseRequest) validateEdgeUpgradeStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeUpgradeStrategy) { // not required
		return nil
	}

	if err := m.EdgeUpgradeStrategy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeUpgradeStrategy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeUpgradeStrategy")
		}
		return err
	}

	return nil
}

// ContextValidate validate this upgrade database request based on the context it is used
func (m *UpgradeDatabaseRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeUpgradeStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeDatabaseRequest) contextValidateEdgeUpgradeStrategy(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeUpgradeStrategy) { // not required
		return nil
	}

	if err := m.EdgeUpgradeStrategy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeUpgradeStrategy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeUpgradeStrategy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeDatabaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeDatabaseRequest) UnmarshalBinary(b []byte) error {
	var res UpgradeDatabaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
