package app

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CallsItemMicrosoftGraphRecordResponseRecordResponsePostRequestBodyable 
type CallsItemMicrosoftGraphRecordResponseRecordResponsePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBargeInAllowed()(*bool)
    GetClientContext()(*string)
    GetInitialSilenceTimeoutInSeconds()(*int32)
    GetMaxRecordDurationInSeconds()(*int32)
    GetMaxSilenceTimeoutInSeconds()(*int32)
    GetPlayBeep()(*bool)
    GetPrompts()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Promptable)
    GetStopTones()([]string)
    GetStreamWhileRecording()(*bool)
    SetBargeInAllowed(value *bool)()
    SetClientContext(value *string)()
    SetInitialSilenceTimeoutInSeconds(value *int32)()
    SetMaxRecordDurationInSeconds(value *int32)()
    SetMaxSilenceTimeoutInSeconds(value *int32)()
    SetPlayBeep(value *bool)()
    SetPrompts(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Promptable)()
    SetStopTones(value []string)()
    SetStreamWhileRecording(value *bool)()
}
