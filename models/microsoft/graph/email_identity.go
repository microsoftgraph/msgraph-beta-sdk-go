package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EmailIdentity 
type EmailIdentity struct {
    Identity
    // Email address of the user.
    email *string;
}
// NewEmailIdentity instantiates a new emailIdentity and sets the default values.
func NewEmailIdentity()(*EmailIdentity) {
    m := &EmailIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// GetEmail gets the email property value. Email address of the user.
func (m *EmailIdentity) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetEmail sets the email property value. Email address of the user.
func (m *EmailIdentity) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
