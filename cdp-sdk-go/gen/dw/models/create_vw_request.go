// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateVwRequest Request object for the createVw method.
//
// swagger:model CreateVwRequest
type CreateVwRequest struct {

	// Autoscaling settings for the Virtual Warehouse.
	Autoscaling *AutoscalingOptionsCreateRequest `json:"autoscaling,omitempty"`

	// This feature works only for AWS cluster type. An availability zone to host compute instances. If not specified, defaults to a randomly selected availability zone inferred from available subnets. In order to query possible options, see "availabilityZones" field of describe-cluster or list-clusters command response.
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// ID of cluster where Virtual Warehouse should be created.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// Configuration settings for the Virtual Warehouse.
	Config *ServiceConfigReq `json:"config,omitempty"`

	// ID of Database Catalog that the Virtual Warehouse should be attached to.
	// Required: true
	DbcID *string `json:"dbcId"`

	// Provides EBS gp3 volume as temporary storage space for Hive LLAP cache, and improves query performance. Configurable only at Virtual Warehouse creation. Using EBS volumes incurs additional costs.
	EbsLLAPSpillGB int32 `json:"ebsLLAPSpillGB,omitempty"`

	// Enable Unified Analytics. In the case of Hive Virtual Warehouses, this cannot be provided, because this value is inferred. In the case of Impala, this can be set. Passing --query-isolation-options will be considered only if this flag is set to true. If Unified Analytics is enabled then the "enableShutdownOfCoordinator" in --impala-ha-settings is explicitly disabled (ignored) and should not be provided, furthermore the "highAvailabilityMode" in --impala-ha-settings cannot be set to ACTIVE_ACTIVE.
	EnableUnifiedAnalytics bool `json:"enableUnifiedAnalytics,omitempty"`

	// DEPRECATED - Sets the authentication mode to use by Hive Server: * `LDAP` * `KERBEROS` Default: `LDAP` if not specified
	HiveAuthenticationMode *string `json:"hiveAuthenticationMode,omitempty"`

	// Set Hive Server High Availability mode in Private Cloud: * `DISABLED` (default) - Disables Hive Server high availability. * `ACTIVE_PASSIVE` - Runs two Hive Server instances, one active and one passive. Hive session failover is not supported in this setup.
	HiveServerHaMode *string `json:"hiveServerHaMode,omitempty"`

	// version of the Virtual Warehouse.
	ImageVersion string `json:"imageVersion,omitempty"`

	// High Availability settings for the Impala Virtual Warehouse. NOTE that in --autoscaling object you should avoid using the same deprecated properties because only the --impala-ha-setting properties will be considered if any of its values are set.
	ImpalaHaSettings *ImpalaHASettingsCreateRequest `json:"impalaHaSettings,omitempty"`

	// Impala specific options. It cannot be provided for Virtual Warehouse types other than Impala.
	ImpalaOptions *ImpalaOptionsCreateRequest `json:"impalaOptions,omitempty"`

	// Denotes whether the Virtual Warehouse has the Impala Query Log enabled or not.
	ImpalaQueryLog bool `json:"impalaQueryLog,omitempty"`

	// Instance type for this Virtual Warehouse. To learn what instance types are allowed to be used for a Hive or an Impala Virtual Warehouse, please use the 'describe-allowed-instance-types' command. The command output will list the usable instance types in its 'hive' and 'impala' fields accordingly.
	InstanceType string `json:"instanceType,omitempty"`

	// Name of the Virtual Warehouse.
	// Required: true
	Name *string `json:"name"`

	// Nodes per compute cluster. If specified, forces 'template' to be 'custom'
	NodeCount int32 `json:"nodeCount,omitempty"`

	// Value of 'true' automatically configures the Virtual Warehouse to support JWTs issues by the CDP JWT token provider.  Value of 'false' does not enable JWT auth on the Virtual Warehouse.  If this field is not specified, it defaults to 'false'.
	PlatformJwtAuth *bool `json:"platformJwtAuth,omitempty"`

	// Query isolation settings for the Virtual Warehouse. For Impala this value will be considered only if Unified Analytics is enabled.
	QueryIsolationOptions *QueryIsolationOptionsRequest `json:"queryIsolationOptions,omitempty"`

	// The Resource Pool of the Virtual Warehouse.
	ResourcePool string `json:"resourcePool,omitempty"`

	// Tags associated with the resources.
	Tags []*TagRequest `json:"tags"`

	// DEPRECATED: It will be replaced by the tShirtSize parameter in an upcoming release. Name of configuration template to use.
	// Enum: ["xsmall","small","medium","large"]
	Template string `json:"template,omitempty"`

	// Type of Virtual Warehouse to be created.
	// Required: true
	VwType *VwType `json:"vwType"`
}

