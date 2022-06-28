package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkPkcsCertificateProfileable 
type AndroidForWorkPkcsCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    AndroidForWorkCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateTemplateName()(*string)
    GetCertificationAuthority()(*string)
    GetCertificationAuthorityName()(*string)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetSubjectAlternativeNameFormatString()(*string)
    SetCertificateTemplateName(value *string)()
    SetCertificationAuthority(value *string)()
    SetCertificationAuthorityName(value *string)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetSubjectAlternativeNameFormatString(value *string)()
}
