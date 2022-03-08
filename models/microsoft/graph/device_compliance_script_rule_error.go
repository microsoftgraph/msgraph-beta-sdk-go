package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScriptRuleError provides operations to call the validateComplianceScript method.
type DeviceComplianceScriptRuleError struct {
    DeviceComplianceScriptError
    // Setting name for the rule with error.
    settingName *string;
}
// NewDeviceComplianceScriptRuleError instantiates a new deviceComplianceScriptRuleError and sets the default values.
func NewDeviceComplianceScriptRuleError()(*DeviceComplianceScriptRuleError) {
    m := &DeviceComplianceScriptRuleError{
        DeviceComplianceScriptError: *NewDeviceComplianceScriptError(),
    }
    return m
}
// CreateDeviceComplianceScriptRuleErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptRuleErrorFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceComplianceScriptRuleError(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptRuleError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DeviceComplianceScriptError.GetFieldDeserializers()
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    return res
}
// GetSettingName gets the settingName property value. Setting name for the rule with error.
func (m *DeviceComplianceScriptRuleError) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
func (m *DeviceComplianceScriptRuleError) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptRuleError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DeviceComplianceScriptError.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettingName sets the settingName property value. Setting name for the rule with error.
func (m *DeviceComplianceScriptRuleError) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
