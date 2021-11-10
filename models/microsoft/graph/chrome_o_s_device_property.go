package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChromeOSDeviceProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the property
    name *string;
    // Whether this property is updatable
    updatable *bool;
    // Value of the property
    value *string;
    // Type of the value
    valueType *string;
}
// Instantiates a new chromeOSDeviceProperty and sets the default values.
func NewChromeOSDeviceProperty()(*ChromeOSDeviceProperty) {
    m := &ChromeOSDeviceProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChromeOSDeviceProperty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the name property value. Name of the property
func (m *ChromeOSDeviceProperty) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the updatable property value. Whether this property is updatable
func (m *ChromeOSDeviceProperty) GetUpdatable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.updatable
    }
}
// Gets the value property value. Value of the property
func (m *ChromeOSDeviceProperty) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Gets the valueType property value. Type of the value
func (m *ChromeOSDeviceProperty) GetValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// The deserialization information for the current model
func (m *ChromeOSDeviceProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["updatable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatable(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val)
        }
        return nil
    }
    return res
}
func (m *ChromeOSDeviceProperty) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ChromeOSDeviceProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("updatable", m.GetUpdatable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("valueType", m.GetValueType())
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
func (m *ChromeOSDeviceProperty) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the name property value. Name of the property
// Parameters:
//  - value : Value to set for the name property.
func (m *ChromeOSDeviceProperty) SetName(value *string)() {
    m.name = value
}
// Sets the updatable property value. Whether this property is updatable
// Parameters:
//  - value : Value to set for the updatable property.
func (m *ChromeOSDeviceProperty) SetUpdatable(value *bool)() {
    m.updatable = value
}
// Sets the value property value. Value of the property
// Parameters:
//  - value : Value to set for the value property.
func (m *ChromeOSDeviceProperty) SetValue(value *string)() {
    m.value = value
}
// Sets the valueType property value. Type of the value
// Parameters:
//  - value : Value to set for the valueType property.
func (m *ChromeOSDeviceProperty) SetValueType(value *string)() {
    m.valueType = value
}
