package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XSCEPCertificateProfile 
type Windows10XSCEPCertificateProfile struct {
    Windows10XCertificateProfile
    // CertificateStore types
    certificateStore *CertificateStore
    // Certificate Validity Period Options.
    certificateValidityPeriodScale *CertificateValidityPeriodScale
    // Value for the Certificate Validity Period
    certificateValidityPeriodValue *int32
    // Extended Key Usage (EKU) settings.
    extendedKeyUsages []ExtendedKeyUsageable
    // SCEP Hash Algorithm.
    hashAlgorithm []HashAlgorithms
    // Key Size Options.
    keySize *KeySize
    // Key Storage Provider (KSP) Import Options.
    keyStorageProvider *KeyStorageProviderOption
    // Key Usage Options.
    keyUsage *KeyUsages
    // Certificate renewal threshold percentage
    renewalThresholdPercentage *int32
    // Trusted Root Certificate ID
    rootCertificateId *string
    // SCEP Server Url(s).
    scepServerUrls []string
    // Custom AAD Attributes.
    subjectAlternativeNameFormats []Windows10XCustomSubjectAlternativeNameable
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewWindows10XSCEPCertificateProfile instantiates a new Windows10XSCEPCertificateProfile and sets the default values.
func NewWindows10XSCEPCertificateProfile()(*Windows10XSCEPCertificateProfile) {
    m := &Windows10XSCEPCertificateProfile{
        Windows10XCertificateProfile: *NewWindows10XCertificateProfile(),
    }
    odataTypeValue := "#microsoft.graph.windows10XSCEPCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10XSCEPCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XSCEPCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XSCEPCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. CertificateStore types
func (m *Windows10XSCEPCertificateProfile) GetCertificateStore()(*CertificateStore) {
    return m.certificateStore
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *Windows10XSCEPCertificateProfile) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodScale
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *Windows10XSCEPCertificateProfile) GetCertificateValidityPeriodValue()(*int32) {
    return m.certificateValidityPeriodValue
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings.
func (m *Windows10XSCEPCertificateProfile) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    return m.extendedKeyUsages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XSCEPCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows10XCertificateProfile.GetFieldDeserializers()
    res["certificateStore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCertificateStore , m.SetCertificateStore)
    res["certificateValidityPeriodScale"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCertificateValidityPeriodScale , m.SetCertificateValidityPeriodScale)
    res["certificateValidityPeriodValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCertificateValidityPeriodValue)
    res["extendedKeyUsages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtendedKeyUsageFromDiscriminatorValue , m.SetExtendedKeyUsages)
    res["hashAlgorithm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseHashAlgorithms , m.SetHashAlgorithm)
    res["keySize"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseKeySize , m.SetKeySize)
    res["keyStorageProvider"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseKeyStorageProviderOption , m.SetKeyStorageProvider)
    res["keyUsage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseKeyUsages , m.SetKeyUsage)
    res["renewalThresholdPercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRenewalThresholdPercentage)
    res["rootCertificateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRootCertificateId)
    res["scepServerUrls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetScepServerUrls)
    res["subjectAlternativeNameFormats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindows10XCustomSubjectAlternativeNameFromDiscriminatorValue , m.SetSubjectAlternativeNameFormats)
    res["subjectNameFormatString"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubjectNameFormatString)
    return res
}
// GetHashAlgorithm gets the hashAlgorithm property value. SCEP Hash Algorithm.
func (m *Windows10XSCEPCertificateProfile) GetHashAlgorithm()([]HashAlgorithms) {
    return m.hashAlgorithm
}
// GetKeySize gets the keySize property value. Key Size Options.
func (m *Windows10XSCEPCertificateProfile) GetKeySize()(*KeySize) {
    return m.keySize
}
// GetKeyStorageProvider gets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10XSCEPCertificateProfile) GetKeyStorageProvider()(*KeyStorageProviderOption) {
    return m.keyStorageProvider
}
// GetKeyUsage gets the keyUsage property value. Key Usage Options.
func (m *Windows10XSCEPCertificateProfile) GetKeyUsage()(*KeyUsages) {
    return m.keyUsage
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage
func (m *Windows10XSCEPCertificateProfile) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// GetRootCertificateId gets the rootCertificateId property value. Trusted Root Certificate ID
func (m *Windows10XSCEPCertificateProfile) GetRootCertificateId()(*string) {
    return m.rootCertificateId
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows10XSCEPCertificateProfile) GetScepServerUrls()([]string) {
    return m.scepServerUrls
}
// GetSubjectAlternativeNameFormats gets the subjectAlternativeNameFormats property value. Custom AAD Attributes.
func (m *Windows10XSCEPCertificateProfile) GetSubjectAlternativeNameFormats()([]Windows10XCustomSubjectAlternativeNameable) {
    return m.subjectAlternativeNameFormats
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows10XSCEPCertificateProfile) GetSubjectNameFormatString()(*string) {
    return m.subjectNameFormatString
}
// Serialize serializes information the current object
func (m *Windows10XSCEPCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Windows10XCertificateProfile.Serialize(writer)
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
    if m.GetCertificateValidityPeriodScale() != nil {
        cast := (*m.GetCertificateValidityPeriodScale()).String()
        err = writer.WriteStringValue("certificateValidityPeriodScale", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateValidityPeriodValue", m.GetCertificateValidityPeriodValue())
        if err != nil {
            return err
        }
    }
    if m.GetExtendedKeyUsages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtendedKeyUsages())
        err = writer.WriteCollectionOfObjectValues("extendedKeyUsages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHashAlgorithm() != nil {
        err = writer.WriteCollectionOfStringValues("hashAlgorithm", SerializeHashAlgorithms(m.GetHashAlgorithm()))
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
    if m.GetKeyStorageProvider() != nil {
        cast := (*m.GetKeyStorageProvider()).String()
        err = writer.WriteStringValue("keyStorageProvider", &cast)
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
        err = writer.WriteInt32Value("renewalThresholdPercentage", m.GetRenewalThresholdPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rootCertificateId", m.GetRootCertificateId())
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
    if m.GetSubjectAlternativeNameFormats() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSubjectAlternativeNameFormats())
        err = writer.WriteCollectionOfObjectValues("subjectAlternativeNameFormats", cast)
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
// SetCertificateStore sets the certificateStore property value. CertificateStore types
func (m *Windows10XSCEPCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    m.certificateStore = value
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *Windows10XSCEPCertificateProfile) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodScale = value
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *Windows10XSCEPCertificateProfile) SetCertificateValidityPeriodValue(value *int32)() {
    m.certificateValidityPeriodValue = value
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings.
func (m *Windows10XSCEPCertificateProfile) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    m.extendedKeyUsages = value
}
// SetHashAlgorithm sets the hashAlgorithm property value. SCEP Hash Algorithm.
func (m *Windows10XSCEPCertificateProfile) SetHashAlgorithm(value []HashAlgorithms)() {
    m.hashAlgorithm = value
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *Windows10XSCEPCertificateProfile) SetKeySize(value *KeySize)() {
    m.keySize = value
}
// SetKeyStorageProvider sets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10XSCEPCertificateProfile) SetKeyStorageProvider(value *KeyStorageProviderOption)() {
    m.keyStorageProvider = value
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *Windows10XSCEPCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    m.keyUsage = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage
func (m *Windows10XSCEPCertificateProfile) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
// SetRootCertificateId sets the rootCertificateId property value. Trusted Root Certificate ID
func (m *Windows10XSCEPCertificateProfile) SetRootCertificateId(value *string)() {
    m.rootCertificateId = value
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows10XSCEPCertificateProfile) SetScepServerUrls(value []string)() {
    m.scepServerUrls = value
}
// SetSubjectAlternativeNameFormats sets the subjectAlternativeNameFormats property value. Custom AAD Attributes.
func (m *Windows10XSCEPCertificateProfile) SetSubjectAlternativeNameFormats(value []Windows10XCustomSubjectAlternativeNameable)() {
    m.subjectAlternativeNameFormats = value
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows10XSCEPCertificateProfile) SetSubjectNameFormatString(value *string)() {
    m.subjectNameFormatString = value
}
