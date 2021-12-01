package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernancePolicyTemplate 
type GovernancePolicyTemplate struct {
    Entity
    // 
    displayName *string;
    // 
    policy *GovernancePolicy;
    // 
    settings *BusinessFlowSettings;
}
// NewGovernancePolicyTemplate instantiates a new governancePolicyTemplate and sets the default values.
func NewGovernancePolicyTemplate()(*GovernancePolicyTemplate) {
    m := &GovernancePolicyTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. 
func (m *GovernancePolicyTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetPolicy gets the policy property value. 
func (m *GovernancePolicyTemplate) GetPolicy()(*GovernancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// GetSettings gets the settings property value. 
func (m *GovernancePolicyTemplate) GetSettings()(*BusinessFlowSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetDisplayName sets the displayName property value. 
func (m *GovernancePolicyTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPolicy sets the policy property value. 
func (m *GovernancePolicyTemplate) SetPolicy(value *GovernancePolicy)() {
    if m != nil {
        m.policy = value
    }
}
// SetSettings sets the settings property value. 
func (m *GovernancePolicyTemplate) SetSettings(value *BusinessFlowSettings)() {
    if m != nil {
        m.settings = value
    }
}
