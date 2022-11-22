package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTaskGroup provides operations to manage the collection of activityStatistics entities.
type OutlookTaskGroup struct {
    Entity
    // The version of the task group.
    changeKey *string
    // The unique GUID identifier for the task group.
    groupKey *string
    // True if the task group is the default task group.
    isDefaultGroup *bool
    // The name of the task group.
    name *string
    // The collection of task folders in the task group. Read-only. Nullable.
    taskFolders []OutlookTaskFolderable
}
// NewOutlookTaskGroup instantiates a new outlookTaskGroup and sets the default values.
func NewOutlookTaskGroup()(*OutlookTaskGroup) {
    m := &OutlookTaskGroup{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOutlookTaskGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookTaskGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookTaskGroup(), nil
}
// GetChangeKey gets the changeKey property value. The version of the task group.
func (m *OutlookTaskGroup) GetChangeKey()(*string) {
    return m.changeKey
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTaskGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["changeKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetChangeKey)
    res["groupKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGroupKey)
    res["isDefaultGroup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDefaultGroup)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["taskFolders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOutlookTaskFolderFromDiscriminatorValue , m.SetTaskFolders)
    return res
}
// GetGroupKey gets the groupKey property value. The unique GUID identifier for the task group.
func (m *OutlookTaskGroup) GetGroupKey()(*string) {
    return m.groupKey
}
// GetIsDefaultGroup gets the isDefaultGroup property value. True if the task group is the default task group.
func (m *OutlookTaskGroup) GetIsDefaultGroup()(*bool) {
    return m.isDefaultGroup
}
// GetName gets the name property value. The name of the task group.
func (m *OutlookTaskGroup) GetName()(*string) {
    return m.name
}
// GetTaskFolders gets the taskFolders property value. The collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskGroup) GetTaskFolders()([]OutlookTaskFolderable) {
    return m.taskFolders
}
// Serialize serializes information the current object
func (m *OutlookTaskGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskFolders())
        err = writer.WriteCollectionOfObjectValues("taskFolders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChangeKey sets the changeKey property value. The version of the task group.
func (m *OutlookTaskGroup) SetChangeKey(value *string)() {
    m.changeKey = value
}
// SetGroupKey sets the groupKey property value. The unique GUID identifier for the task group.
func (m *OutlookTaskGroup) SetGroupKey(value *string)() {
    m.groupKey = value
}
// SetIsDefaultGroup sets the isDefaultGroup property value. True if the task group is the default task group.
func (m *OutlookTaskGroup) SetIsDefaultGroup(value *bool)() {
    m.isDefaultGroup = value
}
// SetName sets the name property value. The name of the task group.
func (m *OutlookTaskGroup) SetName(value *string)() {
    m.name = value
}
// SetTaskFolders sets the taskFolders property value. The collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskGroup) SetTaskFolders(value []OutlookTaskFolderable)() {
    m.taskFolders = value
}
