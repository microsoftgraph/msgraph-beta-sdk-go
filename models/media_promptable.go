package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaPromptable 
type MediaPromptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Promptable
    GetLoop()(*int32)
    GetMediaInfo()(MediaInfoable)
    SetLoop(value *int32)()
    SetMediaInfo(value MediaInfoable)()
}
