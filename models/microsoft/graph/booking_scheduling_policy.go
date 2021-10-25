package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingSchedulingPolicy struct {
    additionalData map[string]interface{};
    allowStaffSelection *bool;
    maximumAdvance *string;
    minimumLeadTime *string;
    sendConfirmationsToOwner *bool;
    timeSlotInterval *string;
}
func NewBookingSchedulingPolicy()(*BookingSchedulingPolicy) {
    m := &BookingSchedulingPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BookingSchedulingPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BookingSchedulingPolicy) GetAllowStaffSelection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowStaffSelection
    }
}
func (m *BookingSchedulingPolicy) GetMaximumAdvance()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumAdvance
    }
}
func (m *BookingSchedulingPolicy) GetMinimumLeadTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumLeadTime
    }
}
func (m *BookingSchedulingPolicy) GetSendConfirmationsToOwner()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendConfirmationsToOwner
    }
}
func (m *BookingSchedulingPolicy) GetTimeSlotInterval()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeSlotInterval
    }
}
func (m *BookingSchedulingPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowStaffSelection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowStaffSelection(val)
        return nil
    }
    res["maximumAdvance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaximumAdvance(val)
        return nil
    }
    res["minimumLeadTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumLeadTime(val)
        return nil
    }
    res["sendConfirmationsToOwner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSendConfirmationsToOwner(val)
        return nil
    }
    res["timeSlotInterval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTimeSlotInterval(val)
        return nil
    }
    return res
}
func (m *BookingSchedulingPolicy) IsNil()(bool) {
    return m == nil
}
func (m *BookingSchedulingPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowStaffSelection", m.GetAllowStaffSelection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maximumAdvance", m.GetMaximumAdvance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("minimumLeadTime", m.GetMinimumLeadTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendConfirmationsToOwner", m.GetSendConfirmationsToOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeSlotInterval", m.GetTimeSlotInterval())
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
func (m *BookingSchedulingPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BookingSchedulingPolicy) SetAllowStaffSelection(value *bool)() {
    m.allowStaffSelection = value
}
func (m *BookingSchedulingPolicy) SetMaximumAdvance(value *string)() {
    m.maximumAdvance = value
}
func (m *BookingSchedulingPolicy) SetMinimumLeadTime(value *string)() {
    m.minimumLeadTime = value
}
func (m *BookingSchedulingPolicy) SetSendConfirmationsToOwner(value *bool)() {
    m.sendConfirmationsToOwner = value
}
func (m *BookingSchedulingPolicy) SetTimeSlotInterval(value *string)() {
    m.timeSlotInterval = value
}
