package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TimeCardBreak struct {
    additionalData map[string]interface{};
    breakId *string;
    end *TimeCardEvent;
    notes *ItemBody;
    start *TimeCardEvent;
}
func NewTimeCardBreak()(*TimeCardBreak) {
    m := &TimeCardBreak{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TimeCardBreak) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TimeCardBreak) GetBreakId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.breakId
    }
}
func (m *TimeCardBreak) GetEnd()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
func (m *TimeCardBreak) GetNotes()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
func (m *TimeCardBreak) GetStart()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
func (m *TimeCardBreak) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["breakId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBreakId(val)
        return nil
    }
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        m.SetEnd(val.(*TimeCardEvent))
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetNotes(val.(*ItemBody))
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEvent() })
        if err != nil {
            return err
        }
        m.SetStart(val.(*TimeCardEvent))
        return nil
    }
    return res
}
func (m *TimeCardBreak) IsNil()(bool) {
    return m == nil
}
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
func (m *TimeCardBreak) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TimeCardBreak) SetBreakId(value *string)() {
    m.breakId = value
}
func (m *TimeCardBreak) SetEnd(value *TimeCardEvent)() {
    m.end = value
}
func (m *TimeCardBreak) SetNotes(value *ItemBody)() {
    m.notes = value
}
func (m *TimeCardBreak) SetStart(value *TimeCardEvent)() {
    m.start = value
}
