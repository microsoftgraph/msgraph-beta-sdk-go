package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// LifecycleWorkflowsContainerable 
type LifecycleWorkflowsContainerable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomTaskExtensions()([]CustomTaskExtensionable)
    GetDeletedItems()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable)
    GetSettings()(LifecycleManagementSettingsable)
    GetTaskDefinitions()([]TaskDefinitionable)
    GetWorkflows()([]Workflowable)
    GetWorkflowTemplates()([]WorkflowTemplateable)
    SetCustomTaskExtensions(value []CustomTaskExtensionable)()
    SetDeletedItems(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable)()
    SetSettings(value LifecycleManagementSettingsable)()
    SetTaskDefinitions(value []TaskDefinitionable)()
    SetWorkflows(value []Workflowable)()
    SetWorkflowTemplates(value []WorkflowTemplateable)()
}
