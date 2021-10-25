package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TimeCardEntry struct {
    additionalData map[string]interface{};
    breaks []TimeCardBreak;
    clockInEvent *TimeCardEvent;
    clockOutEvent *TimeCardEvent;
}
func NewTimeCardEntry()(*TimeCardEntry) {
    m := &TimeCardEntry{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TimeCardEntry) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TimeCardEntry) GetBreaks()([]TimeCardBreak) {
    if m == nil {
        return nil
    } else {
        return m.breaks
    }
}
func (m *TimeCardEntry) GetClockInEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockInEvent
    }
}
func (m *TimeCardEntry) GetClockOutEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockOutEvent
    }
}
func (m *TimeCardEntry) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["breaks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardBreak() })
        if err != nil {
            return err
        }
        res := make([]TimeCardBreak, len(val))
        for i, v := range val {
            res[i] = *(v.(*TimeCardBreak))
        }
        m.SetBreaks(res)
        return nil
    }
    res["clockInEvent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        m.SetClockInEvent(val.(*TimeCardEvent))
        return nil
    }
    res["clockOutEvent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        m.SetClockOutEvent(val.(*TimeCardEvent))
        return nil
    }
    return res
}
func (m *TimeCardEntry) IsNil()(bool) {
    return m == nil
}
func (m *TimeCardEntry) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBreaks()))
        for i, v := range m.GetBreaks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("breaks", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clockInEvent", m.GetClockInEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clockOutEvent", m.GetClockOutEvent())
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
func (m *TimeCardEntry) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TimeCardEntry) SetBreaks(value []TimeCardBreak)() {
    m.breaks = value
}
func (m *TimeCardEntry) SetClockInEvent(value *TimeCardEvent)() {
    m.clockInEvent = value
}
func (m *TimeCardEntry) SetClockOutEvent(value *TimeCardEvent)() {
    m.clockOutEvent = value
}
