package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CrossTenantAccessPolicy 
type CrossTenantAccessPolicy struct {
    TenantRelationshipAccessPolicyBase
    // Defines the default configuration for how your organization interacts with external Azure Active Directory organizations.
    default_escaped *CrossTenantAccessPolicyConfigurationDefault;
    // Defines partner-specific configurations for external Azure Active Directory organizations.
    partners []CrossTenantAccessPolicyConfigurationPartner;
}
// NewCrossTenantAccessPolicy instantiates a new crossTenantAccessPolicy and sets the default values.
func NewCrossTenantAccessPolicy()(*CrossTenantAccessPolicy) {
    m := &CrossTenantAccessPolicy{
        TenantRelationshipAccessPolicyBase: *NewTenantRelationshipAccessPolicyBase(),
    }
    return m
}
// GetDefault gets the default property value. Defines the default configuration for how your organization interacts with external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) GetDefault()(*CrossTenantAccessPolicyConfigurationDefault) {
    if m == nil {
        return nil
    } else {
        return m.default_escaped
    }
}
// GetPartners gets the partners property value. Defines partner-specific configurations for external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) GetPartners()([]CrossTenantAccessPolicyConfigurationPartner) {
    if m == nil {
        return nil
    } else {
        return m.partners
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TenantRelationshipAccessPolicyBase.GetFieldDeserializers()
    res["default"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyConfigurationDefault() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefault(val.(*CrossTenantAccessPolicyConfigurationDefault))
        }
        return nil
    }
    res["partners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCrossTenantAccessPolicyConfigurationPartner() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CrossTenantAccessPolicyConfigurationPartner, len(val))
            for i, v := range val {
                res[i] = *(v.(*CrossTenantAccessPolicyConfigurationPartner))
            }
            m.SetPartners(res)
        }
        return nil
    }
    return res
}
func (m *CrossTenantAccessPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.TenantRelationshipAccessPolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("default", m.GetDefault())
        if err != nil {
            return err
        }
    }
    if m.GetPartners() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPartners()))
        for i, v := range m.GetPartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("partners", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefault sets the default property value. Defines the default configuration for how your organization interacts with external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) SetDefault(value *CrossTenantAccessPolicyConfigurationDefault)() {
    if m != nil {
        m.default_escaped = value
    }
}
// SetPartners sets the partners property value. Defines partner-specific configurations for external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicy) SetPartners(value []CrossTenantAccessPolicyConfigurationPartner)() {
    if m != nil {
        m.partners = value
    }
}
