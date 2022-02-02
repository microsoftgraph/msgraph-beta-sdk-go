package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDateTimeConfiguration 
type TeamworkDateTimeConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date format for the device.
    dateFormat *string;
    // The time of the day when the device is turned off.
    officeHoursEndTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // The time of the day when the device is turned on.
    officeHoursStartTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // The time format for the device.
    timeFormat *string;
    // The time zone to which the office hours apply.
    timeZone *string;
}
// NewTeamworkDateTimeConfiguration instantiates a new teamworkDateTimeConfiguration and sets the default values.
func NewTeamworkDateTimeConfiguration()(*TeamworkDateTimeConfiguration) {
    m := &TeamworkDateTimeConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDateTimeConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDateFormat gets the dateFormat property value. The date format for the device.
func (m *TeamworkDateTimeConfiguration) GetDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dateFormat
    }
}
// GetOfficeHoursEndTime gets the officeHoursEndTime property value. The time of the day when the device is turned off.
func (m *TeamworkDateTimeConfiguration) GetOfficeHoursEndTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.officeHoursEndTime
    }
}
// GetOfficeHoursStartTime gets the officeHoursStartTime property value. The time of the day when the device is turned on.
func (m *TeamworkDateTimeConfiguration) GetOfficeHoursStartTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.officeHoursStartTime
    }
}
// GetTimeFormat gets the timeFormat property value. The time format for the device.
func (m *TeamworkDateTimeConfiguration) GetTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeFormat
    }
}
// GetTimeZone gets the timeZone property value. The time zone to which the office hours apply.
func (m *TeamworkDateTimeConfiguration) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDateTimeConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dateFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateFormat(val)
        }
        return nil
    }
    res["officeHoursEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeHoursEndTime(val)
        }
        return nil
    }
    res["officeHoursStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeHoursStartTime(val)
        }
        return nil
    }
    res["timeFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkDateTimeConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDateTimeConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dateFormat", m.GetDateFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("officeHoursEndTime", m.GetOfficeHoursEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("officeHoursStartTime", m.GetOfficeHoursStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeFormat", m.GetTimeFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
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
func (m *TeamworkDateTimeConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDateFormat sets the dateFormat property value. The date format for the device.
func (m *TeamworkDateTimeConfiguration) SetDateFormat(value *string)() {
    if m != nil {
        m.dateFormat = value
    }
}
// SetOfficeHoursEndTime sets the officeHoursEndTime property value. The time of the day when the device is turned off.
func (m *TeamworkDateTimeConfiguration) SetOfficeHoursEndTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.officeHoursEndTime = value
    }
}
// SetOfficeHoursStartTime sets the officeHoursStartTime property value. The time of the day when the device is turned on.
func (m *TeamworkDateTimeConfiguration) SetOfficeHoursStartTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.officeHoursStartTime = value
    }
}
// SetTimeFormat sets the timeFormat property value. The time format for the device.
func (m *TeamworkDateTimeConfiguration) SetTimeFormat(value *string)() {
    if m != nil {
        m.timeFormat = value
    }
}
// SetTimeZone sets the timeZone property value. The time zone to which the office hours apply.
func (m *TeamworkDateTimeConfiguration) SetTimeZone(value *string)() {
    if m != nil {
        m.timeZone = value
    }
}
