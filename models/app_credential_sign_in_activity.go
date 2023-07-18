package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCredentialSignInActivity 
type AppCredentialSignInActivity struct {
    Entity
}
// NewAppCredentialSignInActivity instantiates a new appCredentialSignInActivity and sets the default values.
func NewAppCredentialSignInActivity()(*AppCredentialSignInActivity) {
    m := &AppCredentialSignInActivity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppCredentialSignInActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCredentialSignInActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCredentialSignInActivity(), nil
}
// GetAppId gets the appId property value. The globally unique appId (also called client ID on the Azure portal) of the credential application.
func (m *AppCredentialSignInActivity) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppObjectId gets the appObjectId property value. The ID of the credential application instance.
func (m *AppCredentialSignInActivity) GetAppObjectId()(*string) {
    val, err := m.GetBackingStore().Get("appObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the credential was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AppCredentialSignInActivity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCredentialOrigin gets the credentialOrigin property value. The credentialOrigin property
func (m *AppCredentialSignInActivity) GetCredentialOrigin()(*ApplicationKeyOrigin) {
    val, err := m.GetBackingStore().Get("credentialOrigin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApplicationKeyOrigin)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The date and time when the credential is set to expire. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AppCredentialSignInActivity) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCredentialSignInActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["appObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppObjectId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["credentialOrigin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationKeyOrigin)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialOrigin(val.(*ApplicationKeyOrigin))
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["keyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationKeyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyType(val.(*ApplicationKeyType))
        }
        return nil
    }
    res["keyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationKeyUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyUsage(val.(*ApplicationKeyUsage))
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
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["servicePrincipalObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalObjectId(val)
        }
        return nil
    }
    res["signInActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInActivity(val.(SignInActivityable))
        }
        return nil
    }
    return res
}
// GetKeyId gets the keyId property value. The key ID of the credential.
func (m *AppCredentialSignInActivity) GetKeyId()(*string) {
    val, err := m.GetBackingStore().Get("keyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKeyType gets the keyType property value. Specifies the key type. The possible values are: clientSecret, certificate, unknownFutureValue.
func (m *AppCredentialSignInActivity) GetKeyType()(*ApplicationKeyType) {
    val, err := m.GetBackingStore().Get("keyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApplicationKeyType)
    }
    return nil
}
// GetKeyUsage gets the keyUsage property value. Specifies what the key was used for. The possible values are: sign, verify, unknownFutureValue.
func (m *AppCredentialSignInActivity) GetKeyUsage()(*ApplicationKeyUsage) {
    val, err := m.GetBackingStore().Get("keyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApplicationKeyUsage)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppCredentialSignInActivity) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceId gets the resourceId property value. The ID of the accessed resource.
func (m *AppCredentialSignInActivity) GetResourceId()(*string) {
    val, err := m.GetBackingStore().Get("resourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePrincipalObjectId gets the servicePrincipalObjectId property value. The ID of the service principal.
func (m *AppCredentialSignInActivity) GetServicePrincipalObjectId()(*string) {
    val, err := m.GetBackingStore().Get("servicePrincipalObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSignInActivity gets the signInActivity property value. The signInActivity property
func (m *AppCredentialSignInActivity) GetSignInActivity()(SignInActivityable) {
    val, err := m.GetBackingStore().Get("signInActivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SignInActivityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppCredentialSignInActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteStringValue("appObjectId", m.GetAppObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCredentialOrigin() != nil {
        cast := (*m.GetCredentialOrigin()).String()
        err = writer.WriteStringValue("credentialOrigin", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    if m.GetKeyType() != nil {
        cast := (*m.GetKeyType()).String()
        err = writer.WriteStringValue("keyType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetKeyUsage() != nil {
        cast := (*m.GetKeyUsage()).String()
        err = writer.WriteStringValue("keyUsage", &cast)
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
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalObjectId", m.GetServicePrincipalObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signInActivity", m.GetSignInActivity())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The globally unique appId (also called client ID on the Azure portal) of the credential application.
func (m *AppCredentialSignInActivity) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppObjectId sets the appObjectId property value. The ID of the credential application instance.
func (m *AppCredentialSignInActivity) SetAppObjectId(value *string)() {
    err := m.GetBackingStore().Set("appObjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the credential was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AppCredentialSignInActivity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCredentialOrigin sets the credentialOrigin property value. The credentialOrigin property
func (m *AppCredentialSignInActivity) SetCredentialOrigin(value *ApplicationKeyOrigin)() {
    err := m.GetBackingStore().Set("credentialOrigin", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The date and time when the credential is set to expire. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AppCredentialSignInActivity) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyId sets the keyId property value. The key ID of the credential.
func (m *AppCredentialSignInActivity) SetKeyId(value *string)() {
    err := m.GetBackingStore().Set("keyId", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyType sets the keyType property value. Specifies the key type. The possible values are: clientSecret, certificate, unknownFutureValue.
func (m *AppCredentialSignInActivity) SetKeyType(value *ApplicationKeyType)() {
    err := m.GetBackingStore().Set("keyType", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyUsage sets the keyUsage property value. Specifies what the key was used for. The possible values are: sign, verify, unknownFutureValue.
func (m *AppCredentialSignInActivity) SetKeyUsage(value *ApplicationKeyUsage)() {
    err := m.GetBackingStore().Set("keyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppCredentialSignInActivity) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. The ID of the accessed resource.
func (m *AppCredentialSignInActivity) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalObjectId sets the servicePrincipalObjectId property value. The ID of the service principal.
func (m *AppCredentialSignInActivity) SetServicePrincipalObjectId(value *string)() {
    err := m.GetBackingStore().Set("servicePrincipalObjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInActivity sets the signInActivity property value. The signInActivity property
func (m *AppCredentialSignInActivity) SetSignInActivity(value SignInActivityable)() {
    err := m.GetBackingStore().Set("signInActivity", value)
    if err != nil {
        panic(err)
    }
}
// AppCredentialSignInActivityable 
type AppCredentialSignInActivityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetAppObjectId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCredentialOrigin()(*ApplicationKeyOrigin)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetKeyId()(*string)
    GetKeyType()(*ApplicationKeyType)
    GetKeyUsage()(*ApplicationKeyUsage)
    GetOdataType()(*string)
    GetResourceId()(*string)
    GetServicePrincipalObjectId()(*string)
    GetSignInActivity()(SignInActivityable)
    SetAppId(value *string)()
    SetAppObjectId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCredentialOrigin(value *ApplicationKeyOrigin)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetKeyId(value *string)()
    SetKeyType(value *ApplicationKeyType)()
    SetKeyUsage(value *ApplicationKeyUsage)()
    SetOdataType(value *string)()
    SetResourceId(value *string)()
    SetServicePrincipalObjectId(value *string)()
    SetSignInActivity(value SignInActivityable)()
}
