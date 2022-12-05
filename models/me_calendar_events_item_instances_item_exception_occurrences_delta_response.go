package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse provides operations to call the delta method.
type MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []Eventable
}
// NewMeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse instantiates a new MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewMeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse()(*MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse) {
    m := &MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateMeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse) GetValue()([]Eventable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeCalendarEventsItemInstancesItemExceptionOccurrencesDeltaResponse) SetValue(value []Eventable)() {
    m.value = value
}
