package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserPFXCertificate entity that encapsulates all information required for a user's PFX certificates.
type UserPFXCertificate struct {
    Entity
}
// NewUserPFXCertificate instantiates a new userPFXCertificate and sets the default values.
func NewUserPFXCertificate()(*UserPFXCertificate) {
    m := &UserPFXCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserPFXCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserPFXCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserPFXCertificate(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date/time when this PFX certificate was imported.
func (m *UserPFXCertificate) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEncryptedPfxBlob gets the encryptedPfxBlob property value. Encrypted PFX blob.
func (m *UserPFXCertificate) GetEncryptedPfxBlob()([]byte) {
    val, err := m.GetBackingStore().Get("encryptedPfxBlob")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetEncryptedPfxPassword gets the encryptedPfxPassword property value. Encrypted PFX password.
func (m *UserPFXCertificate) GetEncryptedPfxPassword()(*string) {
    val, err := m.GetBackingStore().Get("encryptedPfxPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. Certificate's validity expiration date/time.
func (m *UserPFXCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *UserPFXCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["encryptedPfxBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedPfxBlob(val)
        }
        return nil
    }
    res["encryptedPfxPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedPfxPassword(val)
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
    res["intendedPurpose"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPfxIntendedPurpose)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntendedPurpose(val.(*UserPfxIntendedPurpose))
        }
        return nil
    }
    res["keyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["paddingScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPfxPaddingScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaddingScheme(val.(*UserPfxPaddingScheme))
        }
        return nil
    }
    res["providerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["thumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbprint(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetIntendedPurpose gets the intendedPurpose property value. Supported values for the intended purpose of a user PFX certificate.
func (m *UserPFXCertificate) GetIntendedPurpose()(*UserPfxIntendedPurpose) {
    val, err := m.GetBackingStore().Get("intendedPurpose")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserPfxIntendedPurpose)
    }
    return nil
}
// GetKeyName gets the keyName property value. Name of the key (within the provider) used to encrypt the blob.
func (m *UserPFXCertificate) GetKeyName()(*string) {
    val, err := m.GetBackingStore().Get("keyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date/time when this PFX certificate was last modified.
func (m *UserPFXCertificate) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserPFXCertificate) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPaddingScheme gets the paddingScheme property value. Supported values for the padding scheme used by encryption provider.
func (m *UserPFXCertificate) GetPaddingScheme()(*UserPfxPaddingScheme) {
    val, err := m.GetBackingStore().Get("paddingScheme")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserPfxPaddingScheme)
    }
    return nil
}
// GetProviderName gets the providerName property value. Crypto provider used to encrypt this blob.
func (m *UserPFXCertificate) GetProviderName()(*string) {
    val, err := m.GetBackingStore().Get("providerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. Certificate's validity start date/time.
func (m *UserPFXCertificate) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetThumbprint gets the thumbprint property value. SHA-1 thumbprint of the PFX certificate.
func (m *UserPFXCertificate) GetThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("thumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name of the PFX certificate.
func (m *UserPFXCertificate) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserPFXCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("encryptedPfxBlob", m.GetEncryptedPfxBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptedPfxPassword", m.GetEncryptedPfxPassword())
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
    if m.GetIntendedPurpose() != nil {
        cast := (*m.GetIntendedPurpose()).String()
        err = writer.WriteStringValue("intendedPurpose", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyName", m.GetKeyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
    if m.GetPaddingScheme() != nil {
        cast := (*m.GetPaddingScheme()).String()
        err = writer.WriteStringValue("paddingScheme", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerName", m.GetProviderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbprint", m.GetThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Date/time when this PFX certificate was imported.
func (m *UserPFXCertificate) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptedPfxBlob sets the encryptedPfxBlob property value. Encrypted PFX blob.
func (m *UserPFXCertificate) SetEncryptedPfxBlob(value []byte)() {
    err := m.GetBackingStore().Set("encryptedPfxBlob", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptedPfxPassword sets the encryptedPfxPassword property value. Encrypted PFX password.
func (m *UserPFXCertificate) SetEncryptedPfxPassword(value *string)() {
    err := m.GetBackingStore().Set("encryptedPfxPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. Certificate's validity expiration date/time.
func (m *UserPFXCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIntendedPurpose sets the intendedPurpose property value. Supported values for the intended purpose of a user PFX certificate.
func (m *UserPFXCertificate) SetIntendedPurpose(value *UserPfxIntendedPurpose)() {
    err := m.GetBackingStore().Set("intendedPurpose", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyName sets the keyName property value. Name of the key (within the provider) used to encrypt the blob.
func (m *UserPFXCertificate) SetKeyName(value *string)() {
    err := m.GetBackingStore().Set("keyName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date/time when this PFX certificate was last modified.
func (m *UserPFXCertificate) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserPFXCertificate) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPaddingScheme sets the paddingScheme property value. Supported values for the padding scheme used by encryption provider.
func (m *UserPFXCertificate) SetPaddingScheme(value *UserPfxPaddingScheme)() {
    err := m.GetBackingStore().Set("paddingScheme", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderName sets the providerName property value. Crypto provider used to encrypt this blob.
func (m *UserPFXCertificate) SetProviderName(value *string)() {
    err := m.GetBackingStore().Set("providerName", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. Certificate's validity start date/time.
func (m *UserPFXCertificate) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbprint sets the thumbprint property value. SHA-1 thumbprint of the PFX certificate.
func (m *UserPFXCertificate) SetThumbprint(value *string)() {
    err := m.GetBackingStore().Set("thumbprint", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name of the PFX certificate.
func (m *UserPFXCertificate) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// UserPFXCertificateable 
type UserPFXCertificateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEncryptedPfxBlob()([]byte)
    GetEncryptedPfxPassword()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIntendedPurpose()(*UserPfxIntendedPurpose)
    GetKeyName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetPaddingScheme()(*UserPfxPaddingScheme)
    GetProviderName()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetThumbprint()(*string)
    GetUserPrincipalName()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEncryptedPfxBlob(value []byte)()
    SetEncryptedPfxPassword(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIntendedPurpose(value *UserPfxIntendedPurpose)()
    SetKeyName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetPaddingScheme(value *UserPfxPaddingScheme)()
    SetProviderName(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetThumbprint(value *string)()
    SetUserPrincipalName(value *string)()
}
