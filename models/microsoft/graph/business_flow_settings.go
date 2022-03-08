package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BusinessFlowSettings provides operations to manage the collection of approvalWorkflowProvider entities.
type BusinessFlowSettings struct {
    AccessReviewSettings
    // 
    durationInDays *int32;
}
// NewBusinessFlowSettings instantiates a new businessFlowSettings and sets the default values.
func NewBusinessFlowSettings()(*BusinessFlowSettings) {
    m := &BusinessFlowSettings{
        AccessReviewSettings: *NewAccessReviewSettings(),
    }
    return m
}
// CreateBusinessFlowSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessFlowSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBusinessFlowSettings(), nil
}
// GetDurationInDays gets the durationInDays property value. 
func (m *BusinessFlowSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessFlowSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AccessReviewSettings.GetFieldDeserializers()
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    return res
}
func (m *BusinessFlowSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDurationInDays sets the durationInDays property value. 
func (m *BusinessFlowSettings) SetDurationInDays(value *int32)() {
    if m != nil {
        m.durationInDays = value
    }
}
