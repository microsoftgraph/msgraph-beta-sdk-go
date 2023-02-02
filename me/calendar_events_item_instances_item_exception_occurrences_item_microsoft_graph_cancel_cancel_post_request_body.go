package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody 
type CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
}
// NewCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody instantiates a new CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody and sets the default values.
func NewCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody()(*CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody) {
    m := &CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *CalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelCancelPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
