package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserPurpose struct {
    additionalData map[string]interface{};
    value *MailboxRecipientType;
}
func NewUserPurpose()(*UserPurpose) {
    m := &UserPurpose{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserPurpose) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserPurpose) GetValue()(*MailboxRecipientType) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *UserPurpose) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMailboxRecipientType)
        if err != nil {
            return err
        }
        cast := val.(MailboxRecipientType)
        m.SetValue(&cast)
        return nil
    }
    return res
}
func (m *UserPurpose) IsNil()(bool) {
    return m == nil
}
func (m *UserPurpose) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := m.GetValue().String()
        err := writer.WriteStringValue("value", &cast)
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
func (m *UserPurpose) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserPurpose) SetValue(value *MailboxRecipientType)() {
    m.value = value
}
