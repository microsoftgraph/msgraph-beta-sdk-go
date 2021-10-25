package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GovernancePolicyTemplate struct {
    Entity
    displayName *string;
    policy *GovernancePolicy;
    settings *BusinessFlowSettings;
}
func NewGovernancePolicyTemplate()(*GovernancePolicyTemplate) {
    m := &GovernancePolicyTemplate{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GovernancePolicyTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GovernancePolicyTemplate) GetPolicy()(*GovernancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
func (m *GovernancePolicyTemplate) GetSettings()(*BusinessFlowSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *GovernancePolicyTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernancePolicy() })
        if err != nil {
            return err
        }
        m.SetPolicy(val.(*GovernancePolicy))
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlowSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*BusinessFlowSettings))
        return nil
    }
    return res
}
func (m *GovernancePolicyTemplate) IsNil()(bool) {
    return m == nil
}
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
func (m *GovernancePolicyTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GovernancePolicyTemplate) SetPolicy(value *GovernancePolicy)() {
    m.policy = value
}
func (m *GovernancePolicyTemplate) SetSettings(value *BusinessFlowSettings)() {
    m.settings = value
}
