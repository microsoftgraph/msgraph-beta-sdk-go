package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BaseTaskList 
type BaseTaskList struct {
    Entity
    // The name of the task list.
    displayName *string;
    // The collection of open extensions defined for the task list. Nullable.
    extensions []Extensionable;
    // The tasks in this task list. Read-only. Nullable.
    tasks []BaseTaskable;
}
// NewBaseTaskList instantiates a new baseTaskList and sets the default values.
func NewBaseTaskList()(*BaseTaskList) {
    m := &BaseTaskList{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBaseTaskListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBaseTaskListFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *BaseTaskList) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *BaseTaskList) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
