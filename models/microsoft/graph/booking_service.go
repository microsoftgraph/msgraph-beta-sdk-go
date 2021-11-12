package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BookingService struct {
    BookingNamedEntity
    // 
    additionalInformation *string;
    // The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
    defaultDuration *string;
    // The default physical location for the service.
    defaultLocation *Location;
    // The default monetary price for the service.
    defaultPrice *float64;
    // The default way the service is charged. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet.
    defaultPriceType *BookingPriceType;
    // The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
    defaultReminders []BookingReminder;
    // A text description for the service.
    description *string;
    // True means this service is not available to customers for booking.
    isHiddenFromCustomers *bool;
    // 
    isLocationOnline *bool;
    // Additional information about this service.
    notes *string;
    // The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
    postBuffer *string;
    // The time to buffer before an appointment for this service can start.
    preBuffer *string;
    // The set of policies that determine how appointments for this type of service should be created and managed.
    schedulingPolicy *BookingSchedulingPolicy;
    // 
    smsNotificationsEnabled *bool;
    // Represents those staff members who provide this service.
    staffMemberIds []string;
    // The URL of the booking service.
    webUrl *string;
}
// Instantiates a new bookingService and sets the default values.
func NewBookingService()(*BookingService) {
    m := &BookingService{
        BookingNamedEntity: *NewBookingNamedEntity(),
    }
    return m
}
// Gets the additionalInformation property value. 
func (m *BookingService) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// Gets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
func (m *BookingService) GetDefaultDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDuration
    }
}
// Gets the defaultLocation property value. The default physical location for the service.
func (m *BookingService) GetDefaultLocation()(*Location) {
    if m == nil {
        return nil
    } else {
        return m.defaultLocation
    }
}
// Gets the defaultPrice property value. The default monetary price for the service.
func (m *BookingService) GetDefaultPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.defaultPrice
    }
}
// Gets the defaultPriceType property value. The default way the service is charged. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet.
func (m *BookingService) GetDefaultPriceType()(*BookingPriceType) {
    if m == nil {
        return nil
    } else {
        return m.defaultPriceType
    }
}
// Gets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
func (m *BookingService) GetDefaultReminders()([]BookingReminder) {
    if m == nil {
        return nil
    } else {
        return m.defaultReminders
    }
}
// Gets the description property value. A text description for the service.
func (m *BookingService) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the isHiddenFromCustomers property value. True means this service is not available to customers for booking.
func (m *BookingService) GetIsHiddenFromCustomers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHiddenFromCustomers
    }
}
// Gets the isLocationOnline property value. 
func (m *BookingService) GetIsLocationOnline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLocationOnline
    }
}
// Gets the notes property value. Additional information about this service.
func (m *BookingService) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
func (m *BookingService) GetPostBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postBuffer
    }
}
// Gets the preBuffer property value. The time to buffer before an appointment for this service can start.
func (m *BookingService) GetPreBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preBuffer
    }
}
// Gets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
func (m *BookingService) GetSchedulingPolicy()(*BookingSchedulingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.schedulingPolicy
    }
}
// Gets the smsNotificationsEnabled property value. 
func (m *BookingService) GetSmsNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.smsNotificationsEnabled
    }
}
// Gets the staffMemberIds property value. Represents those staff members who provide this service.
func (m *BookingService) GetStaffMemberIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.staffMemberIds
    }
}
// Gets the webUrl property value. The URL of the booking service.
func (m *BookingService) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *BookingService) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingNamedEntity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["defaultDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDuration(val)
        }
        return nil
    }
    res["defaultLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLocation(val.(*Location))
        }
        return nil
    }
    res["defaultPrice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultPrice(val)
        }
        return nil
    }
    res["defaultPriceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(BookingPriceType)
            m.SetDefaultPriceType(&cast)
        }
        return nil
    }
    res["defaultReminders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingReminder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingReminder, len(val))
            for i, v := range val {
                res[i] = *(v.(*BookingReminder))
            }
            m.SetDefaultReminders(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["isHiddenFromCustomers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHiddenFromCustomers(val)
        }
        return nil
    }
    res["isLocationOnline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocationOnline(val)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["postBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBuffer(val)
        }
        return nil
    }
    res["preBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreBuffer(val)
        }
        return nil
    }
    res["schedulingPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingSchedulingPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingPolicy(val.(*BookingSchedulingPolicy))
        }
        return nil
    }
    res["smsNotificationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsNotificationsEnabled(val)
        }
        return nil
    }
    res["staffMemberIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStaffMemberIds(res)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
func (m *BookingService) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BookingService) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BookingNamedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultDuration", m.GetDefaultDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultLocation", m.GetDefaultLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("defaultPrice", m.GetDefaultPrice())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultPriceType() != nil {
        cast := m.GetDefaultPriceType().String()
        err = writer.WriteStringValue("defaultPriceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefaultReminders()))
        for i, v := range m.GetDefaultReminders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("defaultReminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHiddenFromCustomers", m.GetIsHiddenFromCustomers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocationOnline", m.GetIsLocationOnline())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postBuffer", m.GetPostBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preBuffer", m.GetPreBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedulingPolicy", m.GetSchedulingPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smsNotificationsEnabled", m.GetSmsNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("staffMemberIds", m.GetStaffMemberIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the additionalInformation property value. 
// Parameters:
//  - value : Value to set for the additionalInformation property.
func (m *BookingService) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// Sets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
// Parameters:
//  - value : Value to set for the defaultDuration property.
func (m *BookingService) SetDefaultDuration(value *string)() {
    m.defaultDuration = value
}
// Sets the defaultLocation property value. The default physical location for the service.
// Parameters:
//  - value : Value to set for the defaultLocation property.
func (m *BookingService) SetDefaultLocation(value *Location)() {
    m.defaultLocation = value
}
// Sets the defaultPrice property value. The default monetary price for the service.
// Parameters:
//  - value : Value to set for the defaultPrice property.
func (m *BookingService) SetDefaultPrice(value *float64)() {
    m.defaultPrice = value
}
// Sets the defaultPriceType property value. The default way the service is charged. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet.
// Parameters:
//  - value : Value to set for the defaultPriceType property.
func (m *BookingService) SetDefaultPriceType(value *BookingPriceType)() {
    m.defaultPriceType = value
}
// Sets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
// Parameters:
//  - value : Value to set for the defaultReminders property.
func (m *BookingService) SetDefaultReminders(value []BookingReminder)() {
    m.defaultReminders = value
}
// Sets the description property value. A text description for the service.
// Parameters:
//  - value : Value to set for the description property.
func (m *BookingService) SetDescription(value *string)() {
    m.description = value
}
// Sets the isHiddenFromCustomers property value. True means this service is not available to customers for booking.
// Parameters:
//  - value : Value to set for the isHiddenFromCustomers property.
func (m *BookingService) SetIsHiddenFromCustomers(value *bool)() {
    m.isHiddenFromCustomers = value
}
// Sets the isLocationOnline property value. 
// Parameters:
//  - value : Value to set for the isLocationOnline property.
func (m *BookingService) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
// Sets the notes property value. Additional information about this service.
// Parameters:
//  - value : Value to set for the notes property.
func (m *BookingService) SetNotes(value *string)() {
    m.notes = value
}
// Sets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
// Parameters:
//  - value : Value to set for the postBuffer property.
func (m *BookingService) SetPostBuffer(value *string)() {
    m.postBuffer = value
}
// Sets the preBuffer property value. The time to buffer before an appointment for this service can start.
// Parameters:
//  - value : Value to set for the preBuffer property.
func (m *BookingService) SetPreBuffer(value *string)() {
    m.preBuffer = value
}
// Sets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
// Parameters:
//  - value : Value to set for the schedulingPolicy property.
func (m *BookingService) SetSchedulingPolicy(value *BookingSchedulingPolicy)() {
    m.schedulingPolicy = value
}
// Sets the smsNotificationsEnabled property value. 
// Parameters:
//  - value : Value to set for the smsNotificationsEnabled property.
func (m *BookingService) SetSmsNotificationsEnabled(value *bool)() {
    m.smsNotificationsEnabled = value
}
// Sets the staffMemberIds property value. Represents those staff members who provide this service.
// Parameters:
//  - value : Value to set for the staffMemberIds property.
func (m *BookingService) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
// Sets the webUrl property value. The URL of the booking service.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *BookingService) SetWebUrl(value *string)() {
    m.webUrl = value
}
