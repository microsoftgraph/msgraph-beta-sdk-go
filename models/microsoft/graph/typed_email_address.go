package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TypedEmailAddress struct {
    EmailAddress
    otherLabel *string;
    type_escpaped *EmailType;
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
func (m *TypedEmailAddress) GetType_escpaped()(*EmailType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
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
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailType)
        if err != nil {
            return err
        }
        cast := val.(EmailType)
        m.SetType_escpaped(&cast)
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
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err = writer.WriteStringValue("type_escpaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TypedEmailAddress) SetOtherLabel(value *string)() {
    m.otherLabel = value
}
func (m *TypedEmailAddress) SetType_escpaped(value *EmailType)() {
    m.type_escpaped = value
}
