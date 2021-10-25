package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementApplicabilityRuleOsEdition struct {
    additionalData map[string]interface{};
    name *string;
    osEditionTypes []Windows10EditionType;
    ruleType *DeviceManagementApplicabilityRuleType;
}
func NewDeviceManagementApplicabilityRuleOsEdition()(*DeviceManagementApplicabilityRuleOsEdition) {
    m := &DeviceManagementApplicabilityRuleOsEdition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementApplicabilityRuleOsEdition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementApplicabilityRuleOsEdition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *DeviceManagementApplicabilityRuleOsEdition) GetOsEditionTypes()([]Windows10EditionType) {
    if m == nil {
        return nil
    } else {
        return m.osEditionTypes
    }
}
func (m *DeviceManagementApplicabilityRuleOsEdition) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
func (m *DeviceManagementApplicabilityRuleOsEdition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["osEditionTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseWindows10EditionType)
        if err != nil {
            return err
        }
        res := make([]Windows10EditionType, len(val))
        for i, v := range val {
            res[i] = *(v.(*Windows10EditionType))
        }
        m.SetOsEditionTypes(res)
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
func (m *DeviceManagementApplicabilityRuleOsEdition) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementApplicabilityRuleOsEdition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("osEditionTypes", SerializeWindows10EditionType(m.GetOsEditionTypes()))
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
func (m *DeviceManagementApplicabilityRuleOsEdition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementApplicabilityRuleOsEdition) SetName(value *string)() {
    m.name = value
}
func (m *DeviceManagementApplicabilityRuleOsEdition) SetOsEditionTypes(value []Windows10EditionType)() {
    m.osEditionTypes = value
}
func (m *DeviceManagementApplicabilityRuleOsEdition) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    m.ruleType = value
}
