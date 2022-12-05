package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody provides operations to call the snoozeReminder method.
type UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The NewReminderTime property
    newReminderTime DateTimeTimeZoneable
}
// NewUsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody instantiates a new UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody and sets the default values.
func NewUsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody()(*UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody) {
    m := &UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody) GetNewReminderTime()(DateTimeTimeZoneable) {
    return m.newReminderTime
}
// Serialize serializes information the current object
func (m *UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewReminderTime sets the newReminderTime property value. The NewReminderTime property
func (m *UsersItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderPostRequestBody) SetNewReminderTime(value DateTimeTimeZoneable)() {
    m.newReminderTime = value
}
