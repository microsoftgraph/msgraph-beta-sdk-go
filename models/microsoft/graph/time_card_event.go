package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeCardEvent 
type TimeCardEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the entry was recorded at the approved location.
    atApprovedLocation *bool;
    // The time the entry is recorded.
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Notes about the timeCardEvent.
    notes *ItemBody;
}
// NewTimeCardEvent instantiates a new timeCardEvent and sets the default values.
func NewTimeCardEvent()(*TimeCardEvent) {
    m := &TimeCardEvent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEvent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAtApprovedLocation gets the atApprovedLocation property value. Indicates whether the entry was recorded at the approved location.
func (m *TimeCardEvent) GetAtApprovedLocation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.atApprovedLocation
    }
}
// GetDateTime gets the dateTime property value. The time the entry is recorded.
func (m *TimeCardEvent) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
// GetNotes gets the notes property value. Notes about the timeCardEvent.
func (m *TimeCardEvent) GetNotes()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCardEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["atApprovedLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtApprovedLocation(val)
        }
        return nil
    }
    res["dateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val)
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
    return res
}
func (m *TimeCardEvent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeCardEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("atApprovedLocation", m.GetAtApprovedLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("dateTime", m.GetDateTime())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEvent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAtApprovedLocation sets the atApprovedLocation property value. Indicates whether the entry was recorded at the approved location.
func (m *TimeCardEvent) SetAtApprovedLocation(value *bool)() {
    if m != nil {
        m.atApprovedLocation = value
    }
}
// SetDateTime sets the dateTime property value. The time the entry is recorded.
func (m *TimeCardEvent) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.dateTime = value
    }
}
// SetNotes sets the notes property value. Notes about the timeCardEvent.
func (m *TimeCardEvent) SetNotes(value *ItemBody)() {
    if m != nil {
        m.notes = value
    }
}
