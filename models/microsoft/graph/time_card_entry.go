package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new timeCardEntry and sets the default values.
func NewTimeCardEntry()(*TimeCardEntry) {
    m := &TimeCardEntry{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEntry) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCardEntry) GetBreaks()([]TimeCardBreak) {
    if m == nil {
        return nil
    } else {
        return m.breaks
    }
}
// Gets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCardEntry) GetClockInEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockInEvent
    }
}
// Gets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCardEntry) GetClockOutEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockOutEvent
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TimeCardEntry) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the breaks property value. The list of breaks associated with the timeCard.
// Parameters:
//  - value : Value to set for the breaks property.
func (m *TimeCardEntry) SetBreaks(value []TimeCardBreak)() {
    m.breaks = value
}
// Sets the clockInEvent property value. The clock-in event of the timeCard.
// Parameters:
//  - value : Value to set for the clockInEvent property.
func (m *TimeCardEntry) SetClockInEvent(value *TimeCardEvent)() {
    m.clockInEvent = value
}
// Sets the clockOutEvent property value. The clock-out event of the timeCard.
// Parameters:
//  - value : Value to set for the clockOutEvent property.
func (m *TimeCardEntry) SetClockOutEvent(value *TimeCardEvent)() {
    m.clockOutEvent = value
}
