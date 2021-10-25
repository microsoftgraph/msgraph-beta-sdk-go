package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttendanceRecord struct {
    additionalData map[string]interface{};
    attendanceIntervals []AttendanceInterval;
    emailAddress *string;
    identity *Identity;
    role *string;
    totalAttendanceInSeconds *int32;
}
func NewAttendanceRecord()(*AttendanceRecord) {
    m := &AttendanceRecord{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttendanceRecord) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttendanceRecord) GetAttendanceIntervals()([]AttendanceInterval) {
    if m == nil {
        return nil
    } else {
        return m.attendanceIntervals
    }
}
func (m *AttendanceRecord) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
func (m *AttendanceRecord) GetIdentity()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
func (m *AttendanceRecord) GetRole()(*string) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
func (m *AttendanceRecord) GetTotalAttendanceInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalAttendanceInSeconds
    }
}
func (m *AttendanceRecord) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendanceIntervals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendanceInterval() })
        if err != nil {
            return err
        }
        res := make([]AttendanceInterval, len(val))
        for i, v := range val {
            res[i] = *(v.(*AttendanceInterval))
        }
        m.SetAttendanceIntervals(res)
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmailAddress(val)
        return nil
    }
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetIdentity(val.(*Identity))
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRole(val)
        return nil
    }
    res["totalAttendanceInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalAttendanceInSeconds(val)
        return nil
    }
    return res
}
func (m *AttendanceRecord) IsNil()(bool) {
    return m == nil
}
func (m *AttendanceRecord) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendanceIntervals()))
        for i, v := range m.GetAttendanceIntervals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("attendanceIntervals", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalAttendanceInSeconds", m.GetTotalAttendanceInSeconds())
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
func (m *AttendanceRecord) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttendanceRecord) SetAttendanceIntervals(value []AttendanceInterval)() {
    m.attendanceIntervals = value
}
func (m *AttendanceRecord) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
func (m *AttendanceRecord) SetIdentity(value *Identity)() {
    m.identity = value
}
func (m *AttendanceRecord) SetRole(value *string)() {
    m.role = value
}
func (m *AttendanceRecord) SetTotalAttendanceInSeconds(value *int32)() {
    m.totalAttendanceInSeconds = value
}
