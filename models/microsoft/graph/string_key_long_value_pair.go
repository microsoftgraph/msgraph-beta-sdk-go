package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type StringKeyLongValuePair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Key.
    key *string;
    // Value.
    value *int64;
}
// Instantiates a new stringKeyLongValuePair and sets the default values.
func NewStringKeyLongValuePair()(*StringKeyLongValuePair) {
    m := &StringKeyLongValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StringKeyLongValuePair) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the key property value. Key.
func (m *StringKeyLongValuePair) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// Gets the value property value. Value.
func (m *StringKeyLongValuePair) GetValue()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *StringKeyLongValuePair) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *StringKeyLongValuePair) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *StringKeyLongValuePair) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("value", m.GetValue())
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
func (m *StringKeyLongValuePair) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the key property value. Key.
// Parameters:
//  - value : Value to set for the key property.
func (m *StringKeyLongValuePair) SetKey(value *string)() {
    m.key = value
}
// Sets the value property value. Value.
// Parameters:
//  - value : Value to set for the value property.
func (m *StringKeyLongValuePair) SetValue(value *int64)() {
    m.value = value
}
