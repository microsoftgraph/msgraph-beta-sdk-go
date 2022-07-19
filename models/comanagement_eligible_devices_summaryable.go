package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagementEligibleDevicesSummaryable 
type ComanagementEligibleDevicesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComanagedCount()(*int32)
    GetEligibleButNotAzureAdJoinedCount()(*int32)
    GetEligibleCount()(*int32)
    GetIneligibleCount()(*int32)
    GetNeedsOsUpdateCount()(*int32)
    GetOdataType()(*string)
    SetComanagedCount(value *int32)()
    SetEligibleButNotAzureAdJoinedCount(value *int32)()
    SetEligibleCount(value *int32)()
    SetIneligibleCount(value *int32)()
    SetNeedsOsUpdateCount(value *int32)()
    SetOdataType(value *string)()
}
