package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TypedEmailAddress struct {
    EmailAddress
    otherLabel *string;
    type_escaped *EmailType;
}
func NewTypedEmailAddress()(*TypedEmailAddress) {
    m := &TypedEmailAddress{
        EmailAddress: *NewEmailAddress(),
    }
    return m
}
func (m *TypedEmailAddress) GetOtherLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.otherLabel
    }
}
func (m *TypedEmailAddress) GetType_escaped()(*EmailType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *TypedEmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.EmailAddress.GetFieldDeserializers()
    res["otherLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOtherLabel(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailType)
        if err != nil {
            return err
        }
        cast := val.(EmailType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *TypedEmailAddress) IsNil()(bool) {
    return m == nil
}
func (m *TypedEmailAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.EmailAddress.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("otherLabel", m.GetOtherLabel())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TypedEmailAddress) SetOtherLabel(value *string)() {
    m.otherLabel = value
}
func (m *TypedEmailAddress) SetType_escaped(value *EmailType)() {
    m.type_escaped = value
}
