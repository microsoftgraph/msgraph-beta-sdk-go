package haspayloadlinks

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HasPayloadLinksRequestBody 
type HasPayloadLinksRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    payloadIds []string;
}
// NewHasPayloadLinksRequestBody instantiates a new hasPayloadLinksRequestBody and sets the default values.
func NewHasPayloadLinksRequestBody()(*HasPayloadLinksRequestBody) {
    m := &HasPayloadLinksRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HasPayloadLinksRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPayloadIds gets the payloadIds property value. 
func (m *HasPayloadLinksRequestBody) GetPayloadIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.payloadIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HasPayloadLinksRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["payloadIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPayloadIds(res)
        }
        return nil
    }
    return res
}
func (m *HasPayloadLinksRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HasPayloadLinksRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetPayloadIds() != nil {
        err := writer.WriteCollectionOfStringValues("payloadIds", m.GetPayloadIds())
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
func (m *HasPayloadLinksRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPayloadIds sets the payloadIds property value. 
func (m *HasPayloadLinksRequestBody) SetPayloadIds(value []string)() {
    if m != nil {
        m.payloadIds = value
    }
}
