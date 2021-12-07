package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCardBreak 
type TimeCardBreak struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // ID of the timeCardBreak.
    breakId *string;
    // The start event of the timeCardBreak.
    end *TimeCardEvent;
    // Notes about the timeCardBreak.
    notes *ItemBody;
    // 
    start *TimeCardEvent;
}
// NewTimeCardBreak instantiates a new timeCardBreak and sets the default values.
func NewTimeCardBreak()(*TimeCardBreak) {
    m := &TimeCardBreak{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardBreak) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBreakId gets the breakId property value. ID of the timeCardBreak.
func (m *TimeCardBreak) GetBreakId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.breakId
    }
}
// GetEnd gets the end property value. The start event of the timeCardBreak.
func (m *TimeCardBreak) GetEnd()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// GetNotes gets the notes property value. Notes about the timeCardBreak.
func (m *TimeCardBreak) GetNotes()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetStart gets the start property value. 
func (m *TimeCardBreak) GetStart()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCardBreak) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["breakId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBreakId(val)
        }
        return nil
    }
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(*TimeCardEvent))
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(*ItemBody))
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(*TimeCardEvent))
        }
        return nil
    }
    return res
}
func (m *TimeCardBreak) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeCardBreak) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("breakId", m.GetBreakId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("start", m.GetStart())
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
func (m *TimeCardBreak) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBreakId sets the breakId property value. ID of the timeCardBreak.
func (m *TimeCardBreak) SetBreakId(value *string)() {
    if m != nil {
        m.breakId = value
    }
}
// SetEnd sets the end property value. The start event of the timeCardBreak.
func (m *TimeCardBreak) SetEnd(value *TimeCardEvent)() {
    if m != nil {
        m.end = value
    }
}
// SetNotes sets the notes property value. Notes about the timeCardBreak.
func (m *TimeCardBreak) SetNotes(value *ItemBody)() {
    if m != nil {
        m.notes = value
    }
}
// SetStart sets the start property value. 
func (m *TimeCardBreak) SetStart(value *TimeCardEvent)() {
    if m != nil {
        m.start = value
    }
}
