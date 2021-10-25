package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingStaffMember struct {
    BookingPerson
    availabilityIsAffectedByPersonalCalendar *bool;
    colorIndex *int32;
    role *BookingStaffRole;
    useBusinessHours *bool;
    workingHours []BookingWorkHours;
}
func NewBookingStaffMember()(*BookingStaffMember) {
    m := &BookingStaffMember{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
func (m *BookingStaffMember) GetAvailabilityIsAffectedByPersonalCalendar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.availabilityIsAffectedByPersonalCalendar
    }
}
func (m *BookingStaffMember) GetColorIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.colorIndex
    }
}
func (m *BookingStaffMember) GetRole()(*BookingStaffRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
func (m *BookingStaffMember) GetUseBusinessHours()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useBusinessHours
    }
}
func (m *BookingStaffMember) GetWorkingHours()([]BookingWorkHours) {
    if m == nil {
        return nil
    } else {
        return m.workingHours
    }
}
func (m *BookingStaffMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["availabilityIsAffectedByPersonalCalendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAvailabilityIsAffectedByPersonalCalendar(val)
        return nil
    }
    res["colorIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetColorIndex(val)
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffRole)
        if err != nil {
            return err
        }
        cast := val.(BookingStaffRole)
        m.SetRole(&cast)
        return nil
    }
    res["useBusinessHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUseBusinessHours(val)
        return nil
    }
    res["workingHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingWorkHours() })
        if err != nil {
            return err
        }
        res := make([]BookingWorkHours, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingWorkHours))
        }
        m.SetWorkingHours(res)
        return nil
    }
    return res
}
func (m *BookingStaffMember) IsNil()(bool) {
    return m == nil
}
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
func (m *BookingStaffMember) SetAvailabilityIsAffectedByPersonalCalendar(value *bool)() {
    m.availabilityIsAffectedByPersonalCalendar = value
}
func (m *BookingStaffMember) SetColorIndex(value *int32)() {
    m.colorIndex = value
}
func (m *BookingStaffMember) SetRole(value *BookingStaffRole)() {
    m.role = value
}
func (m *BookingStaffMember) SetUseBusinessHours(value *bool)() {
    m.useBusinessHours = value
}
func (m *BookingStaffMember) SetWorkingHours(value []BookingWorkHours)() {
    m.workingHours = value
}
