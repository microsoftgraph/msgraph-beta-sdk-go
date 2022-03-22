package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantAppManagementPolicy 
type TenantAppManagementPolicy struct {
    PolicyBase
    // Restrictions that apply as default to all application objects in the tenant.
    applicationRestrictions AppManagementConfigurationable;
    // Denotes whether the policy is enabled. Default value is false.
    isEnabled *bool;
    // Restrictions that apply as default to all service principal objects in the tenant.
    servicePrincipalRestrictions AppManagementConfigurationable;
}
// NewTenantAppManagementPolicy instantiates a new tenantAppManagementPolicy and sets the default values.
func NewTenantAppManagementPolicy()(*TenantAppManagementPolicy) {
    m := &TenantAppManagementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreateTenantAppManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAppManagementPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTenantAppManagementPolicy(), nil
}
// GetApplicationRestrictions gets the applicationRestrictions property value. Restrictions that apply as default to all application objects in the tenant.
func (m *TenantAppManagementPolicy) GetApplicationRestrictions()(AppManagementConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.applicationRestrictions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAppManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["applicationRestrictions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppManagementConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationRestrictions(val.(AppManagementConfigurationable))
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["servicePrincipalRestrictions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppManagementConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalRestrictions(val.(AppManagementConfigurationable))
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Denotes whether the policy is enabled. Default value is false.
func (m *TenantAppManagementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetServicePrincipalRestrictions gets the servicePrincipalRestrictions property value. Restrictions that apply as default to all service principal objects in the tenant.
func (m *TenantAppManagementPolicy) GetServicePrincipalRestrictions()(AppManagementConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalRestrictions
    }
}
// Serialize serializes information the current object
func (m *TenantAppManagementPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicationRestrictions", m.GetApplicationRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("servicePrincipalRestrictions", m.GetServicePrincipalRestrictions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationRestrictions sets the applicationRestrictions property value. Restrictions that apply as default to all application objects in the tenant.
func (m *TenantAppManagementPolicy) SetApplicationRestrictions(value AppManagementConfigurationable)() {
    if m != nil {
        m.applicationRestrictions = value
    }
}
// SetIsEnabled sets the isEnabled property value. Denotes whether the policy is enabled. Default value is false.
func (m *TenantAppManagementPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetServicePrincipalRestrictions sets the servicePrincipalRestrictions property value. Restrictions that apply as default to all service principal objects in the tenant.
func (m *TenantAppManagementPolicy) SetServicePrincipalRestrictions(value AppManagementConfigurationable)() {
    if m != nil {
        m.servicePrincipalRestrictions = value
    }
}
