package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TimeCard struct {
    ChangeTrackedEntity
    breaks []TimeCardBreak;
    clockInEvent *TimeCardEvent;
    clockOutEvent *TimeCardEvent;
    confirmedBy *ConfirmedBy;
    notes *ItemBody;
    originalEntry *TimeCardEntry;
    state *TimeCardState;
    userId *string;
}
func NewTimeCard()(*TimeCard) {
    m := &TimeCard{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
func (m *TimeCard) GetBreaks()([]TimeCardBreak) {
    if m == nil {
        return nil
    } else {
        return m.breaks
    }
}
func (m *TimeCard) GetClockInEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockInEvent
    }
}
func (m *TimeCard) GetClockOutEvent()(*TimeCardEvent) {
    if m == nil {
        return nil
    } else {
        return m.clockOutEvent
    }
}
func (m *TimeCard) GetConfirmedBy()(*ConfirmedBy) {
    if m == nil {
        return nil
    } else {
        return m.confirmedBy
    }
}
func (m *TimeCard) GetNotes()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
func (m *TimeCard) GetOriginalEntry()(*TimeCardEntry) {
    if m == nil {
        return nil
    } else {
        return m.originalEntry
    }
}
func (m *TimeCard) GetState()(*TimeCardState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *TimeCard) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *TimeCard) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
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
    res["confirmedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfirmedBy)
        if err != nil {
            return err
        }
        cast := val.(ConfirmedBy)
        m.SetConfirmedBy(&cast)
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
    res["originalEntry"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeCardEntry() })
        if err != nil {
            return err
        }
        m.SetOriginalEntry(val.(*TimeCardEntry))
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTimeCardState)
        if err != nil {
            return err
        }
        cast := val.(TimeCardState)
        m.SetState(&cast)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *TimeCard) IsNil()(bool) {
    return m == nil
}
func (m *TimeCard) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBreaks()))
        for i, v := range m.GetBreaks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("breaks", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clockInEvent", m.GetClockInEvent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clockOutEvent", m.GetClockOutEvent())
        if err != nil {
            return err
        }
    }
    if m.GetConfirmedBy() != nil {
        cast := m.GetConfirmedBy().String()
        err = writer.WriteStringValue("confirmedBy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("originalEntry", m.GetOriginalEntry())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TimeCard) SetBreaks(value []TimeCardBreak)() {
    m.breaks = value
}
func (m *TimeCard) SetClockInEvent(value *TimeCardEvent)() {
    m.clockInEvent = value
}
func (m *TimeCard) SetClockOutEvent(value *TimeCardEvent)() {
    m.clockOutEvent = value
}
func (m *TimeCard) SetConfirmedBy(value *ConfirmedBy)() {
    m.confirmedBy = value
}
func (m *TimeCard) SetNotes(value *ItemBody)() {
    m.notes = value
}
func (m *TimeCard) SetOriginalEntry(value *TimeCardEntry)() {
    m.originalEntry = value
}
func (m *TimeCard) SetState(value *TimeCardState)() {
    m.state = value
}
func (m *TimeCard) SetUserId(value *string)() {
    m.userId = value
}
