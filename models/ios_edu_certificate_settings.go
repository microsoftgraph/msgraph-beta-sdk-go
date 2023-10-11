package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// IosEduCertificateSettings trusted Root and PFX certificates for iOS EDU.
type IosEduCertificateSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewIosEduCertificateSettings instantiates a new iosEduCertificateSettings and sets the default values.
func NewIosEduCertificateSettings()(*IosEduCertificateSettings) {
    m := &IosEduCertificateSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIosEduCertificateSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEduCertificateSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEduCertificateSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosEduCertificateSettings) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *IosEduCertificateSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
func (m *IosEduCertificateSettings) GetCertFileName()(*string) {
    val, err := m.GetBackingStore().Get("certFileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateTemplateName gets the certificateTemplateName property value. PKCS Certificate Template Name.
func (m *IosEduCertificateSettings) GetCertificateTemplateName()(*string) {
    val, err := m.GetBackingStore().Get("certificateTemplateName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *IosEduCertificateSettings) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
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
func (m *IosEduCertificateSettings) GetCertificateValidityPeriodValue()(*int32) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCertificationAuthority gets the certificationAuthority property value. PKCS Certification Authority.
func (m *IosEduCertificateSettings) GetCertificationAuthority()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificationAuthorityName gets the certificationAuthorityName property value. PKCS Certification Authority Name.
func (m *IosEduCertificateSettings) GetCertificationAuthorityName()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthorityName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosEduCertificateSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertFileName(val)
        }
        return nil
    }
    res["certificateTemplateName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateTemplateName(val)
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
    res["certificationAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthority(val)
        }
        return nil
    }
    res["certificationAuthorityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    res["trustedRootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustedRootCertificate(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosEduCertificateSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *IosEduCertificateSettings) GetRenewalThresholdPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("renewalThresholdPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate.
func (m *IosEduCertificateSettings) GetTrustedRootCertificate()([]byte) {
    val, err := m.GetBackingStore().Get("trustedRootCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosEduCertificateSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *IosEduCertificateSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCertFileName sets the certFileName property value. File name to display in UI.
func (m *IosEduCertificateSettings) SetCertFileName(value *string)() {
    err := m.GetBackingStore().Set("certFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateTemplateName sets the certificateTemplateName property value. PKCS Certificate Template Name.
func (m *IosEduCertificateSettings) SetCertificateTemplateName(value *string)() {
    err := m.GetBackingStore().Set("certificateTemplateName", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *IosEduCertificateSettings) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodScale", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *IosEduCertificateSettings) SetCertificateValidityPeriodValue(value *int32)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodValue", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthority sets the certificationAuthority property value. PKCS Certification Authority.
func (m *IosEduCertificateSettings) SetCertificationAuthority(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthority", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthorityName sets the certificationAuthorityName property value. PKCS Certification Authority Name.
func (m *IosEduCertificateSettings) SetCertificationAuthorityName(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthorityName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosEduCertificateSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *IosEduCertificateSettings) SetRenewalThresholdPercentage(value *int32)() {
    err := m.GetBackingStore().Set("renewalThresholdPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate.
func (m *IosEduCertificateSettings) SetTrustedRootCertificate(value []byte)() {
    err := m.GetBackingStore().Set("trustedRootCertificate", value)
    if err != nil {
        panic(err)
    }
}
// IosEduCertificateSettingsable 
type IosEduCertificateSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCertFileName()(*string)
    GetCertificateTemplateName()(*string)
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetCertificationAuthority()(*string)
    GetCertificationAuthorityName()(*string)
    GetOdataType()(*string)
    GetRenewalThresholdPercentage()(*int32)
    GetTrustedRootCertificate()([]byte)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCertFileName(value *string)()
    SetCertificateTemplateName(value *string)()
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetCertificationAuthority(value *string)()
    SetCertificationAuthorityName(value *string)()
    SetOdataType(value *string)()
    SetRenewalThresholdPercentage(value *int32)()
    SetTrustedRootCertificate(value []byte)()
}
