package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GovernancePolicyTemplate struct {
    Entity
    // 
    displayName *string;
    // 
    policy *GovernancePolicy;
    // 
    settings *BusinessFlowSettings;
}
// Instantiates a new governancePolicyTemplate and sets the default values.
func NewGovernancePolicyTemplate()(*GovernancePolicyTemplate) {
    m := &GovernancePolicyTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. 
func (m *GovernancePolicyTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the policy property value. 
func (m *GovernancePolicyTemplate) GetPolicy()(*GovernancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// Gets the settings property value. 
func (m *GovernancePolicyTemplate) GetSettings()(*BusinessFlowSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// The deserialization information for the current model
func (m *GovernancePolicyTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernancePolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(*GovernancePolicy))
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlowSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*BusinessFlowSettings))
        }
        return nil
    }
    return res
}
func (m *GovernancePolicyTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GovernancePolicyTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GovernancePolicyTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the policy property value. 
// Parameters:
//  - value : Value to set for the policy property.
func (m *GovernancePolicyTemplate) SetPolicy(value *GovernancePolicy)() {
    m.policy = value
}
// Sets the settings property value. 
// Parameters:
//  - value : Value to set for the settings property.
func (m *GovernancePolicyTemplate) SetSettings(value *BusinessFlowSettings)() {
    m.settings = value
}
