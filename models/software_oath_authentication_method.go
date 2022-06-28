package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SoftwareOathAuthenticationMethod 
type SoftwareOathAuthenticationMethod struct {
    AuthenticationMethod
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The secret key of the method. Always returns null.
    secretKey *string
}
// NewSoftwareOathAuthenticationMethod instantiates a new SoftwareOathAuthenticationMethod and sets the default values.
func NewSoftwareOathAuthenticationMethod()(*SoftwareOathAuthenticationMethod) {
    m := &SoftwareOathAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSoftwareOathAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSoftwareOathAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSoftwareOathAuthenticationMethod(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SoftwareOathAuthenticationMethod) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareOathAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["secretKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetSecretKey gets the secretKey property value. The secret key of the method. Always returns null.
func (m *SoftwareOathAuthenticationMethod) GetSecretKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secretKey
    }
}
// Serialize serializes information the current object
func (m *SoftwareOathAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SoftwareOathAuthenticationMethod) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSecretKey sets the secretKey property value. The secret key of the method. Always returns null.
func (m *SoftwareOathAuthenticationMethod) SetSecretKey(value *string)() {
    if m != nil {
        m.secretKey = value
    }
}
