package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookDocumentTask struct {
    Entity
}
// NewWorkbookDocumentTask instantiates a new WorkbookDocumentTask and sets the default values.
func NewWorkbookDocumentTask()(*WorkbookDocumentTask) {
    m := &WorkbookDocumentTask{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookDocumentTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookDocumentTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookDocumentTask(), nil
}
// GetAssignees gets the assignees property value. A collection of user identities the task is assigned to.
// returns a []WorkbookEmailIdentityable when successful
func (m *WorkbookDocumentTask) GetAssignees()([]WorkbookEmailIdentityable) {
    val, err := m.GetBackingStore().Get("assignees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkbookEmailIdentityable)
    }
    return nil
}
// GetChanges gets the changes property value. A collection of task change histories.
// returns a []WorkbookDocumentTaskChangeable when successful
func (m *WorkbookDocumentTask) GetChanges()([]WorkbookDocumentTaskChangeable) {
    val, err := m.GetBackingStore().Get("changes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkbookDocumentTaskChangeable)
    }
    return nil
}
// GetComment gets the comment property value. The comment that the task is associated with.
// returns a WorkbookCommentable when successful
func (m *WorkbookDocumentTask) GetComment()(WorkbookCommentable) {
    val, err := m.GetBackingStore().Get("comment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookCommentable)
    }
    return nil
}
// GetCompletedBy gets the completedBy property value. The identity of the user who completed the task. Nullable.
// returns a WorkbookEmailIdentityable when successful
func (m *WorkbookDocumentTask) GetCompletedBy()(WorkbookEmailIdentityable) {
    val, err := m.GetBackingStore().Get("completedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookEmailIdentityable)
    }
    return nil
}
// GetCompletedDateTime gets the completedDateTime property value. Date and time when the task was completed. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WorkbookDocumentTask) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("completedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. A user identity that creates the task. Nullable.
// returns a WorkbookEmailIdentityable when successful
func (m *WorkbookDocumentTask) GetCreatedBy()(WorkbookEmailIdentityable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookEmailIdentityable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the task was created. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WorkbookDocumentTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkbookDocumentTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookEmailIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkbookEmailIdentityable)
                }
            }
            m.SetAssignees(res)
        }
        return nil
    }
    res["changes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookDocumentTaskChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookDocumentTaskChangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkbookDocumentTaskChangeable)
                }
            }
            m.SetChanges(res)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookCommentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val.(WorkbookCommentable))
        }
        return nil
    }
    res["completedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedBy(val.(WorkbookEmailIdentityable))
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
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(WorkbookEmailIdentityable))
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
    res["percentComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentComplete(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["startAndDueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookDocumentTaskScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartAndDueDateTime(val.(WorkbookDocumentTaskScheduleable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetPercentComplete gets the percentComplete property value. An integer value from 0 to 100 that represents the percentage of the completion of the task. 100 means that the task is completed. Nullable.
// returns a *int32 when successful
func (m *WorkbookDocumentTask) GetPercentComplete()(*int32) {
    val, err := m.GetBackingStore().Get("percentComplete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPriority gets the priority property value. An integer value from 0 to 10 that represents the priority of the task. A lower value indicates a higher priority. Nullable.
// returns a *int32 when successful
func (m *WorkbookDocumentTask) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStartAndDueDateTime gets the startAndDueDateTime property value. Start and due date of the task. Nullable.
// returns a WorkbookDocumentTaskScheduleable when successful
func (m *WorkbookDocumentTask) GetStartAndDueDateTime()(WorkbookDocumentTaskScheduleable) {
    val, err := m.GetBackingStore().Get("startAndDueDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookDocumentTaskScheduleable)
    }
    return nil
}
// GetTitle gets the title property value. The title of the task.
// returns a *string when successful
func (m *WorkbookDocumentTask) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookDocumentTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignees()))
        for i, v := range m.GetAssignees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChanges()))
        for i, v := range m.GetChanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("changes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completedBy", m.GetCompletedBy())
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
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteInt32Value("percentComplete", m.GetPercentComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startAndDueDateTime", m.GetStartAndDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignees sets the assignees property value. A collection of user identities the task is assigned to.
func (m *WorkbookDocumentTask) SetAssignees(value []WorkbookEmailIdentityable)() {
    err := m.GetBackingStore().Set("assignees", value)
    if err != nil {
        panic(err)
    }
}
// SetChanges sets the changes property value. A collection of task change histories.
func (m *WorkbookDocumentTask) SetChanges(value []WorkbookDocumentTaskChangeable)() {
    err := m.GetBackingStore().Set("changes", value)
    if err != nil {
        panic(err)
    }
}
// SetComment sets the comment property value. The comment that the task is associated with.
func (m *WorkbookDocumentTask) SetComment(value WorkbookCommentable)() {
    err := m.GetBackingStore().Set("comment", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletedBy sets the completedBy property value. The identity of the user who completed the task. Nullable.
func (m *WorkbookDocumentTask) SetCompletedBy(value WorkbookEmailIdentityable)() {
    err := m.GetBackingStore().Set("completedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletedDateTime sets the completedDateTime property value. Date and time when the task was completed. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WorkbookDocumentTask) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("completedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. A user identity that creates the task. Nullable.
func (m *WorkbookDocumentTask) SetCreatedBy(value WorkbookEmailIdentityable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the task was created. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WorkbookDocumentTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPercentComplete sets the percentComplete property value. An integer value from 0 to 100 that represents the percentage of the completion of the task. 100 means that the task is completed. Nullable.
func (m *WorkbookDocumentTask) SetPercentComplete(value *int32)() {
    err := m.GetBackingStore().Set("percentComplete", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. An integer value from 0 to 10 that represents the priority of the task. A lower value indicates a higher priority. Nullable.
func (m *WorkbookDocumentTask) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetStartAndDueDateTime sets the startAndDueDateTime property value. Start and due date of the task. Nullable.
func (m *WorkbookDocumentTask) SetStartAndDueDateTime(value WorkbookDocumentTaskScheduleable)() {
    err := m.GetBackingStore().Set("startAndDueDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. The title of the task.
func (m *WorkbookDocumentTask) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
type WorkbookDocumentTaskable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignees()([]WorkbookEmailIdentityable)
    GetChanges()([]WorkbookDocumentTaskChangeable)
    GetComment()(WorkbookCommentable)
    GetCompletedBy()(WorkbookEmailIdentityable)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(WorkbookEmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPercentComplete()(*int32)
    GetPriority()(*int32)
    GetStartAndDueDateTime()(WorkbookDocumentTaskScheduleable)
    GetTitle()(*string)
    SetAssignees(value []WorkbookEmailIdentityable)()
    SetChanges(value []WorkbookDocumentTaskChangeable)()
    SetComment(value WorkbookCommentable)()
    SetCompletedBy(value WorkbookEmailIdentityable)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value WorkbookEmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPercentComplete(value *int32)()
    SetPriority(value *int32)()
    SetStartAndDueDateTime(value WorkbookDocumentTaskScheduleable)()
    SetTitle(value *string)()
}
