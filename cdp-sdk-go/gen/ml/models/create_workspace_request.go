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

// CreateWorkspaceRequest Request object for the CreateWorkspace method.
//
// swagger:model CreateWorkspaceRequest
type CreateWorkspaceRequest struct {

	// The whitelist of CIDR blocks which can access the API server.
	AuthorizedIPRanges []string `json:"authorizedIPRanges"`

	// Toggle for cdsw migration preflight validation
	CdswMigrationMode string `json:"cdswMigrationMode,omitempty"`

	// The boolean flag to disable TLS setup for workbench. By default, the TLS is enabled.
	DisableTLS bool `json:"disableTLS,omitempty"`

	// Enables Cloudera AI governance by integrating with Cloudera Atlas. By default, this flag is disabled.
	EnableGovernance bool `json:"enableGovernance,omitempty"`

	// Enables the model metrics service for exporting metrics for models to a metrics store.
	EnableModelMetrics bool `json:"enableModelMetrics,omitempty"`

	// The boolean flag is used to enable monitoring. By default, monitoring is disabled.
	EnableMonitoring bool `json:"enableMonitoring,omitempty"`

	// The environment for the workbench to create.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Optional configurations for an existing Postgres to export model metrics to.
	ExistingDatabaseConfig *ExistingDatabaseConfig `json:"existingDatabaseConfig,omitempty"`

	// Optionally use an existing NFS by providing the hostname and desired path (Azure and Private Cloud only).
	ExistingNFS string `json:"existingNFS,omitempty"`

	// The whitelist of IPs for load balancer.
	LoadBalancerIPWhitelists []string `json:"loadBalancerIPWhitelists"`

	// The version of ML workload app to install.
	MlVersion string `json:"mlVersion,omitempty"`

	// The NFS Protocol version of the NFS server we are using for Azure and Private Cloud.
	NfsVersion string `json:"nfsVersion,omitempty"`

	// Outbound Types provided for the workbench.
	OutboundTypes []OutboundTypes `json:"outboundTypes"`

	// Whether to create a private cluster.
	PrivateCluster bool `json:"privateCluster,omitempty"`

	// The request for Kubernetes workbench provision. Required in public cloud.
	ProvisionK8sRequest *ProvisionK8sRequest `json:"provisionK8sRequest,omitempty"`

	// Skip pre-flight validations if requested.
	SkipValidation bool `json:"skipValidation,omitempty"`

	// The static subdomain to be used for the workbench.
	StaticSubdomain string `json:"staticSubdomain,omitempty"`

	// The list of subnets used for the load balancer that Cloudera AI creates.
	SubnetsForLoadBalancers []string `json:"subnetsForLoadBalancers"`

	// The boolean flag to request public load balancer. By default, private load balancer is used.
	UsePublicLoadBalancer bool `json:"usePublicLoadBalancer,omitempty"`

	// Whether to whitelist only 'authorizedIPRanges' given or all public IPs.
	WhitelistAuthorizedIPRanges bool `json:"whitelistAuthorizedIPRanges,omitempty"`

	// The name of the workbench to create.
	// Required: true
	WorkspaceName *string `json:"workspaceName"`
}

// Validate validates this create workspace request
func (m *CreateWorkspaceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingDatabaseConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisionK8sRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaceName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateWorkspaceRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *CreateWorkspaceRequest) validateExistingDatabaseConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ExistingDatabaseConfig) { // not required
		return nil
	}

	if m.ExistingDatabaseConfig != nil {
		if err := m.ExistingDatabaseConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingDatabaseConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingDatabaseConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CreateWorkspaceRequest) validateOutboundTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.OutboundTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.OutboundTypes); i++ {

		if err := m.OutboundTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CreateWorkspaceRequest) validateProvisionK8sRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvisionK8sRequest) { // not required
		return nil
	}

	if m.ProvisionK8sRequest != nil {
		if err := m.ProvisionK8sRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisionK8sRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provisionK8sRequest")
			}
			return err
		}
	}

	return nil
}

func (m *CreateWorkspaceRequest) validateWorkspaceName(formats strfmt.Registry) error {

	if err := validate.Required("workspaceName", "body", m.WorkspaceName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create workspace request based on the context it is used
func (m *CreateWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExistingDatabaseConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutboundTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvisionK8sRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateWorkspaceRequest) contextValidateExistingDatabaseConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ExistingDatabaseConfig != nil {

		if swag.IsZero(m.ExistingDatabaseConfig) { // not required
			return nil
		}

		if err := m.ExistingDatabaseConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingDatabaseConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingDatabaseConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CreateWorkspaceRequest) contextValidateOutboundTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OutboundTypes); i++ {

		if swag.IsZero(m.OutboundTypes[i]) { // not required
			return nil
		}

		if err := m.OutboundTypes[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outboundTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CreateWorkspaceRequest) contextValidateProvisionK8sRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.ProvisionK8sRequest != nil {

		if swag.IsZero(m.ProvisionK8sRequest) { // not required
			return nil
		}

		if err := m.ProvisionK8sRequest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provisionK8sRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provisionK8sRequest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res CreateWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
