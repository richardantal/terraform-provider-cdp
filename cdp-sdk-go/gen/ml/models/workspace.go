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

// Workspace A ML workbench, which includes the cluster and storage.
//
// swagger:model Workspace
type Workspace struct {

	// The whitelist of CIDR blocks which can access the API server.
	AuthorizedIPRanges []string `json:"authorizedIPRanges"`

	// The Backup MetaData for this workbench
	BackupMetadata *BackupMetadata `json:"backupMetadata,omitempty"`

	// The cloud platform of the environment that was used to create this workbench.
	// Required: true
	CloudPlatform *string `json:"cloudPlatform"`

	// The basedomain of the cluster.
	ClusterBaseDomain string `json:"clusterBaseDomain,omitempty"`

	// The Cluster ID for the workbench.
	ClusterID string `json:"clusterID,omitempty"`

	// Creation date of workbench.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The CRN of the creator of the workbench.
	// Required: true
	CreatorCrn *string `json:"creatorCrn"`

	// The CRN of the workbench.
	// Required: true
	Crn *string `json:"crn"`

	// To check if the cluster is publicly accessible or not.
	// Required: true
	EndpointPublicAccess *bool `json:"endpointPublicAccess"`

	// CRN of the environment.
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// The name of the workbench's environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Failure message from the most recent failure that has occurred during workbench provisioning.
	FailureMessage string `json:"failureMessage,omitempty"`

	// filesystemID used by the workbench
	// Required: true
	FilesystemID *string `json:"filesystemID"`

	// Whether governance is enabled.
	GovernanceEnabled bool `json:"governanceEnabled,omitempty"`

	// The health info information of the workbench.
	HealthInfoLists []*HealthInfo `json:"healthInfoLists"`

	// To Display if Https is enabled or not.
	// Required: true
	HTTPSEnabled *bool `json:"httpsEnabled"`

	// The instance groups.
	// Required: true
	InstanceGroups []*WorkspaceInstanceGroup `json:"instanceGroups"`

	// The name of the workbench.
	// Required: true
	InstanceName *string `json:"instanceName"`

	// The workbench's current status.
	// Required: true
	InstanceStatus *string `json:"instanceStatus"`

	// URL of the workbench's user interface.
	// Required: true
	InstanceURL *string `json:"instanceUrl"`

	// The value to indicate if the cluster is private or not.
	IsPrivate bool `json:"isPrivate,omitempty"`

	// The value to indicate if the workbench is restored one or not
	IsRestored bool `json:"isRestored,omitempty"`

	// The Kubernetes cluster name.
	// Required: true
	K8sClusterName *string `json:"k8sClusterName"`

	// IP whitelist for loadBalancer.
	LoadBalancerIPWhitelists []string `json:"loadBalancerIPWhitelists"`

	// Whether model metrics is enabled.
	ModelMetricsEnabled bool `json:"modelMetricsEnabled,omitempty"`

	// If usage monitoring is enabled or not on this workbench.
	// Required: true
	MonitoringEnabled *bool `json:"monitoringEnabled"`

	// NFS Version of the filesystem.
	NfsVersion string `json:"nfsVersion,omitempty"`

	// The subnets of the workbench.
	Subnets []string `json:"subnets"`

	// The list of subnets used for the load balancer that Cloudera AI creates.
	SubnetsForLoadBalancers []string `json:"subnetsForLoadBalancers"`

	// Tags provided by the user at the time of workbench creation.
	// Required: true
	Tags []*Tag `json:"tags"`

	// The upgrade state contains the workbench upgrade information.
	UpgradeState *UpgradeState `json:"upgradeState,omitempty"`

	// The version of Cloudera AI that was installed on the workbench.
	// Required: true
	Version *string `json:"version"`

	// Whether to whitelist only 'authorizedIPRanges' given or all public IPs.
	WhitelistAuthorizedIPRanges bool `json:"whitelistAuthorizedIPRanges,omitempty"`
}

// Validate validates this workspace
func (m *Workspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatorCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointPublicAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthInfoLists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPSEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateK8sClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoringEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspace) validateBackupMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupMetadata) { // not required
		return nil
	}

	if m.BackupMetadata != nil {
		if err := m.BackupMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *Workspace) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.Required("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateCreatorCrn(formats strfmt.Registry) error {

	if err := validate.Required("creatorCrn", "body", m.CreatorCrn); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateEndpointPublicAccess(formats strfmt.Registry) error {

	if err := validate.Required("endpointPublicAccess", "body", m.EndpointPublicAccess); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateFilesystemID(formats strfmt.Registry) error {

	if err := validate.Required("filesystemID", "body", m.FilesystemID); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateHealthInfoLists(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthInfoLists) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthInfoLists); i++ {
		if swag.IsZero(m.HealthInfoLists[i]) { // not required
			continue
		}

		if m.HealthInfoLists[i] != nil {
			if err := m.HealthInfoLists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("healthInfoLists" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("healthInfoLists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workspace) validateHTTPSEnabled(formats strfmt.Registry) error {

	if err := validate.Required("httpsEnabled", "body", m.HTTPSEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateInstanceGroups(formats strfmt.Registry) error {

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

func (m *Workspace) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instanceName", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateInstanceStatus(formats strfmt.Registry) error {

	if err := validate.Required("instanceStatus", "body", m.InstanceStatus); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateInstanceURL(formats strfmt.Registry) error {

	if err := validate.Required("instanceUrl", "body", m.InstanceURL); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateK8sClusterName(formats strfmt.Registry) error {

	if err := validate.Required("k8sClusterName", "body", m.K8sClusterName); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateMonitoringEnabled(formats strfmt.Registry) error {

	if err := validate.Required("monitoringEnabled", "body", m.MonitoringEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workspace) validateUpgradeState(formats strfmt.Registry) error {
	if swag.IsZero(m.UpgradeState) { // not required
		return nil
	}

	if m.UpgradeState != nil {
		if err := m.UpgradeState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upgradeState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upgradeState")
			}
			return err
		}
	}

	return nil
}

func (m *Workspace) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this workspace based on the context it is used
func (m *Workspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthInfoLists(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstanceGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpgradeState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspace) contextValidateBackupMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupMetadata != nil {

		if swag.IsZero(m.BackupMetadata) { // not required
			return nil
		}

		if err := m.BackupMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *Workspace) contextValidateHealthInfoLists(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HealthInfoLists); i++ {

		if m.HealthInfoLists[i] != nil {

			if swag.IsZero(m.HealthInfoLists[i]) { // not required
				return nil
			}

			if err := m.HealthInfoLists[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("healthInfoLists" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("healthInfoLists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workspace) contextValidateInstanceGroups(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Workspace) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Workspace) contextValidateUpgradeState(ctx context.Context, formats strfmt.Registry) error {

	if m.UpgradeState != nil {

		if swag.IsZero(m.UpgradeState) { // not required
			return nil
		}

		if err := m.UpgradeState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upgradeState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upgradeState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workspace) UnmarshalBinary(b []byte) error {
	var res Workspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
