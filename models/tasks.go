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
    return m
}
// CreateTasksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTasksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTasks(), nil
}
// GetAlltasks gets the alltasks property value. All tasks in the users mailbox.
func (m *Tasks) GetAlltasks()([]BaseTaskable) {
    if m == nil {
        return nil
    } else {
        return m.alltasks
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Tasks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alltasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBaseTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BaseTaskable, len(val))
            for i, v := range val {
                res[i] = v.(BaseTaskable)
            }
            m.SetAlltasks(res)
        }
        return nil
    }
    res["lists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBaseTaskListFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BaseTaskListable, len(val))
            for i, v := range val {
                res[i] = v.(BaseTaskListable)
            }
            m.SetLists(res)
        }
        return nil
    }
    return res
}
// GetLists gets the lists property value. The task lists in the users mailbox.
func (m *Tasks) GetLists()([]BaseTaskListable) {
    if m == nil {
        return nil
    } else {
        return m.lists
    }
}
// Serialize serializes information the current object
func (m *Tasks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlltasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlltasks()))
        for i, v := range m.GetAlltasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alltasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLists() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLists()))
        for i, v := range m.GetLists() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlltasks sets the alltasks property value. All tasks in the users mailbox.
func (m *Tasks) SetAlltasks(value []BaseTaskable)() {
    if m != nil {
        m.alltasks = value
    }
}
// SetLists sets the lists property value. The task lists in the users mailbox.
func (m *Tasks) SetLists(value []BaseTaskListable)() {
    if m != nil {
        m.lists = value
    }
}
