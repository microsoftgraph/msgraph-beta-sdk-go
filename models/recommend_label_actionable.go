package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendLabelActionable 
type RecommendLabelActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]InformationProtectionActionable)
    GetActionSource()(*ActionSource)
    GetLabel()(LabelDetailsable)
    GetResponsibleSensitiveTypeIds()([]string)
    SetActions(value []InformationProtectionActionable)()
    SetActionSource(value *ActionSource)()
    SetLabel(value LabelDetailsable)()
    SetResponsibleSensitiveTypeIds(value []string)()
}
