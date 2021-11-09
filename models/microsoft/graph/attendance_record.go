package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttendanceRecord struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of time periods between joining and leaving.
    attendanceIntervals []AttendanceInterval;
    // Email address.
    emailAddress *string;
    // Identifier, such as display name.
    identity *Identity;
    // Role of the attendee. Possible values are None, Attendee, Presenter, and Organizer.
    role *string;
    // Total duration of the attendances in seconds.
    totalAttendanceInSeconds *int32;
}
// Instantiates a new attendanceRecord and sets the default values.
func NewAttendanceRecord()(*AttendanceRecord) {
    m := &AttendanceRecord{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttendanceRecord) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attendanceIntervals property value. List of time periods between joining and leaving.
func (m *AttendanceRecord) GetAttendanceIntervals()([]AttendanceInterval) {
    if m == nil {
        return nil
    } else {
        return m.attendanceIntervals
    }
}
// Gets the emailAddress property value. Email address.
func (m *AttendanceRecord) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// Gets the identity property value. Identifier, such as display name.
func (m *AttendanceRecord) GetIdentity()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
// Gets the role property value. Role of the attendee. Possible values are None, Attendee, Presenter, and Organizer.
func (m *AttendanceRecord) GetRole()(*string) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// Gets the totalAttendanceInSeconds property value. Total duration of the attendances in seconds.
func (m *AttendanceRecord) GetTotalAttendanceInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalAttendanceInSeconds
    }
}
// The deserialization information for the current model
func (m *AttendanceRecord) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendanceIntervals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendanceInterval() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttendanceInterval, len(val))
            for i, v := range val {
                res[i] = *(v.(*AttendanceInterval))
            }
            m.SetAttendanceIntervals(res)
        }
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(*Identity))
        }
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["totalAttendanceInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAttendanceInSeconds(val)
        }
        return nil
    }
    return res
}
func (m *AttendanceRecord) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AttendanceRecord) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attendanceIntervals property value. List of time periods between joining and leaving.
// Parameters:
//  - value : Value to set for the attendanceIntervals property.
func (m *AttendanceRecord) SetAttendanceIntervals(value []AttendanceInterval)() {
    m.attendanceIntervals = value
}
// Sets the emailAddress property value. Email address.
// Parameters:
//  - value : Value to set for the emailAddress property.
func (m *AttendanceRecord) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// Sets the identity property value. Identifier, such as display name.
// Parameters:
//  - value : Value to set for the identity property.
func (m *AttendanceRecord) SetIdentity(value *Identity)() {
    m.identity = value
}
// Sets the role property value. Role of the attendee. Possible values are None, Attendee, Presenter, and Organizer.
// Parameters:
//  - value : Value to set for the role property.
func (m *AttendanceRecord) SetRole(value *string)() {
    m.role = value
}
// Sets the totalAttendanceInSeconds property value. Total duration of the attendances in seconds.
// Parameters:
//  - value : Value to set for the totalAttendanceInSeconds property.
func (m *AttendanceRecord) SetTotalAttendanceInSeconds(value *int32)() {
    m.totalAttendanceInSeconds = value
}
