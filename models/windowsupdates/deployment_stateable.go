package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentStateable 
type DeploymentStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEffectiveValue()(*DeploymentStateValue)
    GetOdataType()(*string)
    GetReasons()([]DeploymentStateReasonable)
    GetRequestedValue()(*RequestedDeploymentStateValue)
    SetEffectiveValue(value *DeploymentStateValue)()
    SetOdataType(value *string)()
    SetReasons(value []DeploymentStateReasonable)()
    SetRequestedValue(value *RequestedDeploymentStateValue)()
}
