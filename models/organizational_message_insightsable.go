package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageInsightsable 
type OrganizationalMessageInsightsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClicks()(*int32)
    GetDismisses()(*int32)
    GetImpressions()(*int32)
    GetOdataType()(*string)
    SetClicks(value *int32)()
    SetDismisses(value *int32)()
    SetImpressions(value *int32)()
    SetOdataType(value *string)()
}
