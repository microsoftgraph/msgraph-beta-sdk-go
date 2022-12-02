package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingBusinessesItemGetStaffAvailabilityPostRequestBody provides operations to call the getStaffAvailability method.
type BookingBusinessesItemGetStaffAvailabilityPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The endDateTime property
    endDateTime DateTimeTimeZoneable
    // The staffIds property
    staffIds []string
    // The startDateTime property
    startDateTime DateTimeTimeZoneable
}
// NewBookingBusinessesItemGetStaffAvailabilityPostRequestBody instantiates a new BookingBusinessesItemGetStaffAvailabilityPostRequestBody and sets the default values.
func NewBookingBusinessesItemGetStaffAvailabilityPostRequestBody()(*BookingBusinessesItemGetStaffAvailabilityPostRequestBody) {
    m := &BookingBusinessesItemGetStaffAvailabilityPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBookingBusinessesItemGetStaffAvailabilityPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingBusinessesItemGetStaffAvailabilityPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingBusinessesItemGetStaffAvailabilityPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetEndDateTime()(DateTimeTimeZoneable) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["staffIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStaffIds(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetStaffIds gets the staffIds property value. The staffIds property
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetStaffIds()([]string) {
    return m.staffIds
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStaffIds() != nil {
        err := writer.WriteCollectionOfStringValues("staffIds", m.GetStaffIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
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
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetEndDateTime(value DateTimeTimeZoneable)() {
    m.endDateTime = value
}
// SetStaffIds sets the staffIds property value. The staffIds property
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetStaffIds(value []string)() {
    m.staffIds = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *BookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
