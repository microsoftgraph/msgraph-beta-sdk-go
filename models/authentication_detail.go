package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AuthenticationDetail 
type AuthenticationDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAuthenticationDetail instantiates a new authenticationDetail and sets the default values.
func NewAuthenticationDetail()(*AuthenticationDetail) {
    m := &AuthenticationDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationDetail) GetAdditionalData()(map[string]any) {
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
// GetAuthenticationMethod gets the authenticationMethod property value. The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
func (m *AuthenticationDetail) GetAuthenticationMethod()(*string) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationMethodDetail gets the authenticationMethodDetail property value. Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (for example, cloud, AD FS, PTA, PHS).
func (m *AuthenticationDetail) GetAuthenticationMethodDetail()(*string) {
    val, err := m.GetBackingStore().Get("authenticationMethodDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationStepDateTime gets the authenticationStepDateTime property value. Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AuthenticationDetail) GetAuthenticationStepDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("authenticationStepDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAuthenticationStepRequirement gets the authenticationStepRequirement property value. The step of authentication that this satisfied. For example, primary authentication, or multifactor authentication.
func (m *AuthenticationDetail) GetAuthenticationStepRequirement()(*string) {
    val, err := m.GetBackingStore().Get("authenticationStepRequirement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticationStepResultDetail gets the authenticationStepResultDetail property value. Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
func (m *AuthenticationDetail) GetAuthenticationStepResultDetail()(*string) {
    val, err := m.GetBackingStore().Get("authenticationStepResultDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AuthenticationDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val)
        }
        return nil
    }
    res["authenticationMethodDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodDetail(val)
        }
        return nil
    }
    res["authenticationStepDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStepDateTime(val)
        }
        return nil
    }
    res["authenticationStepRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStepRequirement(val)
        }
        return nil
    }
    res["authenticationStepResultDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStepResultDetail(val)
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
    res["succeeded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSucceeded(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSucceeded gets the succeeded property value. Indicates the status of the authentication step. Possible values: succeeded, failed.
func (m *AuthenticationDetail) GetSucceeded()(*bool) {
    val, err := m.GetBackingStore().Get("succeeded")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticationMethod", m.GetAuthenticationMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authenticationMethodDetail", m.GetAuthenticationMethodDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("authenticationStepDateTime", m.GetAuthenticationStepDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authenticationStepRequirement", m.GetAuthenticationStepRequirement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authenticationStepResultDetail", m.GetAuthenticationStepResultDetail())
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
        err := writer.WriteBoolValue("succeeded", m.GetSucceeded())
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
func (m *AuthenticationDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
func (m *AuthenticationDetail) SetAuthenticationMethod(value *string)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethodDetail sets the authenticationMethodDetail property value. Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (for example, cloud, AD FS, PTA, PHS).
func (m *AuthenticationDetail) SetAuthenticationMethodDetail(value *string)() {
    err := m.GetBackingStore().Set("authenticationMethodDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationStepDateTime sets the authenticationStepDateTime property value. Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AuthenticationDetail) SetAuthenticationStepDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("authenticationStepDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationStepRequirement sets the authenticationStepRequirement property value. The step of authentication that this satisfied. For example, primary authentication, or multifactor authentication.
func (m *AuthenticationDetail) SetAuthenticationStepRequirement(value *string)() {
    err := m.GetBackingStore().Set("authenticationStepRequirement", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationStepResultDetail sets the authenticationStepResultDetail property value. Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
func (m *AuthenticationDetail) SetAuthenticationStepResultDetail(value *string)() {
    err := m.GetBackingStore().Set("authenticationStepResultDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AuthenticationDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSucceeded sets the succeeded property value. Indicates the status of the authentication step. Possible values: succeeded, failed.
func (m *AuthenticationDetail) SetSucceeded(value *bool)() {
    err := m.GetBackingStore().Set("succeeded", value)
    if err != nil {
        panic(err)
    }
}
// AuthenticationDetailable 
type AuthenticationDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*string)
    GetAuthenticationMethodDetail()(*string)
    GetAuthenticationStepDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAuthenticationStepRequirement()(*string)
    GetAuthenticationStepResultDetail()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetSucceeded()(*bool)
    SetAuthenticationMethod(value *string)()
    SetAuthenticationMethodDetail(value *string)()
    SetAuthenticationStepDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAuthenticationStepRequirement(value *string)()
    SetAuthenticationStepResultDetail(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetSucceeded(value *bool)()
}
