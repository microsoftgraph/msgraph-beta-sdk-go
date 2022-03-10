package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// StringKeyAttributeMappingSourceValuePair provides operations to manage the collection of application entities.
type StringKeyAttributeMappingSourceValuePair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The name of the parameter.
    key *string;
    // The value of the parameter.
    value AttributeMappingSourceable;
}
// NewStringKeyAttributeMappingSourceValuePair instantiates a new stringKeyAttributeMappingSourceValuePair and sets the default values.
func NewStringKeyAttributeMappingSourceValuePair()(*StringKeyAttributeMappingSourceValuePair) {
    m := &StringKeyAttributeMappingSourceValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewStringKeyAttributeMappingSourceValuePair(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StringKeyAttributeMappingSourceValuePair) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyAttributeMappingSourceValuePair) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttributeMappingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(AttributeMappingSourceable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetValue gets the value property value. The value of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) GetValue()(AttributeMappingSourceable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *StringKeyAttributeMappingSourceValuePair) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StringKeyAttributeMappingSourceValuePair) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *StringKeyAttributeMappingSourceValuePair) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKey sets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) SetKey(value *string)() {
    if m != nil {
        m.key = value
    }
}
// SetValue sets the value property value. The value of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) SetValue(value AttributeMappingSourceable)() {
    if m != nil {
        m.value = value
    }
}
