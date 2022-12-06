package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// LifecycleWorkflowsContainer 
type LifecycleWorkflowsContainer struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The customTaskExtension instance.
    customTaskExtensions []CustomTaskExtensionable
    // Deleted workflows in your lifecycle workflows instance.
    deletedItems ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable
    // The settings property
    settings LifecycleManagementSettingsable
    // The definition of tasks within the lifecycle workflows instance.
    taskDefinitions []TaskDefinitionable
    // The workflows in the lifecycle workflows instance.
    workflows []Workflowable
    // The workflow templates in the lifecycle workflow instance.
    workflowTemplates []WorkflowTemplateable
}
// NewLifecycleWorkflowsContainer instantiates a new LifecycleWorkflowsContainer and sets the default values.
func NewLifecycleWorkflowsContainer()(*LifecycleWorkflowsContainer) {
    m := &LifecycleWorkflowsContainer{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateLifecycleWorkflowsContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleWorkflowsContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleWorkflowsContainer(), nil
}
// GetCustomTaskExtensions gets the customTaskExtensions property value. The customTaskExtension instance.
func (m *LifecycleWorkflowsContainer) GetCustomTaskExtensions()([]CustomTaskExtensionable) {
    return m.customTaskExtensions
}
// GetDeletedItems gets the deletedItems property value. Deleted workflows in your lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) GetDeletedItems()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable) {
    return m.deletedItems
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LifecycleWorkflowsContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customTaskExtensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomTaskExtensionFromDiscriminatorValue , m.SetCustomTaskExtensions)
    res["deletedItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeletedItemContainerFromDiscriminatorValue , m.SetDeletedItems)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLifecycleManagementSettingsFromDiscriminatorValue , m.SetSettings)
    res["taskDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaskDefinitionFromDiscriminatorValue , m.SetTaskDefinitions)
    res["workflows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkflowFromDiscriminatorValue , m.SetWorkflows)
    res["workflowTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkflowTemplateFromDiscriminatorValue , m.SetWorkflowTemplates)
    return res
}
// GetSettings gets the settings property value. The settings property
func (m *LifecycleWorkflowsContainer) GetSettings()(LifecycleManagementSettingsable) {
    return m.settings
}
// GetTaskDefinitions gets the taskDefinitions property value. The definition of tasks within the lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) GetTaskDefinitions()([]TaskDefinitionable) {
    return m.taskDefinitions
}
// GetWorkflows gets the workflows property value. The workflows in the lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) GetWorkflows()([]Workflowable) {
    return m.workflows
}
// GetWorkflowTemplates gets the workflowTemplates property value. The workflow templates in the lifecycle workflow instance.
func (m *LifecycleWorkflowsContainer) GetWorkflowTemplates()([]WorkflowTemplateable) {
    return m.workflowTemplates
}
// Serialize serializes information the current object
func (m *LifecycleWorkflowsContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomTaskExtensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomTaskExtensions())
        err = writer.WriteCollectionOfObjectValues("customTaskExtensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deletedItems", m.GetDeletedItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    if m.GetTaskDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskDefinitions())
        err = writer.WriteCollectionOfObjectValues("taskDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflows() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWorkflows())
        err = writer.WriteCollectionOfObjectValues("workflows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflowTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWorkflowTemplates())
        err = writer.WriteCollectionOfObjectValues("workflowTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomTaskExtensions sets the customTaskExtensions property value. The customTaskExtension instance.
func (m *LifecycleWorkflowsContainer) SetCustomTaskExtensions(value []CustomTaskExtensionable)() {
    m.customTaskExtensions = value
}
// SetDeletedItems sets the deletedItems property value. Deleted workflows in your lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) SetDeletedItems(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable)() {
    m.deletedItems = value
}
// SetSettings sets the settings property value. The settings property
func (m *LifecycleWorkflowsContainer) SetSettings(value LifecycleManagementSettingsable)() {
    m.settings = value
}
// SetTaskDefinitions sets the taskDefinitions property value. The definition of tasks within the lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) SetTaskDefinitions(value []TaskDefinitionable)() {
    m.taskDefinitions = value
}
// SetWorkflows sets the workflows property value. The workflows in the lifecycle workflows instance.
func (m *LifecycleWorkflowsContainer) SetWorkflows(value []Workflowable)() {
    m.workflows = value
}
// SetWorkflowTemplates sets the workflowTemplates property value. The workflow templates in the lifecycle workflow instance.
func (m *LifecycleWorkflowsContainer) SetWorkflowTemplates(value []WorkflowTemplateable)() {
    m.workflowTemplates = value
}
