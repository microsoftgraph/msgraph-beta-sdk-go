package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterStatusable 
type PrinterStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDetails()([]string)
    GetProcessingState()(*PrinterProcessingState)
    GetProcessingStateDescription()(*string)
    GetProcessingStateReasons()([]string)
    GetState()(*PrinterProcessingState)
    SetDescription(value *string)()
    SetDetails(value []string)()
    SetProcessingState(value *PrinterProcessingState)()
    SetProcessingStateDescription(value *string)()
    SetProcessingStateReasons(value []string)()
    SetState(value *PrinterProcessingState)()
}
