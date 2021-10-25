package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MentionsPreview struct {
    additionalData map[string]interface{};
    isMentioned *bool;
}
func NewMentionsPreview()(*MentionsPreview) {
    m := &MentionsPreview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MentionsPreview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MentionsPreview) GetIsMentioned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMentioned
    }
}
func (m *MentionsPreview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isMentioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMentioned(val)
        return nil
    }
    return res
}
func (m *MentionsPreview) IsNil()(bool) {
    return m == nil
}
func (m *MentionsPreview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isMentioned", m.GetIsMentioned())
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
func (m *MentionsPreview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MentionsPreview) SetIsMentioned(value *bool)() {
    m.isMentioned = value
}
