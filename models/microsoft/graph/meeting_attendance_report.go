package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingAttendanceReport 
type MeetingAttendanceReport struct {
    Entity
    // The list of attendance records.
    attendanceRecords []AttendanceRecord;
    // Total number of participants.
    totalParticipantCount *int32;
}
// NewMeetingAttendanceReport instantiates a new meetingAttendanceReport and sets the default values.
func NewMeetingAttendanceReport()(*MeetingAttendanceReport) {
    m := &MeetingAttendanceReport{
        Entity: *NewEntity(),
    }
    return m
}
// GetAttendanceRecords gets the attendanceRecords property value. The list of attendance records.
func (m *MeetingAttendanceReport) GetAttendanceRecords()([]AttendanceRecord) {
    if m == nil {
        return nil
    } else {
        return m.attendanceRecords
    }
}
// GetTotalParticipantCount gets the totalParticipantCount property value. Total number of participants.
func (m *MeetingAttendanceReport) GetTotalParticipantCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalParticipantCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingAttendanceReport) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attendanceRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendanceRecord() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttendanceRecord, len(val))
            for i, v := range val {
                res[i] = *(v.(*AttendanceRecord))
            }
            m.SetAttendanceRecords(res)
        }
        return nil
    }
    res["totalParticipantCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalParticipantCount(val)
        }
        return nil
    }
    return res
}
func (m *MeetingAttendanceReport) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    {
        err = writer.WriteInt32Value("totalParticipantCount", m.GetTotalParticipantCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttendanceRecords sets the attendanceRecords property value. The list of attendance records.
func (m *MeetingAttendanceReport) SetAttendanceRecords(value []AttendanceRecord)() {
    m.attendanceRecords = value
}
// SetTotalParticipantCount sets the totalParticipantCount property value. Total number of participants.
func (m *MeetingAttendanceReport) SetTotalParticipantCount(value *int32)() {
    m.totalParticipantCount = value
}
