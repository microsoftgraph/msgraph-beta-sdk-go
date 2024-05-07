package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalAuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
}
// NewExternalAuthenticationMethodConfiguration instantiates a new ExternalAuthenticationMethodConfiguration and sets the default values.
func NewExternalAuthenticationMethodConfiguration()(*ExternalAuthenticationMethodConfiguration) {
    m := &ExternalAuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.externalAuthenticationMethodConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExternalAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalAuthenticationMethodConfiguration(), nil
}
// GetAppId gets the appId property value. appId for the app registration in Microsoft Entra ID representing the integration with the external provider.
// returns a *string when successful
func (m *ExternalAuthenticationMethodConfiguration) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display name for the external authentication method. This name is shown to users during sign-in.
// returns a *string when successful
func (m *ExternalAuthenticationMethodConfiguration) GetDisplayName()(*string) {
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
func (m *ExternalAuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
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
    res["includeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodTargetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationMethodTargetable)
                }
            }
            m.SetIncludeTargets(res)
        }
        return nil
    }
    res["openIdConnectSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOpenIdConnectSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenIdConnectSetting(val.(OpenIdConnectSettingable))
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of groups that are enabled to use an authentication method as part of an authentication method policy in Microsoft Entra ID.
// returns a []AuthenticationMethodTargetable when successful
func (m *ExternalAuthenticationMethodConfiguration) GetIncludeTargets()([]AuthenticationMethodTargetable) {
    val, err := m.GetBackingStore().Get("includeTargets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationMethodTargetable)
    }
    return nil
}
// GetOpenIdConnectSetting gets the openIdConnectSetting property value. The openIdConnectSetting property
// returns a OpenIdConnectSettingable when successful
func (m *ExternalAuthenticationMethodConfiguration) GetOpenIdConnectSetting()(OpenIdConnectSettingable) {
    val, err := m.GetBackingStore().Get("openIdConnectSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OpenIdConnectSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExternalAuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
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
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("openIdConnectSetting", m.GetOpenIdConnectSetting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. appId for the app registration in Microsoft Entra ID representing the integration with the external provider.
func (m *ExternalAuthenticationMethodConfiguration) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name for the external authentication method. This name is shown to users during sign-in.
func (m *ExternalAuthenticationMethodConfiguration) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeTargets sets the includeTargets property value. A collection of groups that are enabled to use an authentication method as part of an authentication method policy in Microsoft Entra ID.
func (m *ExternalAuthenticationMethodConfiguration) SetIncludeTargets(value []AuthenticationMethodTargetable)() {
    err := m.GetBackingStore().Set("includeTargets", value)
    if err != nil {
        panic(err)
    }
}
// SetOpenIdConnectSetting sets the openIdConnectSetting property value. The openIdConnectSetting property
func (m *ExternalAuthenticationMethodConfiguration) SetOpenIdConnectSetting(value OpenIdConnectSettingable)() {
    err := m.GetBackingStore().Set("openIdConnectSetting", value)
    if err != nil {
        panic(err)
    }
}
type ExternalAuthenticationMethodConfigurationable interface {
    AuthenticationMethodConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetDisplayName()(*string)
    GetIncludeTargets()([]AuthenticationMethodTargetable)
    GetOpenIdConnectSetting()(OpenIdConnectSettingable)
    SetAppId(value *string)()
    SetDisplayName(value *string)()
    SetIncludeTargets(value []AuthenticationMethodTargetable)()
    SetOpenIdConnectSetting(value OpenIdConnectSettingable)()
}
