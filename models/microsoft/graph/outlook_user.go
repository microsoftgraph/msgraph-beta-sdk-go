package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookUser 
type OutlookUser struct {
    Entity
    // A list of categories defined for the user.
    masterCategories []OutlookCategory;
    // 
    taskFolders []OutlookTaskFolder;
    // 
    taskGroups []OutlookTaskGroup;
    // 
    tasks []OutlookTask;
}
// NewOutlookUser instantiates a new outlookUser and sets the default values.
func NewOutlookUser()(*OutlookUser) {
    m := &OutlookUser{
        Entity: *NewEntity(),
    }
    return m
}
// GetMasterCategories gets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) GetMasterCategories()([]OutlookCategory) {
    if m == nil {
        return nil
    } else {
        return m.masterCategories
    }
}
// GetTaskFolders gets the taskFolders property value. 
func (m *OutlookUser) GetTaskFolders()([]OutlookTaskFolder) {
    if m == nil {
        return nil
    } else {
        return m.taskFolders
    }
}
// GetTaskGroups gets the taskGroups property value. 
func (m *OutlookUser) GetTaskGroups()([]OutlookTaskGroup) {
    if m == nil {
        return nil
    } else {
        return m.taskGroups
    }
}
// GetTasks gets the tasks property value. 
func (m *OutlookUser) GetTasks()([]OutlookTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["masterCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*OutlookCategory))
            }
            m.SetMasterCategories(res)
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
    res["taskGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookTaskGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*OutlookTaskGroup))
            }
            m.SetTaskGroups(res)
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookTask() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTask, len(val))
            for i, v := range val {
                res[i] = *(v.(*OutlookTask))
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
func (m *OutlookUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OutlookUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMasterCategories() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMasterCategories()))
        for i, v := range m.GetMasterCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("masterCategories", cast)
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
    if m.GetTaskGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskGroups()))
        for i, v := range m.GetTaskGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taskGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMasterCategories sets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) SetMasterCategories(value []OutlookCategory)() {
    if m != nil {
        m.masterCategories = value
    }
}
// SetTaskFolders sets the taskFolders property value. 
func (m *OutlookUser) SetTaskFolders(value []OutlookTaskFolder)() {
    if m != nil {
        m.taskFolders = value
    }
}
// SetTaskGroups sets the taskGroups property value. 
func (m *OutlookUser) SetTaskGroups(value []OutlookTaskGroup)() {
    if m != nil {
        m.taskGroups = value
    }
}
// SetTasks sets the tasks property value. 
func (m *OutlookUser) SetTasks(value []OutlookTask)() {
    if m != nil {
        m.tasks = value
    }
}
