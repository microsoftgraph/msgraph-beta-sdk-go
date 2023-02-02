package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody 
type ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The ProposedNewTime property
    proposedNewTime ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSlotable
    // The SendResponse property
    sendResponse *bool
}
// NewItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody instantiates a new ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody()(*ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) {
    m := &ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSlotable))
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
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) GetProposedNewTime()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSlotable) {
    return m.proposedNewTime
}
// GetSendResponse gets the sendResponse property value. The SendResponse property
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) GetSendResponse()(*bool) {
    return m.sendResponse
}
// Serialize serializes information the current object
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetProposedNewTime sets the proposedNewTime property value. The ProposedNewTime property
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) SetProposedNewTime(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSlotable)() {
    m.proposedNewTime = value
}
// SetSendResponse sets the sendResponse property value. The SendResponse property
func (m *ItemCalendarEventsItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}
