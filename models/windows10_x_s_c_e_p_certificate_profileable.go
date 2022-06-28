package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XSCEPCertificateProfileable 
type Windows10XSCEPCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Windows10XCertificateProfileable
    GetCertificateStore()(*CertificateStore)
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetHashAlgorithm()([]string)
    GetKeySize()(*KeySize)
    GetKeyStorageProvider()(*KeyStorageProviderOption)
    GetKeyUsage()(*KeyUsages)
    GetRenewalThresholdPercentage()(*int32)
    GetRootCertificateId()(*string)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormats()([]Windows10XCustomSubjectAlternativeNameable)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *CertificateStore)()
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetHashAlgorithm(value []string)()
    SetKeySize(value *KeySize)()
    SetKeyStorageProvider(value *KeyStorageProviderOption)()
    SetKeyUsage(value *KeyUsages)()
    SetRenewalThresholdPercentage(value *int32)()
    SetRootCertificateId(value *string)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormats(value []Windows10XCustomSubjectAlternativeNameable)()
    SetSubjectNameFormatString(value *string)()
}
