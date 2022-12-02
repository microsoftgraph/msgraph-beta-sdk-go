package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody provides operations to call the cancel method.
type MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Comment property
    comment *string
}
// NewMeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody instantiates a new MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody()(*MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody) {
    m := &MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
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
func (m *MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *MeCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
