package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BookingWorkTimeSlot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The time of the day when work stops. For example, 17:00:00.0000000.
    end *string;
    // The time of the day when work starts. For example, 08:00:00.0000000.
    start *string;
}
// Instantiates a new bookingWorkTimeSlot and sets the default values.
func NewBookingWorkTimeSlot()(*BookingWorkTimeSlot) {
    m := &BookingWorkTimeSlot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkTimeSlot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the end property value. The time of the day when work stops. For example, 17:00:00.0000000.
func (m *BookingWorkTimeSlot) GetEnd()(*string) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// Gets the start property value. The time of the day when work starts. For example, 08:00:00.0000000.
func (m *BookingWorkTimeSlot) GetStart()(*string) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// The deserialization information for the current model
func (m *BookingWorkTimeSlot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    return res
}
func (m *BookingWorkTimeSlot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BookingWorkTimeSlot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("start", m.GetStart())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *BookingWorkTimeSlot) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the end property value. The time of the day when work stops. For example, 17:00:00.0000000.
// Parameters:
//  - value : Value to set for the end property.
func (m *BookingWorkTimeSlot) SetEnd(value *string)() {
    m.end = value
}
// Sets the start property value. The time of the day when work starts. For example, 08:00:00.0000000.
// Parameters:
//  - value : Value to set for the start property.
func (m *BookingWorkTimeSlot) SetStart(value *string)() {
    m.start = value
}
