package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OutlookUser struct {
    Entity
}
// NewOutlookUser instantiates a new OutlookUser and sets the default values.
func NewOutlookUser()(*OutlookUser) {
    m := &OutlookUser{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOutlookUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutlookUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OutlookUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["masterCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OutlookCategoryable)
                }
            }
            m.SetMasterCategories(res)
        }
        return nil
    }
    res["taskFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskFolderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OutlookTaskFolderable)
                }
            }
            m.SetTaskFolders(res)
        }
        return nil
    }
    res["taskGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OutlookTaskGroupable)
                }
            }
            m.SetTaskGroups(res)
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OutlookTaskable)
                }
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetMasterCategories gets the masterCategories property value. A list of categories defined for the user.
// returns a []OutlookCategoryable when successful
func (m *OutlookUser) GetMasterCategories()([]OutlookCategoryable) {
    val, err := m.GetBackingStore().Get("masterCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OutlookCategoryable)
    }
    return nil
}
// GetTaskFolders gets the taskFolders property value. The taskFolders property
// returns a []OutlookTaskFolderable when successful
func (m *OutlookUser) GetTaskFolders()([]OutlookTaskFolderable) {
    val, err := m.GetBackingStore().Get("taskFolders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OutlookTaskFolderable)
    }
    return nil
}
// GetTaskGroups gets the taskGroups property value. The taskGroups property
// returns a []OutlookTaskGroupable when successful
func (m *OutlookUser) GetTaskGroups()([]OutlookTaskGroupable) {
    val, err := m.GetBackingStore().Get("taskGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OutlookTaskGroupable)
    }
    return nil
}
// GetTasks gets the tasks property value. The tasks property
// returns a []OutlookTaskable when successful
func (m *OutlookUser) GetTasks()([]OutlookTaskable) {
    val, err := m.GetBackingStore().Get("tasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OutlookTaskable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OutlookUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMasterCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMasterCategories()))
        for i, v := range m.GetMasterCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("masterCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskFolders()))
        for i, v := range m.GetTaskFolders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("taskFolders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskGroups()))
        for i, v := range m.GetTaskGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("taskGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("masterCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetTaskFolders sets the taskFolders property value. The taskFolders property
func (m *OutlookUser) SetTaskFolders(value []OutlookTaskFolderable)() {
    err := m.GetBackingStore().Set("taskFolders", value)
    if err != nil {
        panic(err)
    }
}
// SetTaskGroups sets the taskGroups property value. The taskGroups property
func (m *OutlookUser) SetTaskGroups(value []OutlookTaskGroupable)() {
    err := m.GetBackingStore().Set("taskGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetTasks sets the tasks property value. The tasks property
func (m *OutlookUser) SetTasks(value []OutlookTaskable)() {
    err := m.GetBackingStore().Set("tasks", value)
    if err != nil {
        panic(err)
    }
}
type OutlookUserable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMasterCategories()([]OutlookCategoryable)
    GetTaskFolders()([]OutlookTaskFolderable)
    GetTaskGroups()([]OutlookTaskGroupable)
    GetTasks()([]OutlookTaskable)
    SetMasterCategories(value []OutlookCategoryable)()
    SetTaskFolders(value []OutlookTaskFolderable)()
    SetTaskGroups(value []OutlookTaskGroupable)()
    SetTasks(value []OutlookTaskable)()
}
