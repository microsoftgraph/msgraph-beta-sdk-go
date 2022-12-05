package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody provides operations to call the snoozeReminder method.
type MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The NewReminderTime property
    newReminderTime DateTimeTimeZoneable
}
// NewMeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody instantiates a new MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody and sets the default values.
func NewMeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody()(*MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody) {
    m := &MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newReminderTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewReminderTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetNewReminderTime gets the newReminderTime property value. The NewReminderTime property
func (m *MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody) GetNewReminderTime()(DateTimeTimeZoneable) {
    return m.newReminderTime
}
// Serialize serializes information the current object
func (m *MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewReminderTime sets the newReminderTime property value. The NewReminderTime property
func (m *MeCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBody) SetNewReminderTime(value DateTimeTimeZoneable)() {
    m.newReminderTime = value
}
