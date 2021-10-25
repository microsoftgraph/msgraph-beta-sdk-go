package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GovernanceRuleSetting struct {
    additionalData map[string]interface{};
    ruleIdentifier *string;
    setting *string;
}
func NewGovernanceRuleSetting()(*GovernanceRuleSetting) {
    m := &GovernanceRuleSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GovernanceRuleSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GovernanceRuleSetting) GetRuleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleIdentifier
    }
}
func (m *GovernanceRuleSetting) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
func (m *GovernanceRuleSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ruleIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRuleIdentifier(val)
        return nil
    }
    res["setting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSetting(val)
        return nil
    }
    return res
}
func (m *GovernanceRuleSetting) IsNil()(bool) {
    return m == nil
}
func (m *GovernanceRuleSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ruleIdentifier", m.GetRuleIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GovernanceRuleSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GovernanceRuleSetting) SetRuleIdentifier(value *string)() {
    m.ruleIdentifier = value
}
func (m *GovernanceRuleSetting) SetSetting(value *string)() {
    m.setting = value
}
