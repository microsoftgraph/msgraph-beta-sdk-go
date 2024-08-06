package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type WebauthnAuthenticatorSelectionCriteria struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWebauthnAuthenticatorSelectionCriteria instantiates a new WebauthnAuthenticatorSelectionCriteria and sets the default values.
func NewWebauthnAuthenticatorSelectionCriteria()(*WebauthnAuthenticatorSelectionCriteria) {
    m := &WebauthnAuthenticatorSelectionCriteria{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebauthnAuthenticatorSelectionCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebauthnAuthenticatorSelectionCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebauthnAuthenticatorSelectionCriteria(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WebauthnAuthenticatorSelectionCriteria) GetAdditionalData()(map[string]any) {
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
// GetAuthenticatorAttachment gets the authenticatorAttachment property value. Microsoft Entra ID-preferred attachment modality. For more information, see Authenticator Attachment Modality
// returns a *string when successful
func (m *WebauthnAuthenticatorSelectionCriteria) GetAuthenticatorAttachment()(*string) {
    val, err := m.GetBackingStore().Get("authenticatorAttachment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *WebauthnAuthenticatorSelectionCriteria) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebauthnAuthenticatorSelectionCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticatorAttachment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticatorAttachment(val)
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
    res["requireResidentKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireResidentKey(val)
        }
        return nil
    }
    res["userVerification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserVerification(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WebauthnAuthenticatorSelectionCriteria) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequireResidentKey gets the requireResidentKey property value. Microsoft Entra ID-preferred client-side credential discoverability. Currently always true. The WebAuthn authenticator must store the credential identifier on the authenticator.
// returns a *bool when successful
func (m *WebauthnAuthenticatorSelectionCriteria) GetRequireResidentKey()(*bool) {
    val, err := m.GetBackingStore().Get("requireResidentKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserVerification gets the userVerification property value. Microsoft Entra ID requirement to verify the user is present during credential provisioning. Currently always required.
// returns a *string when successful
func (m *WebauthnAuthenticatorSelectionCriteria) GetUserVerification()(*string) {
    val, err := m.GetBackingStore().Get("userVerification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WebauthnAuthenticatorSelectionCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticatorAttachment", m.GetAuthenticatorAttachment())
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
        err := writer.WriteBoolValue("requireResidentKey", m.GetRequireResidentKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userVerification", m.GetUserVerification())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebauthnAuthenticatorSelectionCriteria) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticatorAttachment sets the authenticatorAttachment property value. Microsoft Entra ID-preferred attachment modality. For more information, see Authenticator Attachment Modality
func (m *WebauthnAuthenticatorSelectionCriteria) SetAuthenticatorAttachment(value *string)() {
    err := m.GetBackingStore().Set("authenticatorAttachment", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WebauthnAuthenticatorSelectionCriteria) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebauthnAuthenticatorSelectionCriteria) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireResidentKey sets the requireResidentKey property value. Microsoft Entra ID-preferred client-side credential discoverability. Currently always true. The WebAuthn authenticator must store the credential identifier on the authenticator.
func (m *WebauthnAuthenticatorSelectionCriteria) SetRequireResidentKey(value *bool)() {
    err := m.GetBackingStore().Set("requireResidentKey", value)
    if err != nil {
        panic(err)
    }
}
// SetUserVerification sets the userVerification property value. Microsoft Entra ID requirement to verify the user is present during credential provisioning. Currently always required.
func (m *WebauthnAuthenticatorSelectionCriteria) SetUserVerification(value *string)() {
    err := m.GetBackingStore().Set("userVerification", value)
    if err != nil {
        panic(err)
    }
}
type WebauthnAuthenticatorSelectionCriteriaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticatorAttachment()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetRequireResidentKey()(*bool)
    GetUserVerification()(*string)
    SetAuthenticatorAttachment(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetRequireResidentKey(value *bool)()
    SetUserVerification(value *string)()
}
