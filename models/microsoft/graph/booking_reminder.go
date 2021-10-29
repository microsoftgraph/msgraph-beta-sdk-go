package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BookingReminder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The message in the reminder.
    message *string;
    // The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
    offset *string;
    // The persons who shouold receive the reminder. Possible values are: allAttendees, staff, customer.
    recipients *BookingReminderRecipients;
}
// Instantiates a new bookingReminder and sets the default values.
func NewBookingReminder()(*BookingReminder) {
    m := &BookingReminder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingReminder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the message property value. The message in the reminder.
func (m *BookingReminder) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Gets the offset property value. The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
func (m *BookingReminder) GetOffset()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
// Gets the recipients property value. The persons who shouold receive the reminder. Possible values are: allAttendees, staff, customer.
func (m *BookingReminder) GetRecipients()(*BookingReminderRecipients) {
    if m == nil {
        return nil
    } else {
        return m.recipients
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *BookingReminder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the message property value. The message in the reminder.
// Parameters:
//  - value : Value to set for the message property.
func (m *BookingReminder) SetMessage(value *string)() {
    m.message = value
}
// Sets the offset property value. The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
// Parameters:
//  - value : Value to set for the offset property.
func (m *BookingReminder) SetOffset(value *string)() {
    m.offset = value
}
// Sets the recipients property value. The persons who shouold receive the reminder. Possible values are: allAttendees, staff, customer.
// Parameters:
//  - value : Value to set for the recipients property.
func (m *BookingReminder) SetRecipients(value *BookingReminderRecipients)() {
    m.recipients = value
}
