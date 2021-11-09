package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttendanceInterval struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Duration of the meeting interval in seconds; that is, the difference between joinDateTime and leaveDateTime.
    durationInSeconds *int32;
    // Time attendee joined in UTC.
    joinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Time attendee left in UTC.
    leaveDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new attendanceInterval and sets the default values.
func NewAttendanceInterval()(*AttendanceInterval) {
    m := &AttendanceInterval{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttendanceInterval) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the durationInSeconds property value. Duration of the meeting interval in seconds; that is, the difference between joinDateTime and leaveDateTime.
func (m *AttendanceInterval) GetDurationInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInSeconds
    }
}
// Gets the joinDateTime property value. Time attendee joined in UTC.
func (m *AttendanceInterval) GetJoinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.joinDateTime
    }
}
// Gets the leaveDateTime property value. Time attendee left in UTC.
func (m *AttendanceInterval) GetLeaveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.leaveDateTime
    }
}
// The deserialization information for the current model
func (m *AttendanceInterval) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["durationInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInSeconds(val)
        }
        return nil
    }
    res["joinDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinDateTime(val)
        }
        return nil
    }
    res["leaveDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLeaveDateTime(val)
        }
        return nil
    }
    return res
}
func (m *AttendanceInterval) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttendanceInterval) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("durationInSeconds", m.GetDurationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("joinDateTime", m.GetJoinDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("leaveDateTime", m.GetLeaveDateTime())
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
func (m *AttendanceInterval) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the durationInSeconds property value. Duration of the meeting interval in seconds; that is, the difference between joinDateTime and leaveDateTime.
// Parameters:
//  - value : Value to set for the durationInSeconds property.
func (m *AttendanceInterval) SetDurationInSeconds(value *int32)() {
    m.durationInSeconds = value
}
// Sets the joinDateTime property value. Time attendee joined in UTC.
// Parameters:
//  - value : Value to set for the joinDateTime property.
func (m *AttendanceInterval) SetJoinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.joinDateTime = value
}
// Sets the leaveDateTime property value. Time attendee left in UTC.
// Parameters:
//  - value : Value to set for the leaveDateTime property.
func (m *AttendanceInterval) SetLeaveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.leaveDateTime = value
}
