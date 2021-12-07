package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChromeOSDeviceProperty 
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
// NewChromeOSDeviceProperty instantiates a new chromeOSDeviceProperty and sets the default values.
func NewChromeOSDeviceProperty()(*ChromeOSDeviceProperty) {
    m := &ChromeOSDeviceProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChromeOSDeviceProperty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetName gets the name property value. Name of the property
func (m *ChromeOSDeviceProperty) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetUpdatable gets the updatable property value. Whether this property is updatable
func (m *ChromeOSDeviceProperty) GetUpdatable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.updatable
    }
}
// GetValue gets the value property value. Value of the property
func (m *ChromeOSDeviceProperty) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetValueType gets the valueType property value. Type of the value
func (m *ChromeOSDeviceProperty) GetValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChromeOSDeviceProperty) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Name of the property
func (m *ChromeOSDeviceProperty) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetUpdatable sets the updatable property value. Whether this property is updatable
func (m *ChromeOSDeviceProperty) SetUpdatable(value *bool)() {
    if m != nil {
        m.updatable = value
    }
}
// SetValue sets the value property value. Value of the property
func (m *ChromeOSDeviceProperty) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
// SetValueType sets the valueType property value. Type of the value
func (m *ChromeOSDeviceProperty) SetValueType(value *string)() {
    if m != nil {
        m.valueType = value
    }
}
