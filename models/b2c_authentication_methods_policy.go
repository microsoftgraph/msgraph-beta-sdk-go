package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// B2cAuthenticationMethodsPolicy 
type B2cAuthenticationMethodsPolicy struct {
    Entity
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
    odataTypeValue := "#microsoft.graph.b2cAuthenticationMethodsPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cAuthenticationMethodsPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cAuthenticationMethodsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEmailPasswordAuthenticationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEmailPasswordAuthenticationEnabled)
    res["isPhoneOneTimePasswordAuthenticationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsPhoneOneTimePasswordAuthenticationEnabled)
    res["isUserNameAuthenticationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsUserNameAuthenticationEnabled)
    return res
}
// GetIsEmailPasswordAuthenticationEnabled gets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsEmailPasswordAuthenticationEnabled()(*bool) {
    return m.isEmailPasswordAuthenticationEnabled
}
// GetIsPhoneOneTimePasswordAuthenticationEnabled gets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsPhoneOneTimePasswordAuthenticationEnabled()(*bool) {
    return m.isPhoneOneTimePasswordAuthenticationEnabled
}
// GetIsUserNameAuthenticationEnabled gets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsUserNameAuthenticationEnabled()(*bool) {
    return m.isUserNameAuthenticationEnabled
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
    return nil
}
// SetIsEmailPasswordAuthenticationEnabled sets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsEmailPasswordAuthenticationEnabled(value *bool)() {
    m.isEmailPasswordAuthenticationEnabled = value
}
// SetIsPhoneOneTimePasswordAuthenticationEnabled sets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsPhoneOneTimePasswordAuthenticationEnabled(value *bool)() {
    m.isPhoneOneTimePasswordAuthenticationEnabled = value
}
// SetIsUserNameAuthenticationEnabled sets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsUserNameAuthenticationEnabled(value *bool)() {
    m.isUserNameAuthenticationEnabled = value
}
