package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionPolicySettingable 
type InformationProtectionPolicySettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultLabelId()(*string)
    GetIsDowngradeJustificationRequired()(*bool)
    GetIsMandatory()(*bool)
    GetMoreInfoUrl()(*string)
    SetDefaultLabelId(value *string)()
    SetIsDowngradeJustificationRequired(value *bool)()
    SetIsMandatory(value *bool)()
    SetMoreInfoUrl(value *string)()
}
