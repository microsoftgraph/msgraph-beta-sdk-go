package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceComplianceScriptError struct {
    additionalData map[string]interface{};
    code *Code;
    deviceComplianceScriptRulesValidationError *DeviceComplianceScriptRulesValidationError;
    message *string;
}
func NewDeviceComplianceScriptError()(*DeviceComplianceScriptError) {
    m := &DeviceComplianceScriptError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceComplianceScriptError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceComplianceScriptError) GetCode()(*Code) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
func (m *DeviceComplianceScriptError) GetDeviceComplianceScriptRulesValidationError()(*DeviceComplianceScriptRulesValidationError) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptRulesValidationError
    }
}
func (m *DeviceComplianceScriptError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *DeviceComplianceScriptError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCode)
        if err != nil {
            return err
        }
        cast := val.(Code)
        m.SetCode(&cast)
        return nil
    }
    res["deviceComplianceScriptRulesValidationError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRulesValidationError)
        if err != nil {
            return err
        }
        cast := val.(DeviceComplianceScriptRulesValidationError)
        m.SetDeviceComplianceScriptRulesValidationError(&cast)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptError) IsNil()(bool) {
    return m == nil
}
func (m *DeviceComplianceScriptError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCode() != nil {
        cast := m.GetCode().String()
        err := writer.WriteStringValue("code", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRulesValidationError() != nil {
        cast := m.GetDeviceComplianceScriptRulesValidationError().String()
        err := writer.WriteStringValue("deviceComplianceScriptRulesValidationError", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *DeviceComplianceScriptError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceComplianceScriptError) SetCode(value *Code)() {
    m.code = value
}
func (m *DeviceComplianceScriptError) SetDeviceComplianceScriptRulesValidationError(value *DeviceComplianceScriptRulesValidationError)() {
    m.deviceComplianceScriptRulesValidationError = value
}
func (m *DeviceComplianceScriptError) SetMessage(value *string)() {
    m.message = value
}
