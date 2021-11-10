package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BookingWorkHours struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The day of the week represented by this instance. Possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
    day *DayOfWeek;
    // A list of start/end times during a day.
    timeSlots []BookingWorkTimeSlot;
}
// Instantiates a new bookingWorkHours and sets the default values.
func NewBookingWorkHours()(*BookingWorkHours) {
    m := &BookingWorkHours{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkHours) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the day property value. The day of the week represented by this instance. Possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
func (m *BookingWorkHours) GetDay()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.day
    }
}
// Gets the timeSlots property value. A list of start/end times during a day.
func (m *BookingWorkHours) GetTimeSlots()([]BookingWorkTimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
// The deserialization information for the current model
func (m *BookingWorkHours) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["day"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DayOfWeek)
            m.SetDay(&cast)
        }
        return nil
    }
    res["timeSlots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingWorkTimeSlot() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkTimeSlot, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingWorkTimeSlot))
            }
            m.SetTimeSlots(res)
        }
        return nil
    }
    return res
}
func (m *BookingWorkHours) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BookingWorkHours) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDay() != nil {
        cast := m.GetDay().String()
        err := writer.WriteStringValue("day", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeSlots()))
        for i, v := range m.GetTimeSlots() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("timeSlots", cast)
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
func (m *BookingWorkHours) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the day property value. The day of the week represented by this instance. Possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
// Parameters:
//  - value : Value to set for the day property.
func (m *BookingWorkHours) SetDay(value *DayOfWeek)() {
    m.day = value
}
// Sets the timeSlots property value. A list of start/end times during a day.
// Parameters:
//  - value : Value to set for the timeSlots property.
func (m *BookingWorkHours) SetTimeSlots(value []BookingWorkTimeSlot)() {
    m.timeSlots = value
}
