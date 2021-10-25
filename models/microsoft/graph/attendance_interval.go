package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttendanceInterval struct {
    additionalData map[string]interface{};
    durationInSeconds *int32;
    joinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    leaveDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewAttendanceInterval()(*AttendanceInterval) {
    m := &AttendanceInterval{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttendanceInterval) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttendanceInterval) GetDurationInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInSeconds
    }
}
func (m *AttendanceInterval) GetJoinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.joinDateTime
    }
}
func (m *AttendanceInterval) GetLeaveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.leaveDateTime
    }
}
func (m *AttendanceInterval) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["durationInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDurationInSeconds(val)
        return nil
    }
    res["joinDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetJoinDateTime(val)
        return nil
    }
    res["leaveDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLeaveDateTime(val)
        return nil
    }
    return res
}
func (m *AttendanceInterval) IsNil()(bool) {
    return m == nil
}
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
func (m *AttendanceInterval) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttendanceInterval) SetDurationInSeconds(value *int32)() {
    m.durationInSeconds = value
}
func (m *AttendanceInterval) SetJoinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.joinDateTime = value
}
func (m *AttendanceInterval) SetLeaveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.leaveDateTime = value
}
