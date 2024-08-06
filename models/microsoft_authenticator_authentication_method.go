package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MicrosoftAuthenticatorAuthenticationMethod struct {
    AuthenticationMethod
}
// NewMicrosoftAuthenticatorAuthenticationMethod instantiates a new MicrosoftAuthenticatorAuthenticationMethod and sets the default values.
func NewMicrosoftAuthenticatorAuthenticationMethod()(*MicrosoftAuthenticatorAuthenticationMethod) {
    m := &MicrosoftAuthenticatorAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    odataTypeValue := "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorAuthenticationMethod(), nil
}
// GetClientAppName gets the clientAppName property value. The app that the user has registered to use to approve push notifications. The possible values are: microsoftAuthenticator, outlookMobile, unknownFutureValue.
// returns a *MicrosoftAuthenticatorAuthenticationMethodClientAppName when successful
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetClientAppName()(*MicrosoftAuthenticatorAuthenticationMethodClientAppName) {
    val, err := m.GetBackingStore().Get("clientAppName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftAuthenticatorAuthenticationMethodClientAppName)
    }
    return nil
}
// GetDevice gets the device property value. The registered device on which Microsoft Authenticator resides. This property is null if the device isn't registered for passwordless Phone Sign-In.
// returns a Deviceable when successful
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDevice()(Deviceable) {
    val, err := m.GetBackingStore().Get("device")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Deviceable)
    }
    return nil
}
// GetDeviceTag gets the deviceTag property value. Tags containing app metadata.
// returns a *string when successful
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDeviceTag()(*string) {
    val, err := m.GetBackingStore().Get("deviceTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the device on which this app is registered.
// returns a *string when successful
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["clientAppName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftAuthenticatorAuthenticationMethodClientAppName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppName(val.(*MicrosoftAuthenticatorAuthenticationMethodClientAppName))
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(Deviceable))
        }
        return nil
    }
    res["deviceTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["phoneAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneAppVersion(val)
        }
        return nil
    }
    return res
}
// GetPhoneAppVersion gets the phoneAppVersion property value. Numerical version of this instance of the Authenticator app.
// returns a *string when successful
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetPhoneAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("phoneAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftAuthenticatorAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClientAppName() != nil {
        cast := (*m.GetClientAppName()).String()
        err = writer.WriteStringValue("clientAppName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceTag", m.GetDeviceTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneAppVersion", m.GetPhoneAppVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientAppName sets the clientAppName property value. The app that the user has registered to use to approve push notifications. The possible values are: microsoftAuthenticator, outlookMobile, unknownFutureValue.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetClientAppName(value *MicrosoftAuthenticatorAuthenticationMethodClientAppName)() {
    err := m.GetBackingStore().Set("clientAppName", value)
    if err != nil {
        panic(err)
    }
}
// SetDevice sets the device property value. The registered device on which Microsoft Authenticator resides. This property is null if the device isn't registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDevice(value Deviceable)() {
    err := m.GetBackingStore().Set("device", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceTag sets the deviceTag property value. Tags containing app metadata.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDeviceTag(value *string)() {
    err := m.GetBackingStore().Set("deviceTag", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the device on which this app is registered.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneAppVersion sets the phoneAppVersion property value. Numerical version of this instance of the Authenticator app.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetPhoneAppVersion(value *string)() {
    err := m.GetBackingStore().Set("phoneAppVersion", value)
    if err != nil {
        panic(err)
    }
}
type MicrosoftAuthenticatorAuthenticationMethodable interface {
    AuthenticationMethodable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientAppName()(*MicrosoftAuthenticatorAuthenticationMethodClientAppName)
    GetDevice()(Deviceable)
    GetDeviceTag()(*string)
    GetDisplayName()(*string)
    GetPhoneAppVersion()(*string)
    SetClientAppName(value *MicrosoftAuthenticatorAuthenticationMethodClientAppName)()
    SetDevice(value Deviceable)()
    SetDeviceTag(value *string)()
    SetDisplayName(value *string)()
    SetPhoneAppVersion(value *string)()
}
