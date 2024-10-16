package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CertificateAuthorityDetail struct {
    DirectoryObject
}
// NewCertificateAuthorityDetail instantiates a new CertificateAuthorityDetail and sets the default values.
func NewCertificateAuthorityDetail()(*CertificateAuthorityDetail) {
    m := &CertificateAuthorityDetail{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.certificateAuthorityDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCertificateAuthorityDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCertificateAuthorityDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateAuthorityDetail(), nil
}
// GetCertificate gets the certificate property value. The certificate property
// returns a []byte when successful
func (m *CertificateAuthorityDetail) GetCertificate()([]byte) {
    val, err := m.GetBackingStore().Get("certificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetCertificateAuthorityType gets the certificateAuthorityType property value. The certificateAuthorityType property
// returns a *CertificateAuthorityType when successful
func (m *CertificateAuthorityDetail) GetCertificateAuthorityType()(*CertificateAuthorityType) {
    val, err := m.GetBackingStore().Get("certificateAuthorityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateAuthorityType)
    }
    return nil
}
// GetCertificateRevocationListUrl gets the certificateRevocationListUrl property value. The certificateRevocationListUrl property
// returns a *string when successful
func (m *CertificateAuthorityDetail) GetCertificateRevocationListUrl()(*string) {
    val, err := m.GetBackingStore().Get("certificateRevocationListUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
// returns a *Time when successful
func (m *CertificateAuthorityDetail) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeltacertificateRevocationListUrl gets the deltacertificateRevocationListUrl property value. The deltacertificateRevocationListUrl property
// returns a *string when successful
func (m *CertificateAuthorityDetail) GetDeltacertificateRevocationListUrl()(*string) {
    val, err := m.GetBackingStore().Get("deltacertificateRevocationListUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *CertificateAuthorityDetail) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
// returns a *Time when successful
func (m *CertificateAuthorityDetail) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CertificateAuthorityDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["certificateAuthorityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateAuthorityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateAuthorityType(val.(*CertificateAuthorityType))
        }
        return nil
    }
    res["certificateRevocationListUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevocationListUrl(val)
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
    res["deltacertificateRevocationListUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeltacertificateRevocationListUrl(val)
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
    res["isIssuerHintEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsIssuerHintEnabled(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerSubjectKeyIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerSubjectKeyIdentifier(val)
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
    return res
}
// GetIsIssuerHintEnabled gets the isIssuerHintEnabled property value. The isIssuerHintEnabled property
// returns a *bool when successful
func (m *CertificateAuthorityDetail) GetIsIssuerHintEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isIssuerHintEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIssuer gets the issuer property value. The issuer property
// returns a *string when successful
func (m *CertificateAuthorityDetail) GetIssuer()(*string) {
    val, err := m.GetBackingStore().Get("issuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIssuerSubjectKeyIdentifier gets the issuerSubjectKeyIdentifier property value. The issuerSubjectKeyIdentifier property
// returns a *string when successful
func (m *CertificateAuthorityDetail) GetIssuerSubjectKeyIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("issuerSubjectKeyIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThumbprint gets the thumbprint property value. The thumbprint property
// returns a *string when successful
func (m *CertificateAuthorityDetail) GetThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("thumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CertificateAuthorityDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateAuthorityType() != nil {
        cast := (*m.GetCertificateAuthorityType()).String()
        err = writer.WriteStringValue("certificateAuthorityType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateRevocationListUrl", m.GetCertificateRevocationListUrl())
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
    {
        err = writer.WriteStringValue("deltacertificateRevocationListUrl", m.GetDeltacertificateRevocationListUrl())
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
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isIssuerHintEnabled", m.GetIsIssuerHintEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerSubjectKeyIdentifier", m.GetIssuerSubjectKeyIdentifier())
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
    return nil
}
// SetCertificate sets the certificate property value. The certificate property
func (m *CertificateAuthorityDetail) SetCertificate(value []byte)() {
    err := m.GetBackingStore().Set("certificate", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateAuthorityType sets the certificateAuthorityType property value. The certificateAuthorityType property
func (m *CertificateAuthorityDetail) SetCertificateAuthorityType(value *CertificateAuthorityType)() {
    err := m.GetBackingStore().Set("certificateAuthorityType", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateRevocationListUrl sets the certificateRevocationListUrl property value. The certificateRevocationListUrl property
func (m *CertificateAuthorityDetail) SetCertificateRevocationListUrl(value *string)() {
    err := m.GetBackingStore().Set("certificateRevocationListUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *CertificateAuthorityDetail) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeltacertificateRevocationListUrl sets the deltacertificateRevocationListUrl property value. The deltacertificateRevocationListUrl property
func (m *CertificateAuthorityDetail) SetDeltacertificateRevocationListUrl(value *string)() {
    err := m.GetBackingStore().Set("deltacertificateRevocationListUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CertificateAuthorityDetail) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *CertificateAuthorityDetail) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsIssuerHintEnabled sets the isIssuerHintEnabled property value. The isIssuerHintEnabled property
func (m *CertificateAuthorityDetail) SetIsIssuerHintEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isIssuerHintEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuer sets the issuer property value. The issuer property
func (m *CertificateAuthorityDetail) SetIssuer(value *string)() {
    err := m.GetBackingStore().Set("issuer", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuerSubjectKeyIdentifier sets the issuerSubjectKeyIdentifier property value. The issuerSubjectKeyIdentifier property
func (m *CertificateAuthorityDetail) SetIssuerSubjectKeyIdentifier(value *string)() {
    err := m.GetBackingStore().Set("issuerSubjectKeyIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbprint sets the thumbprint property value. The thumbprint property
func (m *CertificateAuthorityDetail) SetThumbprint(value *string)() {
    err := m.GetBackingStore().Set("thumbprint", value)
    if err != nil {
        panic(err)
    }
}
type CertificateAuthorityDetailable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificate()([]byte)
    GetCertificateAuthorityType()(*CertificateAuthorityType)
    GetCertificateRevocationListUrl()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeltacertificateRevocationListUrl()(*string)
    GetDisplayName()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsIssuerHintEnabled()(*bool)
    GetIssuer()(*string)
    GetIssuerSubjectKeyIdentifier()(*string)
    GetThumbprint()(*string)
    SetCertificate(value []byte)()
    SetCertificateAuthorityType(value *CertificateAuthorityType)()
    SetCertificateRevocationListUrl(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeltacertificateRevocationListUrl(value *string)()
    SetDisplayName(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsIssuerHintEnabled(value *bool)()
    SetIssuer(value *string)()
    SetIssuerSubjectKeyIdentifier(value *string)()
    SetThumbprint(value *string)()
}
