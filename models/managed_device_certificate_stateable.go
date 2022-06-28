package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceCertificateStateable 
type ManagedDeviceCertificateStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
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
