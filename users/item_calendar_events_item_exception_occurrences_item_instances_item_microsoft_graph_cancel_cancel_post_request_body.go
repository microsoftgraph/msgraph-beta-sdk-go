package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody 
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
}
// NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody instantiates a new ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody()(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody) {
    m := &ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
