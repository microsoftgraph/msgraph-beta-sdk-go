package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkEasEmailProfileBaseable 
type AndroidForWorkEasEmailProfileBaseable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*EasAuthenticationMethod)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEmailAddressSource()(*UserEmailSource)
    GetHostName()(*string)
    GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable)
    GetRequireSsl()(*bool)
    GetType()(*string)
    GetUsernameSource()(*AndroidUsernameSource)
    SetAuthenticationMethod(value *EasAuthenticationMethod)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetHostName(value *string)()
    SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)()
    SetRequireSsl(value *bool)()
    SetType(value *string)()
    SetUsernameSource(value *AndroidUsernameSource)()
}
