package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Tasks 
type Tasks struct {
    Entity
    // All tasks in the users mailbox.
    alltasks []BaseTask;
    // The task lists in the users mailbox.
    lists []BaseTaskList;
}
// NewTasks instantiates a new tasks and sets the default values.
func NewTasks()(*Tasks) {
    m := &Tasks{
        Entity: *NewEntity(),
    }
    return m
}
// GetAlltasks gets the alltasks property value. All tasks in the users mailbox.
func (m *Tasks) GetAlltasks()([]BaseTask) {
    if m == nil {
        return nil
    } else {
        return m.alltasks
    }
}
// GetLists gets the lists property value. The task lists in the users mailbox.
func (m *Tasks) GetLists()([]BaseTaskList) {
    if m == nil {
        return nil
    } else {
        return m.lists
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Tasks) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alltasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBaseTask() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BaseTask, len(val))
            for i, v := range val {
                res[i] = *(v.(*BaseTask))
            }
            m.SetAlltasks(res)
        }
        return nil
    }
    res["lists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBaseTaskList() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BaseTaskList, len(val))
            for i, v := range val {
                res[i] = *(v.(*BaseTaskList))
            }
            m.SetLists(res)
        }
        return nil
    }
    return res
}
func (m *Tasks) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Tasks) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlltasks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlltasks()))
        for i, v := range m.GetAlltasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("alltasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLists() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLists()))
        for i, v := range m.GetLists() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlltasks sets the alltasks property value. All tasks in the users mailbox.
func (m *Tasks) SetAlltasks(value []BaseTask)() {
    if m != nil {
        m.alltasks = value
    }
}
// SetLists sets the lists property value. The task lists in the users mailbox.
func (m *Tasks) SetLists(value []BaseTaskList)() {
    if m != nil {
        m.lists = value
    }
}
