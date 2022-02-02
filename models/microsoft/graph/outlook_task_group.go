package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookTaskGroup 
type OutlookTaskGroup struct {
    Entity
    // The version of the task group.
    changeKey *string;
    // The unique GUID identifier for the task group.
    groupKey *string;
    // True if the task group is the default task group.
    isDefaultGroup *bool;
    // The name of the task group.
    name *string;
    // The collection of task folders in the task group. Read-only. Nullable.
    taskFolders []OutlookTaskFolder;
}
// NewOutlookTaskGroup instantiates a new outlookTaskGroup and sets the default values.
func NewOutlookTaskGroup()(*OutlookTaskGroup) {
    m := &OutlookTaskGroup{
        Entity: *NewEntity(),
    }
    return m
}
// GetChangeKey gets the changeKey property value. The version of the task group.
func (m *OutlookTaskGroup) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// GetGroupKey gets the groupKey property value. The unique GUID identifier for the task group.
func (m *OutlookTaskGroup) GetGroupKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupKey
    }
}
// GetIsDefaultGroup gets the isDefaultGroup property value. True if the task group is the default task group.
func (m *OutlookTaskGroup) GetIsDefaultGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultGroup
    }
}
// GetName gets the name property value. The name of the task group.
func (m *OutlookTaskGroup) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetTaskFolders gets the taskFolders property value. The collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskGroup) GetTaskFolders()([]OutlookTaskFolder) {
    if m == nil {
        return nil
    } else {
        return m.taskFolders
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTaskGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["groupKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupKey(val)
        }
        return nil
    }
    res["isDefaultGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultGroup(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["taskFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookTaskFolder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskFolder, len(val))
            for i, v := range val {
                res[i] = *(v.(*OutlookTaskFolder))
            }
            m.SetTaskFolders(res)
        }
        return nil
    }
    return res
}
func (m *OutlookTaskGroup) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OutlookTaskGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupKey", m.GetGroupKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultGroup", m.GetIsDefaultGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetTaskFolders() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskFolders()))
        for i, v := range m.GetTaskFolders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taskFolders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChangeKey sets the changeKey property value. The version of the task group.
func (m *OutlookTaskGroup) SetChangeKey(value *string)() {
    if m != nil {
        m.changeKey = value
    }
}
// SetGroupKey sets the groupKey property value. The unique GUID identifier for the task group.
func (m *OutlookTaskGroup) SetGroupKey(value *string)() {
    if m != nil {
        m.groupKey = value
    }
}
// SetIsDefaultGroup sets the isDefaultGroup property value. True if the task group is the default task group.
func (m *OutlookTaskGroup) SetIsDefaultGroup(value *bool)() {
    if m != nil {
        m.isDefaultGroup = value
    }
}
// SetName sets the name property value. The name of the task group.
func (m *OutlookTaskGroup) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetTaskFolders sets the taskFolders property value. The collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskGroup) SetTaskFolders(value []OutlookTaskFolder)() {
    if m != nil {
        m.taskFolders = value
    }
}
