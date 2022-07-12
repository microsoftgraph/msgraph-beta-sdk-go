package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseTaskList 
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
    typeValue := "#microsoft.graph.baseTaskList";
    m.SetType(&typeValue);
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
                mappingStr := *mappingValue
                switch mappingStr {
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
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *BaseTaskList) GetExtensions()([]Extensionable) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseTaskList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBaseTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BaseTaskable, len(val))
            for i, v := range val {
                res[i] = v.(BaseTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetTasks gets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *BaseTaskList) GetTasks()([]BaseTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the task list.
func (m *BaseTaskList) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *BaseTaskList) SetExtensions(value []Extensionable)() {
    if m != nil {
        m.extensions = value
    }
}
// SetTasks sets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *BaseTaskList) SetTasks(value []BaseTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
