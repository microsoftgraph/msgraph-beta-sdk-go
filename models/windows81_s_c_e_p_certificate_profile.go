package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81SCEPCertificateProfile 
type Windows81SCEPCertificateProfile struct {
    Windows81CertificateProfileBase
    // Target store certificate. Possible values are: user, machine.
    certificateStore *CertificateStore
    // Hash Algorithm Options.
    hashAlgorithm *HashAlgorithms
    // Key Size Options.
    keySize *KeySize
    // Key Usage Options.
    keyUsage *KeyUsages
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Trusted Root Certificate
    rootCertificate Windows81TrustedRootCertificateable
    // SCEP Server Url(s).
    scepServerUrls []string
    // Custom String that defines the AAD Attribute.
    subjectAlternativeNameFormatString *string
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewWindows81SCEPCertificateProfile instantiates a new Windows81SCEPCertificateProfile and sets the default values.
func NewWindows81SCEPCertificateProfile()(*Windows81SCEPCertificateProfile) {
    m := &Windows81SCEPCertificateProfile{
        Windows81CertificateProfileBase: *NewWindows81CertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows81SCEPCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows81SCEPCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81SCEPCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81SCEPCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *Windows81SCEPCertificateProfile) GetCertificateStore()(*CertificateStore) {
    return m.certificateStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81SCEPCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows81CertificateProfileBase.GetFieldDeserializers()
    res["certificateStore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCertificateStore , m.SetCertificateStore)
    res["hashAlgorithm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseHashAlgorithms , m.SetHashAlgorithm)
    res["keySize"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseKeySize , m.SetKeySize)
    res["keyUsage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseKeyUsages , m.SetKeyUsage)
    res["managedDeviceCertificateStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue , m.SetManagedDeviceCertificateStates)
    res["rootCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindows81TrustedRootCertificateFromDiscriminatorValue , m.SetRootCertificate)
    res["scepServerUrls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetScepServerUrls)
    res["subjectAlternativeNameFormatString"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubjectAlternativeNameFormatString)
    res["subjectNameFormatString"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubjectNameFormatString)
    return res
}
// GetHashAlgorithm gets the hashAlgorithm property value. Hash Algorithm Options.
func (m *Windows81SCEPCertificateProfile) GetHashAlgorithm()(*HashAlgorithms) {
    return m.hashAlgorithm
}
// GetKeySize gets the keySize property value. Key Size Options.
func (m *Windows81SCEPCertificateProfile) GetKeySize()(*KeySize) {
    return m.keySize
}
// GetKeyUsage gets the keyUsage property value. Key Usage Options.
func (m *Windows81SCEPCertificateProfile) GetKeyUsage()(*KeyUsages) {
    return m.keyUsage
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *Windows81SCEPCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    return m.managedDeviceCertificateStates
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate
func (m *Windows81SCEPCertificateProfile) GetRootCertificate()(Windows81TrustedRootCertificateable) {
    return m.rootCertificate
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows81SCEPCertificateProfile) GetScepServerUrls()([]string) {
    return m.scepServerUrls
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *Windows81SCEPCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    return m.subjectAlternativeNameFormatString
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows81SCEPCertificateProfile) GetSubjectNameFormatString()(*string) {
    return m.subjectNameFormatString
}
// Serialize serializes information the current object
func (m *Windows81SCEPCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Windows81CertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateStore() != nil {
        cast := (*m.GetCertificateStore()).String()
        err = writer.WriteStringValue("certificateStore", &cast)
        if err != nil {
            return err
        }
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDeviceCertificateStates())
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
    return nil
}
// SetCertificateStore sets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *Windows81SCEPCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    m.certificateStore = value
}
// SetHashAlgorithm sets the hashAlgorithm property value. Hash Algorithm Options.
func (m *Windows81SCEPCertificateProfile) SetHashAlgorithm(value *HashAlgorithms)() {
    m.hashAlgorithm = value
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *Windows81SCEPCertificateProfile) SetKeySize(value *KeySize)() {
    m.keySize = value
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *Windows81SCEPCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    m.keyUsage = value
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *Windows81SCEPCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    m.managedDeviceCertificateStates = value
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate
func (m *Windows81SCEPCertificateProfile) SetRootCertificate(value Windows81TrustedRootCertificateable)() {
    m.rootCertificate = value
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows81SCEPCertificateProfile) SetScepServerUrls(value []string)() {
    m.scepServerUrls = value
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *Windows81SCEPCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    m.subjectAlternativeNameFormatString = value
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows81SCEPCertificateProfile) SetSubjectNameFormatString(value *string)() {
    m.subjectNameFormatString = value
}
