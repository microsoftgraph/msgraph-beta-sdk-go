package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkCertificateProfileBaseable 
type AndroidForWorkCertificateProfileBaseable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetRenewalThresholdPercentage()(*int32)
    GetRootCertificate()(AndroidForWorkTrustedRootCertificateable)
    GetSubjectAlternativeNameType()(*SubjectAlternativeNameType)
    GetSubjectNameFormat()(*SubjectNameFormat)
    GetType()(*string)
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetRenewalThresholdPercentage(value *int32)()
    SetRootCertificate(value AndroidForWorkTrustedRootCertificateable)()
    SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)()
    SetSubjectNameFormat(value *SubjectNameFormat)()
    SetType(value *string)()
}
