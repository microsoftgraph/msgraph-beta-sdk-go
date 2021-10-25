package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingReminder struct {
    additionalData map[string]interface{};
    message *string;
    offset *string;
    recipients *BookingReminderRecipients;
}
func NewBookingReminder()(*BookingReminder) {
    m := &BookingReminder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BookingReminder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BookingReminder) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *BookingReminder) GetOffset()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
func (m *BookingReminder) GetRecipients()(*BookingReminderRecipients) {
    if m == nil {
        return nil
    } else {
        return m.recipients
    }
}
func (m *BookingReminder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["offset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOffset(val)
        return nil
    }
    res["recipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingReminderRecipients)
        if err != nil {
            return err
        }
        cast := val.(BookingReminderRecipients)
        m.SetRecipients(&cast)
        return nil
    }
    return res
}
func (m *BookingReminder) IsNil()(bool) {
    return m == nil
}
func (m *BookingReminder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("offset", m.GetOffset())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := m.GetRecipients().String()
        err := writer.WriteStringValue("recipients", &cast)
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
func (m *BookingReminder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BookingReminder) SetMessage(value *string)() {
    m.message = value
}
func (m *BookingReminder) SetOffset(value *string)() {
    m.offset = value
}
func (m *BookingReminder) SetRecipients(value *BookingReminderRecipients)() {
    m.recipients = value
}
