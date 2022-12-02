package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsItemRedirectPostRequestBodyable 
type AppCallsItemRedirectPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallbackUri()(*string)
    GetMaskCallee()(*bool)
    GetMaskCaller()(*bool)
    GetTargetDisposition()(*CallDisposition)
    GetTargets()([]InvitationParticipantInfoable)
    GetTimeout()(*int32)
    SetCallbackUri(value *string)()
    SetMaskCallee(value *bool)()
    SetMaskCaller(value *bool)()
    SetTargetDisposition(value *CallDisposition)()
    SetTargets(value []InvitationParticipantInfoable)()
    SetTimeout(value *int32)()
}
