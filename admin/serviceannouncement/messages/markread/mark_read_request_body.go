package markread

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MarkReadRequestBody struct {
    additionalData map[string]interface{};
    messageIds []string;
}
func NewMarkReadRequestBody()(*MarkReadRequestBody) {
    m := &MarkReadRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MarkReadRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MarkReadRequestBody) GetMessageIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.messageIds
    }
}
func (m *MarkReadRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["messageIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMessageIds(res)
        return nil
    }
    return res
}
func (m *MarkReadRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *MarkReadRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("messageIds", m.GetMessageIds())
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
func (m *MarkReadRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MarkReadRequestBody) SetMessageIds(value []string)() {
    m.messageIds = value
}
