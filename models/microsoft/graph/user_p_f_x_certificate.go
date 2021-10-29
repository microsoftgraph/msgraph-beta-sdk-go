package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserPFXCertificate struct {
    Entity
    // Date/time when this PFX certificate was imported.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Encrypted PFX blob.
    encryptedPfxBlob []byte;
    // Encrypted PFX password.
    encryptedPfxPassword *string;
    // Certificate's validity expiration date/time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Certificate's intended purpose from the point-of-view of deployment. Possible values are: unassigned, smimeEncryption, smimeSigning, vpn, wifi.
    intendedPurpose *UserPfxIntendedPurpose;
    // Name of the key (within the provider) used to encrypt the blob.
    keyName *string;
    // Date/time when this PFX certificate was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Padding scheme used by the provider during encryption/decryption. Possible values are: none, pkcs1, oaepSha1, oaepSha256, oaepSha384, oaepSha512.
    paddingScheme *UserPfxPaddingScheme;
    // Crypto provider used to encrypt this blob.
    providerName *string;
    // Certificate's validity start date/time.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // SHA-1 thumbprint of the PFX certificate.
    thumbprint *string;
    // User Principal Name of the PFX certificate.
    userPrincipalName *string;
}
// Instantiates a new userPFXCertificate and sets the default values.
func NewUserPFXCertificate()(*UserPFXCertificate) {
    m := &UserPFXCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. Date/time when this PFX certificate was imported.
func (m *UserPFXCertificate) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the encryptedPfxBlob property value. Encrypted PFX blob.
func (m *UserPFXCertificate) GetEncryptedPfxBlob()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.encryptedPfxBlob
    }
}
// Gets the encryptedPfxPassword property value. Encrypted PFX password.
func (m *UserPFXCertificate) GetEncryptedPfxPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.encryptedPfxPassword
    }
}
// Gets the expirationDateTime property value. Certificate's validity expiration date/time.
func (m *UserPFXCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the intendedPurpose property value. Certificate's intended purpose from the point-of-view of deployment. Possible values are: unassigned, smimeEncryption, smimeSigning, vpn, wifi.
func (m *UserPFXCertificate) GetIntendedPurpose()(*UserPfxIntendedPurpose) {
    if m == nil {
        return nil
    } else {
        return m.intendedPurpose
    }
}
// Gets the keyName property value. Name of the key (within the provider) used to encrypt the blob.
func (m *UserPFXCertificate) GetKeyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyName
    }
}
// Gets the lastModifiedDateTime property value. Date/time when this PFX certificate was last modified.
func (m *UserPFXCertificate) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the paddingScheme property value. Padding scheme used by the provider during encryption/decryption. Possible values are: none, pkcs1, oaepSha1, oaepSha256, oaepSha384, oaepSha512.
func (m *UserPFXCertificate) GetPaddingScheme()(*UserPfxPaddingScheme) {
    if m == nil {
        return nil
    } else {
        return m.paddingScheme
    }
}
// Gets the providerName property value. Crypto provider used to encrypt this blob.
func (m *UserPFXCertificate) GetProviderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerName
    }
}
// Gets the startDateTime property value. Certificate's validity start date/time.
func (m *UserPFXCertificate) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the thumbprint property value. SHA-1 thumbprint of the PFX certificate.
func (m *UserPFXCertificate) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
// Gets the userPrincipalName property value. User Principal Name of the PFX certificate.
func (m *UserPFXCertificate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *UserPFXCertificate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["encryptedPfxBlob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetEncryptedPfxBlob(val)
        return nil
    }
    res["encryptedPfxPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEncryptedPfxPassword(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["intendedPurpose"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPfxIntendedPurpose)
        if err != nil {
            return err
        }
        cast := val.(UserPfxIntendedPurpose)
        m.SetIntendedPurpose(&cast)
        return nil
    }
    res["keyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKeyName(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["paddingScheme"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPfxPaddingScheme)
        if err != nil {
            return err
        }
        cast := val.(UserPfxPaddingScheme)
        m.SetPaddingScheme(&cast)
        return nil
    }
    res["providerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProviderName(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["thumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbprint(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *UserPFXCertificate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserPFXCertificate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetIntendedPurpose().String()
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
    if m.GetPaddingScheme() != nil {
        cast := m.GetPaddingScheme().String()
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
// Sets the createdDateTime property value. Date/time when this PFX certificate was imported.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *UserPFXCertificate) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the encryptedPfxBlob property value. Encrypted PFX blob.
// Parameters:
//  - value : Value to set for the encryptedPfxBlob property.
func (m *UserPFXCertificate) SetEncryptedPfxBlob(value []byte)() {
    m.encryptedPfxBlob = value
}
// Sets the encryptedPfxPassword property value. Encrypted PFX password.
// Parameters:
//  - value : Value to set for the encryptedPfxPassword property.
func (m *UserPFXCertificate) SetEncryptedPfxPassword(value *string)() {
    m.encryptedPfxPassword = value
}
// Sets the expirationDateTime property value. Certificate's validity expiration date/time.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *UserPFXCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the intendedPurpose property value. Certificate's intended purpose from the point-of-view of deployment. Possible values are: unassigned, smimeEncryption, smimeSigning, vpn, wifi.
// Parameters:
//  - value : Value to set for the intendedPurpose property.
func (m *UserPFXCertificate) SetIntendedPurpose(value *UserPfxIntendedPurpose)() {
    m.intendedPurpose = value
}
// Sets the keyName property value. Name of the key (within the provider) used to encrypt the blob.
// Parameters:
//  - value : Value to set for the keyName property.
func (m *UserPFXCertificate) SetKeyName(value *string)() {
    m.keyName = value
}
// Sets the lastModifiedDateTime property value. Date/time when this PFX certificate was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *UserPFXCertificate) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the paddingScheme property value. Padding scheme used by the provider during encryption/decryption. Possible values are: none, pkcs1, oaepSha1, oaepSha256, oaepSha384, oaepSha512.
// Parameters:
//  - value : Value to set for the paddingScheme property.
func (m *UserPFXCertificate) SetPaddingScheme(value *UserPfxPaddingScheme)() {
    m.paddingScheme = value
}
// Sets the providerName property value. Crypto provider used to encrypt this blob.
// Parameters:
//  - value : Value to set for the providerName property.
func (m *UserPFXCertificate) SetProviderName(value *string)() {
    m.providerName = value
}
// Sets the startDateTime property value. Certificate's validity start date/time.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *UserPFXCertificate) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the thumbprint property value. SHA-1 thumbprint of the PFX certificate.
// Parameters:
//  - value : Value to set for the thumbprint property.
func (m *UserPFXCertificate) SetThumbprint(value *string)() {
    m.thumbprint = value
}
// Sets the userPrincipalName property value. User Principal Name of the PFX certificate.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *UserPFXCertificate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
