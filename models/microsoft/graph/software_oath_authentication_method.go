package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SoftwareOathAuthenticationMethod 
type SoftwareOathAuthenticationMethod struct {
    AuthenticationMethod
    // The secret key of the method. Always returns null.
    secretKey *string;
}
// NewSoftwareOathAuthenticationMethod instantiates a new softwareOathAuthenticationMethod and sets the default values.
func NewSoftwareOathAuthenticationMethod()(*SoftwareOathAuthenticationMethod) {
    m := &SoftwareOathAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// GetSecretKey gets the secretKey property value. The secret key of the method. Always returns null.
func (m *SoftwareOathAuthenticationMethod) GetSecretKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secretKey
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareOathAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["secretKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretKey(val)
        }
        return nil
    }
    return res
}
func (m *SoftwareOathAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SoftwareOathAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("secretKey", m.GetSecretKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSecretKey sets the secretKey property value. The secret key of the method. Always returns null.
func (m *SoftwareOathAuthenticationMethod) SetSecretKey(value *string)() {
    if m != nil {
        m.secretKey = value
    }
}
