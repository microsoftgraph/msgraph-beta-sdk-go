package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallTranscriptEventMessageDetail 
type CallTranscriptEventMessageDetail struct {
    EventMessageDetail
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Unique identifier of the call.
    callId *string
    // Unique identifier for a call transcript.
    callTranscriptICalUid *string
    // The organizer of the meeting.
    meetingOrganizer IdentitySetable
}
// NewCallTranscriptEventMessageDetail instantiates a new CallTranscriptEventMessageDetail and sets the default values.
func NewCallTranscriptEventMessageDetail()(*CallTranscriptEventMessageDetail) {
    m := &CallTranscriptEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCallTranscriptEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallTranscriptEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallTranscriptEventMessageDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallTranscriptEventMessageDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCallId gets the callId property value. Unique identifier of the call.
func (m *CallTranscriptEventMessageDetail) GetCallId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callId
    }
}
// GetCallTranscriptICalUid gets the callTranscriptICalUid property value. Unique identifier for a call transcript.
func (m *CallTranscriptEventMessageDetail) GetCallTranscriptICalUid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callTranscriptICalUid
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallTranscriptEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
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
    res["callTranscriptICalUid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallTranscriptICalUid(val)
        }
        return nil
    }
    res["meetingOrganizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingOrganizer(val.(IdentitySetable))
        }
        return nil
    }
    return res
}
// GetMeetingOrganizer gets the meetingOrganizer property value. The organizer of the meeting.
func (m *CallTranscriptEventMessageDetail) GetMeetingOrganizer()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.meetingOrganizer
    }
}
// Serialize serializes information the current object
func (m *CallTranscriptEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("callId", m.GetCallId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callTranscriptICalUid", m.GetCallTranscriptICalUid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meetingOrganizer", m.GetMeetingOrganizer())
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
func (m *CallTranscriptEventMessageDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCallId sets the callId property value. Unique identifier of the call.
func (m *CallTranscriptEventMessageDetail) SetCallId(value *string)() {
    if m != nil {
        m.callId = value
    }
}
// SetCallTranscriptICalUid sets the callTranscriptICalUid property value. Unique identifier for a call transcript.
func (m *CallTranscriptEventMessageDetail) SetCallTranscriptICalUid(value *string)() {
    if m != nil {
        m.callTranscriptICalUid = value
    }
}
// SetMeetingOrganizer sets the meetingOrganizer property value. The organizer of the meeting.
func (m *CallTranscriptEventMessageDetail) SetMeetingOrganizer(value IdentitySetable)() {
    if m != nil {
        m.meetingOrganizer = value
    }
}
