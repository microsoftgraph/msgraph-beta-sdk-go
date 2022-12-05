package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody provides operations to call the decline method.
type GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Comment property
    comment *string
    // The ProposedNewTime property
    proposedNewTime TimeSlotable
    // The SendResponse property
    sendResponse *bool
}
// NewGroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody instantiates a new GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody and sets the default values.
func NewGroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody()(*GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) {
    m := &GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["proposedNewTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(TimeSlotable))
        }
        return nil
    }
    res["sendResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendResponse(val)
        }
        return nil
    }
    return res
}
// GetProposedNewTime gets the proposedNewTime property value. The ProposedNewTime property
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) GetProposedNewTime()(TimeSlotable) {
    return m.proposedNewTime
}
// GetSendResponse gets the sendResponse property value. The SendResponse property
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) GetSendResponse()(*bool) {
    return m.sendResponse
}
// Serialize serializes information the current object
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendResponse", m.GetSendResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetProposedNewTime sets the proposedNewTime property value. The ProposedNewTime property
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) SetProposedNewTime(value TimeSlotable)() {
    m.proposedNewTime = value
}
// SetSendResponse sets the sendResponse property value. The SendResponse property
func (m *GroupsItemCalendarEventsItemExceptionOccurrencesItemDeclinePostRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}
