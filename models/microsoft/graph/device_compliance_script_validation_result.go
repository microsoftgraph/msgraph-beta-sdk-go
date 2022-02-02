package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScriptValidationResult 
type DeviceComplianceScriptValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Errors in json for the script for rules.
    ruleErrors []DeviceComplianceScriptRuleError;
    // Parsed rules from json.
    rules []DeviceComplianceScriptRule;
    // Errors in json for the script.
    scriptErrors []DeviceComplianceScriptError;
}
// NewDeviceComplianceScriptValidationResult instantiates a new deviceComplianceScriptValidationResult and sets the default values.
func NewDeviceComplianceScriptValidationResult()(*DeviceComplianceScriptValidationResult) {
    m := &DeviceComplianceScriptValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRuleErrors gets the ruleErrors property value. Errors in json for the script for rules.
func (m *DeviceComplianceScriptValidationResult) GetRuleErrors()([]DeviceComplianceScriptRuleError) {
    if m == nil {
        return nil
    } else {
        return m.ruleErrors
    }
}
// GetRules gets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) GetRules()([]DeviceComplianceScriptRule) {
    if m == nil {
        return nil
    } else {
        return m.rules
    }
}
// GetScriptErrors gets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) GetScriptErrors()([]DeviceComplianceScriptError) {
    if m == nil {
        return nil
    } else {
        return m.scriptErrors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ruleErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptRuleError() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptRuleError, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceScriptRuleError))
            }
            m.SetRuleErrors(res)
        }
        return nil
    }
    res["rules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptRule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptRule, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceScriptRule))
            }
            m.SetRules(res)
        }
        return nil
    }
    res["scriptErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptError() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptError, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceScriptError))
            }
            m.SetScriptErrors(res)
        }
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptValidationResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptValidationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRuleErrors() != nil {
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
    if m.GetRules() != nil {
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
    if m.GetScriptErrors() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptValidationResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRuleErrors sets the ruleErrors property value. Errors in json for the script for rules.
func (m *DeviceComplianceScriptValidationResult) SetRuleErrors(value []DeviceComplianceScriptRuleError)() {
    if m != nil {
        m.ruleErrors = value
    }
}
// SetRules sets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) SetRules(value []DeviceComplianceScriptRule)() {
    if m != nil {
        m.rules = value
    }
}
// SetScriptErrors sets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) SetScriptErrors(value []DeviceComplianceScriptError)() {
    if m != nil {
        m.scriptErrors = value
    }
}
