package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TimeClockSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The aprroved location of the timeClock.
    approvedLocation *GeoCoordinates;
}
// Instantiates a new timeClockSettings and sets the default values.
func NewTimeClockSettings()(*TimeClockSettings) {
    m := &TimeClockSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeClockSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the approvedLocation property value. The aprroved location of the timeClock.
func (m *TimeClockSettings) GetApprovedLocation()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.approvedLocation
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TimeClockSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the approvedLocation property value. The aprroved location of the timeClock.
// Parameters:
//  - value : Value to set for the approvedLocation property.
func (m *TimeClockSettings) SetApprovedLocation(value *GeoCoordinates)() {
    m.approvedLocation = value
}
