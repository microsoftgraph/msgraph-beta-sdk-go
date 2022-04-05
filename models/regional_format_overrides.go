package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegionalFormatOverrides 
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
// NewRegionalFormatOverrides instantiates a new regionalFormatOverrides and sets the default values.
func NewRegionalFormatOverrides()(*RegionalFormatOverrides) {
    m := &RegionalFormatOverrides{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRegionalFormatOverridesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegionalFormatOverridesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegionalFormatOverrides(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegionalFormatOverrides) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCalendar gets the calendar property value. The calendar to use, e.g., Gregorian Calendar.Returned by default.
func (m *RegionalFormatOverrides) GetCalendar()(*string) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegionalFormatOverrides) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calendar"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val)
        }
        return nil
    }
    res["firstDayOfWeek"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstDayOfWeek(val)
        }
        return nil
    }
    res["longDateFormat"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongDateFormat(val)
        }
        return nil
    }
    res["longTimeFormat"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongTimeFormat(val)
        }
        return nil
    }
    res["shortDateFormat"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortDateFormat(val)
        }
        return nil
    }
    res["shortTimeFormat"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFirstDayOfWeek gets the firstDayOfWeek property value. The first day of the week to use, e.g., Sunday.Returned by default.
func (m *RegionalFormatOverrides) GetFirstDayOfWeek()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firstDayOfWeek
    }
}
// GetLongDateFormat gets the longDateFormat property value. The long date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) GetLongDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.longDateFormat
    }
}
// GetLongTimeFormat gets the longTimeFormat property value. The long time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetLongTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.longTimeFormat
    }
}
// GetShortDateFormat gets the shortDateFormat property value. The short date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) GetShortDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortDateFormat
    }
}
// GetShortTimeFormat gets the shortTimeFormat property value. The short time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetShortTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortTimeFormat
    }
}
// GetTimeZone gets the timeZone property value. The timezone to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// Serialize serializes information the current object
func (m *RegionalFormatOverrides) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegionalFormatOverrides) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCalendar sets the calendar property value. The calendar to use, e.g., Gregorian Calendar.Returned by default.
func (m *RegionalFormatOverrides) SetCalendar(value *string)() {
    if m != nil {
        m.calendar = value
    }
}
// SetFirstDayOfWeek sets the firstDayOfWeek property value. The first day of the week to use, e.g., Sunday.Returned by default.
func (m *RegionalFormatOverrides) SetFirstDayOfWeek(value *string)() {
    if m != nil {
        m.firstDayOfWeek = value
    }
}
// SetLongDateFormat sets the longDateFormat property value. The long date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) SetLongDateFormat(value *string)() {
    if m != nil {
        m.longDateFormat = value
    }
}
// SetLongTimeFormat sets the longTimeFormat property value. The long time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) SetLongTimeFormat(value *string)() {
    if m != nil {
        m.longTimeFormat = value
    }
}
// SetShortDateFormat sets the shortDateFormat property value. The short date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) SetShortDateFormat(value *string)() {
    if m != nil {
        m.shortDateFormat = value
    }
}
// SetShortTimeFormat sets the shortTimeFormat property value. The short time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) SetShortTimeFormat(value *string)() {
    if m != nil {
        m.shortTimeFormat = value
    }
}
// SetTimeZone sets the timeZone property value. The timezone to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) SetTimeZone(value *string)() {
    if m != nil {
        m.timeZone = value
    }
}
