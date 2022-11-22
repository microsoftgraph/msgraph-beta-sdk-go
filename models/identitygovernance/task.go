package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Task provides operations to manage the collection of activityStatistics entities.
type Task struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Arguments included within the task.  For guidance to configure this property, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.
    arguments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable
    // The category property
    category *LifecycleTaskCategory
    // A boolean value that determines if the failure of this task stops the subsequent workflows from running. Optional.
    continueOnError *bool
    // A string that describes the purpose of the task for administrative use. Optional.
    description *string
    // A unique string that identifies the task. Required.Supports $filter(eq, ne) and orderBy.
    displayName *string
    // An integer that states in what order the task will run in a workflow.Supports $orderby.
    executionSequence *int32
    // A boolean value that denotes whether the task is set to run or not. Optional.Supports $filter(eq, ne) and orderBy.
    isEnabled *bool
    // A unique template identifier for the task. For more information about the tasks that Lifecycle Workflows currently supports and their unique identifiers, see supported tasks. Required.Supports $filter(eq, ne).
    taskDefinitionId *string
    // The result of processing the task.
    taskProcessingResults []TaskProcessingResultable
}
// NewTask instantiates a new task and sets the default values.
func NewTask()(*Task) {
    m := &Task{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTask(), nil
}
// GetArguments gets the arguments property value. Arguments included within the task.  For guidance to configure this property, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.
func (m *Task) GetArguments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable) {
    return m.arguments
}
// GetCategory gets the category property value. The category property
func (m *Task) GetCategory()(*LifecycleTaskCategory) {
    return m.category
}
// GetContinueOnError gets the continueOnError property value. A boolean value that determines if the failure of this task stops the subsequent workflows from running. Optional.
func (m *Task) GetContinueOnError()(*bool) {
    return m.continueOnError
}
// GetDescription gets the description property value. A string that describes the purpose of the task for administrative use. Optional.
func (m *Task) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. A unique string that identifies the task. Required.Supports $filter(eq, ne) and orderBy.
func (m *Task) GetDisplayName()(*string) {
    return m.displayName
}
// GetExecutionSequence gets the executionSequence property value. An integer that states in what order the task will run in a workflow.Supports $orderby.
func (m *Task) GetExecutionSequence()(*int32) {
    return m.executionSequence
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Task) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["arguments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateKeyValuePairFromDiscriminatorValue , m.SetArguments)
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLifecycleTaskCategory , m.SetCategory)
    res["continueOnError"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContinueOnError)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["executionSequence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetExecutionSequence)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["taskDefinitionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTaskDefinitionId)
    res["taskProcessingResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue , m.SetTaskProcessingResults)
    return res
}
// GetIsEnabled gets the isEnabled property value. A boolean value that denotes whether the task is set to run or not. Optional.Supports $filter(eq, ne) and orderBy.
func (m *Task) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetTaskDefinitionId gets the taskDefinitionId property value. A unique template identifier for the task. For more information about the tasks that Lifecycle Workflows currently supports and their unique identifiers, see supported tasks. Required.Supports $filter(eq, ne).
func (m *Task) GetTaskDefinitionId()(*string) {
    return m.taskDefinitionId
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The result of processing the task.
func (m *Task) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// Serialize serializes information the current object
func (m *Task) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetArguments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetArguments())
        err = writer.WriteCollectionOfObjectValues("arguments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("continueOnError", m.GetContinueOnError())
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
        err = writer.WriteInt32Value("executionSequence", m.GetExecutionSequence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taskDefinitionId", m.GetTaskDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetTaskProcessingResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskProcessingResults())
        err = writer.WriteCollectionOfObjectValues("taskProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArguments sets the arguments property value. Arguments included within the task.  For guidance to configure this property, see Configure the arguments for built-in Lifecycle Workflow tasks. Required.
func (m *Task) SetArguments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable)() {
    m.arguments = value
}
// SetCategory sets the category property value. The category property
func (m *Task) SetCategory(value *LifecycleTaskCategory)() {
    m.category = value
}
// SetContinueOnError sets the continueOnError property value. A boolean value that determines if the failure of this task stops the subsequent workflows from running. Optional.
func (m *Task) SetContinueOnError(value *bool)() {
    m.continueOnError = value
}
// SetDescription sets the description property value. A string that describes the purpose of the task for administrative use. Optional.
func (m *Task) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. A unique string that identifies the task. Required.Supports $filter(eq, ne) and orderBy.
func (m *Task) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExecutionSequence sets the executionSequence property value. An integer that states in what order the task will run in a workflow.Supports $orderby.
func (m *Task) SetExecutionSequence(value *int32)() {
    m.executionSequence = value
}
// SetIsEnabled sets the isEnabled property value. A boolean value that denotes whether the task is set to run or not. Optional.Supports $filter(eq, ne) and orderBy.
func (m *Task) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetTaskDefinitionId sets the taskDefinitionId property value. A unique template identifier for the task. For more information about the tasks that Lifecycle Workflows currently supports and their unique identifiers, see supported tasks. Required.Supports $filter(eq, ne).
func (m *Task) SetTaskDefinitionId(value *string)() {
    m.taskDefinitionId = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The result of processing the task.
func (m *Task) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
