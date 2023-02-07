package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApprovalRuleable 
type ContentApprovalRuleable interface {
    ComplianceChangeRuleable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentFilter()(ContentFilterable)
    GetDurationBeforeDeploymentStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetContentFilter(value ContentFilterable)()
    SetDurationBeforeDeploymentStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}
