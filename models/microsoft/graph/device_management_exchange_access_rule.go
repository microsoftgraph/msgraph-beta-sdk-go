package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementExchangeAccessRule struct {
    accessLevel *DeviceManagementExchangeAccessLevel;
    additionalData map[string]interface{};
    deviceClass *DeviceManagementExchangeDeviceClass;
}
func NewDeviceManagementExchangeAccessRule()(*DeviceManagementExchangeAccessRule) {
    m := &DeviceManagementExchangeAccessRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementExchangeAccessRule) GetAccessLevel()(*DeviceManagementExchangeAccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.accessLevel
    }
}
func (m *DeviceManagementExchangeAccessRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementExchangeAccessRule) GetDeviceClass()(*DeviceManagementExchangeDeviceClass) {
    if m == nil {
        return nil
    } else {
        return m.deviceClass
    }
}
func (m *DeviceManagementExchangeAccessRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessLevel)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeAccessLevel)
        m.SetAccessLevel(&cast)
        return nil
    }
    res["deviceClass"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeDeviceClass() })
        if err != nil {
            return err
        }
        m.SetDeviceClass(val.(*DeviceManagementExchangeDeviceClass))
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeAccessRule) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementExchangeAccessRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessLevel() != nil {
        cast := m.GetAccessLevel().String()
        err := writer.WriteStringValue("accessLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceClass", m.GetDeviceClass())
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
func (m *DeviceManagementExchangeAccessRule) SetAccessLevel(value *DeviceManagementExchangeAccessLevel)() {
    m.accessLevel = value
}
func (m *DeviceManagementExchangeAccessRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementExchangeAccessRule) SetDeviceClass(value *DeviceManagementExchangeDeviceClass)() {
    m.deviceClass = value
}
