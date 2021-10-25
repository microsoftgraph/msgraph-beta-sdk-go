package haspayloadlinks

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type HasPayloadLinksRequestBody struct {
    additionalData map[string]interface{};
    payloadIds []string;
}
func NewHasPayloadLinksRequestBody()(*HasPayloadLinksRequestBody) {
    m := &HasPayloadLinksRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *HasPayloadLinksRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *HasPayloadLinksRequestBody) GetPayloadIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.payloadIds
    }
}
func (m *HasPayloadLinksRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["payloadIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPayloadIds(res)
        return nil
    }
    return res
}
func (m *HasPayloadLinksRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *HasPayloadLinksRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
func (m *HasPayloadLinksRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *HasPayloadLinksRequestBody) SetPayloadIds(value []string)() {
    m.payloadIds = value
}
