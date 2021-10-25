package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DlpNotification struct {
    additionalData map[string]interface{};
    author *string;
}
func NewDlpNotification()(*DlpNotification) {
    m := &DlpNotification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DlpNotification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DlpNotification) GetAuthor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.author
    }
}
func (m *DlpNotification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["author"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthor(val)
        return nil
    }
    return res
}
func (m *DlpNotification) IsNil()(bool) {
    return m == nil
}
func (m *DlpNotification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("author", m.GetAuthor())
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
func (m *DlpNotification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DlpNotification) SetAuthor(value *string)() {
    m.author = value
}
