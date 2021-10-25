package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingBusiness struct {
    BookingNamedEntity
    address *PhysicalAddress;
    appointments []BookingAppointment;
    businessHours []BookingWorkHours;
    businessType *string;
    calendarView []BookingAppointment;
    customers []BookingCustomer;
    defaultCurrencyIso *string;
    email *string;
    isPublished *bool;
    phone *string;
    publicUrl *string;
    schedulingPolicy *BookingSchedulingPolicy;
    services []BookingService;
    staffMembers []BookingStaffMember;
    webSiteUrl *string;
}
func NewBookingBusiness()(*BookingBusiness) {
    m := &BookingBusiness{
        BookingNamedEntity: *NewBookingNamedEntity(),
    }
    return m
}
func (m *BookingBusiness) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *BookingBusiness) GetAppointments()([]BookingAppointment) {
    if m == nil {
        return nil
    } else {
        return m.appointments
    }
}
func (m *BookingBusiness) GetBusinessHours()([]BookingWorkHours) {
    if m == nil {
        return nil
    } else {
        return m.businessHours
    }
}
func (m *BookingBusiness) GetBusinessType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessType
    }
}
func (m *BookingBusiness) GetCalendarView()([]BookingAppointment) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
func (m *BookingBusiness) GetCustomers()([]BookingCustomer) {
    if m == nil {
        return nil
    } else {
        return m.customers
    }
}
func (m *BookingBusiness) GetDefaultCurrencyIso()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultCurrencyIso
    }
}
func (m *BookingBusiness) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
func (m *BookingBusiness) GetIsPublished()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPublished
    }
}
func (m *BookingBusiness) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
func (m *BookingBusiness) GetPublicUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicUrl
    }
}
func (m *BookingBusiness) GetSchedulingPolicy()(*BookingSchedulingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.schedulingPolicy
    }
}
func (m *BookingBusiness) GetServices()([]BookingService) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
func (m *BookingBusiness) GetStaffMembers()([]BookingStaffMember) {
    if m == nil {
        return nil
    } else {
        return m.staffMembers
    }
}
func (m *BookingBusiness) GetWebSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webSiteUrl
    }
}
func (m *BookingBusiness) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingNamedEntity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetAddress(val.(*PhysicalAddress))
        return nil
    }
    res["appointments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingAppointment() })
        if err != nil {
            return err
        }
        res := make([]BookingAppointment, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingAppointment))
        }
        m.SetAppointments(res)
        return nil
    }
    res["businessHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingWorkHours() })
        if err != nil {
            return err
        }
        res := make([]BookingWorkHours, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingWorkHours))
        }
        m.SetBusinessHours(res)
        return nil
    }
    res["businessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBusinessType(val)
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingAppointment() })
        if err != nil {
            return err
        }
        res := make([]BookingAppointment, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingAppointment))
        }
        m.SetCalendarView(res)
        return nil
    }
    res["customers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingCustomer() })
        if err != nil {
            return err
        }
        res := make([]BookingCustomer, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingCustomer))
        }
        m.SetCustomers(res)
        return nil
    }
    res["defaultCurrencyIso"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultCurrencyIso(val)
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["isPublished"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPublished(val)
        return nil
    }
    res["phone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhone(val)
        return nil
    }
    res["publicUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublicUrl(val)
        return nil
    }
    res["schedulingPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingSchedulingPolicy() })
        if err != nil {
            return err
        }
        m.SetSchedulingPolicy(val.(*BookingSchedulingPolicy))
        return nil
    }
    res["services"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingService() })
        if err != nil {
            return err
        }
        res := make([]BookingService, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingService))
        }
        m.SetServices(res)
        return nil
    }
    res["staffMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingStaffMember() })
        if err != nil {
            return err
        }
        res := make([]BookingStaffMember, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingStaffMember))
        }
        m.SetStaffMembers(res)
        return nil
    }
    res["webSiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebSiteUrl(val)
        return nil
    }
    return res
}
func (m *BookingBusiness) IsNil()(bool) {
    return m == nil
}
func (m *BookingBusiness) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BookingNamedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppointments()))
        for i, v := range m.GetAppointments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appointments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBusinessHours()))
        for i, v := range m.GetBusinessHours() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("businessHours", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessType", m.GetBusinessType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultCurrencyIso", m.GetDefaultCurrencyIso())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPublished", m.GetIsPublished())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publicUrl", m.GetPublicUrl())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStaffMembers()))
        for i, v := range m.GetStaffMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("staffMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webSiteUrl", m.GetWebSiteUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *BookingBusiness) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
func (m *BookingBusiness) SetAppointments(value []BookingAppointment)() {
    m.appointments = value
}
func (m *BookingBusiness) SetBusinessHours(value []BookingWorkHours)() {
    m.businessHours = value
}
func (m *BookingBusiness) SetBusinessType(value *string)() {
    m.businessType = value
}
func (m *BookingBusiness) SetCalendarView(value []BookingAppointment)() {
    m.calendarView = value
}
func (m *BookingBusiness) SetCustomers(value []BookingCustomer)() {
    m.customers = value
}
func (m *BookingBusiness) SetDefaultCurrencyIso(value *string)() {
    m.defaultCurrencyIso = value
}
func (m *BookingBusiness) SetEmail(value *string)() {
    m.email = value
}
func (m *BookingBusiness) SetIsPublished(value *bool)() {
    m.isPublished = value
}
func (m *BookingBusiness) SetPhone(value *string)() {
    m.phone = value
}
func (m *BookingBusiness) SetPublicUrl(value *string)() {
    m.publicUrl = value
}
func (m *BookingBusiness) SetSchedulingPolicy(value *BookingSchedulingPolicy)() {
    m.schedulingPolicy = value
}
func (m *BookingBusiness) SetServices(value []BookingService)() {
    m.services = value
}
func (m *BookingBusiness) SetStaffMembers(value []BookingStaffMember)() {
    m.staffMembers = value
}
func (m *BookingBusiness) SetWebSiteUrl(value *string)() {
    m.webSiteUrl = value
}
