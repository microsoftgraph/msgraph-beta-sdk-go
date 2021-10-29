package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SoftwareOathAuthenticationMethod struct {
    AuthenticationMethod
    // The secret key of the method. Always returns null.
    secretKey *string;
}
// Instantiates a new softwareOathAuthenticationMethod and sets the default values.
func NewSoftwareOathAuthenticationMethod()(*SoftwareOathAuthenticationMethod) {
    m := &SoftwareOathAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// Gets the secretKey property value. The secret key of the method. Always returns null.
func (m *SoftwareOathAuthenticationMethod) GetSecretKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secretKey
    }
}
// The deserialization information for the current model
func (m *SoftwareOathAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["secretKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSecretKey(val)
        return nil
    }
    return res
}
func (m *SoftwareOathAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the secretKey property value. The secret key of the method. Always returns null.
// Parameters:
//  - value : Value to set for the secretKey property.
func (m *SoftwareOathAuthenticationMethod) SetSecretKey(value *string)() {
    m.secretKey = value
}
