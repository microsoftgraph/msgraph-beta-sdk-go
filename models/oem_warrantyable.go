package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OemWarrantyable 
type OemWarrantyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalWarranties()([]WarrantyOfferable)
    GetBaseWarranties()([]WarrantyOfferable)
    GetDeviceConfigurationUrl()(*string)
    GetDeviceWarrantyUrl()(*string)
    SetAdditionalWarranties(value []WarrantyOfferable)()
    SetBaseWarranties(value []WarrantyOfferable)()
    SetDeviceConfigurationUrl(value *string)()
    SetDeviceWarrantyUrl(value *string)()
}
