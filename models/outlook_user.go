package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookUser 
type OutlookUser struct {
    Entity
    // A list of categories defined for the user.
    masterCategories []OutlookCategoryable
    // The taskFolders property
    taskFolders []OutlookTaskFolderable
    // The taskGroups property
    taskGroups []OutlookTaskGroupable
    // The tasks property
    tasks []OutlookTaskable
}
// NewOutlookUser instantiates a new outlookUser and sets the default values.
func NewOutlookUser()(*OutlookUser) {
    m := &OutlookUser{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.outlookUser";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOutlookUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["masterCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOutlookCategoryFromDiscriminatorValue , m.SetMasterCategories)
    res["taskFolders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOutlookTaskFolderFromDiscriminatorValue , m.SetTaskFolders)
    res["taskGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOutlookTaskGroupFromDiscriminatorValue , m.SetTaskGroups)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOutlookTaskFromDiscriminatorValue , m.SetTasks)
    return res
}
// GetMasterCategories gets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) GetMasterCategories()([]OutlookCategoryable) {
    return m.masterCategories
}
// GetTaskFolders gets the taskFolders property value. The taskFolders property
func (m *OutlookUser) GetTaskFolders()([]OutlookTaskFolderable) {
    return m.taskFolders
}
// GetTaskGroups gets the taskGroups property value. The taskGroups property
func (m *OutlookUser) GetTaskGroups()([]OutlookTaskGroupable) {
    return m.taskGroups
}
// GetTasks gets the tasks property value. The tasks property
func (m *OutlookUser) GetTasks()([]OutlookTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *OutlookUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMasterCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMasterCategories())
        err = writer.WriteCollectionOfObjectValues("masterCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskFolders() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskFolders())
        err = writer.WriteCollectionOfObjectValues("taskFolders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskGroups() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskGroups())
        err = writer.WriteCollectionOfObjectValues("taskGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTasks())
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMasterCategories sets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) SetMasterCategories(value []OutlookCategoryable)() {
    m.masterCategories = value
}
// SetTaskFolders sets the taskFolders property value. The taskFolders property
func (m *OutlookUser) SetTaskFolders(value []OutlookTaskFolderable)() {
    m.taskFolders = value
}
// SetTaskGroups sets the taskGroups property value. The taskGroups property
func (m *OutlookUser) SetTaskGroups(value []OutlookTaskGroupable)() {
    m.taskGroups = value
}
// SetTasks sets the tasks property value. The tasks property
func (m *OutlookUser) SetTasks(value []OutlookTaskable)() {
    m.tasks = value
}
