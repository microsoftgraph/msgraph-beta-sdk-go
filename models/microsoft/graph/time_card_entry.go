package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCardEntry 
type TimeCardEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of breaks associated with the timeCard.
    breaks []TimeCardBreakable;
    // The clock-in event of the timeCard.
    clockInEvent TimeCardEventable;
    // The clock-out event of the timeCard.
    clockOutEvent TimeCardEventable;
}
// NewTimeCardEntry instantiates a new timeCardEntry and sets the default values.
func NewTimeCardEntry()(*TimeCardEntry) {
    m := &TimeCardEntry{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTimeCardEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardEntryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTimeCardEntry(), nil
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
func (m *TimeCardEntry) GetBreaks()([]TimeCardBreakable) {
    if m == nil {
        return nil
    } else {
        return m.breaks
    }
}
// GetClockInEvent gets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCardEntry) GetClockInEvent()(TimeCardEventable) {
    if m == nil {
        return nil
    } else {
        return m.clockInEvent
    }
}
// GetClockOutEvent gets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCardEntry) GetClockOutEvent()(TimeCardEventable) {
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
        val, err := n.GetCollectionOfObjectValues(CreateTimeCardBreakFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeCardBreakable, len(val))
            for i, v := range val {
                res[i] = v.(TimeCardBreakable)
            }
            m.SetBreaks(res)
        }
        return nil
    }
    res["clockInEvent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockInEvent(val.(TimeCardEventable))
        }
        return nil
    }
    res["clockOutEvent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockOutEvent(val.(TimeCardEventable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TimeCardEntry) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetBreaks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBreaks()))
        for i, v := range m.GetBreaks() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *TimeCardEntry) SetBreaks(value []TimeCardBreakable)() {
    if m != nil {
        m.breaks = value
    }
}
// SetClockInEvent sets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCardEntry) SetClockInEvent(value TimeCardEventable)() {
    if m != nil {
        m.clockInEvent = value
    }
}
// SetClockOutEvent sets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCardEntry) SetClockOutEvent(value TimeCardEventable)() {
    if m != nil {
        m.clockOutEvent = value
    }
}
