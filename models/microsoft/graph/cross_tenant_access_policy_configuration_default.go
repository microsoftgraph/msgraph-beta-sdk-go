package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CrossTenantAccessPolicyConfigurationDefault 
type CrossTenantAccessPolicyConfigurationDefault struct {
    CrossTenantAccessPolicyConfigurationBase
    // If true, the default configuration is set to the system default configuration. If false, the default settings have been customized.
    isServiceDefault *bool;
}
// NewCrossTenantAccessPolicyConfigurationDefault instantiates a new crossTenantAccessPolicyConfigurationDefault and sets the default values.
func NewCrossTenantAccessPolicyConfigurationDefault()(*CrossTenantAccessPolicyConfigurationDefault) {
    m := &CrossTenantAccessPolicyConfigurationDefault{
        CrossTenantAccessPolicyConfigurationBase: *NewCrossTenantAccessPolicyConfigurationBase(),
    }
    return m
}
// GetIsServiceDefault gets the isServiceDefault property value. If true, the default configuration is set to the system default configuration. If false, the default settings have been customized.
func (m *CrossTenantAccessPolicyConfigurationDefault) GetIsServiceDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isServiceDefault
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyConfigurationDefault) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CrossTenantAccessPolicyConfigurationBase.GetFieldDeserializers()
    res["isServiceDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsServiceDefault(val)
        }
        return nil
    }
    return res
}
func (m *CrossTenantAccessPolicyConfigurationDefault) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyConfigurationDefault) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CrossTenantAccessPolicyConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isServiceDefault", m.GetIsServiceDefault())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsServiceDefault sets the isServiceDefault property value. If true, the default configuration is set to the system default configuration. If false, the default settings have been customized.
func (m *CrossTenantAccessPolicyConfigurationDefault) SetIsServiceDefault(value *bool)() {
    if m != nil {
        m.isServiceDefault = value
    }
}
