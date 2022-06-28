package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppDependencyable 
type MobileAppDependencyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    MobileAppRelationshipable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDependencyType()(*MobileAppDependencyType)
    GetDependentAppCount()(*int32)
    GetDependsOnAppCount()(*int32)
    SetDependencyType(value *MobileAppDependencyType)()
    SetDependentAppCount(value *int32)()
    SetDependsOnAppCount(value *int32)()
}
