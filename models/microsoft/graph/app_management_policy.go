package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppManagementPolicy 
type AppManagementPolicy struct {
    PolicyBase
    // Collection of application and service principals to which a policy is applied.
    appliesTo []DirectoryObject;
    // Denotes whether the policy is enabled.
    isEnabled *bool;
    // Restrictions that apply to an application or service principal object.
    restrictions *AppManagementConfiguration;
}
// NewAppManagementPolicy instantiates a new appManagementPolicy and sets the default values.
func NewAppManagementPolicy()(*AppManagementPolicy) {
    m := &AppManagementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// GetAppliesTo gets the appliesTo property value. Collection of application and service principals to which a policy is applied.
func (m *AppManagementPolicy) GetAppliesTo()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetIsEnabled gets the isEnabled property value. Denotes whether the policy is enabled.
func (m *AppManagementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetRestrictions gets the restrictions property value. Restrictions that apply to an application or service principal object.
func (m *AppManagementPolicy) GetRestrictions()(*AppManagementConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.restrictions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppManagementConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictions(val.(*AppManagementConfiguration))
        }
        return nil
    }
    return res
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
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// SetAppliesTo sets the appliesTo property value. Collection of application and service principals to which a policy is applied.
func (m *AppManagementPolicy) SetAppliesTo(value []DirectoryObject)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetIsEnabled sets the isEnabled property value. Denotes whether the policy is enabled.
func (m *AppManagementPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetRestrictions sets the restrictions property value. Restrictions that apply to an application or service principal object.
func (m *AppManagementPolicy) SetRestrictions(value *AppManagementConfiguration)() {
    if m != nil {
        m.restrictions = value
    }
}
