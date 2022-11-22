package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TaskDefinition provides operations to manage the collection of activityStatistics entities.
type TaskDefinition struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The category property
    category *LifecycleTaskCategory
    // The continueOnError property
    continueOnError *bool
    // The description of the taskDefinition.
    description *string
    // The display name of the taskDefinition.Supports $filter(eq, ne) and $orderby.
    displayName *string
    // The parameters that must be supplied when creating a workflow task object.Supports $filter(any).
    parameters []Parameterable
    // The version number of the taskDefinition. New records are pushed when we add support for new parameters.Supports $filter(ge, gt, le, lt, eq, ne) and $orderby.
    version *int32
}
// NewTaskDefinition instantiates a new taskDefinition and sets the default values.
func NewTaskDefinition()(*TaskDefinition) {
    m := &TaskDefinition{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTaskDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskDefinition(), nil
}
// GetCategory gets the category property value. The category property
func (m *TaskDefinition) GetCategory()(*LifecycleTaskCategory) {
    return m.category
}
// GetContinueOnError gets the continueOnError property value. The continueOnError property
func (m *TaskDefinition) GetContinueOnError()(*bool) {
    return m.continueOnError
}
// GetDescription gets the description property value. The description of the taskDefinition.
func (m *TaskDefinition) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the taskDefinition.Supports $filter(eq, ne) and $orderby.
func (m *TaskDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLifecycleTaskCategory , m.SetCategory)
    res["continueOnError"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContinueOnError)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["parameters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateParameterFromDiscriminatorValue , m.SetParameters)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    return res
}
// GetParameters gets the parameters property value. The parameters that must be supplied when creating a workflow task object.Supports $filter(any).
func (m *TaskDefinition) GetParameters()([]Parameterable) {
    return m.parameters
}
// GetVersion gets the version property value. The version number of the taskDefinition. New records are pushed when we add support for new parameters.Supports $filter(ge, gt, le, lt, eq, ne) and $orderby.
func (m *TaskDefinition) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *TaskDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetParameters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetParameters())
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category property
func (m *TaskDefinition) SetCategory(value *LifecycleTaskCategory)() {
    m.category = value
}
// SetContinueOnError sets the continueOnError property value. The continueOnError property
func (m *TaskDefinition) SetContinueOnError(value *bool)() {
    m.continueOnError = value
}
// SetDescription sets the description property value. The description of the taskDefinition.
func (m *TaskDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the taskDefinition.Supports $filter(eq, ne) and $orderby.
func (m *TaskDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetParameters sets the parameters property value. The parameters that must be supplied when creating a workflow task object.Supports $filter(any).
func (m *TaskDefinition) SetParameters(value []Parameterable)() {
    m.parameters = value
}
// SetVersion sets the version property value. The version number of the taskDefinition. New records are pushed when we add support for new parameters.Supports $filter(ge, gt, le, lt, eq, ne) and $orderby.
func (m *TaskDefinition) SetVersion(value *int32)() {
    m.version = value
}
