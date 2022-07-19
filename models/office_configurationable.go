package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeConfigurationable 
type OfficeConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientConfigurations()([]OfficeClientConfigurationable)
    GetOdataType()(*string)
    GetTenantCheckinStatuses()([]OfficeClientCheckinStatusable)
    GetTenantUserCheckinSummary()(OfficeUserCheckinSummaryable)
    SetClientConfigurations(value []OfficeClientConfigurationable)()
    SetOdataType(value *string)()
    SetTenantCheckinStatuses(value []OfficeClientCheckinStatusable)()
    SetTenantUserCheckinSummary(value OfficeUserCheckinSummaryable)()
}
