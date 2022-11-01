package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseTaskList provides operations to manage the collection of accessReviewDecision entities.
type BaseTaskList struct {
    Entity
    // The name of the task list.
    displayName *string
    // The collection of open extensions defined for the task list. Nullable.
    extensions []Extensionable
    // The tasks in this task list. Read-only. Nullable.
    tasks []BaseTaskable
}
// NewBaseTaskList instantiates a new baseTaskList and sets the default values.
func NewBaseTaskList()(*BaseTaskList) {
    m := &BaseTaskList{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.baseTaskList";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBaseTaskListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBaseTaskListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.taskList":
                        return NewTaskList(), nil
                    case "#microsoft.graph.wellKnownTaskList":
                        return NewWellKnownTaskList(), nil
                }
            }
        }
    }
    return NewBaseTaskList(), nil
}
// GetDisplayName gets the displayName property value. The name of the task list.
func (m *BaseTaskList) GetDisplayName()(*string) {
    return m.displayName
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *BaseTaskList) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseTaskList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["extensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue , m.SetExtensions)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBaseTaskFromDiscriminatorValue , m.SetTasks)
    return res
}
// GetTasks gets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *BaseTaskList) GetTasks()([]BaseTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *BaseTaskList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensions())
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
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
// SetDisplayName sets the displayName property value. The name of the task list.
func (m *BaseTaskList) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *BaseTaskList) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetTasks sets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *BaseTaskList) SetTasks(value []BaseTaskable)() {
    m.tasks = value
}
