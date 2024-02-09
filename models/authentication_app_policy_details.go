package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AuthenticationAppPolicyDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAuthenticationAppPolicyDetails instantiates a new AuthenticationAppPolicyDetails and sets the default values.
func NewAuthenticationAppPolicyDetails()(*AuthenticationAppPolicyDetails) {
    m := &AuthenticationAppPolicyDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationAppPolicyDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationAppPolicyDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationAppPolicyDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationAppPolicyDetails) GetAdditionalData()(map[string]any) {
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
// GetAdminConfiguration gets the adminConfiguration property value. The admin configuration of the policy on the user's authentication app. For a policy that does not impact the success/failure of the authentication, the evaluation defaults to notApplicable. The possible values are: notApplicable, enabled, disabled, unknownFutureValue.
// returns a *AuthenticationAppAdminConfiguration when successful
func (m *AuthenticationAppPolicyDetails) GetAdminConfiguration()(*AuthenticationAppAdminConfiguration) {
    val, err := m.GetBackingStore().Get("adminConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationAppAdminConfiguration)
    }
    return nil
}
// GetAuthenticationEvaluation gets the authenticationEvaluation property value. Evaluates the success/failure of the authentication based on the admin configuration of the policy on the user's client authentication app. The possible values are: success, failure, unknownFutureValue.
// returns a *AuthenticationAppEvaluation when successful
func (m *AuthenticationAppPolicyDetails) GetAuthenticationEvaluation()(*AuthenticationAppEvaluation) {
    val, err := m.GetBackingStore().Get("authenticationEvaluation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationAppEvaluation)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AuthenticationAppPolicyDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationAppPolicyDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adminConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationAppAdminConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConfiguration(val.(*AuthenticationAppAdminConfiguration))
        }
        return nil
    }
    res["authenticationEvaluation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationAppEvaluation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationEvaluation(val.(*AuthenticationAppEvaluation))
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
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationAppPolicyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AuthenticationAppPolicyStatus))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AuthenticationAppPolicyDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyName gets the policyName property value. The name of the policy enforced on the user's authentication app.
// returns a *string when successful
func (m *AuthenticationAppPolicyDetails) GetPolicyName()(*string) {
    val, err := m.GetBackingStore().Get("policyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. Refers to whether the policy executed as expected on the user's client authentication app. The possible values are: unknown, appLockOutOfDate, appLockEnabled, appLockDisabled, appContextOutOfDate, appContextShown, appContextNotShown, locationContextOutOfDate, locationContextShown, locationContextNotShown, numberMatchOutOfDate, numberMatchCorrectNumberEntered, numberMatchIncorrectNumberEntered, numberMatchDeny, tamperResistantHardwareOutOfDate, tamperResistantHardwareUsed, tamperResistantHardwareNotUsed, unknownFutureValue.
// returns a *AuthenticationAppPolicyStatus when successful
func (m *AuthenticationAppPolicyDetails) GetStatus()(*AuthenticationAppPolicyStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationAppPolicyStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationAppPolicyDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdminConfiguration() != nil {
        cast := (*m.GetAdminConfiguration()).String()
        err := writer.WriteStringValue("adminConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationEvaluation() != nil {
        cast := (*m.GetAuthenticationEvaluation()).String()
        err := writer.WriteStringValue("authenticationEvaluation", &cast)
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
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *AuthenticationAppPolicyDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminConfiguration sets the adminConfiguration property value. The admin configuration of the policy on the user's authentication app. For a policy that does not impact the success/failure of the authentication, the evaluation defaults to notApplicable. The possible values are: notApplicable, enabled, disabled, unknownFutureValue.
func (m *AuthenticationAppPolicyDetails) SetAdminConfiguration(value *AuthenticationAppAdminConfiguration)() {
    err := m.GetBackingStore().Set("adminConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationEvaluation sets the authenticationEvaluation property value. Evaluates the success/failure of the authentication based on the admin configuration of the policy on the user's client authentication app. The possible values are: success, failure, unknownFutureValue.
func (m *AuthenticationAppPolicyDetails) SetAuthenticationEvaluation(value *AuthenticationAppEvaluation)() {
    err := m.GetBackingStore().Set("authenticationEvaluation", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AuthenticationAppPolicyDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationAppPolicyDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyName sets the policyName property value. The name of the policy enforced on the user's authentication app.
func (m *AuthenticationAppPolicyDetails) SetPolicyName(value *string)() {
    err := m.GetBackingStore().Set("policyName", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Refers to whether the policy executed as expected on the user's client authentication app. The possible values are: unknown, appLockOutOfDate, appLockEnabled, appLockDisabled, appContextOutOfDate, appContextShown, appContextNotShown, locationContextOutOfDate, locationContextShown, locationContextNotShown, numberMatchOutOfDate, numberMatchCorrectNumberEntered, numberMatchIncorrectNumberEntered, numberMatchDeny, tamperResistantHardwareOutOfDate, tamperResistantHardwareUsed, tamperResistantHardwareNotUsed, unknownFutureValue.
func (m *AuthenticationAppPolicyDetails) SetStatus(value *AuthenticationAppPolicyStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type AuthenticationAppPolicyDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminConfiguration()(*AuthenticationAppAdminConfiguration)
    GetAuthenticationEvaluation()(*AuthenticationAppEvaluation)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetPolicyName()(*string)
    GetStatus()(*AuthenticationAppPolicyStatus)
    SetAdminConfiguration(value *AuthenticationAppAdminConfiguration)()
    SetAuthenticationEvaluation(value *AuthenticationAppEvaluation)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetPolicyName(value *string)()
    SetStatus(value *AuthenticationAppPolicyStatus)()
}
