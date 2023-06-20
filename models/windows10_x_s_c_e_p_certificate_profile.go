package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XSCEPCertificateProfile 
type Windows10XSCEPCertificateProfile struct {
    Windows10XCertificateProfile
}
// NewWindows10XSCEPCertificateProfile instantiates a new Windows10XSCEPCertificateProfile and sets the default values.
func NewWindows10XSCEPCertificateProfile()(*Windows10XSCEPCertificateProfile) {
    m := &Windows10XSCEPCertificateProfile{
        Windows10XCertificateProfile: *NewWindows10XCertificateProfile(),
    }
    odataTypeValue := "#microsoft.graph.windows10XSCEPCertificateProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10XSCEPCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XSCEPCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XSCEPCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. CertificateStore types
func (m *Windows10XSCEPCertificateProfile) GetCertificateStore()(*CertificateStore) {
    val, err := m.GetBackingStore().Get("certificateStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateStore)
    }
    return nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *Windows10XSCEPCertificateProfile) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodScale")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateValidityPeriodScale)
    }
    return nil
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *Windows10XSCEPCertificateProfile) GetCertificateValidityPeriodValue()(*int32) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings.
func (m *Windows10XSCEPCertificateProfile) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    val, err := m.GetBackingStore().Get("extendedKeyUsages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExtendedKeyUsageable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XSCEPCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows10XCertificateProfile.GetFieldDeserializers()
    res["certificateStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateStore(val.(*CertificateStore))
        }
        return nil
    }
    res["certificateValidityPeriodScale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateValidityPeriodScale)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodScale(val.(*CertificateValidityPeriodScale))
        }
        return nil
    }
    res["certificateValidityPeriodValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodValue(val)
        }
        return nil
    }
    res["extendedKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtendedKeyUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExtendedKeyUsageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExtendedKeyUsageable)
                }
            }
            m.SetExtendedKeyUsages(res)
        }
        return nil
    }
    res["hashAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseHashAlgorithms)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HashAlgorithms, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*HashAlgorithms))
                }
            }
            m.SetHashAlgorithm(res)
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
    res["keyStorageProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyStorageProviderOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyStorageProvider(val.(*KeyStorageProviderOption))
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
    res["renewalThresholdPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenewalThresholdPercentage(val)
        }
        return nil
    }
    res["rootCertificateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateId(val)
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetScepServerUrls(res)
        }
        return nil
    }
    res["subjectAlternativeNameFormats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10XCustomSubjectAlternativeNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10XCustomSubjectAlternativeNameable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Windows10XCustomSubjectAlternativeNameable)
                }
            }
            m.SetSubjectAlternativeNameFormats(res)
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
// GetHashAlgorithm gets the hashAlgorithm property value. SCEP Hash Algorithm.
func (m *Windows10XSCEPCertificateProfile) GetHashAlgorithm()([]HashAlgorithms) {
    val, err := m.GetBackingStore().Get("hashAlgorithm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HashAlgorithms)
    }
    return nil
}
// GetKeySize gets the keySize property value. Key Size Options.
func (m *Windows10XSCEPCertificateProfile) GetKeySize()(*KeySize) {
    val, err := m.GetBackingStore().Get("keySize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KeySize)
    }
    return nil
}
// GetKeyStorageProvider gets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10XSCEPCertificateProfile) GetKeyStorageProvider()(*KeyStorageProviderOption) {
    val, err := m.GetBackingStore().Get("keyStorageProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KeyStorageProviderOption)
    }
    return nil
}
// GetKeyUsage gets the keyUsage property value. Key Usage Options.
func (m *Windows10XSCEPCertificateProfile) GetKeyUsage()(*KeyUsages) {
    val, err := m.GetBackingStore().Get("keyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KeyUsages)
    }
    return nil
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage
func (m *Windows10XSCEPCertificateProfile) GetRenewalThresholdPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("renewalThresholdPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRootCertificateId gets the rootCertificateId property value. Trusted Root Certificate ID
func (m *Windows10XSCEPCertificateProfile) GetRootCertificateId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("rootCertificateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows10XSCEPCertificateProfile) GetScepServerUrls()([]string) {
    val, err := m.GetBackingStore().Get("scepServerUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSubjectAlternativeNameFormats gets the subjectAlternativeNameFormats property value. Custom AAD Attributes.
func (m *Windows10XSCEPCertificateProfile) GetSubjectAlternativeNameFormats()([]Windows10XCustomSubjectAlternativeNameable) {
    val, err := m.GetBackingStore().Get("subjectAlternativeNameFormats")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Windows10XCustomSubjectAlternativeNameable)
    }
    return nil
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows10XSCEPCertificateProfile) GetSubjectNameFormatString()(*string) {
    val, err := m.GetBackingStore().Get("subjectNameFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtendedKeyUsages()))
        for i, v := range m.GetExtendedKeyUsages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        err = writer.WriteUUIDValue("rootCertificateId", m.GetRootCertificateId())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjectAlternativeNameFormats()))
        for i, v := range m.GetSubjectAlternativeNameFormats() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("certificateStore", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *Windows10XSCEPCertificateProfile) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodScale", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *Windows10XSCEPCertificateProfile) SetCertificateValidityPeriodValue(value *int32)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodValue", value)
    if err != nil {
        panic(err)
    }
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings.
func (m *Windows10XSCEPCertificateProfile) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    err := m.GetBackingStore().Set("extendedKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetHashAlgorithm sets the hashAlgorithm property value. SCEP Hash Algorithm.
func (m *Windows10XSCEPCertificateProfile) SetHashAlgorithm(value []HashAlgorithms)() {
    err := m.GetBackingStore().Set("hashAlgorithm", value)
    if err != nil {
        panic(err)
    }
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *Windows10XSCEPCertificateProfile) SetKeySize(value *KeySize)() {
    err := m.GetBackingStore().Set("keySize", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyStorageProvider sets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10XSCEPCertificateProfile) SetKeyStorageProvider(value *KeyStorageProviderOption)() {
    err := m.GetBackingStore().Set("keyStorageProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *Windows10XSCEPCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    err := m.GetBackingStore().Set("keyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage
func (m *Windows10XSCEPCertificateProfile) SetRenewalThresholdPercentage(value *int32)() {
    err := m.GetBackingStore().Set("renewalThresholdPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateId sets the rootCertificateId property value. Trusted Root Certificate ID
func (m *Windows10XSCEPCertificateProfile) SetRootCertificateId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("rootCertificateId", value)
    if err != nil {
        panic(err)
    }
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows10XSCEPCertificateProfile) SetScepServerUrls(value []string)() {
    err := m.GetBackingStore().Set("scepServerUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectAlternativeNameFormats sets the subjectAlternativeNameFormats property value. Custom AAD Attributes.
func (m *Windows10XSCEPCertificateProfile) SetSubjectAlternativeNameFormats(value []Windows10XCustomSubjectAlternativeNameable)() {
    err := m.GetBackingStore().Set("subjectAlternativeNameFormats", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows10XSCEPCertificateProfile) SetSubjectNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// Windows10XSCEPCertificateProfileable 
type Windows10XSCEPCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Windows10XCertificateProfileable
    GetCertificateStore()(*CertificateStore)
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetHashAlgorithm()([]HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyStorageProvider()(*KeyStorageProviderOption)
    GetKeyUsage()(*KeyUsages)
    GetRenewalThresholdPercentage()(*int32)
    GetRootCertificateId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormats()([]Windows10XCustomSubjectAlternativeNameable)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *CertificateStore)()
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetHashAlgorithm(value []HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyStorageProvider(value *KeyStorageProviderOption)()
    SetKeyUsage(value *KeyUsages)()
    SetRenewalThresholdPercentage(value *int32)()
    SetRootCertificateId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormats(value []Windows10XCustomSubjectAlternativeNameable)()
    SetSubjectNameFormatString(value *string)()
}
