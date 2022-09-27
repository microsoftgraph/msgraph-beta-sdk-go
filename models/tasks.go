package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Tasks 
type Tasks struct {
    Entity
    // All tasks in the users mailbox.
    alltasks []BaseTaskable
    // The task lists in the users mailbox.
    lists []BaseTaskListable
}
// NewTasks instantiates a new tasks and sets the default values.
func NewTasks()(*Tasks) {
    m := &Tasks{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.tasks";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTasksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTasksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTasks(), nil
}
// GetAlltasks gets the alltasks property value. All tasks in the users mailbox.
func (m *Tasks) GetAlltasks()([]BaseTaskable) {
    return m.alltasks
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Tasks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alltasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBaseTaskFromDiscriminatorValue , m.SetAlltasks)
    res["lists"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBaseTaskListFromDiscriminatorValue , m.SetLists)
    return res
}
// GetLists gets the lists property value. The task lists in the users mailbox.
func (m *Tasks) GetLists()([]BaseTaskListable) {
    return m.lists
}
// Serialize serializes information the current object
func (m *Tasks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlltasks() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAlltasks())
        err = writer.WriteCollectionOfObjectValues("alltasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLists() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLists())
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlltasks sets the alltasks property value. All tasks in the users mailbox.
func (m *Tasks) SetAlltasks(value []BaseTaskable)() {
    m.alltasks = value
}
// SetLists sets the lists property value. The task lists in the users mailbox.
func (m *Tasks) SetLists(value []BaseTaskListable)() {
    m.lists = value
}
