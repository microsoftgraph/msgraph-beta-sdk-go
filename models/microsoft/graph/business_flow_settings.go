package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BusinessFlowSettings struct {
    AccessReviewSettings
    // 
    durationInDays *int32;
}
// Instantiates a new businessFlowSettings and sets the default values.
func NewBusinessFlowSettings()(*BusinessFlowSettings) {
    m := &BusinessFlowSettings{
        AccessReviewSettings: *NewAccessReviewSettings(),
    }
    return m
}
// Gets the durationInDays property value. 
func (m *BusinessFlowSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// The deserialization information for the current model
func (m *BusinessFlowSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AccessReviewSettings.GetFieldDeserializers()
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDurationInDays(val)
        return nil
    }
    return res
}
func (m *BusinessFlowSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BusinessFlowSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AccessReviewSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the durationInDays property value. 
// Parameters:
//  - value : Value to set for the durationInDays property.
func (m *BusinessFlowSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