// Validate validates this create vw request
func (m *CreateVwRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoscaling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDbcID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpalaHaSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpalaOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryIsolationOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVwType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVwRequest) validateAutoscaling(formats strfmt.Registry) error {
	if swag.IsZero(m.Autoscaling) { // not required
		return nil
	}

	if m.Autoscaling != nil {
		if err := m.Autoscaling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *CreateVwRequest) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) validateDbcID(formats strfmt.Registry) error {

	if err := validate.Required("dbcId", "body", m.DbcID); err != nil {
		return err
	}

	return nil
}

func (m *CreateVwRequest) validateImpalaHaSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.ImpalaHaSettings) { // not required
		return nil
	}

	if m.ImpalaHaSettings != nil {
		if err := m.ImpalaHaSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaHaSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaHaSettings")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) validateImpalaOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ImpalaOptions) { // not required
		return nil
	}

	if m.ImpalaOptions != nil {
		if err := m.ImpalaOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateVwRequest) validateQueryIsolationOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryIsolationOptions) { // not required
		return nil
	}

	if m.QueryIsolationOptions != nil {
		if err := m.QueryIsolationOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryIsolationOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryIsolationOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
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

var createVwRequestTypeTemplatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["xsmall","small","medium","large"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createVwRequestTypeTemplatePropEnum = append(createVwRequestTypeTemplatePropEnum, v)
	}
}

const (

	// CreateVwRequestTemplateXsmall captures enum value "xsmall"
	CreateVwRequestTemplateXsmall string = "xsmall"

	// CreateVwRequestTemplateSmall captures enum value "small"
	CreateVwRequestTemplateSmall string = "small"

	// CreateVwRequestTemplateMedium captures enum value "medium"
	CreateVwRequestTemplateMedium string = "medium"

	// CreateVwRequestTemplateLarge captures enum value "large"
	CreateVwRequestTemplateLarge string = "large"
)

// prop value enum
func (m *CreateVwRequest) validateTemplateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createVwRequestTypeTemplatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateVwRequest) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	// value enum
	if err := m.validateTemplateEnum("template", "body", m.Template); err != nil {
		return err
	}

	return nil
}

func (m *CreateVwRequest) validateVwType(formats strfmt.Registry) error {

	if err := validate.Required("vwType", "body", m.VwType); err != nil {
		return err
	}

	if err := validate.Required("vwType", "body", m.VwType); err != nil {
		return err
	}

	if m.VwType != nil {
		if err := m.VwType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vwType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vwType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create vw request based on the context it is used
func (m *CreateVwRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoscaling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImpalaHaSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImpalaOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueryIsolationOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVwType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVwRequest) contextValidateAutoscaling(ctx context.Context, formats strfmt.Registry) error {

	if m.Autoscaling != nil {

		if swag.IsZero(m.Autoscaling) { // not required
			return nil
		}

		if err := m.Autoscaling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) contextValidateImpalaHaSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.ImpalaHaSettings != nil {

		if swag.IsZero(m.ImpalaHaSettings) { // not required
			return nil
		}

		if err := m.ImpalaHaSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaHaSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaHaSettings")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) contextValidateImpalaOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ImpalaOptions != nil {

		if swag.IsZero(m.ImpalaOptions) { // not required
			return nil
		}

		if err := m.ImpalaOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) contextValidateQueryIsolationOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryIsolationOptions != nil {

		if swag.IsZero(m.QueryIsolationOptions) { // not required
			return nil
		}

		if err := m.QueryIsolationOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryIsolationOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryIsolationOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateVwRequest) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateVwRequest) contextValidateVwType(ctx context.Context, formats strfmt.Registry) error {

	if m.VwType != nil {

		if err := m.VwType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vwType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vwType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateVwRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateVwRequest) UnmarshalBinary(b []byte) error {
	var res CreateVwRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
