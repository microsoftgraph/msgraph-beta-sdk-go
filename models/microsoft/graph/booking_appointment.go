package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingAppointment struct {
    Entity
    additionalInformation *string;
    customerEmailAddress *string;
    customerId *string;
    customerLocation *Location;
    customerName *string;
    customerNotes *string;
    customerPhone *string;
    duration *string;
    end *DateTimeTimeZone;
    invoiceAmount *float64;
    invoiceDate *DateTimeTimeZone;
    invoiceId *string;
    invoiceStatus *BookingInvoiceStatus;
    invoiceUrl *string;
    isLocationOnline *bool;
    onlineMeetingUrl *string;
    optOutOfCustomerEmail *bool;
    postBuffer *string;
    preBuffer *string;
    price *float64;
    priceType *BookingPriceType;
    reminders []BookingReminder;
    selfServiceAppointmentId *string;
    serviceId *string;
    serviceLocation *Location;
    serviceName *string;
    serviceNotes *string;
    staffMemberIds []string;
    start *DateTimeTimeZone;
}
func NewBookingAppointment()(*BookingAppointment) {
    m := &BookingAppointment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *BookingAppointment) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
func (m *BookingAppointment) GetCustomerEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerEmailAddress
    }
}
func (m *BookingAppointment) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
func (m *BookingAppointment) GetCustomerLocation()(*Location) {
    if m == nil {
        return nil
    } else {
        return m.customerLocation
    }
}
func (m *BookingAppointment) GetCustomerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerName
    }
}
func (m *BookingAppointment) GetCustomerNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNotes
    }
}
func (m *BookingAppointment) GetCustomerPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerPhone
    }
}
func (m *BookingAppointment) GetDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
func (m *BookingAppointment) GetEnd()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
func (m *BookingAppointment) GetInvoiceAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.invoiceAmount
    }
}
func (m *BookingAppointment) GetInvoiceDate()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDate
    }
}
func (m *BookingAppointment) GetInvoiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceId
    }
}
func (m *BookingAppointment) GetInvoiceStatus()(*BookingInvoiceStatus) {
    if m == nil {
        return nil
    } else {
        return m.invoiceStatus
    }
}
func (m *BookingAppointment) GetInvoiceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceUrl
    }
}
func (m *BookingAppointment) GetIsLocationOnline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLocationOnline
    }
}
func (m *BookingAppointment) GetOnlineMeetingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingUrl
    }
}
func (m *BookingAppointment) GetOptOutOfCustomerEmail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.optOutOfCustomerEmail
    }
}
func (m *BookingAppointment) GetPostBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postBuffer
    }
}
func (m *BookingAppointment) GetPreBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preBuffer
    }
}
func (m *BookingAppointment) GetPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.price
    }
}
func (m *BookingAppointment) GetPriceType()(*BookingPriceType) {
    if m == nil {
        return nil
    } else {
        return m.priceType
    }
}
func (m *BookingAppointment) GetReminders()([]BookingReminder) {
    if m == nil {
        return nil
    } else {
        return m.reminders
    }
}
func (m *BookingAppointment) GetSelfServiceAppointmentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.selfServiceAppointmentId
    }
}
func (m *BookingAppointment) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
func (m *BookingAppointment) GetServiceLocation()(*Location) {
    if m == nil {
        return nil
    } else {
        return m.serviceLocation
    }
}
func (m *BookingAppointment) GetServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceName
    }
}
func (m *BookingAppointment) GetServiceNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceNotes
    }
}
func (m *BookingAppointment) GetStaffMemberIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.staffMemberIds
    }
}
func (m *BookingAppointment) GetStart()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
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
    res["staffMemberIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
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
func (m *BookingAppointment) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
func (m *BookingAppointment) SetCustomerEmailAddress(value *string)() {
    m.customerEmailAddress = value
}
func (m *BookingAppointment) SetCustomerId(value *string)() {
    m.customerId = value
}
func (m *BookingAppointment) SetCustomerLocation(value *Location)() {
    m.customerLocation = value
}
func (m *BookingAppointment) SetCustomerName(value *string)() {
    m.customerName = value
}
func (m *BookingAppointment) SetCustomerNotes(value *string)() {
    m.customerNotes = value
}
func (m *BookingAppointment) SetCustomerPhone(value *string)() {
    m.customerPhone = value
}
func (m *BookingAppointment) SetDuration(value *string)() {
    m.duration = value
}
func (m *BookingAppointment) SetEnd(value *DateTimeTimeZone)() {
    m.end = value
}
func (m *BookingAppointment) SetInvoiceAmount(value *float64)() {
    m.invoiceAmount = value
}
func (m *BookingAppointment) SetInvoiceDate(value *DateTimeTimeZone)() {
    m.invoiceDate = value
}
func (m *BookingAppointment) SetInvoiceId(value *string)() {
    m.invoiceId = value
}
func (m *BookingAppointment) SetInvoiceStatus(value *BookingInvoiceStatus)() {
    m.invoiceStatus = value
}
func (m *BookingAppointment) SetInvoiceUrl(value *string)() {
    m.invoiceUrl = value
}
func (m *BookingAppointment) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
func (m *BookingAppointment) SetOnlineMeetingUrl(value *string)() {
    m.onlineMeetingUrl = value
}
func (m *BookingAppointment) SetOptOutOfCustomerEmail(value *bool)() {
    m.optOutOfCustomerEmail = value
}
func (m *BookingAppointment) SetPostBuffer(value *string)() {
    m.postBuffer = value
}
func (m *BookingAppointment) SetPreBuffer(value *string)() {
    m.preBuffer = value
}
func (m *BookingAppointment) SetPrice(value *float64)() {
    m.price = value
}
func (m *BookingAppointment) SetPriceType(value *BookingPriceType)() {
    m.priceType = value
}
func (m *BookingAppointment) SetReminders(value []BookingReminder)() {
    m.reminders = value
}
func (m *BookingAppointment) SetSelfServiceAppointmentId(value *string)() {
    m.selfServiceAppointmentId = value
}
func (m *BookingAppointment) SetServiceId(value *string)() {
    m.serviceId = value
}
func (m *BookingAppointment) SetServiceLocation(value *Location)() {
    m.serviceLocation = value
}
func (m *BookingAppointment) SetServiceName(value *string)() {
    m.serviceName = value
}
func (m *BookingAppointment) SetServiceNotes(value *string)() {
    m.serviceNotes = value
}
func (m *BookingAppointment) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
func (m *BookingAppointment) SetStart(value *DateTimeTimeZone)() {
    m.start = value
}
