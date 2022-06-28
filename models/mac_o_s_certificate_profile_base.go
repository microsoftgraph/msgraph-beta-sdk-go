package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSCertificateProfileBase mac OS certificate profile.
type MacOSCertificateProfileBase struct {
    DeviceConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Scale for the Certificate Validity Period. Possible values are: days, months, years.
    certificateValidityPeriodScale *CertificateValidityPeriodScale
    // Value for the Certificate Validity Period.
    certificateValidityPeriodValue *int32
    // Certificate renewal threshold percentage.
    renewalThresholdPercentage *int32
    // Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
    subjectAlternativeNameType *SubjectAlternativeNameType
    // Certificate Subject Name Format. Possible values are: commonName, commonNameAsEmail, custom, commonNameIncludingEmail, commonNameAsIMEI, commonNameAsSerialNumber.
    subjectNameFormat *AppleSubjectNameFormat
}
// NewMacOSCertificateProfileBase instantiates a new macOSCertificateProfileBase and sets the default values.
func NewMacOSCertificateProfileBase()(*MacOSCertificateProfileBase) {
    m := &MacOSCertificateProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSCertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSCertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.macOSImportedPFXCertificateProfile":
                        return NewMacOSImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.macOSPkcsCertificateProfile":
                        return NewMacOSPkcsCertificateProfile(), nil
                    case "#microsoft.graph.macOSScepCertificateProfile":
                        return NewMacOSScepCertificateProfile(), nil
                }
            }
        }
    }
    return NewMacOSCertificateProfileBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSCertificateProfileBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Scale for the Certificate Validity Period. Possible values are: days, months, years.
func (m *MacOSCertificateProfileBase) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    if m == nil {
        return nil
    } else {
        return m.certificateValidityPeriodScale
    }
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *MacOSCertificateProfileBase) GetCertificateValidityPeriodValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.certificateValidityPeriodValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSCertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseAppleSubjectNameFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormat(val.(*AppleSubjectNameFormat))
        }
        return nil
    }
    return res
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage.
func (m *MacOSCertificateProfileBase) GetRenewalThresholdPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.renewalThresholdPercentage
    }
}
// GetSubjectAlternativeNameType gets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *MacOSCertificateProfileBase) GetSubjectAlternativeNameType()(*SubjectAlternativeNameType) {
    if m == nil {
        return nil
    } else {
        return m.subjectAlternativeNameType
    }
}
// GetSubjectNameFormat gets the subjectNameFormat property value. Certificate Subject Name Format. Possible values are: commonName, commonNameAsEmail, custom, commonNameIncludingEmail, commonNameAsIMEI, commonNameAsSerialNumber.
func (m *MacOSCertificateProfileBase) GetSubjectNameFormat()(*AppleSubjectNameFormat) {
    if m == nil {
        return nil
    } else {
        return m.subjectNameFormat
    }
}
// Serialize serializes information the current object
func (m *MacOSCertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSCertificateProfileBase) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Scale for the Certificate Validity Period. Possible values are: days, months, years.
func (m *MacOSCertificateProfileBase) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    if m != nil {
        m.certificateValidityPeriodScale = value
    }
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *MacOSCertificateProfileBase) SetCertificateValidityPeriodValue(value *int32)() {
    if m != nil {
        m.certificateValidityPeriodValue = value
    }
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage.
func (m *MacOSCertificateProfileBase) SetRenewalThresholdPercentage(value *int32)() {
    if m != nil {
        m.renewalThresholdPercentage = value
    }
}
// SetSubjectAlternativeNameType sets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *MacOSCertificateProfileBase) SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)() {
    if m != nil {
        m.subjectAlternativeNameType = value
    }
}
// SetSubjectNameFormat sets the subjectNameFormat property value. Certificate Subject Name Format. Possible values are: commonName, commonNameAsEmail, custom, commonNameIncludingEmail, commonNameAsIMEI, commonNameAsSerialNumber.
func (m *MacOSCertificateProfileBase) SetSubjectNameFormat(value *AppleSubjectNameFormat)() {
    if m != nil {
        m.subjectNameFormat = value
    }
}
