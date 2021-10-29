package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RegionalFormatOverrides struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The calendar to use, e.g., Gregorian Calendar.Returned by default.
    calendar *string;
    // The first day of the week to use, e.g., Sunday.Returned by default.
    firstDayOfWeek *string;
    // The long date time format to be used for displaying dates.Returned by default.
    longDateFormat *string;
    // The long time format to be used for displaying time.Returned by default.
    longTimeFormat *string;
    // The short date time format to be used for displaying dates.Returned by default.
    shortDateFormat *string;
    // The short time format to be used for displaying time.Returned by default.
    shortTimeFormat *string;
    // The timezone to be used for displaying time.Returned by default.
    timeZone *string;
}
// Instantiates a new regionalFormatOverrides and sets the default values.
func NewRegionalFormatOverrides()(*RegionalFormatOverrides) {
    m := &RegionalFormatOverrides{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegionalFormatOverrides) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the calendar property value. The calendar to use, e.g., Gregorian Calendar.Returned by default.
func (m *RegionalFormatOverrides) GetCalendar()(*string) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// Gets the firstDayOfWeek property value. The first day of the week to use, e.g., Sunday.Returned by default.
func (m *RegionalFormatOverrides) GetFirstDayOfWeek()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firstDayOfWeek
    }
}
// Gets the longDateFormat property value. The long date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) GetLongDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.longDateFormat
    }
}
// Gets the longTimeFormat property value. The long time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetLongTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.longTimeFormat
    }
}
// Gets the shortDateFormat property value. The short date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) GetShortDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortDateFormat
    }
}
// Gets the shortTimeFormat property value. The short time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetShortTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortTimeFormat
    }
}
// Gets the timeZone property value. The timezone to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// The deserialization information for the current model
func (m *RegionalFormatOverrides) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCalendar(val)
        return nil
    }
    res["firstDayOfWeek"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFirstDayOfWeek(val)
        return nil
    }
    res["longDateFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLongDateFormat(val)
        return nil
    }
    res["longTimeFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLongTimeFormat(val)
        return nil
    }
    res["shortDateFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShortDateFormat(val)
        return nil
    }
    res["shortTimeFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShortTimeFormat(val)
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTimeZone(val)
        return nil
    }
    return res
}
func (m *RegionalFormatOverrides) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RegionalFormatOverrides) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firstDayOfWeek", m.GetFirstDayOfWeek())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("longDateFormat", m.GetLongDateFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("longTimeFormat", m.GetLongTimeFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("shortDateFormat", m.GetShortDateFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("shortTimeFormat", m.GetShortTimeFormat())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RegionalFormatOverrides) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the calendar property value. The calendar to use, e.g., Gregorian Calendar.Returned by default.
// Parameters:
//  - value : Value to set for the calendar property.
func (m *RegionalFormatOverrides) SetCalendar(value *string)() {
    m.calendar = value
}
// Sets the firstDayOfWeek property value. The first day of the week to use, e.g., Sunday.Returned by default.
// Parameters:
//  - value : Value to set for the firstDayOfWeek property.
func (m *RegionalFormatOverrides) SetFirstDayOfWeek(value *string)() {
    m.firstDayOfWeek = value
}
// Sets the longDateFormat property value. The long date time format to be used for displaying dates.Returned by default.
// Parameters:
//  - value : Value to set for the longDateFormat property.
func (m *RegionalFormatOverrides) SetLongDateFormat(value *string)() {
    m.longDateFormat = value
}
// Sets the longTimeFormat property value. The long time format to be used for displaying time.Returned by default.
// Parameters:
//  - value : Value to set for the longTimeFormat property.
func (m *RegionalFormatOverrides) SetLongTimeFormat(value *string)() {
    m.longTimeFormat = value
}
// Sets the shortDateFormat property value. The short date time format to be used for displaying dates.Returned by default.
// Parameters:
//  - value : Value to set for the shortDateFormat property.
func (m *RegionalFormatOverrides) SetShortDateFormat(value *string)() {
    m.shortDateFormat = value
}
// Sets the shortTimeFormat property value. The short time format to be used for displaying time.Returned by default.
// Parameters:
//  - value : Value to set for the shortTimeFormat property.
func (m *RegionalFormatOverrides) SetShortTimeFormat(value *string)() {
    m.shortTimeFormat = value
}
// Sets the timeZone property value. The timezone to be used for displaying time.Returned by default.
// Parameters:
//  - value : Value to set for the timeZone property.
func (m *RegionalFormatOverrides) SetTimeZone(value *string)() {
    m.timeZone = value
}
