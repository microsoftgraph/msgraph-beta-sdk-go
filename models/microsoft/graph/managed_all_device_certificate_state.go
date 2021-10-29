package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedAllDeviceCertificateState struct {
    Entity
    // Certificate expiry date
    certificateExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Enhanced Key Usage
    certificateExtendedKeyUsages *string;
    // Issuance date
    certificateIssuanceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Issuer
    certificateIssuerName *string;
    // Key Usage
    certificateKeyUsages *int32;
    // Revoke status. Possible values are: none, pending, issued, failed, revoked.
    certificateRevokeStatus *CertificateRevocationStatus;
    // The time the revoke status was last changed
    certificateRevokeStatusLastChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Serial number
    certificateSerialNumber *string;
    // Certificate subject name
    certificateSubjectName *string;
    // Thumbprint
    certificateThumbprint *string;
    // Device display name
    managedDeviceDisplayName *string;
    // User principal name
    userPrincipalName *string;
}
// Instantiates a new managedAllDeviceCertificateState and sets the default values.
func NewManagedAllDeviceCertificateState()(*ManagedAllDeviceCertificateState) {
    m := &ManagedAllDeviceCertificateState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedAllDeviceCertificateState) GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certificateExpirationDateTime
    }
}
// Gets the certificateExtendedKeyUsages property value. Enhanced Key Usage
func (m *ManagedAllDeviceCertificateState) GetCertificateExtendedKeyUsages()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateExtendedKeyUsages
    }
}
// Gets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certificateIssuanceDateTime
    }
}
// Gets the certificateIssuerName property value. Issuer
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateIssuerName
    }
}
// Gets the certificateKeyUsages property value. Key Usage
func (m *ManagedAllDeviceCertificateState) GetCertificateKeyUsages()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.certificateKeyUsages
    }
}
// Gets the certificateRevokeStatus property value. Revoke status. Possible values are: none, pending, issued, failed, revoked.
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatus()(*CertificateRevocationStatus) {
    if m == nil {
        return nil
    } else {
        return m.certificateRevokeStatus
    }
}
// Gets the certificateRevokeStatusLastChangeDateTime property value. The time the revoke status was last changed
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatusLastChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certificateRevokeStatusLastChangeDateTime
    }
}
// Gets the certificateSerialNumber property value. Serial number
func (m *ManagedAllDeviceCertificateState) GetCertificateSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateSerialNumber
    }
}
// Gets the certificateSubjectName property value. Certificate subject name
func (m *ManagedAllDeviceCertificateState) GetCertificateSubjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateSubjectName
    }
}
// Gets the certificateThumbprint property value. Thumbprint
func (m *ManagedAllDeviceCertificateState) GetCertificateThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateThumbprint
    }
}
// Gets the managedDeviceDisplayName property value. Device display name
func (m *ManagedAllDeviceCertificateState) GetManagedDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceDisplayName
    }
}
// Gets the userPrincipalName property value. User principal name
func (m *ManagedAllDeviceCertificateState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *ManagedAllDeviceCertificateState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCertificateExpirationDateTime(val)
        return nil
    }
    res["certificateExtendedKeyUsages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateExtendedKeyUsages(val)
        return nil
    }
    res["certificateIssuanceDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCertificateIssuanceDateTime(val)
        return nil
    }
    res["certificateIssuerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateIssuerName(val)
        return nil
    }
    res["certificateKeyUsages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCertificateKeyUsages(val)
        return nil
    }
    res["certificateRevokeStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateRevocationStatus)
        if err != nil {
            return err
        }
        cast := val.(CertificateRevocationStatus)
        m.SetCertificateRevokeStatus(&cast)
        return nil
    }
    res["certificateRevokeStatusLastChangeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCertificateRevokeStatusLastChangeDateTime(val)
        return nil
    }
    res["certificateSerialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateSerialNumber(val)
        return nil
    }
    res["certificateSubjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateSubjectName(val)
        return nil
    }
    res["certificateThumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificateThumbprint(val)
        return nil
    }
    res["managedDeviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceDisplayName(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *ManagedAllDeviceCertificateState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedAllDeviceCertificateState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetCertificateRevokeStatus().String()
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
// Sets the certificateExpirationDateTime property value. Certificate expiry date
// Parameters:
//  - value : Value to set for the certificateExpirationDateTime property.
func (m *ManagedAllDeviceCertificateState) SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateExpirationDateTime = value
}
// Sets the certificateExtendedKeyUsages property value. Enhanced Key Usage
// Parameters:
//  - value : Value to set for the certificateExtendedKeyUsages property.
func (m *ManagedAllDeviceCertificateState) SetCertificateExtendedKeyUsages(value *string)() {
    m.certificateExtendedKeyUsages = value
}
// Sets the certificateIssuanceDateTime property value. Issuance date
// Parameters:
//  - value : Value to set for the certificateIssuanceDateTime property.
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateIssuanceDateTime = value
}
// Sets the certificateIssuerName property value. Issuer
// Parameters:
//  - value : Value to set for the certificateIssuerName property.
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuerName(value *string)() {
    m.certificateIssuerName = value
}
// Sets the certificateKeyUsages property value. Key Usage
// Parameters:
//  - value : Value to set for the certificateKeyUsages property.
func (m *ManagedAllDeviceCertificateState) SetCertificateKeyUsages(value *int32)() {
    m.certificateKeyUsages = value
}
// Sets the certificateRevokeStatus property value. Revoke status. Possible values are: none, pending, issued, failed, revoked.
// Parameters:
//  - value : Value to set for the certificateRevokeStatus property.
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatus(value *CertificateRevocationStatus)() {
    m.certificateRevokeStatus = value
}
// Sets the certificateRevokeStatusLastChangeDateTime property value. The time the revoke status was last changed
// Parameters:
//  - value : Value to set for the certificateRevokeStatusLastChangeDateTime property.
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatusLastChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateRevokeStatusLastChangeDateTime = value
}
// Sets the certificateSerialNumber property value. Serial number
// Parameters:
//  - value : Value to set for the certificateSerialNumber property.
func (m *ManagedAllDeviceCertificateState) SetCertificateSerialNumber(value *string)() {
    m.certificateSerialNumber = value
}
// Sets the certificateSubjectName property value. Certificate subject name
// Parameters:
//  - value : Value to set for the certificateSubjectName property.
func (m *ManagedAllDeviceCertificateState) SetCertificateSubjectName(value *string)() {
    m.certificateSubjectName = value
}
// Sets the certificateThumbprint property value. Thumbprint
// Parameters:
//  - value : Value to set for the certificateThumbprint property.
func (m *ManagedAllDeviceCertificateState) SetCertificateThumbprint(value *string)() {
    m.certificateThumbprint = value
}
// Sets the managedDeviceDisplayName property value. Device display name
// Parameters:
//  - value : Value to set for the managedDeviceDisplayName property.
func (m *ManagedAllDeviceCertificateState) SetManagedDeviceDisplayName(value *string)() {
    m.managedDeviceDisplayName = value
}
// Sets the userPrincipalName property value. User principal name
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *ManagedAllDeviceCertificateState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
