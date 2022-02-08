package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TypedEmailAddress 
type TypedEmailAddress struct {
    EmailAddress
    // To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
    otherLabel *string;
    // The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
    type_escaped *EmailType;
}
// NewTypedEmailAddress instantiates a new typedEmailAddress and sets the default values.
func NewTypedEmailAddress()(*TypedEmailAddress) {
    m := &TypedEmailAddress{
        EmailAddress: *NewEmailAddress(),
    }
    return m
}
// GetOtherLabel gets the otherLabel property value. To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
func (m *TypedEmailAddress) GetOtherLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.otherLabel
    }
}
// GetType gets the type property value. The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
func (m *TypedEmailAddress) GetType()(*EmailType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TypedEmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.EmailAddress.GetFieldDeserializers()
    res["otherLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherLabel(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*EmailType))
        }
        return nil
    }
    return res
}
func (m *TypedEmailAddress) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOtherLabel sets the otherLabel property value. To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
func (m *TypedEmailAddress) SetOtherLabel(value *string)() {
    if m != nil {
        m.otherLabel = value
    }
}
// SetType sets the type property value. The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
func (m *TypedEmailAddress) SetType(value *EmailType)() {
    if m != nil {
        m.type_escaped = value
    }
}
