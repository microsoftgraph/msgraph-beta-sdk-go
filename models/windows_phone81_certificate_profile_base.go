package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81CertificateProfileBase 
type WindowsPhone81CertificateProfileBase struct {
    DeviceConfiguration
    // Certificate Validity Period Options.
    certificateValidityPeriodScale *CertificateValidityPeriodScale
    // Value for the Certificate Validtiy Period.
    certificateValidityPeriodValue *int32
    // Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
    extendedKeyUsages []ExtendedKeyUsageable
    // Key Storage Provider (KSP) Import Options.
    keyStorageProvider *KeyStorageProviderOption
    // Certificate renewal threshold percentage.
    renewalThresholdPercentage *int32
    // Subject Alternative Name Options.
    subjectAlternativeNameType *SubjectAlternativeNameType
    // Subject Name Format Options.
    subjectNameFormat *SubjectNameFormat
}
// NewWindowsPhone81CertificateProfileBase instantiates a new windowsPhone81CertificateProfileBase and sets the default values.
func NewWindowsPhone81CertificateProfileBase()(*WindowsPhone81CertificateProfileBase) {
    m := &WindowsPhone81CertificateProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81CertificateProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsPhone81CertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81CertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsPhone81SCEPCertificateProfile":
                        return NewWindowsPhone81SCEPCertificateProfile(), nil
                }
            }
        }
    }
    return NewWindowsPhone81CertificateProfileBase(), nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *WindowsPhone81CertificateProfileBase) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodScale
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validtiy Period.
func (m *WindowsPhone81CertificateProfileBase) GetCertificateValidityPeriodValue()(*int32) {
    return m.certificateValidityPeriodValue
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *WindowsPhone81CertificateProfileBase) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    return m.extendedKeyUsages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81CertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["certificateValidityPeriodScale"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCertificateValidityPeriodScale , m.SetCertificateValidityPeriodScale)
    res["certificateValidityPeriodValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCertificateValidityPeriodValue)
    res["extendedKeyUsages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtendedKeyUsageFromDiscriminatorValue , m.SetExtendedKeyUsages)
    res["keyStorageProvider"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseKeyStorageProviderOption , m.SetKeyStorageProvider)
    res["renewalThresholdPercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRenewalThresholdPercentage)
    res["subjectAlternativeNameType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubjectAlternativeNameType , m.SetSubjectAlternativeNameType)
    res["subjectNameFormat"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubjectNameFormat , m.SetSubjectNameFormat)
    return res
}
// GetKeyStorageProvider gets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *WindowsPhone81CertificateProfileBase) GetKeyStorageProvider()(*KeyStorageProviderOption) {
    return m.keyStorageProvider
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage.
func (m *WindowsPhone81CertificateProfileBase) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// GetSubjectAlternativeNameType gets the subjectAlternativeNameType property value. Subject Alternative Name Options.
func (m *WindowsPhone81CertificateProfileBase) GetSubjectAlternativeNameType()(*SubjectAlternativeNameType) {
    return m.subjectAlternativeNameType
}
// GetSubjectNameFormat gets the subjectNameFormat property value. Subject Name Format Options.
func (m *WindowsPhone81CertificateProfileBase) GetSubjectNameFormat()(*SubjectNameFormat) {
    return m.subjectNameFormat
}
// Serialize serializes information the current object
func (m *WindowsPhone81CertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetKeyStorageProvider() != nil {
        cast := (*m.GetKeyStorageProvider()).String()
        err = writer.WriteStringValue("keyStorageProvider", &cast)
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
    if m.GetSubjectAlternativeNameType() != nil {
        cast := (*m.GetSubjectAlternativeNameType()).String()
        err = writer.WriteStringValue("subjectAlternativeNameType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubjectNameFormat() != nil {
        cast := (*m.GetSubjectNameFormat()).String()
        err = writer.WriteStringValue("subjectNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *WindowsPhone81CertificateProfileBase) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodScale = value
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validtiy Period.
func (m *WindowsPhone81CertificateProfileBase) SetCertificateValidityPeriodValue(value *int32)() {
    m.certificateValidityPeriodValue = value
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *WindowsPhone81CertificateProfileBase) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    m.extendedKeyUsages = value
}
// SetKeyStorageProvider sets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *WindowsPhone81CertificateProfileBase) SetKeyStorageProvider(value *KeyStorageProviderOption)() {
    m.keyStorageProvider = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage.
func (m *WindowsPhone81CertificateProfileBase) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
// SetSubjectAlternativeNameType sets the subjectAlternativeNameType property value. Subject Alternative Name Options.
func (m *WindowsPhone81CertificateProfileBase) SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)() {
    m.subjectAlternativeNameType = value
}
// SetSubjectNameFormat sets the subjectNameFormat property value. Subject Name Format Options.
func (m *WindowsPhone81CertificateProfileBase) SetSubjectNameFormat(value *SubjectNameFormat)() {
    m.subjectNameFormat = value
}
