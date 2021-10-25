package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceComplianceScriptValidationResult struct {
    additionalData map[string]interface{};
    ruleErrors []DeviceComplianceScriptRuleError;
    rules []DeviceComplianceScriptRule;
    scriptErrors []DeviceComplianceScriptError;
}
func NewDeviceComplianceScriptValidationResult()(*DeviceComplianceScriptValidationResult) {
    m := &DeviceComplianceScriptValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceComplianceScriptValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceComplianceScriptValidationResult) GetRuleErrors()([]DeviceComplianceScriptRuleError) {
    if m == nil {
        return nil
    } else {
        return m.ruleErrors
    }
}
func (m *DeviceComplianceScriptValidationResult) GetRules()([]DeviceComplianceScriptRule) {
    if m == nil {
        return nil
    } else {
        return m.rules
    }
}
func (m *DeviceComplianceScriptValidationResult) GetScriptErrors()([]DeviceComplianceScriptError) {
    if m == nil {
        return nil
    } else {
        return m.scriptErrors
    }
}
func (m *DeviceComplianceScriptValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ruleErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptRuleError() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceScriptRuleError, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceScriptRuleError))
        }
        m.SetRuleErrors(res)
        return nil
    }
    res["rules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptRule() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceScriptRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceScriptRule))
        }
        m.SetRules(res)
        return nil
    }
    res["scriptErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptError() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceScriptError, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceScriptError))
        }
        m.SetScriptErrors(res)
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptValidationResult) IsNil()(bool) {
    return m == nil
}
func (m *DeviceComplianceScriptValidationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRuleErrors()))
        for i, v := range m.GetRuleErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("ruleErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScriptErrors()))
        for i, v := range m.GetScriptErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("scriptErrors", cast)
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
func (m *DeviceComplianceScriptValidationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceComplianceScriptValidationResult) SetRuleErrors(value []DeviceComplianceScriptRuleError)() {
    m.ruleErrors = value
}
func (m *DeviceComplianceScriptValidationResult) SetRules(value []DeviceComplianceScriptRule)() {
    m.rules = value
}
func (m *DeviceComplianceScriptValidationResult) SetScriptErrors(value []DeviceComplianceScriptError)() {
    m.scriptErrors = value
}
