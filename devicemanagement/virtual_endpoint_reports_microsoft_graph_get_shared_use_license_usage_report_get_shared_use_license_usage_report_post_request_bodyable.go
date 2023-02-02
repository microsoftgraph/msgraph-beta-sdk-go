package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportPostRequestBodyable 
type VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilter()(*string)
    GetGroupBy()([]string)
    GetOrderBy()([]string)
    GetReportName()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName)
    GetSearch()(*string)
    GetSelect()([]string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetFilter(value *string)()
    SetGroupBy(value []string)()
    SetOrderBy(value []string)()
    SetReportName(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName)()
    SetSearch(value *string)()
    SetSelect(value []string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
