package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantGroupable 
type TenantGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllTenantsIncluded()(*bool)
    GetDisplayName()(*string)
    GetManagementActions()([]ManagementActionInfoable)
    GetManagementIntents()([]ManagementIntentInfoable)
    GetTenantIds()([]string)
    SetAllTenantsIncluded(value *bool)()
    SetDisplayName(value *string)()
    SetManagementActions(value []ManagementActionInfoable)()
    SetManagementIntents(value []ManagementIntentInfoable)()
    SetTenantIds(value []string)()
}
