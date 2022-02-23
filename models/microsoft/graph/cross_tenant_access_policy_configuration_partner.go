package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CrossTenantAccessPolicyConfigurationPartner 
type CrossTenantAccessPolicyConfigurationPartner struct {
    CrossTenantAccessPolicyConfigurationBase
    // Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
    isServiceProvider *bool;
    // The tenant identifier for the partner Azure AD organization. Read-only.
    tenantId *string;
}
// NewCrossTenantAccessPolicyConfigurationPartner instantiates a new crossTenantAccessPolicyConfigurationPartner and sets the default values.
func NewCrossTenantAccessPolicyConfigurationPartner()(*CrossTenantAccessPolicyConfigurationPartner) {
    m := &CrossTenantAccessPolicyConfigurationPartner{
        CrossTenantAccessPolicyConfigurationBase: *NewCrossTenantAccessPolicyConfigurationBase(),
    }
    return m
}
// GetIsServiceProvider gets the isServiceProvider property value. Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetIsServiceProvider()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isServiceProvider
    }
}
// GetTenantId gets the tenantId property value. The tenant identifier for the partner Azure AD organization. Read-only.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyConfigurationPartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CrossTenantAccessPolicyConfigurationBase.GetFieldDeserializers()
    res["isServiceProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsServiceProvider(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
func (m *CrossTenantAccessPolicyConfigurationPartner) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyConfigurationPartner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CrossTenantAccessPolicyConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isServiceProvider", m.GetIsServiceProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsServiceProvider sets the isServiceProvider property value. Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetIsServiceProvider(value *bool)() {
    if m != nil {
        m.isServiceProvider = value
    }
}
// SetTenantId sets the tenantId property value. The tenant identifier for the partner Azure AD organization. Read-only.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
