package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody 
type CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The NewReminderTime property
    newReminderTime ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable
}
// NewCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody instantiates a new CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody and sets the default values.
func NewCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody()(*CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) {
    m := &CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newReminderTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewReminderTime(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetNewReminderTime gets the newReminderTime property value. The NewReminderTime property
func (m *CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetNewReminderTime()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable) {
    return m.newReminderTime
}
// Serialize serializes information the current object
func (m *CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newReminderTime", m.GetNewReminderTime())
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
func (m *CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewReminderTime sets the newReminderTime property value. The NewReminderTime property
func (m *CalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) SetNewReminderTime(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable)() {
    m.newReminderTime = value
}
