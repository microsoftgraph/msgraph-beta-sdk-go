package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedAllDeviceCertificateState struct {
    Entity
    certificateExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    certificateExtendedKeyUsages *string;
    certificateIssuanceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    certificateIssuerName *string;
    certificateKeyUsages *int32;
    certificateRevokeStatus *CertificateRevocationStatus;
    certificateRevokeStatusLastChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    certificateSerialNumber *string;
    certificateSubjectName *string;
    certificateThumbprint *string;
    managedDeviceDisplayName *string;
    userPrincipalName *string;
}
func NewManagedAllDeviceCertificateState()(*ManagedAllDeviceCertificateState) {
    m := &ManagedAllDeviceCertificateState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ManagedAllDeviceCertificateState) GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certificateExpirationDateTime
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateExtendedKeyUsages()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateExtendedKeyUsages
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certificateIssuanceDateTime
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateIssuerName
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateKeyUsages()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.certificateKeyUsages
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatus()(*CertificateRevocationStatus) {
    if m == nil {
        return nil
    } else {
        return m.certificateRevokeStatus
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatusLastChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certificateRevokeStatusLastChangeDateTime
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateSerialNumber
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateSubjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateSubjectName
    }
}
func (m *ManagedAllDeviceCertificateState) GetCertificateThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateThumbprint
    }
}
func (m *ManagedAllDeviceCertificateState) GetManagedDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceDisplayName
    }
}
func (m *ManagedAllDeviceCertificateState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
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
func (m *ManagedAllDeviceCertificateState) SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateExpirationDateTime = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateExtendedKeyUsages(value *string)() {
    m.certificateExtendedKeyUsages = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateIssuanceDateTime = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuerName(value *string)() {
    m.certificateIssuerName = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateKeyUsages(value *int32)() {
    m.certificateKeyUsages = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatus(value *CertificateRevocationStatus)() {
    m.certificateRevokeStatus = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatusLastChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateRevokeStatusLastChangeDateTime = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateSerialNumber(value *string)() {
    m.certificateSerialNumber = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateSubjectName(value *string)() {
    m.certificateSubjectName = value
}
func (m *ManagedAllDeviceCertificateState) SetCertificateThumbprint(value *string)() {
    m.certificateThumbprint = value
}
func (m *ManagedAllDeviceCertificateState) SetManagedDeviceDisplayName(value *string)() {
    m.managedDeviceDisplayName = value
}
func (m *ManagedAllDeviceCertificateState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
