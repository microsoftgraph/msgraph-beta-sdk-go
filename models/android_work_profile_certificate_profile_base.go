package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileCertificateProfileBase 
type AndroidWorkProfileCertificateProfileBase struct {
    DeviceConfiguration
}
// NewAndroidWorkProfileCertificateProfileBase instantiates a new AndroidWorkProfileCertificateProfileBase and sets the default values.
func NewAndroidWorkProfileCertificateProfileBase()(*AndroidWorkProfileCertificateProfileBase) {
    m := &AndroidWorkProfileCertificateProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileCertificateProfileBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidWorkProfileCertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileCertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidWorkProfilePkcsCertificateProfile":
                        return NewAndroidWorkProfilePkcsCertificateProfile(), nil
                    case "#microsoft.graph.androidWorkProfileScepCertificateProfile":
                        return NewAndroidWorkProfileScepCertificateProfile(), nil
                }
            }
        }
    }
    return NewAndroidWorkProfileCertificateProfileBase(), nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *AndroidWorkProfileCertificateProfileBase) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodScale")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateValidityPeriodScale)
    }
    return nil
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *AndroidWorkProfileCertificateProfileBase) GetCertificateValidityPeriodValue()(*int32) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileCertificateProfileBase) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
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
func (m *AndroidWorkProfileCertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
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
    res["rootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidWorkProfileTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificate(val.(AndroidWorkProfileTrustedRootCertificateable))
        }
        return nil
    }
    res["subjectAlternativeNameType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectAlternativeNameType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectAlternativeNameType(val.(*SubjectAlternativeNameType))
        }
        return nil
    }
    res["subjectNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectNameFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormat(val.(*SubjectNameFormat))
        }
        return nil
    }
    return res
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *AndroidWorkProfileCertificateProfileBase) GetRenewalThresholdPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("renewalThresholdPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate.
func (m *AndroidWorkProfileCertificateProfileBase) GetRootCertificate()(AndroidWorkProfileTrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidWorkProfileTrustedRootCertificateable)
    }
    return nil
}
// GetSubjectAlternativeNameType gets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *AndroidWorkProfileCertificateProfileBase) GetSubjectAlternativeNameType()(*SubjectAlternativeNameType) {
    val, err := m.GetBackingStore().Get("subjectAlternativeNameType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubjectAlternativeNameType)
    }
    return nil
}
// GetSubjectNameFormat gets the subjectNameFormat property value. Subject Name Format Options.
func (m *AndroidWorkProfileCertificateProfileBase) GetSubjectNameFormat()(*SubjectNameFormat) {
    val, err := m.GetBackingStore().Get("subjectNameFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubjectNameFormat)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileCertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteInt32Value("renewalThresholdPercentage", m.GetRenewalThresholdPercentage())
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
func (m *AndroidWorkProfileCertificateProfileBase) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodScale", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *AndroidWorkProfileCertificateProfileBase) SetCertificateValidityPeriodValue(value *int32)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodValue", value)
    if err != nil {
        panic(err)
    }
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfileCertificateProfileBase) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    err := m.GetBackingStore().Set("extendedKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *AndroidWorkProfileCertificateProfileBase) SetRenewalThresholdPercentage(value *int32)() {
    err := m.GetBackingStore().Set("renewalThresholdPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate.
func (m *AndroidWorkProfileCertificateProfileBase) SetRootCertificate(value AndroidWorkProfileTrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectAlternativeNameType sets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *AndroidWorkProfileCertificateProfileBase) SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)() {
    err := m.GetBackingStore().Set("subjectAlternativeNameType", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectNameFormat sets the subjectNameFormat property value. Subject Name Format Options.
func (m *AndroidWorkProfileCertificateProfileBase) SetSubjectNameFormat(value *SubjectNameFormat)() {
    err := m.GetBackingStore().Set("subjectNameFormat", value)
    if err != nil {
        panic(err)
    }
}
// AndroidWorkProfileCertificateProfileBaseable 
type AndroidWorkProfileCertificateProfileBaseable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetRenewalThresholdPercentage()(*int32)
    GetRootCertificate()(AndroidWorkProfileTrustedRootCertificateable)
    GetSubjectAlternativeNameType()(*SubjectAlternativeNameType)
    GetSubjectNameFormat()(*SubjectNameFormat)
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetRenewalThresholdPercentage(value *int32)()
    SetRootCertificate(value AndroidWorkProfileTrustedRootCertificateable)()
    SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)()
    SetSubjectNameFormat(value *SubjectNameFormat)()
}
