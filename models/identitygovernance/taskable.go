package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Taskable 
type Taskable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArguments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable)
    GetCategory()(*LifecycleTaskCategory)
    GetContinueOnError()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExecutionSequence()(*int32)
    GetIsEnabled()(*bool)
    GetTaskDefinitionId()(*string)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    SetArguments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable)()
    SetCategory(value *LifecycleTaskCategory)()
    SetContinueOnError(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExecutionSequence(value *int32)()
    SetIsEnabled(value *bool)()
    SetTaskDefinitionId(value *string)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
}
