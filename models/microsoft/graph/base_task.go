package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BaseTask 
type BaseTask struct {
    Entity
    // The task body that typically contains information about the task.
    body *ItemBody;
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    bodyLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A collection of checklistItems linked to a task.
    checklistItems []ChecklistItem;
    // The date when the task was finished.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The name of the task.
    displayName *string;
    // The date in the specified time zone that the task is to be finished.
    dueDateTime *DateTimeTimeZone;
    // The collection of open extensions defined for the task .
    extensions []Extension;
    // The importance of the task. Possible values are: low, normal, high.  The possible values are: low, normal, high.
    importance *Importance;
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A collection of resources linked to the task.
    linkedResources []LinkedResource_v2;
    // The list which contains the task.
    parentList *BaseTaskList;
    // 
    personalProperties *PersonalTaskProperties;
    // The recurrence pattern for the task.
    recurrence *PatternedRecurrence;
    // The date in the specified time zone when the task is to begin.
    startDateTime *DateTimeTimeZone;
    // Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed,unknownFutureValue.
    status *TaskStatus_v2;
}
// NewBaseTask instantiates a new baseTask and sets the default values.
func NewBaseTask()(*BaseTask) {
    m := &BaseTask{
        Entity: *NewEntity(),
    }
    return m
}
// GetBody gets the body property value. The task body that typically contains information about the task.
func (m *BaseTask) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
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
// GetChecklistItems gets the checklistItems property value. A collection of checklistItems linked to a task.
func (m *BaseTask) GetChecklistItems()([]ChecklistItem) {
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
func (m *BaseTask) GetDueDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task .
func (m *BaseTask) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
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
func (m *BaseTask) GetLinkedResources()([]LinkedResource_v2) {
    if m == nil {
        return nil
    } else {
        return m.linkedResources
    }
}
// GetParentList gets the parentList property value. The list which contains the task.
func (m *BaseTask) GetParentList()(*BaseTaskList) {
    if m == nil {
        return nil
    } else {
        return m.parentList
    }
}
// GetPersonalProperties gets the personalProperties property value. 
func (m *BaseTask) GetPersonalProperties()(*PersonalTaskProperties) {
    if m == nil {
        return nil
    } else {
        return m.personalProperties
    }
}
// GetRecurrence gets the recurrence property value. The recurrence pattern for the task.
func (m *BaseTask) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetStartDateTime gets the startDateTime property value. The date in the specified time zone when the task is to begin.
func (m *BaseTask) GetStartDateTime()(*DateTimeTimeZone) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(*ItemBody))
        }
        return nil
    }
    res["bodyLastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyLastModifiedDateTime(val)
        }
        return nil
    }
    res["checklistItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChecklistItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChecklistItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChecklistItem))
            }
            m.SetChecklistItems(res)
        }
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["linkedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLinkedResource_v2() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LinkedResource_v2, len(val))
            for i, v := range val {
                res[i] = *(v.(*LinkedResource_v2))
            }
            m.SetLinkedResources(res)
        }
        return nil
    }
    res["parentList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBaseTaskList() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentList(val.(*BaseTaskList))
        }
        return nil
    }
    res["personalProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonalTaskProperties() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProperties(val.(*PersonalTaskProperties))
        }
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(*PatternedRecurrence))
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaskStatus_v2)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TaskStatus_v2))
        }
        return nil
    }
    return res
}
func (m *BaseTask) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BaseTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("bodyLastModifiedDateTime", m.GetBodyLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetChecklistItems() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChecklistItems()))
        for i, v := range m.GetChecklistItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLinkedResources()))
        for i, v := range m.GetLinkedResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err = writer.WriteObjectValue("personalProperties", m.GetPersonalProperties())
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
    return nil
}
// SetBody sets the body property value. The task body that typically contains information about the task.
func (m *BaseTask) SetBody(value *ItemBody)() {
    if m != nil {
        m.body = value
    }
}
// SetBodyLastModifiedDateTime sets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *BaseTask) SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.bodyLastModifiedDateTime = value
    }
}
// SetChecklistItems sets the checklistItems property value. A collection of checklistItems linked to a task.
func (m *BaseTask) SetChecklistItems(value []ChecklistItem)() {
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
func (m *BaseTask) SetDueDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.dueDateTime = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task .
func (m *BaseTask) SetExtensions(value []Extension)() {
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
func (m *BaseTask) SetLinkedResources(value []LinkedResource_v2)() {
    if m != nil {
        m.linkedResources = value
    }
}
// SetParentList sets the parentList property value. The list which contains the task.
func (m *BaseTask) SetParentList(value *BaseTaskList)() {
    if m != nil {
        m.parentList = value
    }
}
// SetPersonalProperties sets the personalProperties property value. 
func (m *BaseTask) SetPersonalProperties(value *PersonalTaskProperties)() {
    if m != nil {
        m.personalProperties = value
    }
}
// SetRecurrence sets the recurrence property value. The recurrence pattern for the task.
func (m *BaseTask) SetRecurrence(value *PatternedRecurrence)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetStartDateTime sets the startDateTime property value. The date in the specified time zone when the task is to begin.
func (m *BaseTask) SetStartDateTime(value *DateTimeTimeZone)() {
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
