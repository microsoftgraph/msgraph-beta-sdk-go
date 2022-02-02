package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCardEntry 
type TimeCardEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of breaks associated with the timeCard.
    breaks []TimeCardBreak;
    // The clock-in event of the timeCard.
    clockInEvent *TimeCardEvent;
    // The clock-out event of the timeCard.
    clockOutEvent *TimeCardEvent;
}
// NewTimeCardEntry instantiates a new timeCardEntry and sets the default values.
func NewTimeCardEntry()(*TimeCardEntry) {
    m := &TimeCardEntry{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEntry) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBreaks gets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCardEntry) GetBreaks()([]TimeCardBreak) {
    if m == nil {
        return nil
    } else {
        return m.breaks
    }
}
// GetClockInEvent gets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCardEntry) GetClockInEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockInEvent
    }
}
// GetClockOutEvent gets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCardEntry) GetClockOutEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockOutEvent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCardEntry) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["breaks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardBreak() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeCardBreak, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeCardBreak))
            }
            m.SetBreaks(res)
        }
        return nil
    }
    res["clockInEvent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockInEvent(val.(*TimeCardEvent))
        }
        return nil
    }
    res["clockOutEvent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockOutEvent(val.(*TimeCardEvent))
        }
        return nil
    }
    return res
}
func (m *TimeCardEntry) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeCardEntry) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetBreaks() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEntry) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBreaks sets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCardEntry) SetBreaks(value []TimeCardBreak)() {
    if m != nil {
        m.breaks = value
    }
}
// SetClockInEvent sets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCardEntry) SetClockInEvent(value *TimeCardEvent)() {
    if m != nil {
        m.clockInEvent = value
    }
}
// SetClockOutEvent sets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCardEntry) SetClockOutEvent(value *TimeCardEvent)() {
    if m != nil {
        m.clockOutEvent = value
    }
}
