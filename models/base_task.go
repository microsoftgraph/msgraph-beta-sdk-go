package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseTask provides operations to manage the collection of accessReview entities.
type BaseTask struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    // The importance of the task. Possible values are: low, normal, high.  The possible values are: low, normal, high.
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
    // Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed,unknownFutureValue.
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
    m.SetAdditionalData(make(map[string]interface{}));
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.task":
                        return NewTask(), nil
                }
            }
        }
    }
    return NewBaseTask(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BaseTask) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBodyLastModifiedDateTime gets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) GetBodyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.bodyLastModifiedDateTime
    }
}
// GetChecklistItems gets the checklistItems property value. A collection of smaller subtasks linked to the more complex parent task.
func (m *BaseTask) GetChecklistItems()([]ChecklistItemable) {
    if m == nil {
        return nil
    } else {
        return m.checklistItems
    }
}
// GetCompletedDateTime gets the completedDateTime property value. The date when the task was finished.
func (m *BaseTask) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. The name of the task.
func (m *BaseTask) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDueDateTime gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *BaseTask) GetDueDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task .
func (m *BaseTask) GetExtensions()([]Extensionable) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bodyLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyLastModifiedDateTime(val)
        }
        return nil
    }
    res["checklistItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChecklistItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChecklistItemable, len(val))
            for i, v := range val {
                res[i] = v.(ChecklistItemable)
            }
            m.SetChecklistItems(res)
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["linkedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLinkedResource_v2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LinkedResource_v2able, len(val))
            for i, v := range val {
                res[i] = v.(LinkedResource_v2able)
            }
            m.SetLinkedResources(res)
        }
        return nil
    }
    res["parentList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBaseTaskListFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentList(val.(BaseTaskListable))
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaskStatus_v2)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TaskStatus_v2))
        }
        return nil
    }
    res["textBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTextBody(val)
        }
        return nil
    }
    res["viewpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTaskViewpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewpoint(val.(TaskViewpointable))
        }
        return nil
    }
    return res
}
// GetImportance gets the importance property value. The importance of the task. Possible values are: low, normal, high.  The possible values are: low, normal, high.
func (m *BaseTask) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLinkedResources gets the linkedResources property value. A collection of resources linked to the task.
func (m *BaseTask) GetLinkedResources()([]LinkedResource_v2able) {
    if m == nil {
        return nil
    } else {
        return m.linkedResources
    }
}
// GetParentList gets the parentList property value. The list which contains the task.
func (m *BaseTask) GetParentList()(BaseTaskListable) {
    if m == nil {
        return nil
    } else {
        return m.parentList
    }
}
// GetRecurrence gets the recurrence property value. The recurrence pattern for the task.
func (m *BaseTask) GetRecurrence()(PatternedRecurrenceable) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetStartDateTime gets the startDateTime property value. The date in the specified time zone when the task is to begin.
func (m *BaseTask) GetStartDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetStatus gets the status property value. Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed,unknownFutureValue.
func (m *BaseTask) GetStatus()(*TaskStatus_v2) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTextBody gets the textBody property value. The task body in text format that typically contains information about the task.
func (m *BaseTask) GetTextBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.textBody
    }
}
// GetViewpoint gets the viewpoint property value. The viewpoint property
func (m *BaseTask) GetViewpoint()(TaskViewpointable) {
    if m == nil {
        return nil
    } else {
        return m.viewpoint
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChecklistItems()))
        for i, v := range m.GetChecklistItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinkedResources()))
        for i, v := range m.GetLinkedResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BaseTask) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBodyLastModifiedDateTime sets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.bodyLastModifiedDateTime = value
    }
}
// SetChecklistItems sets the checklistItems property value. A collection of smaller subtasks linked to the more complex parent task.
func (m *BaseTask) SetChecklistItems(value []ChecklistItemable)() {
    if m != nil {
        m.checklistItems = value
    }
}
// SetCompletedDateTime sets the completedDateTime property value. The date when the task was finished.
func (m *BaseTask) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The name of the task.
func (m *BaseTask) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDueDateTime sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *BaseTask) SetDueDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.dueDateTime = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task .
func (m *BaseTask) SetExtensions(value []Extensionable)() {
    if m != nil {
        m.extensions = value
    }
}
// SetImportance sets the importance property value. The importance of the task. Possible values are: low, normal, high.  The possible values are: low, normal, high.
func (m *BaseTask) SetImportance(value *Importance)() {
    if m != nil {
        m.importance = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLinkedResources sets the linkedResources property value. A collection of resources linked to the task.
func (m *BaseTask) SetLinkedResources(value []LinkedResource_v2able)() {
    if m != nil {
        m.linkedResources = value
    }
}
// SetParentList sets the parentList property value. The list which contains the task.
func (m *BaseTask) SetParentList(value BaseTaskListable)() {
    if m != nil {
        m.parentList = value
    }
}
// SetRecurrence sets the recurrence property value. The recurrence pattern for the task.
func (m *BaseTask) SetRecurrence(value PatternedRecurrenceable)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetStartDateTime sets the startDateTime property value. The date in the specified time zone when the task is to begin.
func (m *BaseTask) SetStartDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetStatus sets the status property value. Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed,unknownFutureValue.
func (m *BaseTask) SetStatus(value *TaskStatus_v2)() {
    if m != nil {
        m.status = value
    }
}
// SetTextBody sets the textBody property value. The task body in text format that typically contains information about the task.
func (m *BaseTask) SetTextBody(value *string)() {
    if m != nil {
        m.textBody = value
    }
}
// SetViewpoint sets the viewpoint property value. The viewpoint property
func (m *BaseTask) SetViewpoint(value TaskViewpointable)() {
    if m != nil {
        m.viewpoint = value
    }
}
