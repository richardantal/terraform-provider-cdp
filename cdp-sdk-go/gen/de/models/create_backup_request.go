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

// CreateBackupRequest Request object for Create Backup command.
//
// swagger:model CreateBackupRequest
type CreateBackupRequest struct {

	// Backup file relative path. The path should not include file name, and be relative to the service backup location. The service backup location can be found in the service description obtained via the 'describe-service' command.
	BackupPath string `json:"backupPath,omitempty"`

	// Backup virtual cluster content options.
	BackupVcContentOptions *BackupContentOptions `json:"backupVcContentOptions,omitempty"`

	// Backup description.
	Description string `json:"description,omitempty"`

	// Whether to backup service virtual clusters. By default, virtual clusters are backed up.
	IncludeVc *bool `json:"includeVc,omitempty"`

	// Service ID of the service to backup.
	// Required: true
	ServiceID *string `json:"serviceId"`

	// Backup operation timeout in seconds.
	Timeout int32 `json:"timeout,omitempty"`
}

// Validate validates this create backup request
func (m *CreateBackupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupVcContentOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBackupRequest) validateBackupVcContentOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupVcContentOptions) { // not required
		return nil
	}

	if m.BackupVcContentOptions != nil {
		if err := m.BackupVcContentOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupVcContentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupVcContentOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateBackupRequest) validateServiceID(formats strfmt.Registry) error {

	if err := validate.Required("serviceId", "body", m.ServiceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create backup request based on the context it is used
func (m *CreateBackupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupVcContentOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBackupRequest) contextValidateBackupVcContentOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupVcContentOptions != nil {

		if swag.IsZero(m.BackupVcContentOptions) { // not required
			return nil
		}

		if err := m.BackupVcContentOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupVcContentOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupVcContentOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateBackupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBackupRequest) UnmarshalBinary(b []byte) error {
	var res CreateBackupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
