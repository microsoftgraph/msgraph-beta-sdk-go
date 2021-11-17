package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BookingStaffMember struct {
    BookingPerson
    // True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
    availabilityIsAffectedByPersonalCalendar *bool;
    // Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
    colorIndex *int32;
    // The role of the staff member in the business. Possible values are: guest, administrator, viewer, externalGuest. Required.
    role *BookingStaffRole;
    // The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
    timeZone *string;
    // True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
    useBusinessHours *bool;
    // The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
    workingHours []BookingWorkHours;
}
// Instantiates a new bookingStaffMember and sets the default values.
func NewBookingStaffMember()(*BookingStaffMember) {
    m := &BookingStaffMember{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
// Gets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) GetAvailabilityIsAffectedByPersonalCalendar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.availabilityIsAffectedByPersonalCalendar
    }
}
// Gets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
func (m *BookingStaffMember) GetColorIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.colorIndex
    }
}
// Gets the role property value. The role of the staff member in the business. Possible values are: guest, administrator, viewer, externalGuest. Required.
func (m *BookingStaffMember) GetRole()(*BookingStaffRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// Gets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// Gets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) GetUseBusinessHours()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useBusinessHours
    }
}
// Gets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) GetWorkingHours()([]BookingWorkHours) {
    if m == nil {
        return nil
    } else {
        return m.workingHours
    }
}
// The deserialization information for the current model
func (m *BookingStaffMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["availabilityIsAffectedByPersonalCalendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityIsAffectedByPersonalCalendar(val)
        }
        return nil
    }
    res["colorIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorIndex(val)
        }
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffRole)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(BookingStaffRole)
            m.SetRole(&cast)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["useBusinessHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseBusinessHours(val)
        }
        return nil
    }
    res["workingHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingWorkHours() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHours, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingWorkHours))
            }
            m.SetWorkingHours(res)
        }
        return nil
    }
    return res
}
func (m *BookingStaffMember) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BookingStaffMember) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BookingPerson.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("availabilityIsAffectedByPersonalCalendar", m.GetAvailabilityIsAffectedByPersonalCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("colorIndex", m.GetColorIndex())
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := m.GetRole().String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useBusinessHours", m.GetUseBusinessHours())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkingHours()))
        for i, v := range m.GetWorkingHours() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("workingHours", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
// Parameters:
//  - value : Value to set for the availabilityIsAffectedByPersonalCalendar property.
func (m *BookingStaffMember) SetAvailabilityIsAffectedByPersonalCalendar(value *bool)() {
    m.availabilityIsAffectedByPersonalCalendar = value
}
// Sets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
// Parameters:
//  - value : Value to set for the colorIndex property.
func (m *BookingStaffMember) SetColorIndex(value *int32)() {
    m.colorIndex = value
}
// Sets the role property value. The role of the staff member in the business. Possible values are: guest, administrator, viewer, externalGuest. Required.
// Parameters:
//  - value : Value to set for the role property.
func (m *BookingStaffMember) SetRole(value *BookingStaffRole)() {
    m.role = value
}
// Sets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
// Parameters:
//  - value : Value to set for the timeZone property.
func (m *BookingStaffMember) SetTimeZone(value *string)() {
    m.timeZone = value
}
// Sets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
// Parameters:
//  - value : Value to set for the useBusinessHours property.
func (m *BookingStaffMember) SetUseBusinessHours(value *bool)() {
    m.useBusinessHours = value
}
// Sets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
// Parameters:
//  - value : Value to set for the workingHours property.
func (m *BookingStaffMember) SetWorkingHours(value []BookingWorkHours)() {
    m.workingHours = value
}
