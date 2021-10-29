package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementApplicabilityRuleOsVersion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Max OS version for Applicability Rule.
    maxOSVersion *string;
    // Min OS version for Applicability Rule.
    minOSVersion *string;
    // Name for object.
    name *string;
    // Applicability Rule type. Possible values are: include, exclude.
    ruleType *DeviceManagementApplicabilityRuleType;
}
// Instantiates a new deviceManagementApplicabilityRuleOsVersion and sets the default values.
func NewDeviceManagementApplicabilityRuleOsVersion()(*DeviceManagementApplicabilityRuleOsVersion) {
    m := &DeviceManagementApplicabilityRuleOsVersion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the maxOSVersion property value. Max OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMaxOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maxOSVersion
    }
}
// Gets the minOSVersion property value. Min OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMinOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minOSVersion
    }
}
// Gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
// The deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleOsVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maxOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaxOSVersion(val)
        return nil
    }
    res["minOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinOSVersion(val)
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
func (m *DeviceManagementApplicabilityRuleOsVersion) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementApplicabilityRuleOsVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("maxOSVersion", m.GetMaxOSVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("minOSVersion", m.GetMinOSVersion())
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
func (m *DeviceManagementApplicabilityRuleOsVersion) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the maxOSVersion property value. Max OS version for Applicability Rule.
// Parameters:
//  - value : Value to set for the maxOSVersion property.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMaxOSVersion(value *string)() {
    m.maxOSVersion = value
}
// Sets the minOSVersion property value. Min OS version for Applicability Rule.
// Parameters:
//  - value : Value to set for the minOSVersion property.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMinOSVersion(value *string)() {
    m.minOSVersion = value
}
// Sets the name property value. Name for object.
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetName(value *string)() {
    m.name = value
}
// Sets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
// Parameters:
//  - value : Value to set for the ruleType property.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    m.ruleType = value
}
