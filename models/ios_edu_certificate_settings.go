package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEduCertificateSettings trusted Root and PFX certificates for iOS EDU.
type IosEduCertificateSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // File name to display in UI.
    certFileName *string
    // PKCS Certificate Template Name.
    certificateTemplateName *string
    // Certificate Validity Period Options.
    certificateValidityPeriodScale *CertificateValidityPeriodScale
    // Value for the Certificate Validity Period.
    certificateValidityPeriodValue *int32
    // PKCS Certification Authority.
    certificationAuthority *string
    // PKCS Certification Authority Name.
    certificationAuthorityName *string
    // The OdataType property
    odataType *string
    // Certificate renewal threshold percentage. Valid values 1 to 99
    renewalThresholdPercentage *int32
    // Trusted Root Certificate.
    trustedRootCertificate []byte
}
// NewIosEduCertificateSettings instantiates a new iosEduCertificateSettings and sets the default values.
func NewIosEduCertificateSettings()(*IosEduCertificateSettings) {
    m := &IosEduCertificateSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.iosEduCertificateSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosEduCertificateSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEduCertificateSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEduCertificateSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosEduCertificateSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
func (m *IosEduCertificateSettings) GetCertFileName()(*string) {
    return m.certFileName
}
// GetCertificateTemplateName gets the certificateTemplateName property value. PKCS Certificate Template Name.
func (m *IosEduCertificateSettings) GetCertificateTemplateName()(*string) {
    return m.certificateTemplateName
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *IosEduCertificateSettings) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodScale
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *IosEduCertificateSettings) GetCertificateValidityPeriodValue()(*int32) {
    return m.certificateValidityPeriodValue
}
// GetCertificationAuthority gets the certificationAuthority property value. PKCS Certification Authority.
func (m *IosEduCertificateSettings) GetCertificationAuthority()(*string) {
    return m.certificationAuthority
}
// GetCertificationAuthorityName gets the certificationAuthorityName property value. PKCS Certification Authority Name.
func (m *IosEduCertificateSettings) GetCertificationAuthorityName()(*string) {
    return m.certificationAuthorityName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosEduCertificateSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certFileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCertFileName)
    res["certificateTemplateName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCertificateTemplateName)
    res["certificateValidityPeriodScale"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCertificateValidityPeriodScale , m.SetCertificateValidityPeriodScale)
    res["certificateValidityPeriodValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCertificateValidityPeriodValue)
    res["certificationAuthority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCertificationAuthority)
    res["certificationAuthorityName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCertificationAuthorityName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["renewalThresholdPercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRenewalThresholdPercentage)
    res["trustedRootCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetTrustedRootCertificate)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosEduCertificateSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *IosEduCertificateSettings) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate.
func (m *IosEduCertificateSettings) GetTrustedRootCertificate()([]byte) {
    return m.trustedRootCertificate
}
// Serialize serializes information the current object
func (m *IosEduCertificateSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("certFileName", m.GetCertFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("certificateTemplateName", m.GetCertificateTemplateName())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateValidityPeriodScale() != nil {
        cast := (*m.GetCertificateValidityPeriodScale()).String()
        err := writer.WriteStringValue("certificateValidityPeriodScale", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("certificateValidityPeriodValue", m.GetCertificateValidityPeriodValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("certificationAuthority", m.GetCertificationAuthority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("certificationAuthorityName", m.GetCertificationAuthorityName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("renewalThresholdPercentage", m.GetRenewalThresholdPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("trustedRootCertificate", m.GetTrustedRootCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosEduCertificateSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCertFileName sets the certFileName property value. File name to display in UI.
func (m *IosEduCertificateSettings) SetCertFileName(value *string)() {
    m.certFileName = value
}
// SetCertificateTemplateName sets the certificateTemplateName property value. PKCS Certificate Template Name.
func (m *IosEduCertificateSettings) SetCertificateTemplateName(value *string)() {
    m.certificateTemplateName = value
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *IosEduCertificateSettings) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodScale = value
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *IosEduCertificateSettings) SetCertificateValidityPeriodValue(value *int32)() {
    m.certificateValidityPeriodValue = value
}
// SetCertificationAuthority sets the certificationAuthority property value. PKCS Certification Authority.
func (m *IosEduCertificateSettings) SetCertificationAuthority(value *string)() {
    m.certificationAuthority = value
}
// SetCertificationAuthorityName sets the certificationAuthorityName property value. PKCS Certification Authority Name.
func (m *IosEduCertificateSettings) SetCertificationAuthorityName(value *string)() {
    m.certificationAuthorityName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosEduCertificateSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *IosEduCertificateSettings) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate.
func (m *IosEduCertificateSettings) SetTrustedRootCertificate(value []byte)() {
    m.trustedRootCertificate = value
}
