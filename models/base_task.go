package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseTask provides operations to manage the collection of accessReviewDecision entities.
type BaseTask struct {
    Entity
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    bodyLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of smaller subtasks linked to the more complex parent task.
    checklistItems []ChecklistItemable
    // The date when the task was finished.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the task.
    displayName *string
    // The date in the specified time zone that the task is to be finished.
    dueDateTime DateTimeTimeZoneable
    // The collection of open extensions defined for the task .
    extensions []Extensionable
    // The importance property
    importance *Importance
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of resources linked to the task.
    linkedResources []LinkedResource_v2able
    // The list which contains the task.
    parentList BaseTaskListable
    // The recurrence pattern for the task.
    recurrence PatternedRecurrenceable
    // The date in the specified time zone when the task is to begin.
    startDateTime DateTimeTimeZoneable
    // The status property
    status *TaskStatus_v2
    // The task body in text format that typically contains information about the task.
    textBody *string
    // The viewpoint property
    viewpoint TaskViewpointable
}
// NewBaseTask instantiates a new baseTask and sets the default values.
func NewBaseTask()(*BaseTask) {
    m := &BaseTask{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.baseTask";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBaseTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBaseTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.task":
                        return NewTask(), nil
                }
            }
        }
    }
    return NewBaseTask(), nil
}
// GetBodyLastModifiedDateTime gets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) GetBodyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.bodyLastModifiedDateTime
}
// GetChecklistItems gets the checklistItems property value. A collection of smaller subtasks linked to the more complex parent task.
func (m *BaseTask) GetChecklistItems()([]ChecklistItemable) {
    return m.checklistItems
}
// GetCompletedDateTime gets the completedDateTime property value. The date when the task was finished.
func (m *BaseTask) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of the task.
func (m *BaseTask) GetDisplayName()(*string) {
    return m.displayName
}
// GetDueDateTime gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *BaseTask) GetDueDateTime()(DateTimeTimeZoneable) {
    return m.dueDateTime
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task .
func (m *BaseTask) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bodyLastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetBodyLastModifiedDateTime)
    res["checklistItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateChecklistItemFromDiscriminatorValue , m.SetChecklistItems)
    res["completedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompletedDateTime)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["dueDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetDueDateTime)
    res["extensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue , m.SetExtensions)
    res["importance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseImportance , m.SetImportance)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["linkedResources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLinkedResource_v2FromDiscriminatorValue , m.SetLinkedResources)
    res["parentList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBaseTaskListFromDiscriminatorValue , m.SetParentList)
    res["recurrence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue , m.SetRecurrence)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetStartDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTaskStatus_v2 , m.SetStatus)
    res["textBody"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTextBody)
    res["viewpoint"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTaskViewpointFromDiscriminatorValue , m.SetViewpoint)
    return res
}
// GetImportance gets the importance property value. The importance property
func (m *BaseTask) GetImportance()(*Importance) {
    return m.importance
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLinkedResources gets the linkedResources property value. A collection of resources linked to the task.
func (m *BaseTask) GetLinkedResources()([]LinkedResource_v2able) {
    return m.linkedResources
}
// GetParentList gets the parentList property value. The list which contains the task.
func (m *BaseTask) GetParentList()(BaseTaskListable) {
    return m.parentList
}
// GetRecurrence gets the recurrence property value. The recurrence pattern for the task.
func (m *BaseTask) GetRecurrence()(PatternedRecurrenceable) {
    return m.recurrence
}
// GetStartDateTime gets the startDateTime property value. The date in the specified time zone when the task is to begin.
func (m *BaseTask) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status property
func (m *BaseTask) GetStatus()(*TaskStatus_v2) {
    return m.status
}
// GetTextBody gets the textBody property value. The task body in text format that typically contains information about the task.
func (m *BaseTask) GetTextBody()(*string) {
    return m.textBody
}
// GetViewpoint gets the viewpoint property value. The viewpoint property
func (m *BaseTask) GetViewpoint()(TaskViewpointable) {
    return m.viewpoint
}
// Serialize serializes information the current object
func (m *BaseTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("bodyLastModifiedDateTime", m.GetBodyLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetChecklistItems() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChecklistItems())
        err = writer.WriteCollectionOfObjectValues("checklistItems", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensions())
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLinkedResources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLinkedResources())
        err = writer.WriteCollectionOfObjectValues("linkedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentList", m.GetParentList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("textBody", m.GetTextBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("viewpoint", m.GetViewpoint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBodyLastModifiedDateTime sets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.bodyLastModifiedDateTime = value
}
// SetChecklistItems sets the checklistItems property value. A collection of smaller subtasks linked to the more complex parent task.
func (m *BaseTask) SetChecklistItems(value []ChecklistItemable)() {
    m.checklistItems = value
}
// SetCompletedDateTime sets the completedDateTime property value. The date when the task was finished.
func (m *BaseTask) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the task.
func (m *BaseTask) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDueDateTime sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *BaseTask) SetDueDateTime(value DateTimeTimeZoneable)() {
    m.dueDateTime = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task .
func (m *BaseTask) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetImportance sets the importance property value. The importance property
func (m *BaseTask) SetImportance(value *Importance)() {
    m.importance = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLinkedResources sets the linkedResources property value. A collection of resources linked to the task.
func (m *BaseTask) SetLinkedResources(value []LinkedResource_v2able)() {
    m.linkedResources = value
}
// SetParentList sets the parentList property value. The list which contains the task.
func (m *BaseTask) SetParentList(value BaseTaskListable)() {
    m.parentList = value
}
// SetRecurrence sets the recurrence property value. The recurrence pattern for the task.
func (m *BaseTask) SetRecurrence(value PatternedRecurrenceable)() {
    m.recurrence = value
}
// SetStartDateTime sets the startDateTime property value. The date in the specified time zone when the task is to begin.
func (m *BaseTask) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *BaseTask) SetStatus(value *TaskStatus_v2)() {
    m.status = value
}
// SetTextBody sets the textBody property value. The task body in text format that typically contains information about the task.
func (m *BaseTask) SetTextBody(value *string)() {
    m.textBody = value
}
// SetViewpoint sets the viewpoint property value. The viewpoint property
func (m *BaseTask) SetViewpoint(value TaskViewpointable)() {
    m.viewpoint = value
}
