package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTaskFolder provides operations to manage the collection of accessReview entities.
type OutlookTaskFolder struct {
    Entity
    // The version of the task folder.
    changeKey *string
    // True if the folder is the default task folder.
    isDefaultFolder *bool
    // The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The name of the task folder.
    name *string
    // The unique GUID identifier for the task folder's parent group.
    parentGroupKey *string
    // The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The tasks in this task folder. Read-only. Nullable.
    tasks []OutlookTaskable
}
// NewOutlookTaskFolder instantiates a new outlookTaskFolder and sets the default values.
func NewOutlookTaskFolder()(*OutlookTaskFolder) {
    m := &OutlookTaskFolder{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.outlookTaskFolder";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOutlookTaskFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookTaskFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookTaskFolder(), nil
}
// GetChangeKey gets the changeKey property value. The version of the task folder.
func (m *OutlookTaskFolder) GetChangeKey()(*string) {
    return m.changeKey
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTaskFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["changeKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetChangeKey)
    res["isDefaultFolder"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDefaultFolder)
    res["multiValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetMultiValueExtendedProperties)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["parentGroupKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetParentGroupKey)
    res["singleValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetSingleValueExtendedProperties)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOutlookTaskFromDiscriminatorValue , m.SetTasks)
    return res
}
// GetIsDefaultFolder gets the isDefaultFolder property value. True if the folder is the default task folder.
func (m *OutlookTaskFolder) GetIsDefaultFolder()(*bool) {
    return m.isDefaultFolder
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetName gets the name property value. The name of the task folder.
func (m *OutlookTaskFolder) GetName()(*string) {
    return m.name
}
// GetParentGroupKey gets the parentGroupKey property value. The unique GUID identifier for the task folder's parent group.
func (m *OutlookTaskFolder) GetParentGroupKey()(*string) {
    return m.parentGroupKey
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetTasks gets the tasks property value. The tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetTasks()([]OutlookTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *OutlookTaskFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isDefaultFolder", m.GetIsDefaultFolder())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMultiValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
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
    {
        err = writer.WriteStringValue("parentGroupKey", m.GetParentGroupKey())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSingleValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
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
// SetChangeKey sets the changeKey property value. The version of the task folder.
func (m *OutlookTaskFolder) SetChangeKey(value *string)() {
    m.changeKey = value
}
// SetIsDefaultFolder sets the isDefaultFolder property value. True if the folder is the default task folder.
func (m *OutlookTaskFolder) SetIsDefaultFolder(value *bool)() {
    m.isDefaultFolder = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetName sets the name property value. The name of the task folder.
func (m *OutlookTaskFolder) SetName(value *string)() {
    m.name = value
}
// SetParentGroupKey sets the parentGroupKey property value. The unique GUID identifier for the task folder's parent group.
func (m *OutlookTaskFolder) SetParentGroupKey(value *string)() {
    m.parentGroupKey = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetTasks sets the tasks property value. The tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetTasks(value []OutlookTaskable)() {
    m.tasks = value
}
