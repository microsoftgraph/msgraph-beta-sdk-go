package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantInformationable 
type TenantInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultDomainName()(*string)
    GetDisplayName()(*string)
    GetFederationBrandName()(*string)
    GetOdataType()(*string)
    GetTenantId()(*string)
    SetDefaultDomainName(value *string)()
    SetDisplayName(value *string)()
    SetFederationBrandName(value *string)()
    SetOdataType(value *string)()
    SetTenantId(value *string)()
}
