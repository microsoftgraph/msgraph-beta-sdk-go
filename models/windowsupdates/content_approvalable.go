package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApprovalable 
type ContentApprovalable interface {
    ComplianceChangeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(DeployableContentable)
    GetDeployments()([]Deploymentable)
    GetDeploymentSettings()(DeploymentSettingsable)
    SetContent(value DeployableContentable)()
    SetDeployments(value []Deploymentable)()
    SetDeploymentSettings(value DeploymentSettingsable)()
}
