package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ManagedDeviceCertificateState struct {
    Entity
}
// NewManagedDeviceCertificateState instantiates a new ManagedDeviceCertificateState and sets the default values.
func NewManagedDeviceCertificateState()(*ManagedDeviceCertificateState) {
    m := &ManagedDeviceCertificateState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceCertificateStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedDeviceCertificateStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceCertificateState(), nil
}
// GetCertificateEnhancedKeyUsage gets the certificateEnhancedKeyUsage property value. Extended key usage
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetCertificateEnhancedKeyUsage()(*string) {
    val, err := m.GetBackingStore().Get("certificateEnhancedKeyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateErrorCode gets the certificateErrorCode property value. Error code
// returns a *int32 when successful
func (m *ManagedDeviceCertificateState) GetCertificateErrorCode()(*int32) {
    val, err := m.GetBackingStore().Get("certificateErrorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCertificateExpirationDateTime gets the certificateExpirationDateTime property value. Certificate expiry date
// returns a *Time when successful
func (m *ManagedDeviceCertificateState) GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("certificateExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCertificateIssuanceDateTime gets the certificateIssuanceDateTime property value. Issuance date
// returns a *Time when successful
func (m *ManagedDeviceCertificateState) GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("certificateIssuanceDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCertificateIssuanceState gets the certificateIssuanceState property value. Certificate Issuance State Options.
// returns a *CertificateIssuanceStates when successful
func (m *ManagedDeviceCertificateState) GetCertificateIssuanceState()(*CertificateIssuanceStates) {
    val, err := m.GetBackingStore().Get("certificateIssuanceState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateIssuanceStates)
    }
    return nil
}
// GetCertificateIssuer gets the certificateIssuer property value. Issuer
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetCertificateIssuer()(*string) {
    val, err := m.GetBackingStore().Get("certificateIssuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateKeyLength gets the certificateKeyLength property value. Key length
// returns a *int32 when successful
func (m *ManagedDeviceCertificateState) GetCertificateKeyLength()(*int32) {
    val, err := m.GetBackingStore().Get("certificateKeyLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCertificateKeyStorageProvider gets the certificateKeyStorageProvider property value. Key Storage Provider (KSP) Import Options.
// returns a *KeyStorageProviderOption when successful
func (m *ManagedDeviceCertificateState) GetCertificateKeyStorageProvider()(*KeyStorageProviderOption) {
    val, err := m.GetBackingStore().Get("certificateKeyStorageProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KeyStorageProviderOption)
    }
    return nil
}
// GetCertificateKeyUsage gets the certificateKeyUsage property value. Key Usage Options.
// returns a *KeyUsages when successful
func (m *ManagedDeviceCertificateState) GetCertificateKeyUsage()(*KeyUsages) {
    val, err := m.GetBackingStore().Get("certificateKeyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KeyUsages)
    }
    return nil
}
// GetCertificateLastIssuanceStateChangedDateTime gets the certificateLastIssuanceStateChangedDateTime property value. Last certificate issuance state change
// returns a *Time when successful
func (m *ManagedDeviceCertificateState) GetCertificateLastIssuanceStateChangedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("certificateLastIssuanceStateChangedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCertificateProfileDisplayName gets the certificateProfileDisplayName property value. Certificate profile display name
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetCertificateProfileDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("certificateProfileDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateRevokeStatus gets the certificateRevokeStatus property value. Certificate Revocation Status.
// returns a *CertificateRevocationStatus when successful
func (m *ManagedDeviceCertificateState) GetCertificateRevokeStatus()(*CertificateRevocationStatus) {
    val, err := m.GetBackingStore().Get("certificateRevokeStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateRevocationStatus)
    }
    return nil
}
// GetCertificateSerialNumber gets the certificateSerialNumber property value. Serial number
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetCertificateSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("certificateSerialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateSubjectAlternativeNameFormat gets the certificateSubjectAlternativeNameFormat property value. Subject Alternative Name Options.
// returns a *SubjectAlternativeNameType when successful
func (m *ManagedDeviceCertificateState) GetCertificateSubjectAlternativeNameFormat()(*SubjectAlternativeNameType) {
    val, err := m.GetBackingStore().Get("certificateSubjectAlternativeNameFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubjectAlternativeNameType)
    }
    return nil
}
// GetCertificateSubjectAlternativeNameFormatString gets the certificateSubjectAlternativeNameFormatString property value. Subject alternative name format string for custom formats
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetCertificateSubjectAlternativeNameFormatString()(*string) {
    val, err := m.GetBackingStore().Get("certificateSubjectAlternativeNameFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateSubjectNameFormat gets the certificateSubjectNameFormat property value. Subject Name Format Options.
// returns a *SubjectNameFormat when successful
func (m *ManagedDeviceCertificateState) GetCertificateSubjectNameFormat()(*SubjectNameFormat) {
    val, err := m.GetBackingStore().Get("certificateSubjectNameFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubjectNameFormat)
    }
    return nil
}
// GetCertificateSubjectNameFormatString gets the certificateSubjectNameFormatString property value. Subject name format string for custom subject name formats
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetCertificateSubjectNameFormatString()(*string) {
    val, err := m.GetBackingStore().Get("certificateSubjectNameFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateThumbprint gets the certificateThumbprint property value. Thumbprint
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetCertificateThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("certificateThumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateValidityPeriod gets the certificateValidityPeriod property value. Validity period
// returns a *int32 when successful
func (m *ManagedDeviceCertificateState) GetCertificateValidityPeriod()(*int32) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCertificateValidityPeriodUnits gets the certificateValidityPeriodUnits property value. Certificate Validity Period Options.
// returns a *CertificateValidityPeriodScale when successful
func (m *ManagedDeviceCertificateState) GetCertificateValidityPeriodUnits()(*CertificateValidityPeriodScale) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodUnits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateValidityPeriodScale)
    }
    return nil
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Device display name
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetDeviceDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("deviceDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDevicePlatform gets the devicePlatform property value. Supported platform types.
// returns a *DevicePlatformType when successful
func (m *ManagedDeviceCertificateState) GetDevicePlatform()(*DevicePlatformType) {
    val, err := m.GetBackingStore().Get("devicePlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DevicePlatformType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagedDeviceCertificateState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateEnhancedKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateEnhancedKeyUsage(val)
        }
        return nil
    }
    res["certificateErrorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateErrorCode(val)
        }
        return nil
    }
    res["certificateExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateExpirationDateTime(val)
        }
        return nil
    }
    res["certificateIssuanceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuanceDateTime(val)
        }
        return nil
    }
    res["certificateIssuanceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateIssuanceStates)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuanceState(val.(*CertificateIssuanceStates))
        }
        return nil
    }
    res["certificateIssuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuer(val)
        }
        return nil
    }
    res["certificateKeyLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyLength(val)
        }
        return nil
    }
    res["certificateKeyStorageProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyStorageProviderOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyStorageProvider(val.(*KeyStorageProviderOption))
        }
        return nil
    }
    res["certificateKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyUsages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyUsage(val.(*KeyUsages))
        }
        return nil
    }
    res["certificateLastIssuanceStateChangedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateLastIssuanceStateChangedDateTime(val)
        }
        return nil
    }
    res["certificateProfileDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateProfileDisplayName(val)
        }
        return nil
    }
    res["certificateRevokeStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateRevocationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevokeStatus(val.(*CertificateRevocationStatus))
        }
        return nil
    }
    res["certificateSerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSerialNumber(val)
        }
        return nil
    }
    res["certificateSubjectAlternativeNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectAlternativeNameType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectAlternativeNameFormat(val.(*SubjectAlternativeNameType))
        }
        return nil
    }
    res["certificateSubjectAlternativeNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectAlternativeNameFormatString(val)
        }
        return nil
    }
    res["certificateSubjectNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectNameFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectNameFormat(val.(*SubjectNameFormat))
        }
        return nil
    }
    res["certificateSubjectNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectNameFormatString(val)
        }
        return nil
    }
    res["certificateThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateThumbprint(val)
        }
        return nil
    }
    res["certificateValidityPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriod(val)
        }
        return nil
    }
    res["certificateValidityPeriodUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateValidityPeriodScale)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodUnits(val.(*CertificateValidityPeriodScale))
        }
        return nil
    }
    res["deviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
    res["devicePlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["lastCertificateStateChangeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCertificateStateChangeDateTime(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    return res
}
// GetLastCertificateStateChangeDateTime gets the lastCertificateStateChangeDateTime property value. Last certificate issuance state change
// returns a *Time when successful
func (m *ManagedDeviceCertificateState) GetLastCertificateStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCertificateStateChangeDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. User display name
// returns a *string when successful
func (m *ManagedDeviceCertificateState) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceCertificateState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificateEnhancedKeyUsage", m.GetCertificateEnhancedKeyUsage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateErrorCode", m.GetCertificateErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateExpirationDateTime", m.GetCertificateExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateIssuanceDateTime", m.GetCertificateIssuanceDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateIssuanceState() != nil {
        cast := (*m.GetCertificateIssuanceState()).String()
        err = writer.WriteStringValue("certificateIssuanceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateIssuer", m.GetCertificateIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateKeyLength", m.GetCertificateKeyLength())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateKeyStorageProvider() != nil {
        cast := (*m.GetCertificateKeyStorageProvider()).String()
        err = writer.WriteStringValue("certificateKeyStorageProvider", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertificateKeyUsage() != nil {
        cast := (*m.GetCertificateKeyUsage()).String()
        err = writer.WriteStringValue("certificateKeyUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateLastIssuanceStateChangedDateTime", m.GetCertificateLastIssuanceStateChangedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateProfileDisplayName", m.GetCertificateProfileDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateRevokeStatus() != nil {
        cast := (*m.GetCertificateRevokeStatus()).String()
        err = writer.WriteStringValue("certificateRevokeStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSerialNumber", m.GetCertificateSerialNumber())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateSubjectAlternativeNameFormat() != nil {
        cast := (*m.GetCertificateSubjectAlternativeNameFormat()).String()
        err = writer.WriteStringValue("certificateSubjectAlternativeNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSubjectAlternativeNameFormatString", m.GetCertificateSubjectAlternativeNameFormatString())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateSubjectNameFormat() != nil {
        cast := (*m.GetCertificateSubjectNameFormat()).String()
        err = writer.WriteStringValue("certificateSubjectNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSubjectNameFormatString", m.GetCertificateSubjectNameFormatString())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateThumbprint", m.GetCertificateThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateValidityPeriod", m.GetCertificateValidityPeriod())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateValidityPeriodUnits() != nil {
        cast := (*m.GetCertificateValidityPeriodUnits()).String()
        err = writer.WriteStringValue("certificateValidityPeriodUnits", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDevicePlatform() != nil {
        cast := (*m.GetDevicePlatform()).String()
        err = writer.WriteStringValue("devicePlatform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCertificateStateChangeDateTime", m.GetLastCertificateStateChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateEnhancedKeyUsage sets the certificateEnhancedKeyUsage property value. Extended key usage
func (m *ManagedDeviceCertificateState) SetCertificateEnhancedKeyUsage(value *string)() {
    err := m.GetBackingStore().Set("certificateEnhancedKeyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateErrorCode sets the certificateErrorCode property value. Error code
func (m *ManagedDeviceCertificateState) SetCertificateErrorCode(value *int32)() {
    err := m.GetBackingStore().Set("certificateErrorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateExpirationDateTime sets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedDeviceCertificateState) SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("certificateExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateIssuanceDateTime sets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedDeviceCertificateState) SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("certificateIssuanceDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateIssuanceState sets the certificateIssuanceState property value. Certificate Issuance State Options.
func (m *ManagedDeviceCertificateState) SetCertificateIssuanceState(value *CertificateIssuanceStates)() {
    err := m.GetBackingStore().Set("certificateIssuanceState", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateIssuer sets the certificateIssuer property value. Issuer
func (m *ManagedDeviceCertificateState) SetCertificateIssuer(value *string)() {
    err := m.GetBackingStore().Set("certificateIssuer", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateKeyLength sets the certificateKeyLength property value. Key length
func (m *ManagedDeviceCertificateState) SetCertificateKeyLength(value *int32)() {
    err := m.GetBackingStore().Set("certificateKeyLength", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateKeyStorageProvider sets the certificateKeyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *ManagedDeviceCertificateState) SetCertificateKeyStorageProvider(value *KeyStorageProviderOption)() {
    err := m.GetBackingStore().Set("certificateKeyStorageProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateKeyUsage sets the certificateKeyUsage property value. Key Usage Options.
func (m *ManagedDeviceCertificateState) SetCertificateKeyUsage(value *KeyUsages)() {
    err := m.GetBackingStore().Set("certificateKeyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateLastIssuanceStateChangedDateTime sets the certificateLastIssuanceStateChangedDateTime property value. Last certificate issuance state change
func (m *ManagedDeviceCertificateState) SetCertificateLastIssuanceStateChangedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("certificateLastIssuanceStateChangedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateProfileDisplayName sets the certificateProfileDisplayName property value. Certificate profile display name
func (m *ManagedDeviceCertificateState) SetCertificateProfileDisplayName(value *string)() {
    err := m.GetBackingStore().Set("certificateProfileDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateRevokeStatus sets the certificateRevokeStatus property value. Certificate Revocation Status.
func (m *ManagedDeviceCertificateState) SetCertificateRevokeStatus(value *CertificateRevocationStatus)() {
    err := m.GetBackingStore().Set("certificateRevokeStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSerialNumber sets the certificateSerialNumber property value. Serial number
func (m *ManagedDeviceCertificateState) SetCertificateSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("certificateSerialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSubjectAlternativeNameFormat sets the certificateSubjectAlternativeNameFormat property value. Subject Alternative Name Options.
func (m *ManagedDeviceCertificateState) SetCertificateSubjectAlternativeNameFormat(value *SubjectAlternativeNameType)() {
    err := m.GetBackingStore().Set("certificateSubjectAlternativeNameFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSubjectAlternativeNameFormatString sets the certificateSubjectAlternativeNameFormatString property value. Subject alternative name format string for custom formats
func (m *ManagedDeviceCertificateState) SetCertificateSubjectAlternativeNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("certificateSubjectAlternativeNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSubjectNameFormat sets the certificateSubjectNameFormat property value. Subject Name Format Options.
func (m *ManagedDeviceCertificateState) SetCertificateSubjectNameFormat(value *SubjectNameFormat)() {
    err := m.GetBackingStore().Set("certificateSubjectNameFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSubjectNameFormatString sets the certificateSubjectNameFormatString property value. Subject name format string for custom subject name formats
func (m *ManagedDeviceCertificateState) SetCertificateSubjectNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("certificateSubjectNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateThumbprint sets the certificateThumbprint property value. Thumbprint
func (m *ManagedDeviceCertificateState) SetCertificateThumbprint(value *string)() {
    err := m.GetBackingStore().Set("certificateThumbprint", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriod sets the certificateValidityPeriod property value. Validity period
func (m *ManagedDeviceCertificateState) SetCertificateValidityPeriod(value *int32)() {
    err := m.GetBackingStore().Set("certificateValidityPeriod", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriodUnits sets the certificateValidityPeriodUnits property value. Certificate Validity Period Options.
func (m *ManagedDeviceCertificateState) SetCertificateValidityPeriodUnits(value *CertificateValidityPeriodScale)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodUnits", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Device display name
func (m *ManagedDeviceCertificateState) SetDeviceDisplayName(value *string)() {
    err := m.GetBackingStore().Set("deviceDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDevicePlatform sets the devicePlatform property value. Supported platform types.
func (m *ManagedDeviceCertificateState) SetDevicePlatform(value *DevicePlatformType)() {
    err := m.GetBackingStore().Set("devicePlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetLastCertificateStateChangeDateTime sets the lastCertificateStateChangeDateTime property value. Last certificate issuance state change
func (m *ManagedDeviceCertificateState) SetLastCertificateStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCertificateStateChangeDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. User display name
func (m *ManagedDeviceCertificateState) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
type ManagedDeviceCertificateStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateEnhancedKeyUsage()(*string)
    GetCertificateErrorCode()(*int32)
    GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateIssuanceState()(*CertificateIssuanceStates)
    GetCertificateIssuer()(*string)
    GetCertificateKeyLength()(*int32)
    GetCertificateKeyStorageProvider()(*KeyStorageProviderOption)
    GetCertificateKeyUsage()(*KeyUsages)
    GetCertificateLastIssuanceStateChangedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateProfileDisplayName()(*string)
    GetCertificateRevokeStatus()(*CertificateRevocationStatus)
    GetCertificateSerialNumber()(*string)
    GetCertificateSubjectAlternativeNameFormat()(*SubjectAlternativeNameType)
    GetCertificateSubjectAlternativeNameFormatString()(*string)
    GetCertificateSubjectNameFormat()(*SubjectNameFormat)
    GetCertificateSubjectNameFormatString()(*string)
    GetCertificateThumbprint()(*string)
    GetCertificateValidityPeriod()(*int32)
    GetCertificateValidityPeriodUnits()(*CertificateValidityPeriodScale)
    GetDeviceDisplayName()(*string)
    GetDevicePlatform()(*DevicePlatformType)
    GetLastCertificateStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUserDisplayName()(*string)
    SetCertificateEnhancedKeyUsage(value *string)()
    SetCertificateErrorCode(value *int32)()
    SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateIssuanceState(value *CertificateIssuanceStates)()
    SetCertificateIssuer(value *string)()
    SetCertificateKeyLength(value *int32)()
    SetCertificateKeyStorageProvider(value *KeyStorageProviderOption)()
    SetCertificateKeyUsage(value *KeyUsages)()
    SetCertificateLastIssuanceStateChangedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateProfileDisplayName(value *string)()
    SetCertificateRevokeStatus(value *CertificateRevocationStatus)()
    SetCertificateSerialNumber(value *string)()
    SetCertificateSubjectAlternativeNameFormat(value *SubjectAlternativeNameType)()
    SetCertificateSubjectAlternativeNameFormatString(value *string)()
    SetCertificateSubjectNameFormat(value *SubjectNameFormat)()
    SetCertificateSubjectNameFormatString(value *string)()
    SetCertificateThumbprint(value *string)()
    SetCertificateValidityPeriod(value *int32)()
    SetCertificateValidityPeriodUnits(value *CertificateValidityPeriodScale)()
    SetDeviceDisplayName(value *string)()
    SetDevicePlatform(value *DevicePlatformType)()
    SetLastCertificateStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUserDisplayName(value *string)()
}
