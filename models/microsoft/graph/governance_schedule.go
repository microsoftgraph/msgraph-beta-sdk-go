package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceSchedule 
type GovernanceSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The duration of a role assignment. It is in format of a TimeSpan.
    duration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The end time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Note: if the value is null, it indicates a permanent assignment.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The role assignment schedule type. Only Once is supported for now.
    type_escaped *string;
}
// NewGovernanceSchedule instantiates a new governanceSchedule and sets the default values.
func NewGovernanceSchedule()(*GovernanceSchedule) {
    m := &GovernanceSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceSchedule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDuration gets the duration property value. The duration of a role assignment. It is in format of a TimeSpan.
func (m *GovernanceSchedule) GetDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetEndDateTime gets the endDateTime property value. The end time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Note: if the value is null, it indicates a permanent assignment.
func (m *GovernanceSchedule) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetStartDateTime gets the startDateTime property value. The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetType gets the type property value. The role assignment schedule type. Only Once is supported for now.
func (m *GovernanceSchedule) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *GovernanceSchedule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernanceSchedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *GovernanceSchedule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDuration sets the duration property value. The duration of a role assignment. It is in format of a TimeSpan.
func (m *GovernanceSchedule) SetDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.duration = value
    }
}
// SetEndDateTime sets the endDateTime property value. The end time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Note: if the value is null, it indicates a permanent assignment.
func (m *GovernanceSchedule) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetStartDateTime sets the startDateTime property value. The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetType sets the type property value. The role assignment schedule type. Only Once is supported for now.
func (m *GovernanceSchedule) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
