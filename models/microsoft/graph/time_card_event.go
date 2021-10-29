package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new timeCardEvent and sets the default values.
func NewTimeCardEvent()(*TimeCardEvent) {
    m := &TimeCardEvent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEvent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the atApprovedLocation property value. Indicates whether the entry was recorded at the approved location.
func (m *TimeCardEvent) GetAtApprovedLocation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.atApprovedLocation
    }
}
// Gets the dateTime property value. The time the entry is recorded.
func (m *TimeCardEvent) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
// Gets the notes property value. Notes about the timeCardEvent.
func (m *TimeCardEvent) GetNotes()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// The deserialization information for the current model
func (m *TimeCardEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["atApprovedLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAtApprovedLocation(val)
        return nil
    }
    res["dateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDateTime(val)
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
    return res
}
func (m *TimeCardEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TimeCardEvent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the atApprovedLocation property value. Indicates whether the entry was recorded at the approved location.
// Parameters:
//  - value : Value to set for the atApprovedLocation property.
func (m *TimeCardEvent) SetAtApprovedLocation(value *bool)() {
    m.atApprovedLocation = value
}
// Sets the dateTime property value. The time the entry is recorded.
// Parameters:
//  - value : Value to set for the dateTime property.
func (m *TimeCardEvent) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// Sets the notes property value. Notes about the timeCardEvent.
// Parameters:
//  - value : Value to set for the notes property.
func (m *TimeCardEvent) SetNotes(value *ItemBody)() {
    m.notes = value
}
