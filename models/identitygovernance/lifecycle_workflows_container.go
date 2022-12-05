package identitygovernance

import (
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
    res["customTaskExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomTaskExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomTaskExtensionable, len(val))
            for i, v := range val {
                res[i] = v.(CustomTaskExtensionable)
            }
            m.SetCustomTaskExtensions(res)
        }
        return nil
    }
    res["deletedItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeletedItemContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedItems(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLifecycleManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(LifecycleManagementSettingsable))
        }
        return nil
    }
    res["taskDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(TaskDefinitionable)
            }
            m.SetTaskDefinitions(res)
        }
        return nil
    }
    res["workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkflowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Workflowable, len(val))
            for i, v := range val {
                res[i] = v.(Workflowable)
            }
            m.SetWorkflows(res)
        }
        return nil
    }
    res["workflowTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkflowTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkflowTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(WorkflowTemplateable)
            }
            m.SetWorkflowTemplates(res)
        }
        return nil
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomTaskExtensions()))
        for i, v := range m.GetCustomTaskExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskDefinitions()))
        for i, v := range m.GetTaskDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkflows()))
        for i, v := range m.GetWorkflows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("workflows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflowTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkflowTemplates()))
        for i, v := range m.GetWorkflowTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
