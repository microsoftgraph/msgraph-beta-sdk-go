package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TypedEmailAddress struct {
    EmailAddress
    // To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
    otherLabel *string;
    // The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
    type_escaped *EmailType;
}
// Instantiates a new typedEmailAddress and sets the default values.
func NewTypedEmailAddress()(*TypedEmailAddress) {
    m := &TypedEmailAddress{
        EmailAddress: *NewEmailAddress(),
    }
    return m
}
// Gets the otherLabel property value. To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
func (m *TypedEmailAddress) GetOtherLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.otherLabel
    }
}
// Gets the type_escaped property value. The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
func (m *TypedEmailAddress) GetType_escaped()(*EmailType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the otherLabel property value. To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
// Parameters:
//  - value : Value to set for the otherLabel property.
func (m *TypedEmailAddress) SetOtherLabel(value *string)() {
    m.otherLabel = value
}
// Sets the type_escaped property value. The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *TypedEmailAddress) SetType_escaped(value *EmailType)() {
    m.type_escaped = value
}
