package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81SCEPCertificateProfile 
type WindowsPhone81SCEPCertificateProfile struct {
    WindowsPhone81CertificateProfileBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // SCEP Hash Algorithm. Possible values are: sha1, sha2.
    hashAlgorithm *HashAlgorithms
    // SCEP Key Size. Possible values are: size1024, size2048, size4096.
    keySize *KeySize
    // SCEP Key Usage. Possible values are: keyEncipherment, digitalSignature.
    keyUsage *KeyUsages
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Trusted Root Certificate.
    rootCertificate WindowsPhone81TrustedRootCertificateable
    // SCEP Server Url(s).
    scepServerUrls []string
    // Custom String that defines the AAD Attribute.
    subjectAlternativeNameFormatString *string
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewWindowsPhone81SCEPCertificateProfile instantiates a new WindowsPhone81SCEPCertificateProfile and sets the default values.
func NewWindowsPhone81SCEPCertificateProfile()(*WindowsPhone81SCEPCertificateProfile) {
    m := &WindowsPhone81SCEPCertificateProfile{
        WindowsPhone81CertificateProfileBase: *NewWindowsPhone81CertificateProfileBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsPhone81SCEPCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81SCEPCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81SCEPCertificateProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPhone81SCEPCertificateProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81SCEPCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsPhone81CertificateProfileBase.GetFieldDeserializers()
    res["hashAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHashAlgorithms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashAlgorithm(val.(*HashAlgorithms))
        }
        return nil
    }
    res["keySize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeySize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeySize(val.(*KeySize))
        }
        return nil
    }
    res["keyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyUsages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyUsage(val.(*KeyUsages))
        }
        return nil
    }
    res["managedDeviceCertificateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceCertificateStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceCertificateStateable)
            }
            m.SetManagedDeviceCertificateStates(res)
        }
        return nil
    }
    res["rootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsPhone81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificate(val.(WindowsPhone81TrustedRootCertificateable))
        }
        return nil
    }
    res["scepServerUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetScepServerUrls(res)
        }
        return nil
    }
    res["subjectAlternativeNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectAlternativeNameFormatString(val)
        }
        return nil
    }
    res["subjectNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormatString(val)
        }
        return nil
    }
    return res
}
// GetHashAlgorithm gets the hashAlgorithm property value. SCEP Hash Algorithm. Possible values are: sha1, sha2.
func (m *WindowsPhone81SCEPCertificateProfile) GetHashAlgorithm()(*HashAlgorithms) {
    if m == nil {
        return nil
    } else {
        return m.hashAlgorithm
    }
}
// GetKeySize gets the keySize property value. SCEP Key Size. Possible values are: size1024, size2048, size4096.
func (m *WindowsPhone81SCEPCertificateProfile) GetKeySize()(*KeySize) {
    if m == nil {
        return nil
    } else {
        return m.keySize
    }
}
// GetKeyUsage gets the keyUsage property value. SCEP Key Usage. Possible values are: keyEncipherment, digitalSignature.
func (m *WindowsPhone81SCEPCertificateProfile) GetKeyUsage()(*KeyUsages) {
    if m == nil {
        return nil
    } else {
        return m.keyUsage
    }
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *WindowsPhone81SCEPCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCertificateStates
    }
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate.
func (m *WindowsPhone81SCEPCertificateProfile) GetRootCertificate()(WindowsPhone81TrustedRootCertificateable) {
    if m == nil {
        return nil
    } else {
        return m.rootCertificate
    }
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *WindowsPhone81SCEPCertificateProfile) GetScepServerUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.scepServerUrls
    }
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *WindowsPhone81SCEPCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectAlternativeNameFormatString
    }
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *WindowsPhone81SCEPCertificateProfile) GetSubjectNameFormatString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectNameFormatString
    }
}
// Serialize serializes information the current object
func (m *WindowsPhone81SCEPCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsPhone81CertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetHashAlgorithm() != nil {
        cast := (*m.GetHashAlgorithm()).String()
        err = writer.WriteStringValue("hashAlgorithm", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetKeySize() != nil {
        cast := (*m.GetKeySize()).String()
        err = writer.WriteStringValue("keySize", &cast)
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
    if m.GetManagedDeviceCertificateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceCertificateStates()))
        for i, v := range m.GetManagedDeviceCertificateStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificate", m.GetRootCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetScepServerUrls() != nil {
        err = writer.WriteCollectionOfStringValues("scepServerUrls", m.GetScepServerUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectAlternativeNameFormatString", m.GetSubjectAlternativeNameFormatString())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectNameFormatString", m.GetSubjectNameFormatString())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPhone81SCEPCertificateProfile) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHashAlgorithm sets the hashAlgorithm property value. SCEP Hash Algorithm. Possible values are: sha1, sha2.
func (m *WindowsPhone81SCEPCertificateProfile) SetHashAlgorithm(value *HashAlgorithms)() {
    if m != nil {
        m.hashAlgorithm = value
    }
}
// SetKeySize sets the keySize property value. SCEP Key Size. Possible values are: size1024, size2048, size4096.
func (m *WindowsPhone81SCEPCertificateProfile) SetKeySize(value *KeySize)() {
    if m != nil {
        m.keySize = value
    }
}
// SetKeyUsage sets the keyUsage property value. SCEP Key Usage. Possible values are: keyEncipherment, digitalSignature.
func (m *WindowsPhone81SCEPCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    if m != nil {
        m.keyUsage = value
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *WindowsPhone81SCEPCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    if m != nil {
        m.managedDeviceCertificateStates = value
    }
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate.
func (m *WindowsPhone81SCEPCertificateProfile) SetRootCertificate(value WindowsPhone81TrustedRootCertificateable)() {
    if m != nil {
        m.rootCertificate = value
    }
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *WindowsPhone81SCEPCertificateProfile) SetScepServerUrls(value []string)() {
    if m != nil {
        m.scepServerUrls = value
    }
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *WindowsPhone81SCEPCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    if m != nil {
        m.subjectAlternativeNameFormatString = value
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *WindowsPhone81SCEPCertificateProfile) SetSubjectNameFormatString(value *string)() {
    if m != nil {
        m.subjectNameFormatString = value
    }
}
