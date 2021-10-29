package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TenantAppManagementPolicy struct {
    PolicyBase
    // Restrictions that apply as default to all application objects in the tenant.
    applicationRestrictions *AppManagementConfiguration;
    // Denotes whether the policy is enabled. Default value is false.
    isEnabled *bool;
    // Restrictions that apply as default to all service principal objects in the tenant.
    servicePrincipalRestrictions *AppManagementConfiguration;
}
// Instantiates a new tenantAppManagementPolicy and sets the default values.
func NewTenantAppManagementPolicy()(*TenantAppManagementPolicy) {
    m := &TenantAppManagementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// Gets the applicationRestrictions property value. Restrictions that apply as default to all application objects in the tenant.
func (m *TenantAppManagementPolicy) GetApplicationRestrictions()(*AppManagementConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.applicationRestrictions
    }
}
// Gets the isEnabled property value. Denotes whether the policy is enabled. Default value is false.
func (m *TenantAppManagementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the servicePrincipalRestrictions property value. Restrictions that apply as default to all service principal objects in the tenant.
func (m *TenantAppManagementPolicy) GetServicePrincipalRestrictions()(*AppManagementConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalRestrictions
    }
}
// The deserialization information for the current model
func (m *TenantAppManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["applicationRestrictions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppManagementConfiguration() })
        if err != nil {
            return err
        }
        m.SetApplicationRestrictions(val.(*AppManagementConfiguration))
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["servicePrincipalRestrictions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppManagementConfiguration() })
        if err != nil {
            return err
        }
        m.SetServicePrincipalRestrictions(val.(*AppManagementConfiguration))
        return nil
    }
    return res
}
func (m *TenantAppManagementPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the applicationRestrictions property value. Restrictions that apply as default to all application objects in the tenant.
// Parameters:
//  - value : Value to set for the applicationRestrictions property.
func (m *TenantAppManagementPolicy) SetApplicationRestrictions(value *AppManagementConfiguration)() {
    m.applicationRestrictions = value
}
// Sets the isEnabled property value. Denotes whether the policy is enabled. Default value is false.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *TenantAppManagementPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the servicePrincipalRestrictions property value. Restrictions that apply as default to all service principal objects in the tenant.
// Parameters:
//  - value : Value to set for the servicePrincipalRestrictions property.
func (m *TenantAppManagementPolicy) SetServicePrincipalRestrictions(value *AppManagementConfiguration)() {
    m.servicePrincipalRestrictions = value
}
