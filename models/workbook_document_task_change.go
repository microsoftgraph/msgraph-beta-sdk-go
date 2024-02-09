package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkbookDocumentTaskChange struct {
    Entity
}
// NewWorkbookDocumentTaskChange instantiates a new WorkbookDocumentTaskChange and sets the default values.
func NewWorkbookDocumentTaskChange()(*WorkbookDocumentTaskChange) {
    m := &WorkbookDocumentTaskChange{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookDocumentTaskChangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkbookDocumentTaskChangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookDocumentTaskChange(), nil
}
// GetAssignee gets the assignee property value. The user identity the task is assigned to. Only present when the type property is assign. Nullable.
// returns a WorkbookEmailIdentityable when successful
func (m *WorkbookDocumentTaskChange) GetAssignee()(WorkbookEmailIdentityable) {
    val, err := m.GetBackingStore().Get("assignee")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookEmailIdentityable)
    }
    return nil
}
// GetChangedBy gets the changedBy property value. The changedBy property
// returns a WorkbookEmailIdentityable when successful
func (m *WorkbookDocumentTaskChange) GetChangedBy()(WorkbookEmailIdentityable) {
    val, err := m.GetBackingStore().Get("changedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookEmailIdentityable)
    }
    return nil
}
// GetCommentId gets the commentId property value. The identifier of the associated comment.
// returns a *string when successful
func (m *WorkbookDocumentTaskChange) GetCommentId()(*string) {
    val, err := m.GetBackingStore().Get("commentId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the task was changed. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WorkbookDocumentTaskChange) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDueDateTime gets the dueDateTime property value. The due date and time for the task. Only present when the type property is setSchedule. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WorkbookDocumentTaskChange) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("dueDateTime")
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
func (m *WorkbookDocumentTaskChange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val.(WorkbookEmailIdentityable))
        }
        return nil
    }
    res["changedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangedBy(val.(WorkbookEmailIdentityable))
        }
        return nil
    }
    res["commentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentId(val)
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
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val)
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["undoChangeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUndoChangeId(val)
        }
        return nil
    }
    return res
}
// GetPercentComplete gets the percentComplete property value. An integer value from 0 to 100 that represents the percentage of the completion of the task and associated comment. 100 means that the task and associated comment are completed. If you change the completion from 100 to a lower value, the associated task and comment are reactivated. Only present when the type property is setPercentComplete. Nullable.
// returns a *int32 when successful
func (m *WorkbookDocumentTaskChange) GetPercentComplete()(*int32) {
    val, err := m.GetBackingStore().Get("percentComplete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPriority gets the priority property value. An integer value from 0 to 10 that represents the priority of the task. A lower value indicates a higher priority. 5 indicates the default priority if not set. Only present when the type property is setPriority. Nullable.
// returns a *int32 when successful
func (m *WorkbookDocumentTaskChange) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The start date and time for the task. Only present when the type property is setSchedule. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WorkbookDocumentTaskChange) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTitle gets the title property value. The title of the task. Only present when the type property is setTitle. Nullable.
// returns a *string when successful
func (m *WorkbookDocumentTaskChange) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTypeEscaped gets the type property value. The type of the change history. Possible values are: create, assign, unassign, unassignAll, setPriority, setTitle, setPercentComplete, setSchedule, remove, restore, undo.
// returns a *string when successful
func (m *WorkbookDocumentTaskChange) GetTypeEscaped()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUndoChangeId gets the undoChangeId property value. The ID of the workbookDocumentTaskChange that was undone for the undo change action. Only exists on an undo change history. Nullable.
// returns a *string when successful
func (m *WorkbookDocumentTaskChange) GetUndoChangeId()(*string) {
    val, err := m.GetBackingStore().Get("undoChangeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookDocumentTaskChange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("assignee", m.GetAssignee())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("changedBy", m.GetChangedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("commentId", m.GetCommentId())
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
        err = writer.WriteTimeValue("dueDateTime", m.GetDueDateTime())
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
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
    {
        err = writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("undoChangeId", m.GetUndoChangeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignee sets the assignee property value. The user identity the task is assigned to. Only present when the type property is assign. Nullable.
func (m *WorkbookDocumentTaskChange) SetAssignee(value WorkbookEmailIdentityable)() {
    err := m.GetBackingStore().Set("assignee", value)
    if err != nil {
        panic(err)
    }
}
// SetChangedBy sets the changedBy property value. The changedBy property
func (m *WorkbookDocumentTaskChange) SetChangedBy(value WorkbookEmailIdentityable)() {
    err := m.GetBackingStore().Set("changedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCommentId sets the commentId property value. The identifier of the associated comment.
func (m *WorkbookDocumentTaskChange) SetCommentId(value *string)() {
    err := m.GetBackingStore().Set("commentId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the task was changed. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WorkbookDocumentTaskChange) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDueDateTime sets the dueDateTime property value. The due date and time for the task. Only present when the type property is setSchedule. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WorkbookDocumentTaskChange) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("dueDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPercentComplete sets the percentComplete property value. An integer value from 0 to 100 that represents the percentage of the completion of the task and associated comment. 100 means that the task and associated comment are completed. If you change the completion from 100 to a lower value, the associated task and comment are reactivated. Only present when the type property is setPercentComplete. Nullable.
func (m *WorkbookDocumentTaskChange) SetPercentComplete(value *int32)() {
    err := m.GetBackingStore().Set("percentComplete", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. An integer value from 0 to 10 that represents the priority of the task. A lower value indicates a higher priority. 5 indicates the default priority if not set. Only present when the type property is setPriority. Nullable.
func (m *WorkbookDocumentTaskChange) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The start date and time for the task. Only present when the type property is setSchedule. Nullable. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WorkbookDocumentTaskChange) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. The title of the task. Only present when the type property is setTitle. Nullable.
func (m *WorkbookDocumentTaskChange) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. The type of the change history. Possible values are: create, assign, unassign, unassignAll, setPriority, setTitle, setPercentComplete, setSchedule, remove, restore, undo.
func (m *WorkbookDocumentTaskChange) SetTypeEscaped(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetUndoChangeId sets the undoChangeId property value. The ID of the workbookDocumentTaskChange that was undone for the undo change action. Only exists on an undo change history. Nullable.
func (m *WorkbookDocumentTaskChange) SetUndoChangeId(value *string)() {
    err := m.GetBackingStore().Set("undoChangeId", value)
    if err != nil {
        panic(err)
    }
}
type WorkbookDocumentTaskChangeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignee()(WorkbookEmailIdentityable)
    GetChangedBy()(WorkbookEmailIdentityable)
    GetCommentId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPercentComplete()(*int32)
    GetPriority()(*int32)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    GetUndoChangeId()(*string)
    SetAssignee(value WorkbookEmailIdentityable)()
    SetChangedBy(value WorkbookEmailIdentityable)()
    SetCommentId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPercentComplete(value *int32)()
    SetPriority(value *int32)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
    SetUndoChangeId(value *string)()
}
