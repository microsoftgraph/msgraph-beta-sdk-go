package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse provides operations to call the delta method.
type UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []Eventable
}
// NewUsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse instantiates a new UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewUsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse()(*UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse) {
    m := &UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateUsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                res[i] = v.(Eventable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse) GetValue()([]Eventable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UsersItemCalendarsItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse) SetValue(value []Eventable)() {
    m.value = value
}
