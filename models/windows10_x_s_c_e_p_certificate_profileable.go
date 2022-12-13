package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XSCEPCertificateProfileable 
type Windows10XSCEPCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Windows10XCertificateProfileable
    GetCertificateStore()(*CertificateStore)
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetHashAlgorithm()([]HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyStorageProvider()(*KeyStorageProviderOption)
    GetKeyUsage()(*KeyUsages)
    GetRenewalThresholdPercentage()(*int32)
    GetRootCertificateId()(*UUID)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormats()([]Windows10XCustomSubjectAlternativeNameable)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *CertificateStore)()
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetHashAlgorithm(value []HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyStorageProvider(value *KeyStorageProviderOption)()
    SetKeyUsage(value *KeyUsages)()
    SetRenewalThresholdPercentage(value *int32)()
    SetRootCertificateId(value *UUID)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormats(value []Windows10XCustomSubjectAlternativeNameable)()
    SetSubjectNameFormatString(value *string)()
}
