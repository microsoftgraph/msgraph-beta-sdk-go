package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScriptValidationResult 
type DeviceComplianceScriptValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Errors in json for the script for rules.
    ruleErrors []DeviceComplianceScriptRuleErrorable;
    // Parsed rules from json.
    rules []DeviceComplianceScriptRuleable;
    // Errors in json for the script.
    scriptErrors []DeviceComplianceScriptErrorable;
}
// NewDeviceComplianceScriptValidationResult instantiates a new deviceComplianceScriptValidationResult and sets the default values.
func NewDeviceComplianceScriptValidationResult()(*DeviceComplianceScriptValidationResult) {
    m := &DeviceComplianceScriptValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceComplianceScriptValidationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptValidationResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceComplianceScriptValidationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ruleErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptRuleErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptRuleErrorable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptRuleErrorable)
            }
            m.SetRuleErrors(res)
        }
        return nil
    }
    res["rules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptRuleable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    res["scriptErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptErrorable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptErrorable)
            }
            m.SetScriptErrors(res)
        }
        return nil
    }
    return res
}
// GetRuleErrors gets the ruleErrors property value. Errors in json for the script for rules.
func (m *DeviceComplianceScriptValidationResult) GetRuleErrors()([]DeviceComplianceScriptRuleErrorable) {
    if m == nil {
        return nil
    } else {
        return m.ruleErrors
    }
}
// GetRules gets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) GetRules()([]DeviceComplianceScriptRuleable) {
    if m == nil {
        return nil
    } else {
        return m.rules
    }
}
// GetScriptErrors gets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) GetScriptErrors()([]DeviceComplianceScriptErrorable) {
    if m == nil {
        return nil
    } else {
        return m.scriptErrors
    }
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptValidationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRuleErrors() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRuleErrors()))
        for i, v := range m.GetRuleErrors() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("ruleErrors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetScriptErrors() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScriptErrors()))
        for i, v := range m.GetScriptErrors() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *DeviceComplianceScriptValidationResult) SetRuleErrors(value []DeviceComplianceScriptRuleErrorable)() {
    if m != nil {
        m.ruleErrors = value
    }
}
// SetRules sets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) SetRules(value []DeviceComplianceScriptRuleable)() {
    if m != nil {
        m.rules = value
    }
}
// SetScriptErrors sets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) SetScriptErrors(value []DeviceComplianceScriptErrorable)() {
    if m != nil {
        m.scriptErrors = value
    }
}
