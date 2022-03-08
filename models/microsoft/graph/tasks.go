package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Tasks provides operations to manage the compliance singleton.
type Tasks struct {
    Entity
    // All tasks in the users mailbox.
    alltasks []BaseTaskable;
    // The task lists in the users mailbox.
    lists []BaseTaskListable;
}
// NewTasks instantiates a new tasks and sets the default values.
func NewTasks()(*Tasks) {
    m := &Tasks{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTasksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTasksFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *Tasks) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alltasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["lists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alltasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLists() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLists()))
        for i, v := range m.GetLists() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
