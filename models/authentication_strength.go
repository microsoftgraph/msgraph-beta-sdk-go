package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AuthenticationStrength 
type AuthenticationStrength struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAuthenticationStrength instantiates a new authenticationStrength and sets the default values.
func NewAuthenticationStrength()(*AuthenticationStrength) {
    m := &AuthenticationStrength{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationStrengthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationStrengthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationStrength(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationStrength) GetAdditionalData()(map[string]any) {
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
// GetAuthenticationStrengthId gets the authenticationStrengthId property value. Identifier of the authentication strength.
func (m *AuthenticationStrength) GetAuthenticationStrengthId()(*string) {
    val, err := m.GetBackingStore().Get("authenticationStrengthId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationStrengthResult gets the authenticationStrengthResult property value. The result of the authentication strength. The possible values are: notSet, skippedForProofUp, satisfied, singleChallengeRequired, multipleChallengesRequired, singleRegistrationRequired, multipleRegistrationsRequired, cannotSatisfyDueToCombinationConfiguration, cannotSatisfy, unknownFutureValue.
func (m *AuthenticationStrength) GetAuthenticationStrengthResult()(*AuthenticationStrength_authenticationStrengthResult) {
    val, err := m.GetBackingStore().Get("authenticationStrengthResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationStrength_authenticationStrengthResult)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AuthenticationStrength) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The name of the authentication strength.
func (m *AuthenticationStrength) GetDisplayName()(*string) {
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
func (m *AuthenticationStrength) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationStrengthId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStrengthId(val)
        }
        return nil
    }
    res["authenticationStrengthResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationStrength_authenticationStrengthResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStrengthResult(val.(*AuthenticationStrength_authenticationStrengthResult))
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationStrength) GetOdataType()(*string) {
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
func (m *AuthenticationStrength) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticationStrengthId", m.GetAuthenticationStrengthId())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationStrengthResult() != nil {
        cast := (*m.GetAuthenticationStrengthResult()).String()
        err := writer.WriteStringValue("authenticationStrengthResult", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationStrength) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationStrengthId sets the authenticationStrengthId property value. Identifier of the authentication strength.
func (m *AuthenticationStrength) SetAuthenticationStrengthId(value *string)() {
    err := m.GetBackingStore().Set("authenticationStrengthId", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationStrengthResult sets the authenticationStrengthResult property value. The result of the authentication strength. The possible values are: notSet, skippedForProofUp, satisfied, singleChallengeRequired, multipleChallengesRequired, singleRegistrationRequired, multipleRegistrationsRequired, cannotSatisfyDueToCombinationConfiguration, cannotSatisfy, unknownFutureValue.
func (m *AuthenticationStrength) SetAuthenticationStrengthResult(value *AuthenticationStrength_authenticationStrengthResult)() {
    err := m.GetBackingStore().Set("authenticationStrengthResult", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AuthenticationStrength) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The name of the authentication strength.
func (m *AuthenticationStrength) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationStrength) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AuthenticationStrengthable 
type AuthenticationStrengthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationStrengthId()(*string)
    GetAuthenticationStrengthResult()(*AuthenticationStrength_authenticationStrengthResult)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    SetAuthenticationStrengthId(value *string)()
    SetAuthenticationStrengthResult(value *AuthenticationStrength_authenticationStrengthResult)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
}
