package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingService struct {
    BookingNamedEntity
    additionalInformation *string;
    defaultDuration *string;
    defaultLocation *Location;
    defaultPrice *float64;
    defaultPriceType *BookingPriceType;
    defaultReminders []BookingReminder;
    description *string;
    isHiddenFromCustomers *bool;
    isLocationOnline *bool;
    notes *string;
    postBuffer *string;
    preBuffer *string;
    schedulingPolicy *BookingSchedulingPolicy;
    staffMemberIds []string;
}
func NewBookingService()(*BookingService) {
    m := &BookingService{
        BookingNamedEntity: *NewBookingNamedEntity(),
    }
    return m
}
func (m *BookingService) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
func (m *BookingService) GetDefaultDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDuration
    }
}
func (m *BookingService) GetDefaultLocation()(*Location) {
    if m == nil {
        return nil
    } else {
        return m.defaultLocation
    }
}
func (m *BookingService) GetDefaultPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.defaultPrice
    }
}
func (m *BookingService) GetDefaultPriceType()(*BookingPriceType) {
    if m == nil {
        return nil
    } else {
        return m.defaultPriceType
    }
}
func (m *BookingService) GetDefaultReminders()([]BookingReminder) {
    if m == nil {
        return nil
    } else {
        return m.defaultReminders
    }
}
func (m *BookingService) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *BookingService) GetIsHiddenFromCustomers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHiddenFromCustomers
    }
}
func (m *BookingService) GetIsLocationOnline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLocationOnline
    }
}
func (m *BookingService) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
func (m *BookingService) GetPostBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postBuffer
    }
}
func (m *BookingService) GetPreBuffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preBuffer
    }
}
func (m *BookingService) GetSchedulingPolicy()(*BookingSchedulingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.schedulingPolicy
    }
}
func (m *BookingService) GetStaffMemberIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.staffMemberIds
    }
}
func (m *BookingService) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BookingNamedEntity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalInformation(val)
        return nil
    }
    res["defaultDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultDuration(val)
        return nil
    }
    res["defaultLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        m.SetDefaultLocation(val.(*Location))
        return nil
    }
    res["defaultPrice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetDefaultPrice(val)
        return nil
    }
    res["defaultPriceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        cast := val.(BookingPriceType)
        m.SetDefaultPriceType(&cast)
        return nil
    }
    res["defaultReminders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingReminder() })
        if err != nil {
            return err
        }
        res := make([]BookingReminder, len(val))
        for i, v := range val {
            res[i] = *(v.(*BookingReminder))
        }
        m.SetDefaultReminders(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["isHiddenFromCustomers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsHiddenFromCustomers(val)
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
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotes(val)
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
    res["schedulingPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookingSchedulingPolicy() })
        if err != nil {
            return err
        }
        m.SetSchedulingPolicy(val.(*BookingSchedulingPolicy))
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
    return res
}
func (m *BookingService) IsNil()(bool) {
    return m == nil
}
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
        err = writer.WriteCollectionOfStringValues("staffMemberIds", m.GetStaffMemberIds())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *BookingService) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
func (m *BookingService) SetDefaultDuration(value *string)() {
    m.defaultDuration = value
}
func (m *BookingService) SetDefaultLocation(value *Location)() {
    m.defaultLocation = value
}
func (m *BookingService) SetDefaultPrice(value *float64)() {
    m.defaultPrice = value
}
func (m *BookingService) SetDefaultPriceType(value *BookingPriceType)() {
    m.defaultPriceType = value
}
func (m *BookingService) SetDefaultReminders(value []BookingReminder)() {
    m.defaultReminders = value
}
func (m *BookingService) SetDescription(value *string)() {
    m.description = value
}
func (m *BookingService) SetIsHiddenFromCustomers(value *bool)() {
    m.isHiddenFromCustomers = value
}
func (m *BookingService) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
func (m *BookingService) SetNotes(value *string)() {
    m.notes = value
}
func (m *BookingService) SetPostBuffer(value *string)() {
    m.postBuffer = value
}
func (m *BookingService) SetPreBuffer(value *string)() {
    m.preBuffer = value
}
func (m *BookingService) SetSchedulingPolicy(value *BookingSchedulingPolicy)() {
    m.schedulingPolicy = value
}
func (m *BookingService) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
