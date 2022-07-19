package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TaskList 
type TaskList struct {
    BaseTaskList
}
// NewTaskList instantiates a new TaskList and sets the default values.
func NewTaskList()(*TaskList) {
    m := &TaskList{
        BaseTaskList: *NewBaseTaskList(),
    }
    odataTypeValue := "#microsoft.graph.taskList";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTaskListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskList(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseTaskList.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TaskList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseTaskList.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
