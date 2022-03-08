package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAllDeviceCertificateStateable 
type ManagedAllDeviceCertificateStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
