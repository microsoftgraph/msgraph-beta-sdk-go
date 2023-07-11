package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingStaffMember represents a staff member who provides services in a business.
type BookingStaffMember struct {
    BookingPerson
}
// NewBookingStaffMember instantiates a new bookingStaffMember and sets the default values.
func NewBookingStaffMember()(*BookingStaffMember) {
    m := &BookingStaffMember{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
// CreateBookingStaffMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingStaffMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingStaffMember(), nil
}
// GetAvailabilityIsAffectedByPersonalCalendar gets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) GetAvailabilityIsAffectedByPersonalCalendar()(*bool) {
    val, err := m.GetBackingStore().Get("availabilityIsAffectedByPersonalCalendar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetColorIndex gets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
func (m *BookingStaffMember) GetColorIndex()(*int32) {
    val, err := m.GetBackingStore().Get("colorIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingStaffMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["availabilityIsAffectedByPersonalCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityIsAffectedByPersonalCalendar(val)
        }
        return nil
    }
    res["colorIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorIndex(val)
        }
        return nil
    }
    res["isEmailNotificationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEmailNotificationEnabled(val)
        }
        return nil
    }
    res["membershipStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffMembershipStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipStatus(val.(*BookingStaffMembershipStatus))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*BookingStaffRole))
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["useBusinessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseBusinessHours(val)
        }
        return nil
    }
    res["workingHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHoursable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingWorkHoursable)
                }
            }
            m.SetWorkingHours(res)
        }
        return nil
    }
    return res
}
// GetIsEmailNotificationEnabled gets the isEmailNotificationEnabled property value. True indicates that a staff member will be notified via email when a booking assigned to them is created or changed.
func (m *BookingStaffMember) GetIsEmailNotificationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEmailNotificationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMembershipStatus gets the membershipStatus property value. The membershipStatus property
func (m *BookingStaffMember) GetMembershipStatus()(*BookingStaffMembershipStatus) {
    val, err := m.GetBackingStore().Get("membershipStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*BookingStaffMembershipStatus)
    }
    return nil
}
// GetRole gets the role property value. The role property
func (m *BookingStaffMember) GetRole()(*BookingStaffRole) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*BookingStaffRole)
    }
    return nil
}
// GetTimeZone gets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) GetTimeZone()(*string) {
    val, err := m.GetBackingStore().Get("timeZone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUseBusinessHours gets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) GetUseBusinessHours()(*bool) {
    val, err := m.GetBackingStore().Get("useBusinessHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkingHours gets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) GetWorkingHours()([]BookingWorkHoursable) {
    val, err := m.GetBackingStore().Get("workingHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]BookingWorkHoursable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BookingStaffMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteBoolValue("isEmailNotificationEnabled", m.GetIsEmailNotificationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetMembershipStatus() != nil {
        cast := (*m.GetMembershipStatus()).String()
        err = writer.WriteStringValue("membershipStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
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
    if m.GetWorkingHours() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkingHours()))
        for i, v := range m.GetWorkingHours() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workingHours", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityIsAffectedByPersonalCalendar sets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) SetAvailabilityIsAffectedByPersonalCalendar(value *bool)() {
    err := m.GetBackingStore().Set("availabilityIsAffectedByPersonalCalendar", value)
    if err != nil {
        panic(err)
    }
}
// SetColorIndex sets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
func (m *BookingStaffMember) SetColorIndex(value *int32)() {
    err := m.GetBackingStore().Set("colorIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEmailNotificationEnabled sets the isEmailNotificationEnabled property value. True indicates that a staff member will be notified via email when a booking assigned to them is created or changed.
func (m *BookingStaffMember) SetIsEmailNotificationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEmailNotificationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMembershipStatus sets the membershipStatus property value. The membershipStatus property
func (m *BookingStaffMember) SetMembershipStatus(value *BookingStaffMembershipStatus)() {
    err := m.GetBackingStore().Set("membershipStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. The role property
func (m *BookingStaffMember) SetRole(value *BookingStaffRole)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeZone sets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) SetTimeZone(value *string)() {
    err := m.GetBackingStore().Set("timeZone", value)
    if err != nil {
        panic(err)
    }
}
// SetUseBusinessHours sets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) SetUseBusinessHours(value *bool)() {
    err := m.GetBackingStore().Set("useBusinessHours", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkingHours sets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) SetWorkingHours(value []BookingWorkHoursable)() {
    err := m.GetBackingStore().Set("workingHours", value)
    if err != nil {
        panic(err)
    }
}
// BookingStaffMemberable 
type BookingStaffMemberable interface {
    BookingPersonable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityIsAffectedByPersonalCalendar()(*bool)
    GetColorIndex()(*int32)
    GetIsEmailNotificationEnabled()(*bool)
    GetMembershipStatus()(*BookingStaffMembershipStatus)
    GetRole()(*BookingStaffRole)
    GetTimeZone()(*string)
    GetUseBusinessHours()(*bool)
    GetWorkingHours()([]BookingWorkHoursable)
    SetAvailabilityIsAffectedByPersonalCalendar(value *bool)()
    SetColorIndex(value *int32)()
    SetIsEmailNotificationEnabled(value *bool)()
    SetMembershipStatus(value *BookingStaffMembershipStatus)()
    SetRole(value *BookingStaffRole)()
    SetTimeZone(value *string)()
    SetUseBusinessHours(value *bool)()
    SetWorkingHours(value []BookingWorkHoursable)()
}
