package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GovernanceRuleSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The id of the rule. For example, ExpirationRule and MfaRule.
    ruleIdentifier *string;
    // The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
    setting *string;
}
// Instantiates a new governanceRuleSetting and sets the default values.
func NewGovernanceRuleSetting()(*GovernanceRuleSetting) {
    m := &GovernanceRuleSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceRuleSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the ruleIdentifier property value. The id of the rule. For example, ExpirationRule and MfaRule.
func (m *GovernanceRuleSetting) GetRuleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleIdentifier
    }
}
// Gets the setting property value. The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
func (m *GovernanceRuleSetting) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *GovernanceRuleSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the ruleIdentifier property value. The id of the rule. For example, ExpirationRule and MfaRule.
// Parameters:
//  - value : Value to set for the ruleIdentifier property.
func (m *GovernanceRuleSetting) SetRuleIdentifier(value *string)() {
    m.ruleIdentifier = value
}
// Sets the setting property value. The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
// Parameters:
//  - value : Value to set for the setting property.
func (m *GovernanceRuleSetting) SetSetting(value *string)() {
    m.setting = value
}
