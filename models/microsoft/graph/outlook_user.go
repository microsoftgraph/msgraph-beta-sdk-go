package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookUser provides operations to manage the compliance singleton.
type OutlookUser struct {
    Entity
    // A list of categories defined for the user.
    masterCategories []OutlookCategoryable;
    // 
    taskFolders []OutlookTaskFolderable;
    // 
    taskGroups []OutlookTaskGroupable;
    // 
    tasks []OutlookTaskable;
}
// NewOutlookUser instantiates a new outlookUser and sets the default values.
func NewOutlookUser()(*OutlookUser) {
    m := &OutlookUser{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOutlookUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookUserFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOutlookUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["masterCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(OutlookCategoryable)
            }
            m.SetMasterCategories(res)
        }
        return nil
    }
    res["taskFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskFolderable, len(val))
            for i, v := range val {
                res[i] = v.(OutlookTaskFolderable)
            }
            m.SetTaskFolders(res)
        }
        return nil
    }
    res["taskGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskGroupable, len(val))
            for i, v := range val {
                res[i] = v.(OutlookTaskGroupable)
            }
            m.SetTaskGroups(res)
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskable, len(val))
            for i, v := range val {
                res[i] = v.(OutlookTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetMasterCategories gets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) GetMasterCategories()([]OutlookCategoryable) {
    if m == nil {
        return nil
    } else {
        return m.masterCategories
    }
}
// GetTaskFolders gets the taskFolders property value. 
func (m *OutlookUser) GetTaskFolders()([]OutlookTaskFolderable) {
    if m == nil {
        return nil
    } else {
        return m.taskFolders
    }
}
// GetTaskGroups gets the taskGroups property value. 
func (m *OutlookUser) GetTaskGroups()([]OutlookTaskGroupable) {
    if m == nil {
        return nil
    } else {
        return m.taskGroups
    }
}
// GetTasks gets the tasks property value. 
func (m *OutlookUser) GetTasks()([]OutlookTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("masterCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskFolders() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskFolders()))
        for i, v := range m.GetTaskFolders() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskFolders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskGroups()))
        for i, v := range m.GetTaskGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskGroups", cast)
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
// SetMasterCategories sets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) SetMasterCategories(value []OutlookCategoryable)() {
    if m != nil {
        m.masterCategories = value
    }
}
// SetTaskFolders sets the taskFolders property value. 
func (m *OutlookUser) SetTaskFolders(value []OutlookTaskFolderable)() {
    if m != nil {
        m.taskFolders = value
    }
}
// SetTaskGroups sets the taskGroups property value. 
func (m *OutlookUser) SetTaskGroups(value []OutlookTaskGroupable)() {
    if m != nil {
        m.taskGroups = value
    }
}
// SetTasks sets the tasks property value. 
func (m *OutlookUser) SetTasks(value []OutlookTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
