package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorAuthenticationMethodTargetable 
type MicrosoftAuthenticatorAuthenticationMethodTargetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    AuthenticationMethodTargetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMode()(*MicrosoftAuthenticatorAuthenticationMode)
    GetDisplayAppInformationRequiredState()(*AdvancedConfigState)
    GetNumberMatchingRequiredState()(*AdvancedConfigState)
    SetAuthenticationMode(value *MicrosoftAuthenticatorAuthenticationMode)()
    SetDisplayAppInformationRequiredState(value *AdvancedConfigState)()
    SetNumberMatchingRequiredState(value *AdvancedConfigState)()
}
