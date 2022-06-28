package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// B2cAuthenticationMethodsPolicy 
type B2cAuthenticationMethodsPolicy struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
    isEmailPasswordAuthenticationEnabled *bool
    // The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
    isPhoneOneTimePasswordAuthenticationEnabled *bool
    // The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
    isUserNameAuthenticationEnabled *bool
}
// NewB2cAuthenticationMethodsPolicy instantiates a new B2cAuthenticationMethodsPolicy and sets the default values.
func NewB2cAuthenticationMethodsPolicy()(*B2cAuthenticationMethodsPolicy) {
    m := &B2cAuthenticationMethodsPolicy{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cAuthenticationMethodsPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *B2cAuthenticationMethodsPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cAuthenticationMethodsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEmailPasswordAuthenticationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEmailPasswordAuthenticationEnabled(val)
        }
        return nil
    }
    res["isPhoneOneTimePasswordAuthenticationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPhoneOneTimePasswordAuthenticationEnabled(val)
        }
        return nil
    }
    res["isUserNameAuthenticationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUserNameAuthenticationEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsEmailPasswordAuthenticationEnabled gets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsEmailPasswordAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEmailPasswordAuthenticationEnabled
    }
}
// GetIsPhoneOneTimePasswordAuthenticationEnabled gets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsPhoneOneTimePasswordAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPhoneOneTimePasswordAuthenticationEnabled
    }
}
// GetIsUserNameAuthenticationEnabled gets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsUserNameAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUserNameAuthenticationEnabled
    }
}
// Serialize serializes information the current object
func (m *B2cAuthenticationMethodsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEmailPasswordAuthenticationEnabled", m.GetIsEmailPasswordAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPhoneOneTimePasswordAuthenticationEnabled", m.GetIsPhoneOneTimePasswordAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isUserNameAuthenticationEnabled", m.GetIsUserNameAuthenticationEnabled())
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
func (m *B2cAuthenticationMethodsPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsEmailPasswordAuthenticationEnabled sets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsEmailPasswordAuthenticationEnabled(value *bool)() {
    if m != nil {
        m.isEmailPasswordAuthenticationEnabled = value
    }
}
// SetIsPhoneOneTimePasswordAuthenticationEnabled sets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsPhoneOneTimePasswordAuthenticationEnabled(value *bool)() {
    if m != nil {
        m.isPhoneOneTimePasswordAuthenticationEnabled = value
    }
}
// SetIsUserNameAuthenticationEnabled sets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsUserNameAuthenticationEnabled(value *bool)() {
    if m != nil {
        m.isUserNameAuthenticationEnabled = value
    }
}
