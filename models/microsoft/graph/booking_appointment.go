package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BookingAppointment struct {
    Entity
    // 
    additionalInformation *string;
    // The SMTP address of the bookingCustomer who is booking the appointment.
    customerEmailAddress *string;
    // The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
    customerId *string;
    // Represents location information for the bookingCustomer who is booking the appointment.
    customerLocation *Location;
    // The customer's name.
    customerName *string;
    // Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID.  You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by customerId.
    customerNotes *string;
    // The customer's phone number.
    customerPhone *string;
    // The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
    customerTimeZone *string;
    // The length of the appointment, denoted in ISO8601 format.
    duration *string;
    // 
    end *DateTimeTimeZone;
    // The billed amount on the invoice.
    invoiceAmount *float64;
    // The date, time, and time zone of the invoice for this appointment.
    invoiceDate *DateTimeTimeZone;
    // The ID of the invoice.
    invoiceId *string;
    // The status of the invoice. Possible values are: draft, reviewing, open, canceled, paid, corrective.
    invoiceStatus *BookingInvoiceStatus;
    // The URL of the invoice in Microsoft Bookings.
    invoiceUrl *string;
    // True indicates that the appointment will be held online. Default value is false.
    isLocationOnline *bool;
    // The URL of the online meeting for the appointment.
    joinWebUrl *string;
    // 
    onlineMeetingUrl *string;
    // True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
    optOutOfCustomerEmail *bool;
    // The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
    postBuffer *string;
    // The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
    preBuffer *string;
    // The regular price for an appointment for the specified bookingService.
    price *float64;
    // A setting to provide flexibility for the pricing structure of services. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet.
    priceType *BookingPriceType;
    // The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
    reminders []BookingReminder;
    // An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
    selfServiceAppointmentId *string;
    // The ID of the bookingService associated with this appointment.
    serviceId *string;
    // The location where the service is delivered.
    serviceLocation *Location;
    // The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
    serviceName *string;
    // Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
    serviceNotes *string;
    // True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
    smsNotificationsEnabled *bool;
    // The ID of each bookingStaffMember who is scheduled in this appointment.
    staffMemberIds []string;
    // 
    start *DateTimeTimeZone;
}
// Instantiates a new bookingAppointment and sets the default values.
func NewBookingAppointment()(*BookingAppointment) {
    m := &BookingAppointment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the additionalInformation property value. 
func (m *BookingAppointment) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// Gets the customerEmailAddress property value. The SMTP address of the bookingCustomer who is booking the appointment.
func (m *BookingAppointment) GetCustomerEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerEmailAddress
    }
}
// Gets the customerId property value. The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
func (m *BookingAppointment) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// Gets the customerLocation property value. Represents location information for the bookingCustomer who is booking the appointment.
func (m *BookingAppointment) GetCustomerLocation()(*Location) {
    if m == nil {
        return nil
    } else {
        return m.customerLocation
    }
}
// Gets the customerName property value. The customer's name.
func (m *BookingAppointment) GetCustomerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerName
    }
}
// Gets the customerNotes property value. Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID.  You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by customerId.
func (m *BookingAppointment) GetCustomerNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNotes
    }
}
// Gets the customerPhone property value. The customer's phone number.
func (m *BookingAppointment) GetCustomerPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerPhone
    }
}
// Gets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingAppointment) GetCustomerTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerTimeZone
    }
}
// Gets the duration property value. The length of the appointment, denoted in ISO8601 format.
func (m *BookingAppointment) GetDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// Gets the end property value. 
func (m *BookingAppointment) GetEnd()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// Gets the invoiceAmount property value. The billed amount on the invoice.
func (m *BookingAppointment) GetInvoiceAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.invoiceAmount
    }
}
// Gets the invoiceDate property value. The date, time, and time zone of the invoice for this appointment.
func (m *BookingAppointment) GetInvoiceDate()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDate
    }
}
// Gets the invoiceId property value. The ID of the invoice.
func (m *BookingAppointment) GetInvoiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceId
    }
}
// Gets the invoiceStatus property value. The status of the invoice. Possible values are: draft, reviewing, open, canceled, paid, corrective.
func (m *BookingAppointment) GetInvoiceStatus()(*BookingInvoiceStatus) {
    if m == nil {
        return nil
    } else {
        return m.invoiceStatus
    }
}
// Gets the invoiceUrl property value. The URL of the invoice in Microsoft Bookings.
func (m *BookingAppointment) GetInvoiceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceUrl
    }
}
// Gets the isLocationOnline property value. True indicates that the appointment will be held online. Default value is false.
func (m *BookingAppointment) GetIsLocationOnline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLocationOnline
    }
}
// Gets the joinWebUrl property value. The URL of the online meeting for the appointment.
func (m *BookingAppointment) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// Gets the onlineMeetingUrl property value. 
func (m *BookingAppointment) GetOnlineMeetingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingUrl
    }
}
// Gets the optOutOfCustomerEmail property value. True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
func (m *BookingAppointment) GetOptOutOfCustomerEmail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.optOutOfCustomerEmail
    }
}
// Gets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) GetPostBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postBuffer
    }
}
// Gets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) GetPreBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preBuffer
    }
}
// Gets the price property value. The regular price for an appointment for the specified bookingService.
func (m *BookingAppointment) GetPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.price
    }
}
// Gets the priceType property value. A setting to provide flexibility for the pricing structure of services. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet.
func (m *BookingAppointment) GetPriceType()(*BookingPriceType) {
    if m == nil {
        return nil
    } else {
        return m.priceType
    }
}
// Gets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) GetReminders()([]BookingReminder) {
    if m == nil {
        return nil
    } else {
        return m.reminders
    }
}
// Gets the selfServiceAppointmentId property value. An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
func (m *BookingAppointment) GetSelfServiceAppointmentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.selfServiceAppointmentId
    }
}
// Gets the serviceId property value. The ID of the bookingService associated with this appointment.
func (m *BookingAppointment) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
// Gets the serviceLocation property value. The location where the service is delivered.
func (m *BookingAppointment) GetServiceLocation()(*Location) {
    if m == nil {
        return nil
    } else {
        return m.serviceLocation
    }
}
// Gets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
func (m *BookingAppointment) GetServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceName
    }
}
// Gets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) GetServiceNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceNotes
    }
}
// Gets the smsNotificationsEnabled property value. True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
func (m *BookingAppointment) GetSmsNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.smsNotificationsEnabled
    }
}
// Gets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
func (m *BookingAppointment) GetStaffMemberIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.staffMemberIds
    }
}
// Gets the start property value. 
func (m *BookingAppointment) GetStart()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// The deserialization information for the current model
func (m *BookingAppointment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalInformation(val)
        return nil
    }
    res["customerEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerEmailAddress(val)
        return nil
    }
    res["customerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerId(val)
        return nil
    }
    res["customerLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        m.SetCustomerLocation(val.(*Location))
        return nil
    }
    res["customerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerName(val)
        return nil
    }
    res["customerNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerNotes(val)
        return nil
    }
    res["customerPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerPhone(val)
        return nil
    }
    res["customerTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomerTimeZone(val)
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDuration(val)
        return nil
    }
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetEnd(val.(*DateTimeTimeZone))
        return nil
    }
    res["invoiceAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetInvoiceAmount(val)
        return nil
    }
    res["invoiceDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetInvoiceDate(val.(*DateTimeTimeZone))
        return nil
    }
    res["invoiceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvoiceId(val)
        return nil
    }
    res["invoiceStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingInvoiceStatus)
        if err != nil {
            return err
        }
        cast := val.(BookingInvoiceStatus)
        m.SetInvoiceStatus(&cast)
        return nil
    }
    res["invoiceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvoiceUrl(val)
        return nil
    }
    res["isLocationOnline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsLocationOnline(val)
        return nil
    }
    res["joinWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJoinWebUrl(val)
        return nil
    }
    res["onlineMeetingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnlineMeetingUrl(val)
        return nil
    }
    res["optOutOfCustomerEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOptOutOfCustomerEmail(val)
        return nil
    }
    res["postBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostBuffer(val)
        return nil
    }
    res["preBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreBuffer(val)
        return nil
    }
    res["price"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetPrice(val)
        return nil
    }
    res["priceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        cast := val.(BookingPriceType)
        m.SetPriceType(&cast)
        return nil
    }
    res["reminders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingReminder() })
        if err != nil {
            return err
        }
        res := make([]BookingReminder, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingReminder))
        }
        m.SetReminders(res)
        return nil
    }
    res["selfServiceAppointmentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSelfServiceAppointmentId(val)
        return nil
    }
    res["serviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServiceId(val)
        return nil
    }
    res["serviceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        m.SetServiceLocation(val.(*Location))
        return nil
    }
    res["serviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServiceName(val)
        return nil
    }
    res["serviceNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServiceNotes(val)
        return nil
    }
    res["smsNotificationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSmsNotificationsEnabled(val)
        return nil
    }
    res["staffMemberIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetStaffMemberIds(res)
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetStart(val.(*DateTimeTimeZone))
        return nil
    }
    return res
}
func (m *BookingAppointment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BookingAppointment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteStringValue("customerEmailAddress", m.GetCustomerEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customerLocation", m.GetCustomerLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerName", m.GetCustomerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerNotes", m.GetCustomerNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerPhone", m.GetCustomerPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerTimeZone", m.GetCustomerTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("invoiceAmount", m.GetInvoiceAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("invoiceDate", m.GetInvoiceDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invoiceId", m.GetInvoiceId())
        if err != nil {
            return err
        }
    }
    if m.GetInvoiceStatus() != nil {
        cast := m.GetInvoiceStatus().String()
        err = writer.WriteStringValue("invoiceStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invoiceUrl", m.GetInvoiceUrl())
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
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineMeetingUrl", m.GetOnlineMeetingUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("optOutOfCustomerEmail", m.GetOptOutOfCustomerEmail())
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
        err = writer.WriteFloat64Value("price", m.GetPrice())
        if err != nil {
            return err
        }
    }
    if m.GetPriceType() != nil {
        cast := m.GetPriceType().String()
        err = writer.WriteStringValue("priceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReminders()))
        for i, v := range m.GetReminders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("selfServiceAppointmentId", m.GetSelfServiceAppointmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("serviceLocation", m.GetServiceLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceName", m.GetServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceNotes", m.GetServiceNotes())
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
        err = writer.WriteObjectValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the additionalInformation property value. 
// Parameters:
//  - value : Value to set for the additionalInformation property.
func (m *BookingAppointment) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// Sets the customerEmailAddress property value. The SMTP address of the bookingCustomer who is booking the appointment.
// Parameters:
//  - value : Value to set for the customerEmailAddress property.
func (m *BookingAppointment) SetCustomerEmailAddress(value *string)() {
    m.customerEmailAddress = value
}
// Sets the customerId property value. The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
// Parameters:
//  - value : Value to set for the customerId property.
func (m *BookingAppointment) SetCustomerId(value *string)() {
    m.customerId = value
}
// Sets the customerLocation property value. Represents location information for the bookingCustomer who is booking the appointment.
// Parameters:
//  - value : Value to set for the customerLocation property.
func (m *BookingAppointment) SetCustomerLocation(value *Location)() {
    m.customerLocation = value
}
// Sets the customerName property value. The customer's name.
// Parameters:
//  - value : Value to set for the customerName property.
func (m *BookingAppointment) SetCustomerName(value *string)() {
    m.customerName = value
}
// Sets the customerNotes property value. Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID.  You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by customerId.
// Parameters:
//  - value : Value to set for the customerNotes property.
func (m *BookingAppointment) SetCustomerNotes(value *string)() {
    m.customerNotes = value
}
// Sets the customerPhone property value. The customer's phone number.
// Parameters:
//  - value : Value to set for the customerPhone property.
func (m *BookingAppointment) SetCustomerPhone(value *string)() {
    m.customerPhone = value
}
// Sets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
// Parameters:
//  - value : Value to set for the customerTimeZone property.
func (m *BookingAppointment) SetCustomerTimeZone(value *string)() {
    m.customerTimeZone = value
}
// Sets the duration property value. The length of the appointment, denoted in ISO8601 format.
// Parameters:
//  - value : Value to set for the duration property.
func (m *BookingAppointment) SetDuration(value *string)() {
    m.duration = value
}
// Sets the end property value. 
// Parameters:
//  - value : Value to set for the end property.
func (m *BookingAppointment) SetEnd(value *DateTimeTimeZone)() {
    m.end = value
}
// Sets the invoiceAmount property value. The billed amount on the invoice.
// Parameters:
//  - value : Value to set for the invoiceAmount property.
func (m *BookingAppointment) SetInvoiceAmount(value *float64)() {
    m.invoiceAmount = value
}
// Sets the invoiceDate property value. The date, time, and time zone of the invoice for this appointment.
// Parameters:
//  - value : Value to set for the invoiceDate property.
func (m *BookingAppointment) SetInvoiceDate(value *DateTimeTimeZone)() {
    m.invoiceDate = value
}
// Sets the invoiceId property value. The ID of the invoice.
// Parameters:
//  - value : Value to set for the invoiceId property.
func (m *BookingAppointment) SetInvoiceId(value *string)() {
    m.invoiceId = value
}
// Sets the invoiceStatus property value. The status of the invoice. Possible values are: draft, reviewing, open, canceled, paid, corrective.
// Parameters:
//  - value : Value to set for the invoiceStatus property.
func (m *BookingAppointment) SetInvoiceStatus(value *BookingInvoiceStatus)() {
    m.invoiceStatus = value
}
// Sets the invoiceUrl property value. The URL of the invoice in Microsoft Bookings.
// Parameters:
//  - value : Value to set for the invoiceUrl property.
func (m *BookingAppointment) SetInvoiceUrl(value *string)() {
    m.invoiceUrl = value
}
// Sets the isLocationOnline property value. True indicates that the appointment will be held online. Default value is false.
// Parameters:
//  - value : Value to set for the isLocationOnline property.
func (m *BookingAppointment) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
// Sets the joinWebUrl property value. The URL of the online meeting for the appointment.
// Parameters:
//  - value : Value to set for the joinWebUrl property.
func (m *BookingAppointment) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// Sets the onlineMeetingUrl property value. 
// Parameters:
//  - value : Value to set for the onlineMeetingUrl property.
func (m *BookingAppointment) SetOnlineMeetingUrl(value *string)() {
    m.onlineMeetingUrl = value
}
// Sets the optOutOfCustomerEmail property value. True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
// Parameters:
//  - value : Value to set for the optOutOfCustomerEmail property.
func (m *BookingAppointment) SetOptOutOfCustomerEmail(value *bool)() {
    m.optOutOfCustomerEmail = value
}
// Sets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
// Parameters:
//  - value : Value to set for the postBuffer property.
func (m *BookingAppointment) SetPostBuffer(value *string)() {
    m.postBuffer = value
}
// Sets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
// Parameters:
//  - value : Value to set for the preBuffer property.
func (m *BookingAppointment) SetPreBuffer(value *string)() {
    m.preBuffer = value
}
// Sets the price property value. The regular price for an appointment for the specified bookingService.
// Parameters:
//  - value : Value to set for the price property.
func (m *BookingAppointment) SetPrice(value *float64)() {
    m.price = value
}
// Sets the priceType property value. A setting to provide flexibility for the pricing structure of services. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet.
// Parameters:
//  - value : Value to set for the priceType property.
func (m *BookingAppointment) SetPriceType(value *BookingPriceType)() {
    m.priceType = value
}
// Sets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
// Parameters:
//  - value : Value to set for the reminders property.
func (m *BookingAppointment) SetReminders(value []BookingReminder)() {
    m.reminders = value
}
// Sets the selfServiceAppointmentId property value. An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
// Parameters:
//  - value : Value to set for the selfServiceAppointmentId property.
func (m *BookingAppointment) SetSelfServiceAppointmentId(value *string)() {
    m.selfServiceAppointmentId = value
}
// Sets the serviceId property value. The ID of the bookingService associated with this appointment.
// Parameters:
//  - value : Value to set for the serviceId property.
func (m *BookingAppointment) SetServiceId(value *string)() {
    m.serviceId = value
}
// Sets the serviceLocation property value. The location where the service is delivered.
// Parameters:
//  - value : Value to set for the serviceLocation property.
func (m *BookingAppointment) SetServiceLocation(value *Location)() {
    m.serviceLocation = value
}
// Sets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
// Parameters:
//  - value : Value to set for the serviceName property.
func (m *BookingAppointment) SetServiceName(value *string)() {
    m.serviceName = value
}
// Sets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
// Parameters:
//  - value : Value to set for the serviceNotes property.
func (m *BookingAppointment) SetServiceNotes(value *string)() {
    m.serviceNotes = value
}
// Sets the smsNotificationsEnabled property value. True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
// Parameters:
//  - value : Value to set for the smsNotificationsEnabled property.
func (m *BookingAppointment) SetSmsNotificationsEnabled(value *bool)() {
    m.smsNotificationsEnabled = value
}
// Sets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
// Parameters:
//  - value : Value to set for the staffMemberIds property.
func (m *BookingAppointment) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
// Sets the start property value. 
// Parameters:
//  - value : Value to set for the start property.
func (m *BookingAppointment) SetStart(value *DateTimeTimeZone)() {
    m.start = value
}
