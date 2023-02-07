package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphGetAssignmentFiltersStatusDetailsGetAssignmentFiltersStatusDetailsPostRequestBodyable 
type MicrosoftGraphGetAssignmentFiltersStatusDetailsGetAssignmentFiltersStatusDetailsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterIds()([]string)
    GetManagedDeviceId()(*string)
    GetPayloadId()(*string)
    GetSkip()(*int32)
    GetTop()(*int32)
    GetUserId()(*string)
    SetAssignmentFilterIds(value []string)()
    SetManagedDeviceId(value *string)()
    SetPayloadId(value *string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
    SetUserId(value *string)()
}
