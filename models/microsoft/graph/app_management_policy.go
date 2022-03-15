package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppManagementPolicy provides operations to call the instantiate method.
type AppManagementPolicy struct {
    PolicyBase
    // 
    appliesTo []DirectoryObjectable;
    // 
    isEnabled *bool;
    // 
    restrictions AppManagementConfigurationable;
}
// NewAppManagementPolicy instantiates a new appManagementPolicy and sets the default values.
func NewAppManagementPolicy()(*AppManagementPolicy) {
    m := &AppManagementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreateAppManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppManagementPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAppManagementPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. 
func (m *AppManagementPolicy) GetAppliesTo()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetAppliesTo(res)
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
    res["restrictions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppManagementConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictions(val.(AppManagementConfigurationable))
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. 
func (m *AppManagementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetRestrictions gets the restrictions property value. 
func (m *AppManagementPolicy) GetRestrictions()(AppManagementConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.restrictions
    }
}
func (m *AppManagementPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppManagementPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliesTo", cast)
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
        err = writer.WriteObjectValue("restrictions", m.GetRestrictions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. 
func (m *AppManagementPolicy) SetAppliesTo(value []DirectoryObjectable)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetIsEnabled sets the isEnabled property value. 
func (m *AppManagementPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetRestrictions sets the restrictions property value. 
func (m *AppManagementPolicy) SetRestrictions(value AppManagementConfigurationable)() {
    if m != nil {
        m.restrictions = value
    }
}
