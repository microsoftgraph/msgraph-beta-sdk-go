package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExchangeDeviceClass 
type DeviceManagementExchangeDeviceClass struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the device class which will be impacted by this rule.
    name *string;
    // Type of device which is impacted by this rule e.g. Model, Family. Possible values are: family, model.
    type_escaped *DeviceManagementExchangeAccessRuleType;
}
// NewDeviceManagementExchangeDeviceClass instantiates a new deviceManagementExchangeDeviceClass and sets the default values.
func NewDeviceManagementExchangeDeviceClass()(*DeviceManagementExchangeDeviceClass) {
    m := &DeviceManagementExchangeDeviceClass{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeDeviceClass) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetName gets the name property value. Name of the device class which will be impacted by this rule.
func (m *DeviceManagementExchangeDeviceClass) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetType gets the type property value. Type of device which is impacted by this rule e.g. Model, Family. Possible values are: family, model.
func (m *DeviceManagementExchangeDeviceClass) GetType()(*DeviceManagementExchangeAccessRuleType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeDeviceClass) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*DeviceManagementExchangeAccessRuleType))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeDeviceClass) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeDeviceClass) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *DeviceManagementExchangeDeviceClass) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Name of the device class which will be impacted by this rule.
func (m *DeviceManagementExchangeDeviceClass) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetType sets the type property value. Type of device which is impacted by this rule e.g. Model, Family. Possible values are: family, model.
func (m *DeviceManagementExchangeDeviceClass) SetType(value *DeviceManagementExchangeAccessRuleType)() {
    if m != nil {
        m.type_escaped = value
    }
}
