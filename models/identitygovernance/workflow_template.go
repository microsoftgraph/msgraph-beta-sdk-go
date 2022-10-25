package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// WorkflowTemplate provides operations to manage the collection of accessReview entities.
type WorkflowTemplate struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The category property
    category *LifecycleWorkflowCategory
    // The description of the workflowTemplate.
    description *string
    // The display name of the workflowTemplate. Supports  orderby.
    displayName *string
    // Conditions describing when to execute the workflow and the criteria to identify in-scope subject set.
    executionConditions WorkflowExecutionConditionsable
    // Represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
    tasks []Taskable
}
// NewWorkflowTemplate instantiates a new workflowTemplate and sets the default values.
func NewWorkflowTemplate()(*WorkflowTemplate) {
    m := &WorkflowTemplate{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.workflowTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkflowTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowTemplate(), nil
}
// GetCategory gets the category property value. The category property
func (m *WorkflowTemplate) GetCategory()(*LifecycleWorkflowCategory) {
    return m.category
}
// GetDescription gets the description property value. The description of the workflowTemplate.
func (m *WorkflowTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the workflowTemplate. Supports  orderby.
func (m *WorkflowTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetExecutionConditions gets the executionConditions property value. Conditions describing when to execute the workflow and the criteria to identify in-scope subject set.
func (m *WorkflowTemplate) GetExecutionConditions()(WorkflowExecutionConditionsable) {
    return m.executionConditions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkflowTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLifecycleWorkflowCategory , m.SetCategory)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["executionConditions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWorkflowExecutionConditionsFromDiscriminatorValue , m.SetExecutionConditions)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaskFromDiscriminatorValue , m.SetTasks)
    return res
}
// GetTasks gets the tasks property value. Represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
func (m *WorkflowTemplate) GetTasks()([]Taskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *WorkflowTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("executionConditions", m.GetExecutionConditions())
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTasks())
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category property
func (m *WorkflowTemplate) SetCategory(value *LifecycleWorkflowCategory)() {
    m.category = value
}
// SetDescription sets the description property value. The description of the workflowTemplate.
func (m *WorkflowTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the workflowTemplate. Supports  orderby.
func (m *WorkflowTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExecutionConditions sets the executionConditions property value. Conditions describing when to execute the workflow and the criteria to identify in-scope subject set.
func (m *WorkflowTemplate) SetExecutionConditions(value WorkflowExecutionConditionsable)() {
    m.executionConditions = value
}
// SetTasks sets the tasks property value. Represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
func (m *WorkflowTemplate) SetTasks(value []Taskable)() {
    m.tasks = value
}
