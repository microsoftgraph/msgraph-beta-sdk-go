package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RegionalFormatOverrides struct {
    additionalData map[string]interface{};
    calendar *string;
    firstDayOfWeek *string;
    longDateFormat *string;
    longTimeFormat *string;
    shortDateFormat *string;
    shortTimeFormat *string;
    timeZone *string;
}
func NewRegionalFormatOverrides()(*RegionalFormatOverrides) {
    m := &RegionalFormatOverrides{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RegionalFormatOverrides) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RegionalFormatOverrides) GetCalendar()(*string) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
func (m *RegionalFormatOverrides) GetFirstDayOfWeek()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firstDayOfWeek
    }
}
func (m *RegionalFormatOverrides) GetLongDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.longDateFormat
    }
}
func (m *RegionalFormatOverrides) GetLongTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.longTimeFormat
    }
}
func (m *RegionalFormatOverrides) GetShortDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortDateFormat
    }
}
func (m *RegionalFormatOverrides) GetShortTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortTimeFormat
    }
}
func (m *RegionalFormatOverrides) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
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
func (m *RegionalFormatOverrides) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RegionalFormatOverrides) SetCalendar(value *string)() {
    m.calendar = value
}
func (m *RegionalFormatOverrides) SetFirstDayOfWeek(value *string)() {
    m.firstDayOfWeek = value
}
func (m *RegionalFormatOverrides) SetLongDateFormat(value *string)() {
    m.longDateFormat = value
}
func (m *RegionalFormatOverrides) SetLongTimeFormat(value *string)() {
    m.longTimeFormat = value
}
func (m *RegionalFormatOverrides) SetShortDateFormat(value *string)() {
    m.shortDateFormat = value
}
func (m *RegionalFormatOverrides) SetShortTimeFormat(value *string)() {
    m.shortTimeFormat = value
}
func (m *RegionalFormatOverrides) SetTimeZone(value *string)() {
    m.timeZone = value
}
