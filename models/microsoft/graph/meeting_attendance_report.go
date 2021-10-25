package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MeetingAttendanceReport struct {
    Entity
    attendanceRecords []AttendanceRecord;
}
func NewMeetingAttendanceReport()(*MeetingAttendanceReport) {
    m := &MeetingAttendanceReport{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MeetingAttendanceReport) GetAttendanceRecords()([]AttendanceRecord) {
    if m == nil {
        return nil
    } else {
        return m.attendanceRecords
    }
}
func (m *MeetingAttendanceReport) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attendanceRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendanceRecord() })
        if err != nil {
            return err
        }
        res := make([]AttendanceRecord, len(val))
        for i, v := range val {
            res[i] = *(v.(*AttendanceRecord))
        }
        m.SetAttendanceRecords(res)
        return nil
    }
    return res
}
func (m *MeetingAttendanceReport) IsNil()(bool) {
    return m == nil
}
func (m *MeetingAttendanceReport) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendanceRecords()))
        for i, v := range m.GetAttendanceRecords() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attendanceRecords", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MeetingAttendanceReport) SetAttendanceRecords(value []AttendanceRecord)() {
    m.attendanceRecords = value
}
