package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceRuleSetting 
type GovernanceRuleSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The id of the rule. For example, ExpirationRule and MfaRule.
    ruleIdentifier *string;
    // The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
    setting *string;
}
// NewGovernanceRuleSetting instantiates a new governanceRuleSetting and sets the default values.
func NewGovernanceRuleSetting()(*GovernanceRuleSetting) {
    m := &GovernanceRuleSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceRuleSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRuleIdentifier gets the ruleIdentifier property value. The id of the rule. For example, ExpirationRule and MfaRule.
func (m *GovernanceRuleSetting) GetRuleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleIdentifier
    }
}
// GetSetting gets the setting property value. The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
func (m *GovernanceRuleSetting) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRuleSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ruleIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleIdentifier(val)
        }
        return nil
    }
    res["setting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceRuleSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceRuleSetting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRuleIdentifier sets the ruleIdentifier property value. The id of the rule. For example, ExpirationRule and MfaRule.
func (m *GovernanceRuleSetting) SetRuleIdentifier(value *string)() {
    if m != nil {
        m.ruleIdentifier = value
    }
}
// SetSetting sets the setting property value. The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
func (m *GovernanceRuleSetting) SetSetting(value *string)() {
    if m != nil {
        m.setting = value
    }
}
