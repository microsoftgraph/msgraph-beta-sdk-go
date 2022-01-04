package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonalTaskProperties 
type PersonalTaskProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date and time for a reminder alert of the task to occur.
    reminderDatetime *DateTimeTimeZone;
}
// NewPersonalTaskProperties instantiates a new personalTaskProperties and sets the default values.
func NewPersonalTaskProperties()(*PersonalTaskProperties) {
    m := &PersonalTaskProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonalTaskProperties) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetReminderDatetime gets the reminderDatetime property value. The date and time for a reminder alert of the task to occur.
func (m *PersonalTaskProperties) GetReminderDatetime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.reminderDatetime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonalTaskProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reminderDatetime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderDatetime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    return res
}
func (m *PersonalTaskProperties) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PersonalTaskProperties) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("reminderDatetime", m.GetReminderDatetime())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonalTaskProperties) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReminderDatetime sets the reminderDatetime property value. The date and time for a reminder alert of the task to occur.
func (m *PersonalTaskProperties) SetReminderDatetime(value *DateTimeTimeZone)() {
    if m != nil {
        m.reminderDatetime = value
    }
}
