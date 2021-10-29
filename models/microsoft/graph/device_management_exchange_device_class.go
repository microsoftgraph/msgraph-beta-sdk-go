package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementExchangeDeviceClass struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the device class which will be impacted by this rule.
    name *string;
    // Type of device which is impacted by this rule e.g. Model, Family. Possible values are: family, model.
    type_escaped *DeviceManagementExchangeAccessRuleType;
}
// Instantiates a new deviceManagementExchangeDeviceClass and sets the default values.
func NewDeviceManagementExchangeDeviceClass()(*DeviceManagementExchangeDeviceClass) {
    m := &DeviceManagementExchangeDeviceClass{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeDeviceClass) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the name property value. Name of the device class which will be impacted by this rule.
func (m *DeviceManagementExchangeDeviceClass) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the type_escaped property value. Type of device which is impacted by this rule e.g. Model, Family. Possible values are: family, model.
func (m *DeviceManagementExchangeDeviceClass) GetType_escaped()(*DeviceManagementExchangeAccessRuleType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessRuleType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeAccessRuleType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeDeviceClass) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementExchangeDeviceClass) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *DeviceManagementExchangeDeviceClass) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the name property value. Name of the device class which will be impacted by this rule.
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceManagementExchangeDeviceClass) SetName(value *string)() {
    m.name = value
}
// Sets the type_escaped property value. Type of device which is impacted by this rule e.g. Model, Family. Possible values are: family, model.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *DeviceManagementExchangeDeviceClass) SetType_escaped(value *DeviceManagementExchangeAccessRuleType)() {
    m.type_escaped = value
}
