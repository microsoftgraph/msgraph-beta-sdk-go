package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementApplicabilityRuleDeviceMode 
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
// NewDeviceManagementApplicabilityRuleDeviceMode instantiates a new deviceManagementApplicabilityRuleDeviceMode and sets the default values.
func NewDeviceManagementApplicabilityRuleDeviceMode()(*DeviceManagementApplicabilityRuleDeviceMode) {
    m := &DeviceManagementApplicabilityRuleDeviceMode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceMode gets the deviceMode property value. Applicability rule for device mode. Possible values are: standardConfiguration, sModeConfiguration.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetDeviceMode()(*Windows10DeviceModeType) {
    if m == nil {
        return nil
    } else {
        return m.deviceMode
    }
}
// GetName gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetRuleType gets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10DeviceModeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMode(val.(*Windows10DeviceModeType))
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["ruleType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementApplicabilityRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleType(val.(*DeviceManagementApplicabilityRuleType))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementApplicabilityRuleDeviceMode) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementApplicabilityRuleDeviceMode) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDeviceMode() != nil {
        cast := (*m.GetDeviceMode()).String()
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
        cast := (*m.GetRuleType()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceMode sets the deviceMode property value. Applicability rule for device mode. Possible values are: standardConfiguration, sModeConfiguration.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetDeviceMode(value *Windows10DeviceModeType)() {
    if m != nil {
        m.deviceMode = value
    }
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetRuleType sets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    if m != nil {
        m.ruleType = value
    }
}
