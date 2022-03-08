package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCard provides operations to manage the compliance singleton.
type TimeCard struct {
    ChangeTrackedEntity
    // The list of breaks associated with the timeCard.
    breaks []TimeCardBreakable;
    // The clock-in event of the timeCard.
    clockInEvent TimeCardEventable;
    // The clock-out event of the timeCard.
    clockOutEvent TimeCardEventable;
    // Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
    confirmedBy *ConfirmedBy;
    // Notes about the timeCard.
    notes ItemBodyable;
    // The original timeCardEntry of the timeCard, before user edits.
    originalEntry TimeCardEntryable;
    // The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
    state *TimeCardState;
    // User ID to which  the timeCard belongs.
    userId *string;
}
// NewTimeCard instantiates a new timeCard and sets the default values.
func NewTimeCard()(*TimeCard) {
    m := &TimeCard{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// CreateTimeCardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTimeCard(), nil
}
// GetBreaks gets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCard) GetBreaks()([]TimeCardBreakable) {
    if m == nil {
        return nil
    } else {
        return m.breaks
    }
}
// GetClockInEvent gets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCard) GetClockInEvent()(TimeCardEventable) {
    if m == nil {
        return nil
    } else {
        return m.clockInEvent
    }
}
// GetClockOutEvent gets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCard) GetClockOutEvent()(TimeCardEventable) {
    if m == nil {
        return nil
    } else {
        return m.clockOutEvent
    }
}
// GetConfirmedBy gets the confirmedBy property value. Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
func (m *TimeCard) GetConfirmedBy()(*ConfirmedBy) {
    if m == nil {
        return nil
    } else {
        return m.confirmedBy
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCard) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
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
    res["confirmedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfirmedBy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfirmedBy(val.(*ConfirmedBy))
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    res["originalEntry"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalEntry(val.(TimeCardEntryable))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTimeCardState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*TimeCardState))
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. Notes about the timeCard.
func (m *TimeCard) GetNotes()(ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetOriginalEntry gets the originalEntry property value. The original timeCardEntry of the timeCard, before user edits.
func (m *TimeCard) GetOriginalEntry()(TimeCardEntryable) {
    if m == nil {
        return nil
    } else {
        return m.originalEntry
    }
}
// GetState gets the state property value. The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
func (m *TimeCard) GetState()(*TimeCardState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUserId gets the userId property value. User ID to which  the timeCard belongs.
func (m *TimeCard) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *TimeCard) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeCard) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBreaks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBreaks()))
        for i, v := range m.GetBreaks() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := (*m.GetConfirmedBy()).String()
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
        cast := (*m.GetState()).String()
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
// SetBreaks sets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCard) SetBreaks(value []TimeCardBreakable)() {
    if m != nil {
        m.breaks = value
    }
}
// SetClockInEvent sets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCard) SetClockInEvent(value TimeCardEventable)() {
    if m != nil {
        m.clockInEvent = value
    }
}
// SetClockOutEvent sets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCard) SetClockOutEvent(value TimeCardEventable)() {
    if m != nil {
        m.clockOutEvent = value
    }
}
// SetConfirmedBy sets the confirmedBy property value. Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
func (m *TimeCard) SetConfirmedBy(value *ConfirmedBy)() {
    if m != nil {
        m.confirmedBy = value
    }
}
// SetNotes sets the notes property value. Notes about the timeCard.
func (m *TimeCard) SetNotes(value ItemBodyable)() {
    if m != nil {
        m.notes = value
    }
}
// SetOriginalEntry sets the originalEntry property value. The original timeCardEntry of the timeCard, before user edits.
func (m *TimeCard) SetOriginalEntry(value TimeCardEntryable)() {
    if m != nil {
        m.originalEntry = value
    }
}
// SetState sets the state property value. The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
func (m *TimeCard) SetState(value *TimeCardState)() {
    if m != nil {
        m.state = value
    }
}
// SetUserId sets the userId property value. User ID to which  the timeCard belongs.
func (m *TimeCard) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
