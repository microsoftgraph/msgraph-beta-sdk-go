package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementExchangeDeviceClass struct {
    additionalData map[string]interface{};
    name *string;
    type_escpaped *DeviceManagementExchangeAccessRuleType;
}
func NewDeviceManagementExchangeDeviceClass()(*DeviceManagementExchangeDeviceClass) {
    m := &DeviceManagementExchangeDeviceClass{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementExchangeDeviceClass) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementExchangeDeviceClass) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *DeviceManagementExchangeDeviceClass) GetType_escpaped()(*DeviceManagementExchangeAccessRuleType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *DeviceManagementExchangeDeviceClass) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessRuleType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeAccessRuleType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeDeviceClass) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementExchangeDeviceClass) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err := writer.WriteStringValue("type_escpaped", &cast)
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
func (m *DeviceManagementExchangeDeviceClass) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementExchangeDeviceClass) SetName(value *string)() {
    m.name = value
}
func (m *DeviceManagementExchangeDeviceClass) SetType_escpaped(value *DeviceManagementExchangeAccessRuleType)() {
    m.type_escpaped = value
}
