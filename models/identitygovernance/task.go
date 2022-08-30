package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Task provides operations to manage the collection of activityStatistics entities.
type Task struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The arguments property
    arguments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable
    // The category property
    category *LifecycleTaskCategory
    // The continueOnError property
    continueOnError *bool
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The executionSequence property
    executionSequence *int32
    // The isEnabled property
    isEnabled *bool
    // The taskDefinitionId property
    taskDefinitionId *string
    // The taskProcessingResults property
    taskProcessingResults []TaskProcessingResultable
}
// NewTask instantiates a new task and sets the default values.
func NewTask()(*Task) {
    m := &Task{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.task";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTask(), nil
}
// GetArguments gets the arguments property value. The arguments property
func (m *Task) GetArguments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable) {
    return m.arguments
}
// GetCategory gets the category property value. The category property
func (m *Task) GetCategory()(*LifecycleTaskCategory) {
    return m.category
}
// GetContinueOnError gets the continueOnError property value. The continueOnError property
func (m *Task) GetContinueOnError()(*bool) {
    return m.continueOnError
}
// GetDescription gets the description property value. The description property
func (m *Task) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Task) GetDisplayName()(*string) {
    return m.displayName
}
// GetExecutionSequence gets the executionSequence property value. The executionSequence property
func (m *Task) GetExecutionSequence()(*int32) {
    return m.executionSequence
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Task) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["arguments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable)
            }
            m.SetArguments(res)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLifecycleTaskCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*LifecycleTaskCategory))
        }
        return nil
    }
    res["continueOnError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContinueOnError(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["executionSequence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionSequence(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["taskDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskDefinitionId(val)
        }
        return nil
    }
    res["taskProcessingResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskProcessingResultable, len(val))
            for i, v := range val {
                res[i] = v.(TaskProcessingResultable)
            }
            m.SetTaskProcessingResults(res)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
func (m *Task) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetTaskDefinitionId gets the taskDefinitionId property value. The taskDefinitionId property
func (m *Task) GetTaskDefinitionId()(*string) {
    return m.taskDefinitionId
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The taskProcessingResults property
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArguments()))
        for i, v := range m.GetArguments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskProcessingResults()))
        for i, v := range m.GetTaskProcessingResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArguments sets the arguments property value. The arguments property
func (m *Task) SetArguments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.KeyValuePairable)() {
    m.arguments = value
}
// SetCategory sets the category property value. The category property
func (m *Task) SetCategory(value *LifecycleTaskCategory)() {
    m.category = value
}
// SetContinueOnError sets the continueOnError property value. The continueOnError property
func (m *Task) SetContinueOnError(value *bool)() {
    m.continueOnError = value
}
// SetDescription sets the description property value. The description property
func (m *Task) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Task) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExecutionSequence sets the executionSequence property value. The executionSequence property
func (m *Task) SetExecutionSequence(value *int32)() {
    m.executionSequence = value
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *Task) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetTaskDefinitionId sets the taskDefinitionId property value. The taskDefinitionId property
func (m *Task) SetTaskDefinitionId(value *string)() {
    m.taskDefinitionId = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The taskProcessingResults property
func (m *Task) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
