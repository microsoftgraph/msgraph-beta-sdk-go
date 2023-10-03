package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MicrosoftAuthenticatorFeatureSettings 
type MicrosoftAuthenticatorFeatureSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMicrosoftAuthenticatorFeatureSettings instantiates a new microsoftAuthenticatorFeatureSettings and sets the default values.
func NewMicrosoftAuthenticatorFeatureSettings()(*MicrosoftAuthenticatorFeatureSettings) {
    m := &MicrosoftAuthenticatorFeatureSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftAuthenticatorFeatureSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAuthenticatorFeatureSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorFeatureSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftAuthenticatorFeatureSettings) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *MicrosoftAuthenticatorFeatureSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompanionAppAllowedState gets the companionAppAllowedState property value. Determines whether users are able to approve push notifications on other Microsoft applications such as Outlook Mobile.
func (m *MicrosoftAuthenticatorFeatureSettings) GetCompanionAppAllowedState()(AuthenticationMethodFeatureConfigurationable) {
    val, err := m.GetBackingStore().Get("companionAppAllowedState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationMethodFeatureConfigurationable)
    }
    return nil
}
// GetDisplayAppInformationRequiredState gets the displayAppInformationRequiredState property value. Determines whether the user's Authenticator app shows them the client app they're signing into.
func (m *MicrosoftAuthenticatorFeatureSettings) GetDisplayAppInformationRequiredState()(AuthenticationMethodFeatureConfigurationable) {
    val, err := m.GetBackingStore().Get("displayAppInformationRequiredState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationMethodFeatureConfigurationable)
    }
    return nil
}
// GetDisplayLocationInformationRequiredState gets the displayLocationInformationRequiredState property value. Determines whether the user's Authenticator app shows them the geographic location of where the authentication request originated from.
func (m *MicrosoftAuthenticatorFeatureSettings) GetDisplayLocationInformationRequiredState()(AuthenticationMethodFeatureConfigurationable) {
    val, err := m.GetBackingStore().Get("displayLocationInformationRequiredState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationMethodFeatureConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAuthenticatorFeatureSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["companionAppAllowedState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanionAppAllowedState(val.(AuthenticationMethodFeatureConfigurationable))
        }
        return nil
    }
    res["displayAppInformationRequiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAppInformationRequiredState(val.(AuthenticationMethodFeatureConfigurationable))
        }
        return nil
    }
    res["displayLocationInformationRequiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayLocationInformationRequiredState(val.(AuthenticationMethodFeatureConfigurationable))
        }
        return nil
    }
    res["numberMatchingRequiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberMatchingRequiredState(val.(AuthenticationMethodFeatureConfigurationable))
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
// GetNumberMatchingRequiredState gets the numberMatchingRequiredState property value. Specifies whether the user needs to enter a number in the Authenticator app from the login screen to complete their login. Value is ignored for phone sign-in notifications.
func (m *MicrosoftAuthenticatorFeatureSettings) GetNumberMatchingRequiredState()(AuthenticationMethodFeatureConfigurationable) {
    val, err := m.GetBackingStore().Get("numberMatchingRequiredState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationMethodFeatureConfigurationable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MicrosoftAuthenticatorFeatureSettings) GetOdataType()(*string) {
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
func (m *MicrosoftAuthenticatorFeatureSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("companionAppAllowedState", m.GetCompanionAppAllowedState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("displayAppInformationRequiredState", m.GetDisplayAppInformationRequiredState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("displayLocationInformationRequiredState", m.GetDisplayLocationInformationRequiredState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberMatchingRequiredState", m.GetNumberMatchingRequiredState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftAuthenticatorFeatureSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *MicrosoftAuthenticatorFeatureSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompanionAppAllowedState sets the companionAppAllowedState property value. Determines whether users are able to approve push notifications on other Microsoft applications such as Outlook Mobile.
func (m *MicrosoftAuthenticatorFeatureSettings) SetCompanionAppAllowedState(value AuthenticationMethodFeatureConfigurationable)() {
    err := m.GetBackingStore().Set("companionAppAllowedState", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayAppInformationRequiredState sets the displayAppInformationRequiredState property value. Determines whether the user's Authenticator app shows them the client app they're signing into.
func (m *MicrosoftAuthenticatorFeatureSettings) SetDisplayAppInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)() {
    err := m.GetBackingStore().Set("displayAppInformationRequiredState", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayLocationInformationRequiredState sets the displayLocationInformationRequiredState property value. Determines whether the user's Authenticator app shows them the geographic location of where the authentication request originated from.
func (m *MicrosoftAuthenticatorFeatureSettings) SetDisplayLocationInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)() {
    err := m.GetBackingStore().Set("displayLocationInformationRequiredState", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberMatchingRequiredState sets the numberMatchingRequiredState property value. Specifies whether the user needs to enter a number in the Authenticator app from the login screen to complete their login. Value is ignored for phone sign-in notifications.
func (m *MicrosoftAuthenticatorFeatureSettings) SetNumberMatchingRequiredState(value AuthenticationMethodFeatureConfigurationable)() {
    err := m.GetBackingStore().Set("numberMatchingRequiredState", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MicrosoftAuthenticatorFeatureSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftAuthenticatorFeatureSettingsable 
type MicrosoftAuthenticatorFeatureSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompanionAppAllowedState()(AuthenticationMethodFeatureConfigurationable)
    GetDisplayAppInformationRequiredState()(AuthenticationMethodFeatureConfigurationable)
    GetDisplayLocationInformationRequiredState()(AuthenticationMethodFeatureConfigurationable)
    GetNumberMatchingRequiredState()(AuthenticationMethodFeatureConfigurationable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompanionAppAllowedState(value AuthenticationMethodFeatureConfigurationable)()
    SetDisplayAppInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)()
    SetDisplayLocationInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)()
    SetNumberMatchingRequiredState(value AuthenticationMethodFeatureConfigurationable)()
    SetOdataType(value *string)()
}
