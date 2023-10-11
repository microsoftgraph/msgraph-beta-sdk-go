package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// RegionalFormatOverrides 
type RegionalFormatOverrides struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRegionalFormatOverrides instantiates a new regionalFormatOverrides and sets the default values.
func NewRegionalFormatOverrides()(*RegionalFormatOverrides) {
    m := &RegionalFormatOverrides{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRegionalFormatOverridesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegionalFormatOverridesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegionalFormatOverrides(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegionalFormatOverrides) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *RegionalFormatOverrides) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCalendar gets the calendar property value. The calendar to use, e.g., Gregorian Calendar.Returned by default.
func (m *RegionalFormatOverrides) GetCalendar()(*string) {
    val, err := m.GetBackingStore().Get("calendar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegionalFormatOverrides) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val)
        }
        return nil
    }
    res["firstDayOfWeek"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstDayOfWeek(val)
        }
        return nil
    }
    res["longDateFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongDateFormat(val)
        }
        return nil
    }
    res["longTimeFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongTimeFormat(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["shortDateFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortDateFormat(val)
        }
        return nil
    }
    res["shortTimeFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    val, err := m.GetBackingStore().Get("firstDayOfWeek")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLongDateFormat gets the longDateFormat property value. The long date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) GetLongDateFormat()(*string) {
    val, err := m.GetBackingStore().Get("longDateFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLongTimeFormat gets the longTimeFormat property value. The long time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetLongTimeFormat()(*string) {
    val, err := m.GetBackingStore().Get("longTimeFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RegionalFormatOverrides) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetShortDateFormat gets the shortDateFormat property value. The short date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) GetShortDateFormat()(*string) {
    val, err := m.GetBackingStore().Get("shortDateFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetShortTimeFormat gets the shortTimeFormat property value. The short time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetShortTimeFormat()(*string) {
    val, err := m.GetBackingStore().Get("shortTimeFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTimeZone gets the timeZone property value. The timezone to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) GetTimeZone()(*string) {
    val, err := m.GetBackingStore().Get("timeZone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegionalFormatOverrides) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *RegionalFormatOverrides) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCalendar sets the calendar property value. The calendar to use, e.g., Gregorian Calendar.Returned by default.
func (m *RegionalFormatOverrides) SetCalendar(value *string)() {
    err := m.GetBackingStore().Set("calendar", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstDayOfWeek sets the firstDayOfWeek property value. The first day of the week to use, e.g., Sunday.Returned by default.
func (m *RegionalFormatOverrides) SetFirstDayOfWeek(value *string)() {
    err := m.GetBackingStore().Set("firstDayOfWeek", value)
    if err != nil {
        panic(err)
    }
}
// SetLongDateFormat sets the longDateFormat property value. The long date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) SetLongDateFormat(value *string)() {
    err := m.GetBackingStore().Set("longDateFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetLongTimeFormat sets the longTimeFormat property value. The long time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) SetLongTimeFormat(value *string)() {
    err := m.GetBackingStore().Set("longTimeFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RegionalFormatOverrides) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetShortDateFormat sets the shortDateFormat property value. The short date time format to be used for displaying dates.Returned by default.
func (m *RegionalFormatOverrides) SetShortDateFormat(value *string)() {
    err := m.GetBackingStore().Set("shortDateFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetShortTimeFormat sets the shortTimeFormat property value. The short time format to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) SetShortTimeFormat(value *string)() {
    err := m.GetBackingStore().Set("shortTimeFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeZone sets the timeZone property value. The timezone to be used for displaying time.Returned by default.
func (m *RegionalFormatOverrides) SetTimeZone(value *string)() {
    err := m.GetBackingStore().Set("timeZone", value)
    if err != nil {
        panic(err)
    }
}
// RegionalFormatOverridesable 
type RegionalFormatOverridesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCalendar()(*string)
    GetFirstDayOfWeek()(*string)
    GetLongDateFormat()(*string)
    GetLongTimeFormat()(*string)
    GetOdataType()(*string)
    GetShortDateFormat()(*string)
    GetShortTimeFormat()(*string)
    GetTimeZone()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCalendar(value *string)()
    SetFirstDayOfWeek(value *string)()
    SetLongDateFormat(value *string)()
    SetLongTimeFormat(value *string)()
    SetOdataType(value *string)()
    SetShortDateFormat(value *string)()
    SetShortTimeFormat(value *string)()
    SetTimeZone(value *string)()
}
