package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// B2cAuthenticationMethodsPolicy 
type B2cAuthenticationMethodsPolicy struct {
    Entity
}
// NewB2cAuthenticationMethodsPolicy instantiates a new b2cAuthenticationMethodsPolicy and sets the default values.
func NewB2cAuthenticationMethodsPolicy()(*B2cAuthenticationMethodsPolicy) {
    m := &B2cAuthenticationMethodsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cAuthenticationMethodsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cAuthenticationMethodsPolicy(), nil
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetIsEmailPasswordAuthenticationEnabled gets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsEmailPasswordAuthenticationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEmailPasswordAuthenticationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPhoneOneTimePasswordAuthenticationEnabled gets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsPhoneOneTimePasswordAuthenticationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isPhoneOneTimePasswordAuthenticationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsUserNameAuthenticationEnabled gets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsUserNameAuthenticationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isUserNameAuthenticationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *B2cAuthenticationMethodsPolicy) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsEmailPasswordAuthenticationEnabled sets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsEmailPasswordAuthenticationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEmailPasswordAuthenticationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPhoneOneTimePasswordAuthenticationEnabled sets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsPhoneOneTimePasswordAuthenticationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isPhoneOneTimePasswordAuthenticationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsUserNameAuthenticationEnabled sets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsUserNameAuthenticationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isUserNameAuthenticationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *B2cAuthenticationMethodsPolicy) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// B2cAuthenticationMethodsPolicyable 
type B2cAuthenticationMethodsPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsEmailPasswordAuthenticationEnabled()(*bool)
    GetIsPhoneOneTimePasswordAuthenticationEnabled()(*bool)
    GetIsUserNameAuthenticationEnabled()(*bool)
    GetOdataType()(*string)
    SetIsEmailPasswordAuthenticationEnabled(value *bool)()
    SetIsPhoneOneTimePasswordAuthenticationEnabled(value *bool)()
    SetIsUserNameAuthenticationEnabled(value *bool)()
    SetOdataType(value *string)()
}
