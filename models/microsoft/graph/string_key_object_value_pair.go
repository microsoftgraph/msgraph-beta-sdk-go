package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// StringKeyObjectValuePair provides operations to call the parseExpression method.
type StringKeyObjectValuePair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Key.
    key *string;
}
// NewStringKeyObjectValuePair instantiates a new stringKeyObjectValuePair and sets the default values.
func NewStringKeyObjectValuePair()(*StringKeyObjectValuePair) {
    m := &StringKeyObjectValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStringKeyObjectValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyObjectValuePairFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewStringKeyObjectValuePair(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StringKeyObjectValuePair) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyObjectValuePair) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
// GetKey gets the key property value. Key.
func (m *StringKeyObjectValuePair) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *StringKeyObjectValuePair) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StringKeyObjectValuePair) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *StringKeyObjectValuePair) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKey sets the key property value. Key.
func (m *StringKeyObjectValuePair) SetKey(value *string)() {
    if m != nil {
        m.key = value
    }
}
