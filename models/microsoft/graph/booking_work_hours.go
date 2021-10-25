package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingWorkHours struct {
    additionalData map[string]interface{};
    day *DayOfWeek;
    timeSlots []BookingWorkTimeSlot;
}
func NewBookingWorkHours()(*BookingWorkHours) {
    m := &BookingWorkHours{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BookingWorkHours) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BookingWorkHours) GetDay()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.day
    }
}
func (m *BookingWorkHours) GetTimeSlots()([]BookingWorkTimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
func (m *BookingWorkHours) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["day"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        cast := val.(DayOfWeek)
        m.SetDay(&cast)
        return nil
    }
    res["timeSlots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingWorkTimeSlot() })
        if err != nil {
            return err
        }
        res := make([]BookingWorkTimeSlot, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingWorkTimeSlot))
        }
        m.SetTimeSlots(res)
        return nil
    }
    return res
}
func (m *BookingWorkHours) IsNil()(bool) {
    return m == nil
}
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
func (m *BookingWorkHours) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BookingWorkHours) SetDay(value *DayOfWeek)() {
    m.day = value
}
func (m *BookingWorkHours) SetTimeSlots(value []BookingWorkTimeSlot)() {
    m.timeSlots = value
}
