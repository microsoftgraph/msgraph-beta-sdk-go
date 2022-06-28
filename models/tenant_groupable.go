package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantGroupable 
type TenantGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllTenantsIncluded()(*bool)
    GetDisplayName()(*string)
    GetTenantIds()([]string)
    SetAllTenantsIncluded(value *bool)()
    SetDisplayName(value *string)()
    SetTenantIds(value []string)()
}
