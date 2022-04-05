package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantDetailedInformationable 
type TenantDetailedInformationable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryCode()(*string)
    GetCountryName()(*string)
    GetDefaultDomainName()(*string)
    GetDisplayName()(*string)
    GetIndustryName()(*string)
    GetRegion()(*string)
    GetSegmentName()(*string)
    GetTenantId()(*string)
    GetVerticalName()(*string)
    SetCity(value *string)()
    SetCountryCode(value *string)()
    SetCountryName(value *string)()
    SetDefaultDomainName(value *string)()
    SetDisplayName(value *string)()
    SetIndustryName(value *string)()
    SetRegion(value *string)()
    SetSegmentName(value *string)()
    SetTenantId(value *string)()
    SetVerticalName(value *string)()
}
