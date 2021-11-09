package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EmailIdentity struct {
    Identity
    // Email address of the user.
    email *string;
}
// Instantiates a new emailIdentity and sets the default values.
func NewEmailIdentity()(*EmailIdentity) {
    m := &EmailIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// Gets the email property value. Email address of the user.
func (m *EmailIdentity) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// The deserialization information for the current model
func (m *EmailIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    return res
}
func (m *EmailIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EmailIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the email property value. Email address of the user.
// Parameters:
//  - value : Value to set for the email property.
func (m *EmailIdentity) SetEmail(value *string)() {
    m.email = value
}
