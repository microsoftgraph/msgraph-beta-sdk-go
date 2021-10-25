package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingWorkTimeSlot struct {
    additionalData map[string]interface{};
    end *string;
    start *string;
}
func NewBookingWorkTimeSlot()(*BookingWorkTimeSlot) {
    m := &BookingWorkTimeSlot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BookingWorkTimeSlot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BookingWorkTimeSlot) GetEnd()(*string) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
func (m *BookingWorkTimeSlot) GetStart()(*string) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
func (m *BookingWorkTimeSlot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnd(val)
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStart(val)
        return nil
    }
    return res
}
func (m *BookingWorkTimeSlot) IsNil()(bool) {
    return m == nil
}
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
func (m *BookingWorkTimeSlot) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BookingWorkTimeSlot) SetEnd(value *string)() {
    m.end = value
}
func (m *BookingWorkTimeSlot) SetStart(value *string)() {
    m.start = value
}
