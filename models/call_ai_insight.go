package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CallAiInsight struct {
    Entity
}
// NewCallAiInsight instantiates a new CallAiInsight and sets the default values.
func NewCallAiInsight()(*CallAiInsight) {
    m := &CallAiInsight{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallAiInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCallAiInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallAiInsight(), nil
}
// GetActionItems gets the actionItems property value. The actionItems property
// returns a []ActionItemable when successful
func (m *CallAiInsight) GetActionItems()([]ActionItemable) {
    val, err := m.GetBackingStore().Get("actionItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActionItemable)
    }
    return nil
}
// GetCallId gets the callId property value. The callId property
// returns a *string when successful
func (m *CallAiInsight) GetCallId()(*string) {
    val, err := m.GetBackingStore().Get("callId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContentCorrelationId gets the contentCorrelationId property value. The contentCorrelationId property
// returns a *string when successful
func (m *CallAiInsight) GetContentCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("contentCorrelationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
// returns a *Time when successful
func (m *CallAiInsight) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
// returns a *Time when successful
func (m *CallAiInsight) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
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
func (m *CallAiInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActionItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActionItemable)
                }
            }
            m.SetActionItems(res)
        }
        return nil
    }
    res["callId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallId(val)
        }
        return nil
    }
    res["contentCorrelationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCorrelationId(val)
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["meetingNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingNoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingNoteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MeetingNoteable)
                }
            }
            m.SetMeetingNotes(res)
        }
        return nil
    }
    res["viewpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCallAiInsightViewPointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewpoint(val.(CallAiInsightViewPointable))
        }
        return nil
    }
    return res
}
// GetMeetingNotes gets the meetingNotes property value. The meetingNotes property
// returns a []MeetingNoteable when successful
func (m *CallAiInsight) GetMeetingNotes()([]MeetingNoteable) {
    val, err := m.GetBackingStore().Get("meetingNotes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MeetingNoteable)
    }
    return nil
}
// GetViewpoint gets the viewpoint property value. The viewpoint property
// returns a CallAiInsightViewPointable when successful
func (m *CallAiInsight) GetViewpoint()(CallAiInsightViewPointable) {
    val, err := m.GetBackingStore().Get("viewpoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CallAiInsightViewPointable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallAiInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActionItems()))
        for i, v := range m.GetActionItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("actionItems", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callId", m.GetCallId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentCorrelationId", m.GetContentCorrelationId())
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
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMeetingNotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMeetingNotes()))
        for i, v := range m.GetMeetingNotes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("meetingNotes", cast)
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
// SetActionItems sets the actionItems property value. The actionItems property
func (m *CallAiInsight) SetActionItems(value []ActionItemable)() {
    err := m.GetBackingStore().Set("actionItems", value)
    if err != nil {
        panic(err)
    }
}
// SetCallId sets the callId property value. The callId property
func (m *CallAiInsight) SetCallId(value *string)() {
    err := m.GetBackingStore().Set("callId", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCorrelationId sets the contentCorrelationId property value. The contentCorrelationId property
func (m *CallAiInsight) SetContentCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("contentCorrelationId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *CallAiInsight) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *CallAiInsight) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMeetingNotes sets the meetingNotes property value. The meetingNotes property
func (m *CallAiInsight) SetMeetingNotes(value []MeetingNoteable)() {
    err := m.GetBackingStore().Set("meetingNotes", value)
    if err != nil {
        panic(err)
    }
}
// SetViewpoint sets the viewpoint property value. The viewpoint property
func (m *CallAiInsight) SetViewpoint(value CallAiInsightViewPointable)() {
    err := m.GetBackingStore().Set("viewpoint", value)
    if err != nil {
        panic(err)
    }
}
type CallAiInsightable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionItems()([]ActionItemable)
    GetCallId()(*string)
    GetContentCorrelationId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMeetingNotes()([]MeetingNoteable)
    GetViewpoint()(CallAiInsightViewPointable)
    SetActionItems(value []ActionItemable)()
    SetCallId(value *string)()
    SetContentCorrelationId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMeetingNotes(value []MeetingNoteable)()
    SetViewpoint(value CallAiInsightViewPointable)()
}
