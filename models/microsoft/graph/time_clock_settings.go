package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TimeClockSettings struct {
    additionalData map[string]interface{};
    approvedLocation *GeoCoordinates;
}
func NewTimeClockSettings()(*TimeClockSettings) {
    m := &TimeClockSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TimeClockSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TimeClockSettings) GetApprovedLocation()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.approvedLocation
    }
}
func (m *TimeClockSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approvedLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeoCoordinates() })
        if err != nil {
            return err
        }
        m.SetApprovedLocation(val.(*GeoCoordinates))
        return nil
    }
    return res
}
func (m *TimeClockSettings) IsNil()(bool) {
    return m == nil
}
func (m *TimeClockSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("approvedLocation", m.GetApprovedLocation())
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
func (m *TimeClockSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TimeClockSettings) SetApprovedLocation(value *GeoCoordinates)() {
    m.approvedLocation = value
}
