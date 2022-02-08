package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementApplicabilityRuleOsVersion 
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
// NewDeviceManagementApplicabilityRuleOsVersion instantiates a new deviceManagementApplicabilityRuleOsVersion and sets the default values.
func NewDeviceManagementApplicabilityRuleOsVersion()(*DeviceManagementApplicabilityRuleOsVersion) {
    m := &DeviceManagementApplicabilityRuleOsVersion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMaxOSVersion gets the maxOSVersion property value. Max OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMaxOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maxOSVersion
    }
}
// GetMinOSVersion gets the minOSVersion property value. Min OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMinOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minOSVersion
    }
}
// GetName gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetRuleType gets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleOsVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maxOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxOSVersion(val)
        }
        return nil
    }
    res["minOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinOSVersion(val)
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
func (m *DeviceManagementApplicabilityRuleOsVersion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
func (m *DeviceManagementApplicabilityRuleOsVersion) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaxOSVersion sets the maxOSVersion property value. Max OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMaxOSVersion(value *string)() {
    if m != nil {
        m.maxOSVersion = value
    }
}
// SetMinOSVersion sets the minOSVersion property value. Min OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMinOSVersion(value *string)() {
    if m != nil {
        m.minOSVersion = value
    }
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetRuleType sets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    if m != nil {
        m.ruleType = value
    }
}
