package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExchangeAccessRule 
type DeviceManagementExchangeAccessRule struct {
    // Access Level for Exchange granted by this rule. Possible values are: none, allow, block, quarantine.
    accessLevel *DeviceManagementExchangeAccessLevel;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Device Class which will be impacted by this rule.
    deviceClass *DeviceManagementExchangeDeviceClass;
}
// NewDeviceManagementExchangeAccessRule instantiates a new deviceManagementExchangeAccessRule and sets the default values.
func NewDeviceManagementExchangeAccessRule()(*DeviceManagementExchangeAccessRule) {
    m := &DeviceManagementExchangeAccessRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessLevel gets the accessLevel property value. Access Level for Exchange granted by this rule. Possible values are: none, allow, block, quarantine.
func (m *DeviceManagementExchangeAccessRule) GetAccessLevel()(*DeviceManagementExchangeAccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.accessLevel
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeAccessRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceClass gets the deviceClass property value. Device Class which will be impacted by this rule.
func (m *DeviceManagementExchangeAccessRule) GetDeviceClass()(*DeviceManagementExchangeDeviceClass) {
    if m == nil {
        return nil
    } else {
        return m.deviceClass
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeAccessRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessLevel(val.(*DeviceManagementExchangeAccessLevel))
        }
        return nil
    }
    res["deviceClass"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeDeviceClass() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceClass(val.(*DeviceManagementExchangeDeviceClass))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeAccessRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeAccessRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessLevel() != nil {
        cast := (*m.GetAccessLevel()).String()
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
// SetAccessLevel sets the accessLevel property value. Access Level for Exchange granted by this rule. Possible values are: none, allow, block, quarantine.
func (m *DeviceManagementExchangeAccessRule) SetAccessLevel(value *DeviceManagementExchangeAccessLevel)() {
    if m != nil {
        m.accessLevel = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeAccessRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceClass sets the deviceClass property value. Device Class which will be impacted by this rule.
func (m *DeviceManagementExchangeAccessRule) SetDeviceClass(value *DeviceManagementExchangeDeviceClass)() {
    if m != nil {
        m.deviceClass = value
    }
}
