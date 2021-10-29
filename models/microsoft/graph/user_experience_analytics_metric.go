package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsMetric struct {
    Entity
    // The unit of the user experience analytics metric.
    unit *string;
    // The value of the user experience analytics metric.
    value *float64;
}
// Instantiates a new userExperienceAnalyticsMetric and sets the default values.
func NewUserExperienceAnalyticsMetric()(*UserExperienceAnalyticsMetric) {
    m := &UserExperienceAnalyticsMetric{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the unit property value. The unit of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) GetUnit()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unit
    }
}
// Gets the value property value. The value of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) GetValue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsMetric) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["unit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUnit(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsMetric) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsMetric) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("unit", m.GetUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the unit property value. The unit of the user experience analytics metric.
// Parameters:
//  - value : Value to set for the unit property.
func (m *UserExperienceAnalyticsMetric) SetUnit(value *string)() {
    m.unit = value
}
// Sets the value property value. The value of the user experience analytics metric.
// Parameters:
//  - value : Value to set for the value property.
func (m *UserExperienceAnalyticsMetric) SetValue(value *float64)() {
    m.value = value
}
