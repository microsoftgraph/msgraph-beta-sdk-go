package getzebrafotadeploymentreport

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetZebraFotaDeploymentReportPostRequestBodyable 
type GetZebraFotaDeploymentReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilter()(*string)
    GetGroupBy()([]string)
    GetName()(*string)
    GetOrderBy()([]string)
    GetSearch()(*string)
    GetSelect()([]string)
    GetSessionId()(*int32)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetFilter(value *string)()
    SetGroupBy(value []string)()
    SetName(value *string)()
    SetOrderBy(value []string)()
    SetSearch(value *string)()
    SetSelect(value []string)()
    SetSessionId(value *int32)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
