package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CallsItemMicrosoftGraphRedirectRedirectPostRequestBodyable 
type CallsItemMicrosoftGraphRedirectRedirectPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallbackUri()(*string)
    GetMaskCallee()(*bool)
    GetMaskCaller()(*bool)
    GetTargetDisposition()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition)
    GetTargets()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable)
    GetTimeout()(*int32)
    SetCallbackUri(value *string)()
    SetMaskCallee(value *bool)()
    SetMaskCaller(value *bool)()
    SetTargetDisposition(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition)()
    SetTargets(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable)()
    SetTimeout(value *int32)()
}
