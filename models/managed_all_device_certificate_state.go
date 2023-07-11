package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAllDeviceCertificateState 
type ManagedAllDeviceCertificateState struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewManagedAllDeviceCertificateState instantiates a new managedAllDeviceCertificateState and sets the default values.
func NewManagedAllDeviceCertificateState()(*ManagedAllDeviceCertificateState) {
    m := &ManagedAllDeviceCertificateState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedAllDeviceCertificateStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAllDeviceCertificateStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAllDeviceCertificateState(), nil
}
// GetCertificateExpirationDateTime gets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedAllDeviceCertificateState) GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("certificateExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCertificateExtendedKeyUsages gets the certificateExtendedKeyUsages property value. Enhanced Key Usage
func (m *ManagedAllDeviceCertificateState) GetCertificateExtendedKeyUsages()(*string) {
    val, err := m.GetBackingStore().Get("certificateExtendedKeyUsages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateIssuanceDateTime gets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("certificateIssuanceDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCertificateIssuerName gets the certificateIssuerName property value. Issuer
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuerName()(*string) {
    val, err := m.GetBackingStore().Get("certificateIssuerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateKeyUsages gets the certificateKeyUsages property value. Key Usage
func (m *ManagedAllDeviceCertificateState) GetCertificateKeyUsages()(*int32) {
    val, err := m.GetBackingStore().Get("certificateKeyUsages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCertificateRevokeStatus gets the certificateRevokeStatus property value. Certificate Revocation Status.
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatus()(*CertificateRevocationStatus) {
    val, err := m.GetBackingStore().Get("certificateRevokeStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateRevocationStatus)
    }
    return nil
}
// GetCertificateRevokeStatusLastChangeDateTime gets the certificateRevokeStatusLastChangeDateTime property value. The time the revoke status was last changed
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatusLastChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("certificateRevokeStatusLastChangeDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCertificateSerialNumber gets the certificateSerialNumber property value. Serial number
func (m *ManagedAllDeviceCertificateState) GetCertificateSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("certificateSerialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateSubjectName gets the certificateSubjectName property value. Certificate subject name
func (m *ManagedAllDeviceCertificateState) GetCertificateSubjectName()(*string) {
    val, err := m.GetBackingStore().Get("certificateSubjectName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateThumbprint gets the certificateThumbprint property value. Thumbprint
func (m *ManagedAllDeviceCertificateState) GetCertificateThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("certificateThumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAllDeviceCertificateState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["certificateExtendedKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateExtendedKeyUsages(val)
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
    res["certificateIssuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuerName(val)
        }
        return nil
    }
    res["certificateKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyUsages(val)
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
    res["certificateRevokeStatusLastChangeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevokeStatusLastChangeDateTime(val)
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
    res["certificateSubjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectName(val)
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
    res["managedDeviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceDisplayName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceDisplayName gets the managedDeviceDisplayName property value. Device display name
func (m *ManagedAllDeviceCertificateState) GetManagedDeviceDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name
func (m *ManagedAllDeviceCertificateState) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedAllDeviceCertificateState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("certificateExpirationDateTime", m.GetCertificateExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateExtendedKeyUsages", m.GetCertificateExtendedKeyUsages())
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
    {
        err = writer.WriteStringValue("certificateIssuerName", m.GetCertificateIssuerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateKeyUsages", m.GetCertificateKeyUsages())
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
        err = writer.WriteTimeValue("certificateRevokeStatusLastChangeDateTime", m.GetCertificateRevokeStatusLastChangeDateTime())
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
    {
        err = writer.WriteStringValue("certificateSubjectName", m.GetCertificateSubjectName())
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
        err = writer.WriteStringValue("managedDeviceDisplayName", m.GetManagedDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateExpirationDateTime sets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedAllDeviceCertificateState) SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("certificateExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateExtendedKeyUsages sets the certificateExtendedKeyUsages property value. Enhanced Key Usage
func (m *ManagedAllDeviceCertificateState) SetCertificateExtendedKeyUsages(value *string)() {
    err := m.GetBackingStore().Set("certificateExtendedKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateIssuanceDateTime sets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("certificateIssuanceDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateIssuerName sets the certificateIssuerName property value. Issuer
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuerName(value *string)() {
    err := m.GetBackingStore().Set("certificateIssuerName", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateKeyUsages sets the certificateKeyUsages property value. Key Usage
func (m *ManagedAllDeviceCertificateState) SetCertificateKeyUsages(value *int32)() {
    err := m.GetBackingStore().Set("certificateKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateRevokeStatus sets the certificateRevokeStatus property value. Certificate Revocation Status.
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatus(value *CertificateRevocationStatus)() {
    err := m.GetBackingStore().Set("certificateRevokeStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateRevokeStatusLastChangeDateTime sets the certificateRevokeStatusLastChangeDateTime property value. The time the revoke status was last changed
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatusLastChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("certificateRevokeStatusLastChangeDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSerialNumber sets the certificateSerialNumber property value. Serial number
func (m *ManagedAllDeviceCertificateState) SetCertificateSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("certificateSerialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSubjectName sets the certificateSubjectName property value. Certificate subject name
func (m *ManagedAllDeviceCertificateState) SetCertificateSubjectName(value *string)() {
    err := m.GetBackingStore().Set("certificateSubjectName", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateThumbprint sets the certificateThumbprint property value. Thumbprint
func (m *ManagedAllDeviceCertificateState) SetCertificateThumbprint(value *string)() {
    err := m.GetBackingStore().Set("certificateThumbprint", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceDisplayName sets the managedDeviceDisplayName property value. Device display name
func (m *ManagedAllDeviceCertificateState) SetManagedDeviceDisplayName(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name
func (m *ManagedAllDeviceCertificateState) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// ManagedAllDeviceCertificateStateable 
type ManagedAllDeviceCertificateStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateExtendedKeyUsages()(*string)
    GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateIssuerName()(*string)
    GetCertificateKeyUsages()(*int32)
    GetCertificateRevokeStatus()(*CertificateRevocationStatus)
    GetCertificateRevokeStatusLastChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateSerialNumber()(*string)
    GetCertificateSubjectName()(*string)
    GetCertificateThumbprint()(*string)
    GetManagedDeviceDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateExtendedKeyUsages(value *string)()
    SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateIssuerName(value *string)()
    SetCertificateKeyUsages(value *int32)()
    SetCertificateRevokeStatus(value *CertificateRevocationStatus)()
    SetCertificateRevokeStatusLastChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateSerialNumber(value *string)()
    SetCertificateSubjectName(value *string)()
    SetCertificateThumbprint(value *string)()
    SetManagedDeviceDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
