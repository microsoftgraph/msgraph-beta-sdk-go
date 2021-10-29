package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementApplicabilityRuleDeviceMode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Applicability rule for device mode. Possible values are: standardConfiguration, sModeConfiguration.
    deviceMode *Windows10DeviceModeType;
    // Name for object.
    name *string;
    // Applicability Rule type. Possible values are: include, exclude.
    ruleType *DeviceManagementApplicabilityRuleType;
}
// Instantiates a new deviceManagementApplicabilityRuleDeviceMode and sets the default values.
func NewDeviceManagementApplicabilityRuleDeviceMode()(*DeviceManagementApplicabilityRuleDeviceMode) {
    m := &DeviceManagementApplicabilityRuleDeviceMode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceMode property value. Applicability rule for device mode. Possible values are: standardConfiguration, sModeConfiguration.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetDeviceMode()(*Windows10DeviceModeType) {
    if m == nil {
        return nil
    } else {
        return m.deviceMode
    }
}
// Gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
// The deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10DeviceModeType)
        if err != nil {
            return err
        }
        cast := val.(Windows10DeviceModeType)
        m.SetDeviceMode(&cast)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["ruleType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementApplicabilityRuleType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementApplicabilityRuleType)
        m.SetRuleType(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementApplicabilityRuleDeviceMode) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementApplicabilityRuleDeviceMode) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDeviceMode() != nil {
        cast := m.GetDeviceMode().String()
        err := writer.WriteStringValue("deviceMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRuleType() != nil {
        cast := m.GetRuleType().String()
        err := writer.WriteStringValue("ruleType", &cast)
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
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceMode property value. Applicability rule for device mode. Possible values are: standardConfiguration, sModeConfiguration.
// Parameters:
//  - value : Value to set for the deviceMode property.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetDeviceMode(value *Windows10DeviceModeType)() {
    m.deviceMode = value
}
// Sets the name property value. Name for object.
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetName(value *string)() {
    m.name = value
}
// Sets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
// Parameters:
//  - value : Value to set for the ruleType property.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    m.ruleType = value
}
