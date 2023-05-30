package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// SignInPreferences 
type SignInPreferences struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSignInPreferences instantiates a new signInPreferences and sets the default values.
func NewSignInPreferences()(*SignInPreferences) {
    m := &SignInPreferences{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSignInPreferencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSignInPreferencesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignInPreferences(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SignInPreferences) GetAdditionalData()(map[string]any) {
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
func (m *SignInPreferences) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignInPreferences) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isSystemPreferredAuthenticationMethodEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSystemPreferredAuthenticationMethodEnabled(val)
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
    res["userPreferredMethodForSecondaryAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserDefaultAuthenticationMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPreferredMethodForSecondaryAuthentication(val.(*UserDefaultAuthenticationMethodType))
        }
        return nil
    }
    return res
}
// GetIsSystemPreferredAuthenticationMethodEnabled gets the isSystemPreferredAuthenticationMethodEnabled property value. Indicates whether the credential preferences of the system are enabled.
func (m *SignInPreferences) GetIsSystemPreferredAuthenticationMethodEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSystemPreferredAuthenticationMethodEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SignInPreferences) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPreferredMethodForSecondaryAuthentication gets the userPreferredMethodForSecondaryAuthentication property value. The userPreferredMethodForSecondaryAuthentication property
func (m *SignInPreferences) GetUserPreferredMethodForSecondaryAuthentication()(*UserDefaultAuthenticationMethodType) {
    val, err := m.GetBackingStore().Get("userPreferredMethodForSecondaryAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserDefaultAuthenticationMethodType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SignInPreferences) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isSystemPreferredAuthenticationMethodEnabled", m.GetIsSystemPreferredAuthenticationMethodEnabled())
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
    if m.GetUserPreferredMethodForSecondaryAuthentication() != nil {
        cast := (*m.GetUserPreferredMethodForSecondaryAuthentication()).String()
        err := writer.WriteStringValue("userPreferredMethodForSecondaryAuthentication", &cast)
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
func (m *SignInPreferences) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *SignInPreferences) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsSystemPreferredAuthenticationMethodEnabled sets the isSystemPreferredAuthenticationMethodEnabled property value. Indicates whether the credential preferences of the system are enabled.
func (m *SignInPreferences) SetIsSystemPreferredAuthenticationMethodEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSystemPreferredAuthenticationMethodEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SignInPreferences) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPreferredMethodForSecondaryAuthentication sets the userPreferredMethodForSecondaryAuthentication property value. The userPreferredMethodForSecondaryAuthentication property
func (m *SignInPreferences) SetUserPreferredMethodForSecondaryAuthentication(value *UserDefaultAuthenticationMethodType)() {
    err := m.GetBackingStore().Set("userPreferredMethodForSecondaryAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SignInPreferencesable 
type SignInPreferencesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsSystemPreferredAuthenticationMethodEnabled()(*bool)
    GetOdataType()(*string)
    GetUserPreferredMethodForSecondaryAuthentication()(*UserDefaultAuthenticationMethodType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsSystemPreferredAuthenticationMethodEnabled(value *bool)()
    SetOdataType(value *string)()
    SetUserPreferredMethodForSecondaryAuthentication(value *UserDefaultAuthenticationMethodType)()
}
